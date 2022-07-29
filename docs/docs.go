// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
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
        "/auth/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "inputCredentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserWithTokensResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageUnauthorized"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Recieve new tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh",
                "parameters": [
                    {
                        "description": "RefreshToken",
                        "name": "RefreshToken",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tokens"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Registration",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "inputUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegistrationUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RegistrationUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "GetEvents by selected page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "EventsList",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page of events",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Event"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            }
        },
        "/events/{id}": {
            "get": {
                "description": "Get Event by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "One Event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Find a user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "GetUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RegistrationUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update user profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user information",
                        "name": "newUserInformation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserProfile"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageUnauthorized"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageUnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessageInternalServer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorMessageBadRequest": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Bad request"
                }
            }
        },
        "models.ErrorMessageInternalServer": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Server problems"
                }
            }
        },
        "models.ErrorMessageUnauthorized": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "User is not authorized"
                }
            }
        },
        "models.ErrorMessageUnprocessableEntity": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Wrong Json Request"
                }
            }
        },
        "models.Event": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "specialInfo": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.LoginUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "artyom@mail.ru"
                },
                "password": {
                    "type": "string",
                    "example": "12345678"
                }
            }
        },
        "models.RefreshTokenRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string",
                    "example": "4ffc5f18-99d8-47f6-8141-faf2c2f5a24e"
                }
            }
        },
        "models.RegistrationUserRequest": {
            "type": "object",
            "required": [
                "name",
                "surname"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "artyom@mail.ru"
                },
                "name": {
                    "type": "string",
                    "example": "Artyom"
                },
                "password": {
                    "type": "string",
                    "example": "12345678"
                },
                "surname": {
                    "type": "string",
                    "example": "Shirshov"
                }
            }
        },
        "models.RegistrationUserResponse": {
            "type": "object",
            "required": [
                "name",
                "surname"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "artyom@mail.ru"
                },
                "name": {
                    "type": "string",
                    "example": "Artyom"
                },
                "surname": {
                    "type": "string",
                    "example": "Shirshov"
                }
            }
        },
        "models.Tokens": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "22f37ea5-2d12-4309-afbe-17783b44e24f"
                },
                "refresh_token": {
                    "type": "string",
                    "example": "4ffc5f18-99d8-47f6-8141-faf2c2f5a24e"
                }
            }
        },
        "models.UserProfile": {
            "type": "object",
            "required": [
                "name",
                "surname"
            ],
            "properties": {
                "about": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "example": "artyom@mail.ru"
                },
                "imgUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "example": "Artyom"
                },
                "surname": {
                    "type": "string",
                    "example": "Shirshov"
                }
            }
        },
        "models.UserWithTokensResponse": {
            "type": "object",
            "properties": {
                "tokens": {
                    "$ref": "#/definitions/models.Tokens"
                },
                "user": {
                    "$ref": "#/definitions/models.UserProfile"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "http://45.141.102.243:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Diploma API",
	Description:      "Documentation for Diploma Api\nFor Authorization:\nPut Access token in ApiKey with Bearer. Example: \"Bearer access_token\"",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
