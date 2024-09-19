package controller

import (
	"log"
	"time"

	"todolist-api/src/model"
	"todolist-api/src/service"

	"github.com/gin-gonic/gin"
)

//* Change o controller content with ALT+C(Case sensitive) then CTRL+D o:
//! and don't forget to register the controller in main.go
//? checklist and Checklist

type ChecklistController struct {
	router               *gin.RouterGroup
	service              *service.ChecklistService
	checkListItemService *service.ChecklistItemService
}

func NewChecklistController(router *gin.RouterGroup) *ChecklistController {
	o := &ChecklistController{router: router, service: service.NewChecklistService(), checkListItemService: service.NewChecklistItemService()}

	checklist := o.router.Group("/checklist")
	checklist.GET("/", o.GetAll)
	checklist.GET("/:checklistId", o.GetOne)
	checklist.POST("/", o.Add)
	checklist.DELETE("/:checklistId", o.DeleteOne)
	checklist.GET("/:checklistId/item", o.GetAllChecklistItem)
	checklist.GET("/:checklistId/item/:checklistitemId", o.GetOneCheckListItem)
	checklist.POST("/:checklistId/item", o.AddCheckListItem)
	checklist.PUT("/:checklistId/item/:checklistitemId", o.UpdateNameItem)
	checklist.PUT("/:checklistId/item/rename/:checklistitemId", o.UpdateStatusItem)
	checklist.DELETE("/:checklistId/item/:checklistitemId", o.DeleteOneCheckListItem)

	return o
}

// @Tags Checklist
// @Accept json
// @Produce json
// @Success 200 {object} object{data=[]model.Checklist_View,meta_data=model.MetadataResponse} "OK"
// @Router /checklist [get]
// @Security JWT
func (o *ChecklistController) GetAll(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	param := model.Checklist_Search{}

	resp.Data, resp.Metadata = o.service.GetAll(param)
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Produce json
// @Success 200 {object} object{data=[]model.Checklist_View,meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId}/item [get]
// @Security JWT
func (o *ChecklistController) GetAllChecklistItem(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	param := model.ChecklistItem_Search{}
	param.CheckListId = ctx.Param("checklistId")

	resp.Data, resp.Metadata = o.checkListItemService.GetAll(param)
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Produce json
// @Success 200 {object} object{data=model.Checklist_View,meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId} [get]
// @Security JWT
func (o *ChecklistController) GetOne(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	resp.Data, resp.Metadata.Message = o.service.GetOne("_id", ctx.Param("checklistId"))
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Param checklistitemId path string true "ID"
// @Produce json
// @Success 200 {object} object{data=model.Checklist_View,meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId}/item/{checklistitemId} [get]
// @Security JWT
func (o *ChecklistController) GetOneCheckListItem(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	resp.Data, resp.Metadata.Message = o.checkListItemService.GetOne(ctx.Param("checklistId"), ctx.Param("checklistitemId"))
}

// @Tags Checklist
// @Accept json
// @Param parameter body model.Checklist true "PARAM"
// @Produce json
// @Success 201 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /checklist [post]
// @Security JWT
func (o *ChecklistController) Add(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.Checklist

	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}

	resp = o.service.Upsert(param, false)
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Param checklistitemId path string true "ID"
// @Param parameter body model.CheklistItem_Update_name true "PARAM"
// @Produce json
// @Success 201 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId}/item/{checklistitemId} [put]
// @Security JWT
func (o *ChecklistController) UpdateStatusItem(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.CheklistItem_Update_name
	// param.Status =

	var upserData model.ChecklistItem
	upserData.ItemName = param.Itemname

	oldData, _ := o.checkListItemService.GetOne(ctx.Param("checklistId"), ctx.Param("checklistitemId"))
	upserData.CheckListId = oldData.CheckListId
	upserData.IdDocument = oldData.IdDocument
	upserData.Description = oldData.Description
	upserData.ItemName = oldData.ItemName
	upserData.CreatedAt = oldData.CreatedAt
	upserData.UpdatedAt = oldData.UpdatedAt

	// upserData.CheckListId = oldData.

	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}

	resp = o.checkListItemService.Upsert(upserData, false)
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Param checklistitemId path string true "ID"
// @Param parameter body model.CheklistItem_Update_status true "PARAM"
// @Produce json
// @Success 201 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId}/item/rename/{checklistitemId} [put]
// @Security JWT
func (o *ChecklistController) UpdateNameItem(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.CheklistItem_Update_status
	// param.Status =

	var upserData model.ChecklistItem
	upserData.Status = param.Status

	oldData, _ := o.checkListItemService.GetOne(ctx.Param("checklistId"), ctx.Param("checklistitemId"))
	upserData.CheckListId = oldData.CheckListId
	upserData.IdDocument = oldData.IdDocument
	upserData.Description = oldData.Description
	upserData.ItemName = oldData.ItemName
	upserData.CreatedAt = oldData.CreatedAt
	upserData.UpdatedAt = oldData.UpdatedAt

	// upserData.CheckListId = oldData.

	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}

	resp = o.checkListItemService.Upsert(upserData, false)
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Param parameter body model.ChecklistItem true "PARAM"
// @Produce json
// @Success 201 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId}/item [post]
// @Security JWT
func (o *ChecklistController) AddCheckListItem(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.ChecklistItem
	param.CheckListId = ctx.Param("checklistId")

	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}

	resp = o.checkListItemService.Upsert(param, false)
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId} [delete]
// @Security JWT
func (o *ChecklistController) DeleteOne(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	resp.Metadata.Message = o.service.DeleteOne("_id", ctx.Param("checklistId"))
}

// @Tags Checklist
// @Accept json
// @Param checklistId path string true "ID"
// @Param checklistitemId path string true "ID"
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /checklist/{checklistId}/item/{checklistitemId} [delete]
// @Security JWT
func (o *ChecklistController) DeleteOneCheckListItem(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	resp.Metadata.Message = o.checkListItemService.DeleteOne(ctx.Param("checklistId"), ctx.Param("checklistitemId"))
}
