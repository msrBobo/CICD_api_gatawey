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
        "/v1/patient/delete/:key": {
            "delete": {
                "description": "Patients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "Delete Patient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.DeleteStatus"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    }
                }
            }
        },
        "/v1/patients/create": {
            "post": {
                "description": "Patients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "Create Patient",
                "parameters": [
                    {
                        "description": "Create Patient",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.CreatePatientReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Patient"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    }
                }
            }
        },
        "/v1/patients/get": {
            "get": {
                "description": "Patients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "Create Patient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Patient"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    }
                }
            }
        },
        "/v1/patients/list": {
            "get": {
                "description": "Patients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "List Patient",
                "parameters": [
                    {
                        "type": "string",
                        "example": "first_name",
                        "name": "field",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "example": true,
                        "name": "is_active",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "last_name",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "A",
                        "name": "value",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Patients"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    }
                }
            }
        },
        "/v1/patients/phone": {
            "put": {
                "description": "Patients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "Update Patient",
                "parameters": [
                    {
                        "description": "Update Patient",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.UpdatePhoneNumber"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Patient"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    }
                }
            }
        },
        "/v1/patients/update": {
            "put": {
                "description": "Patients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "Update Patient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "field",
                        "name": "field",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "value",
                        "name": "value",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Update Patient",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.UpdatePatientReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Patient"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models_booking_service.Errors"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models_booking_service.CreatePatientReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "blood_group": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "patient_problem": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "models_booking_service.DeleteStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "boolean"
                }
            }
        },
        "models_booking_service.Errors": {
            "type": "object",
            "properties": {
                "error_res": {},
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "models_booking_service.Patient": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "blood_group": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "patient_problem": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models_booking_service.Patients": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "patients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models_booking_service.Patient"
                    }
                }
            }
        },
        "models_booking_service.UpdatePatientReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "blood_group": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
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
                "patient_problem": {
                    "type": "string"
                }
            }
        },
        "models_booking_service.UpdatePhoneNumber": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
