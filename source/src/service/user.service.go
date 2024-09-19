package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"todolist-api/src/model"
	"todolist-api/src/model/enum"
	"todolist-api/src/util/umongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

//* Change o service content with ALT+C(Case sensitive) then CTRL+D o:
//? user and User

type UserService struct {
	collectionName string
	ctx            context.Context
	dbUtil         *umongo.MongoDbUtil
	dbUtilView     *umongo.MongoDbUtil
	passwordCost   int
}

func NewUserService() *UserService {
	service := &UserService{
		collectionName: enum.MongoCollection_User.String(),
		ctx:            context.Background(),
		passwordCost:   15,
	}
	service.dbUtil = umongo.NewMongoDbUtilUseEnv(service.collectionName)
	service.dbUtilView = umongo.NewMongoDbUtilUseEnv("v_" + service.collectionName)

	return service
}

func (us *UserService) BaseGetAll(param model.User_Search, collection *umongo.MongoDbUtil) (data []model.User_View,
	metadata model.MetadataResponse,
) {
	filter := bson.M{}
	listFilterAnd := []bson.M{}
	param.HandleFilter(&listFilterAnd)

	if len(listFilterAnd) > 0 {
		filter["$and"] = listFilterAnd
	}
	metadata.Pagination, metadata.Message = collection.Find(filter,
		param.Request, &data)
	return
}

func (us *UserService) GetAll(param model.User_Search) (data []model.User_View, metadata model.MetadataResponse) {
	return us.BaseGetAll(param, us.dbUtilView)
}

func (us *UserService) GetOne(key, value string) (res model.User_View, errMessage string) {
	errMessage = OverrideError(us.dbUtilView.FindOne(key, value, &res))
	return
}

func (us *UserService) getValidUsername(username string) string {
	return strings.ToLower(username)
}

func (us *UserService) preparePassword(param *model.User) (err error) {
	if param.Password == "" {
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), us.passwordCost)
	if err != nil {
		log.Println(err)
		return errors.New("fail to hash password")
	}

	param.Password = string(hashPassword)
	return
}

func (service *UserService) UpsertWithHashingPassword(param model.User, isUpdate bool) (resp model.Response) {
	param.Username = service.getValidUsername(param.Username)
	if err := service.preparePassword(&param); err != nil {
		resp.Metadata.Message = err.Error()
		return
	}

	if err := service.dbUtil.CheckDuplicate(param.IdDocument, []bson.M{
		{"email": *param.Email},
	}); err != nil {
		resp.Metadata.Message = err.Error()
		return
	}

	upsertId, err := service.dbUtil.UpsertAndGetId(isUpdate, &param)
	if !isUpdate && err != nil {
		return
	}

	resp.Metadata.Message = umongo.GetErrForResponse(err)
	resp.Data = model.Response_Data_Upsert{
		ID: upsertId,
	}
	return
}

func (us *UserService) EditProfile(ctx *gin.Context, param model.EditProfile) (resp model.Response) {
	uid, _ := ctx.Get("uid")
	param.IdDocument = fmt.Sprint(uid)
	param.Username = us.getValidUsername(param.Username)

	_, err := us.dbUtil.UpsertAndGetId(true, &param)
	resp.Metadata.Message = umongo.GetErrForResponse(err)
	resp.Data = nil

	return
}

func (service *UserService) DeleteOne(key, value string) (errMessage string) {
	errMessage = umongo.GetErrForResponse(service.dbUtil.HardDelete(bson.M{key: value}))
	return
}

func (us *UserService) GetUserWithValidatePassword(key, value,
	passwordToTest string, ptrUser *model.User,
) (resp model.Response, ok bool) {
	if err := us.dbUtil.BaseFindOne(bson.M{key: value}, &ptrUser); err != nil {
		log.Println(err)
		resp.Metadata.Message = "Data not found"
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(ptrUser.Password), []byte(passwordToTest)); err != nil {
		log.Println(err)
		resp.Metadata.Message = "Password doesn't match"
		return
	}

	ok = true
	//? Resp by view, due to look up topic status that calculated in view
	if err := us.dbUtilView.BaseFindOne(bson.M{key: value}, &ptrUser); err != nil {
		log.Println(err)
		resp.Metadata.Message = "Data not found"
		return
	}
	return
}
