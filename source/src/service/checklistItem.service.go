package service

import (
	"context"
	"log"
	"todolist-api/src/model"
	"todolist-api/src/model/enum"
	"todolist-api/src/util/umongo"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
)

//* Change s service content with ALT+C(Case sensitive) then CTRL+D s:
//? ChecklistItem and ChecklistItem

type ChecklistItemService struct {
	collectionName string
	ctx            context.Context
	dbUtil         *umongo.MongoDbUtil
	dbUtilView     *umongo.MongoDbUtil
}

func NewChecklistItemService() *ChecklistItemService {
	s := &ChecklistItemService{
		collectionName: enum.MongoCollection_ChecklistItem.String(),
		ctx:            context.Background(),
	}
	s.dbUtil = umongo.NewMongoDbUtilUseEnv(s.collectionName)
	s.dbUtilView = umongo.NewMongoDbUtilUseEnv("v_" + s.collectionName)
	s.dbUtilView = s.dbUtil

	return s
}

func (s *ChecklistItemService) BaseGetAll(param model.ChecklistItem_Search, collection *umongo.MongoDbUtil) (data []model.ChecklistItem_View,
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

func (s *ChecklistItemService) GetAll(param model.ChecklistItem_Search) (data []model.ChecklistItem_View, metadata model.MetadataResponse) {
	return s.BaseGetAll(param, s.dbUtilView)
}

func (s *ChecklistItemService) GetOne(checklistId, checklistItemId string) (res model.Checklist_View, errMessage string) {
	log.Println(aurora.Green(checklistItemId))
	errMessage = OverrideError(s.dbUtilView.BaseFindOne(bson.M{"checklist_id": checklistId, "_id": checklistItemId}, &res))
	return
}

func (s *ChecklistItemService) Upsert(param model.ChecklistItem, isUpdate bool) (resp model.Response) {
	upsertId, err := s.dbUtil.UpsertAndGetId(isUpdate, &param)
	resp.Metadata.Message = umongo.GetErrForResponse(err)
	resp.Data = model.Response_Data_Upsert{
		ID: upsertId,
	}

	return
}

func (s *ChecklistItemService) DeleteOne(checklistId, checklistItemId string) (errMessage string) {
	errMessage = umongo.GetErrForResponse(s.dbUtil.HardDelete(bson.M{"checklist_id": checklistId, "_id": checklistItemId}))
	return
}
