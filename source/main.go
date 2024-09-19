package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"todolist-api/src/config"
	"todolist-api/src/controller"
	"todolist-api/src/middleware"

	"todolist-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/logrusorgru/aurora"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	for _, arg := range os.Args {
		if arg == "--parseDoc" || arg == "-pd" {
			swag := exec.Command("swag", "i", "--parseVendor")
			swag.Stdout, swag.Stderr = os.Stdout, os.Stderr
			swag.Run()
		}
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Using timezone:", aurora.Green(time.Now().Location().String()))
}

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description E.g. Bearer Your.Token
func main() {
	appPort := ":" + os.Getenv(config.ENV_KEY_PORT)
	if os.Getenv("DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println(aurora.Red("DEBUG"))
	}
	router := gin.Default()

	docs.SwaggerInfo.Title = "Todo List Management API"
	docs.SwaggerInfo.Description = "API documentation for todo list aplication"

	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(cors.Default())

	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.StaticFile("/rapidoc.html", "docs/rapidoc.html")

	apiV1 := router.Group(basePath)
	apiV1.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Imusegipo"}) })

	jwtMiddleware, err := middleware.InitJwt()
	if err != nil {
		log.Println(err)
		return
	}

	controller.NewPublicController(apiV1, jwtMiddleware)

	// controller.NewTopicEventController(apiV1, jwtMiddleware)
	if os.Getenv("DEBUG") != "true" {
		apiV1.Use(jwtMiddleware.MiddlewareFunc())
		apiV1.Use(middleware.SetUserInContextByJwt())
	}

	//* ------------------------ REGISTER CRUD CONTROLLER ------------------------ */
	controller.NewUserController(apiV1)
	controller.NewChecklistController(apiV1)

	//? API DOC
	log.Println(aurora.Green(fmt.Sprintf("http://localhost%s/swagger/index.html", appPort)))
	log.Println(aurora.Yellow(fmt.Sprintf("http://localhost%s/rapidoc.html\n", appPort)))

	log.Fatalln(router.Run(appPort))
}
