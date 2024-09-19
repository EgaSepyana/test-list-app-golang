// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/checklist": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "order",
                        "name": "order",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "orderBy",
                        "name": "orderBy",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "data": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/model.Checklist_View"
                                    }
                                },
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Checklist"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/checklist/{checklistId}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "data": {
                                    "$ref": "#/definitions/model.Checklist_View"
                                },
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/checklist/{checklistId}/item": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "order",
                        "name": "order",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "orderBy",
                        "name": "orderBy",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "data": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/model.Checklist_View"
                                    }
                                },
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChecklistItem"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/checklist/{checklistId}/item/rename/{checklistitemId}": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistitemId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CheklistItem_Update_name"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/checklist/{checklistId}/item/{checklistitemId}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistitemId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "data": {
                                    "$ref": "#/definitions/model.Checklist_View"
                                },
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistitemId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CheklistItem_Update_status"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checklist"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "checklistitemId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User_Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/refresh": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/user/add": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/user/edit-profile": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EditProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/user/get-all": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User_Search"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "data": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/model.User_View"
                                    }
                                },
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/user/get-one": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "data": {
                                    "$ref": "#/definitions/model.User_View"
                                },
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "data": {
                                    "$ref": "#/definitions/model.User_View"
                                },
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/user/update": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "meta_data": {
                                    "$ref": "#/definitions/model.MetadataResponse"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Checklist": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.ChecklistItem": {
            "type": "object",
            "properties": {
                "Itemname": {
                    "type": "string"
                },
                "checklist_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "model.Checklist_View": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.CheklistItem_Update_name": {
            "type": "object",
            "properties": {
                "Itemname": {
                    "type": "string"
                }
            }
        },
        "model.CheklistItem_Update_status": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "boolean"
                }
            }
        },
        "model.EditProfile": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "dateOfBirth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.MetadataResponse": {
            "type": "object",
            "properties": {
                "additional_data": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "message": {
                    "type": "string"
                },
                "pagination": {
                    "$ref": "#/definitions/umongo.PaginationResponse"
                },
                "status": {
                    "type": "boolean"
                },
                "timeExecution": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.User_Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.User_Search": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "string",
                    "example": "DESC"
                },
                "orderBy": {
                    "type": "string",
                    "example": "createdAt"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "search": {
                    "description": "? Regex",
                    "type": "string"
                },
                "size": {
                    "type": "integer",
                    "example": 11
                }
            }
        },
        "model.User_View": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "umongo.PaginationResponse": {
            "type": "object",
            "properties": {
                "size": {
                    "type": "integer"
                },
                "totalElements": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "description": "E.g. Bearer Your.Token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
