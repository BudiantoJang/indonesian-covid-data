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
        "/berita-covid": {
            "get": {
                "description": "An API that used to get Indonesia news regarding covid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Berita Covid",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/data-harian-provinsi": {
            "get": {
                "description": "An API used to get Indonesian daily covid case update based on provinces",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Data Harian Provinsi",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/kasus-seluruh-provinsi": {
            "get": {
                "description": "An API that used to get data of Indonesian covid based on provinces",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "KasusSeluruhProvinsi",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/labolatorium-rujukan": {
            "get": {
                "description": "An API used to get Indonesian labolatorium that treats covid patients data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Labolatorium",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/pemeriksaan-dan-vaksinasi": {
            "get": {
                "description": "An API that used to get data of Indonesian covid checkup and vaccination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Pemeriksaan dan Vaksinasi",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/risk-score-kecamatan": {
            "get": {
                "description": "An API that used to get Indonesia risk score data based on districs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Risk Score Kecamatan",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/risk-score-provinsi": {
            "get": {
                "description": "An API that used to get Indonesia risk score data based on provinces",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Risk Score Provinsi",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/rumah-sakit-rujukan": {
            "get": {
                "description": "An API used to get Indonesian hospital that treats covid patients data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Rumah Sakit Rujukan",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        },
        "/update-harian-indonesia": {
            "get": {
                "description": "An API that used to get data of Indonesian daily covid case",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Update Harian Indonesia",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerSuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swaggerresponse.SwaggerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "swaggerresponse.SwaggerErrorResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "swaggerresponse.SwaggerSuccessResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Indonesian Covid Data API Documentation",
	Description:      "This is an API documentation for Indonesian Covid Data API that is generated using OpenAPI 2.0 specification",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
