// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-10-09 10:49:21.680457 +0800 CST m=+0.099211474

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "description": "根据条件获取信息",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users表操作"
                ],
                "summary": "【GetAll】根据条件获取信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页数，默认1",
                        "name": "page_num",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每夜条数，默认50",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序。id desc, time asc",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SgrResp"
                        }
                    }
                }
            },
            "post": {
                "description": "创建users信息",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users表操作"
                ],
                "summary": "【create】创建users信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SgrResp"
                        }
                    }
                }
            }
        },
        "/users/update-by-id": {
            "put": {
                "description": "根据id更新数据",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users表操作"
                ],
                "summary": "【update】根据id更新数据",
                "parameters": [
                    {
                        "description": "主键更新依据此id",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SgrResp"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "根据id获取信息",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users表操作"
                ],
                "summary": "【GetOne】根据id获取信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "主键id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SgrResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.SgrResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 1000
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string",
                    "example": "请求成功"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "work_report文档平台",
	Description: "work_report自动文档",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
