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
        "/v1/auth/forger_password": {
            "get": {
                "description": "ForgerPassword - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "ForgerPassword",
                "parameters": [
                    {
                        "type": "string",
                        "description": "phone_number",
                        "name": "phone_number",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model_user_service.MessageRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    }
                }
            }
        },
        "/v1/auth/forger_password_verify": {
            "post": {
                "description": "ForgerPasswordVerify - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "ForgerPasswordVerify",
                "parameters": [
                    {
                        "description": "RegisterModelReq",
                        "name": "Verify",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model_user_service.ForgetPasswordVerify"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model_user_service.MessageRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "description": "Login - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login Req",
                        "name": "Login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model_user_service.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model_user_service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    }
                }
            }
        },
        "/v1/auth/register/": {
            "post": {
                "description": "Register - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "RegisterRequest",
                        "name": "Register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model_user_service.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model_user_service.MessageRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    }
                }
            }
        },
        "/v1/auth/verify": {
            "post": {
                "description": "Authorization - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "Verify",
                "parameters": [
                    {
                        "description": "RegisterModelReq",
                        "name": "Verify",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model_user_service.Verify"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model_user_service.Verify"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model_common.StandardErrorModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model_common.ResponseError": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model_common.StandardErrorModel": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/model_common.ResponseError"
                }
            }
        },
        "model_user_service.ForgetPasswordVerify": {
            "type": "object",
            "properties": {
                "cade": {
                    "type": "integer"
                },
                "new_password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "model_user_service.LoginReq": {
            "type": "object",
            "properties": {
                "fcm_token": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "phone_number": {
                    "type": "string",
                    "example": "+998950230605"
                },
                "platform_name": {
                    "type": "string"
                },
                "platform_type": {
                    "type": "string"
                }
            }
        },
        "model_user_service.MessageRes": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model_user_service.RegisterRequest": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string",
                    "example": "2000-01-01"
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "phone_number": {
                    "type": "string",
                    "example": "+998950230605"
                }
            }
        },
        "model_user_service.Response": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "model_user_service.Verify": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 7777
                },
                "fcm_token": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string",
                    "example": "+998950230605"
                },
                "platform_name": {
                    "type": "string"
                },
                "platform_type": {
                    "type": "string"
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
	Version:          "1.7",
	Host:             "localhost:9050",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Admin API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
