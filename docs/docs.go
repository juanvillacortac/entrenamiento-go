// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/auth/register": {
            "post": {
                "description": "Responds with the list of all books as JSON.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Payload",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserLogin"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/search": {
            "get": {
                "description": "Responds with the list of all books as JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Get songs array",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Album name",
                        "name": "album",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Artist name",
                        "name": "artist",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Bypass cache",
                        "name": "force",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Song"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Song": {
            "type": "object",
            "properties": {
                "album": {
                    "type": "string"
                },
                "artist": {
                    "type": "string"
                },
                "artwork": {
                    "type": "string"
                },
                "duration": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                }
            }
        },
        "entities.UserLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
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
	Title:            "Songs Indexer",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}