package model

import (
	"todolist-api/src/util/umongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generated by https://quicktype.io

type Checklist struct {
	MetadataWithID `bson:",inline"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type Checklist_View struct {
	Checklist `bson:",inline"`
}

type Checklist_Search struct {
	//? Regex
	Search string `json:"search"`

	umongo.Request
}

func (o *Checklist_Search) HandleFilter(listFilterAnd *[]bson.M) {
	if search := o.Search; search != "" {
		*listFilterAnd = append(*listFilterAnd, bson.M{"name": primitive.Regex{Pattern: search, Options: "i"}})
	}
}
