// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-17 12:11:00.262395249 +0000 UTC m=+0.049679893

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "List APIs of MovieManagement Service",
        "title": "MovieManagement Service API Document",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "107.113.53.47:9000",
    "basePath": "/api/v1",
    "paths": {
        "/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Log in to the service",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Log in to the service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/movies": {
            "post": {
                "description": "Add a new movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Add a new movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Add Movie",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.AddMovie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/movies/list": {
            "get": {
                "description": "List all existing Movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "List all existing Movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddMovie": {
            "type": "object",
            "properties": {
                "coverImage": {
                    "type": "string",
                    "example": "Movie Cover Image"
                },
                "description": {
                    "type": "string",
                    "example": "Movie Description"
                },
                "name": {
                    "type": "string",
                    "example": "Movie Name"
                },
                "url": {
                    "type": "string",
                    "example": "Movie URL"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 27
                },
                "message": {
                    "type": "string",
                    "example": "Error message"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "coverImage": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "5bb3695b82ebac0f76e1cafa"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}