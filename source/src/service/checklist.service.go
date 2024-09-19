package service

import (
	"context"
	"todolist-api/src/model"
	"todolist-api/src/model/enum"
	"todolist-api/src/util/umongo"

	"go.mongodb.org/mongo-driver/bson"
)

//* Change s service content with ALT+C(Case sensitive) then CTRL+D s:
//? Checklist and Checklist

type ChecklistService struct {
	collectionName string
	ctx            context.Context
	dbUtil         *umongo.MongoDbUtil
	dbUtilView     *umongo.MongoDbUtil
}

func NewChecklistService() *ChecklistService {
	s := &ChecklistService{
		collectionName: enum.MongoCollection_Checklist.String(),
		ctx:            context.Background(),
	}
	s.dbUtil = umongo.NewMongoDbUtilUseEnv(s.collectionName)
	s.dbUtilView = umongo.NewMongoDbUtilUseEnv("v_" + s.collectionName)
	s.dbUtilView = s.dbUtil

	return s
}

func (s *ChecklistService) BaseGetAll(param model.Checklist_Search, collection *umongo.MongoDbUtil) (data []model.Checklist_View,
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

func (s *ChecklistService) GetAll(param model.Checklist_Search) (data []model.Checklist_View, metadata model.MetadataResponse) {
	return s.BaseGetAll(param, s.dbUtilView)
}

func (s *ChecklistService) GetOne(key, value string) (res model.Checklist_View, errMessage string) {
	errMessage = OverrideError(s.dbUtilView.FindOne(key, value, &res))
	return
}

func (s *ChecklistService) Upsert(param model.Checklist, isUpdate bool) (resp model.Response) {
	upsertId, err := s.dbUtil.UpsertAndGetId(isUpdate, &param)
	resp.Metadata.Message = umongo.GetErrForResponse(err)
	resp.Data = model.Response_Data_Upsert{
		ID: upsertId,
	}

	return
}

func (s *ChecklistService) DeleteOne(key, value string) (errMessage string) {
	errMessage = umongo.GetErrForResponse(s.dbUtil.HardDelete(bson.M{key: value}))
	return
}
