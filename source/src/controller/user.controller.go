package controller

import (
	"fmt"
	"log"
	"time"

	"todolist-api/src/model"
	"todolist-api/src/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//* Change o controller content with ALT+C(Case sensitive) then CTRL+D o:
//! and don't forget to register the controller in main.go
//? user and User

type UserController struct {
	router  *gin.RouterGroup
	service *service.UserService
}

func NewUserController(router *gin.RouterGroup) *UserController {
	controller := &UserController{router: router, service: service.NewUserService()}

	user := controller.router.Group("/user")
	user.POST("/get-all", controller.GetAll)
	user.GET("/get-one", controller.GetOne)
	user.POST("/add", controller.Add)
	user.PUT("/update", controller.Update)
	user.PUT("/edit-profile", controller.EditProfile)
	user.DELETE("/delete", controller.DeleteOne)
	user.GET("/profile", controller.GetOneByBearer)

	return controller
}

// @Tags user
// @Accept json
// @Param parameter body model.User_Search true "PARAM"
// @Produce json
// @Success 200 {object} object{data=[]model.User_View,meta_data=model.MetadataResponse} "OK"
// @Router /user/get-all [post]
// @Security JWT
func (o *UserController) GetAll(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.User_Search
	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}

	resp.Data, resp.Metadata = o.service.GetAll(param)
}

// @Tags user
// @Accept json
// @Param id query string true "ID"
// @Produce json
// @Success 200 {object} object{data=model.User_View,meta_data=model.MetadataResponse} "OK"
// @Router /user/get-one [get]
// @Security JWT
func (o *UserController) GetOne(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	resp.Data, resp.Metadata.Message = o.service.GetOne("_id", ctx.Query("id"))
}

// @Tags user
// @Accept json
// @Param parameter body model.User true "PARAM"
// @Produce json
// @Success 201 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /user/add [post]
// @Security JWT
func (controller UserController) Add(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.User
	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}
	resp = controller.service.UpsertWithHashingPassword(param, false)
}

// @Tags user
// @Accept json
// @Param parameter body model.User true "PARAM"
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /user/update [put]
// @Security JWT
func (controller *UserController) Update(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.User
	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}

	resp = controller.service.UpsertWithHashingPassword(param, true)
}

// @Tags user
// @Accept json
// @Param parameter body model.EditProfile true "PARAM"
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /user/edit-profile [put]
// @Security JWT
func (controller *UserController) EditProfile(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.EditProfile
	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}

	resp = controller.service.EditProfile(ctx, param)
}

// @Tags user
// @Accept json
// @Param id query string true "ID"
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /user/delete [delete]
// @Security JWT
func (controller *UserController) DeleteOne(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)
	_, errMessage := controller.service.GetOne("_id", ctx.Query("id"))
	if errMessage != "" {
		resp.Metadata.Message = errMessage
		return
	}
	resp.Metadata.Message = controller.service.DeleteOne("_id", ctx.Query("id"))
}

// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} object{data=model.User_View,meta_data=model.MetadataResponse} "OK"
// @Router /user/profile [get]
// @Security JWT
func (controller *UserController) GetOneByBearer(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var idUser string
	if claim, found := ctx.Get("JWT_PAYLOAD"); found {
		asMap, _ := claim.(jwt.MapClaims)
		idUser = fmt.Sprint(asMap["id"])
	}
	resp.Data, resp.Metadata.Message = controller.service.GetOne("_id", idUser)
}
