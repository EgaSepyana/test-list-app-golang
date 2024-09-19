package controller

import (
	"todolist-api/src/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//* Change o controller content with ALT+C(Case sensitive) then CTRL+D o:
//! and don't forget to register the controller in main.go
//? public and Public

type PublicController struct {
	router *gin.RouterGroup
}

func NewPublicController(router *gin.RouterGroup, jwtMiddleware *jwt.GinJWTMiddleware) *PublicController {
	controller := &PublicController{router: router}
	router.POST("/login", jwtMiddleware.LoginHandler)
	router.GET("/refresh", jwtMiddleware.RefreshHandler)
	router.POST("/register", controller.Register)
	return controller
}

// @Tags public
// @Accept json
// @Param parameter body model.User_Login true "PARAM"
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /login [post]
func (o *PublicController) Login(ctx *gin.Context) {}

// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /refresh [get]
// @Security JWT
func (o *PublicController) Refresh(ctx *gin.Context) {}

// @Tags public
// @Accept json
// @Param parameter body model.User true "PARAM"
// @Produce json
// @Success 201 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /register [post]
func (o *PublicController) Register(ctx *gin.Context) {
	UserController{service: service.NewUserService()}.Add(ctx)
}
