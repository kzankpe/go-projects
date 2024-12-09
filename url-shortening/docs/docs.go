// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "kzankpe",
            "url": "https://github.com/kzankpe"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shorten": {
            "post": {
                "description": "Create new short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortening URL"
                ],
                "summary": "Create new short url",
                "parameters": [
                    {
                        "description": "Long Url",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LongUrl"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/shorten/{shortcode}": {
            "get": {
                "description": "Retrieve  short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortening URL"
                ],
                "summary": "Retrieve  short url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short Code",
                        "name": "shortcode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update  short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortening URL"
                ],
                "summary": "Update  short url",
                "parameters": [
                    {
                        "description": "Short Code",
                        "name": "shortcode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UrlData"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete  short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortening URL"
                ],
                "summary": "Delete  short url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short Code",
                        "name": "shortcode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/shorten/{shortcode}/stats": {
            "get": {
                "description": "Get  short url statistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortening URL"
                ],
                "summary": "Get   short url stats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short Code",
                        "name": "url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.LongUrl": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "models.UrlData": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "shortCode": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Shorten Url Service",
	Description:      "A simple RESTful API that allows users to shorten long URLs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}