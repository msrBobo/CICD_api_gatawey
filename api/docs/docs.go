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
        "/v1/customer/forget_password": {
            "post": {
                "description": "ForgetPassword - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "ForgetPassword",
                "parameters": [
                    {
                        "description": "RegisterModelReq",
                        "name": "ForgetPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model_user_service.PhoneNumberReq"
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
        "/v1/customer/forget_password_verify": {
            "post": {
                "description": "ForgetPasswordVerify - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "ForgetPasswordVerify",
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
        "/v1/customer/login": {
            "post": {
                "description": "Login - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
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
        "/v1/customer/logout": {
            "get": {
                "security": [
                    {
                        "ApiKeycustomer": []
                    }
                ],
                "description": "LogOut - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "LogOut",
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
        "/v1/customer/register": {
            "post": {
                "description": "Register - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
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
        "/v1/customer/verify": {
            "post": {
                "description": "customer - Api for registering users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
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
                    "type": "integer",
                    "example": 7777
                },
                "new_password": {
                    "type": "string",
                    "example": "new_password"
                },
                "phone_number": {
                    "type": "string",
                    "example": "+998950230605"
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
                    "type": "string",
                    "example": "mobile"
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
        "model_user_service.PhoneNumberReq": {
            "type": "object",
            "properties": {
                "phone_number": {
                    "type": "string",
                    "example": "+998950230605"
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
                    "example": "Ali"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "last_name": {
                    "type": "string",
                    "example": "Jo'raxonov'"
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
                    "type": "string",
                    "example": "mobile"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeycustomer": {
            "type": "apiKey",
            "name": "customerorization",
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
	Title:            "API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
