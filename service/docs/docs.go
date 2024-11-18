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
        "/artists/": {
            "post": {
                "description": "Добавить нового исполнителя. Название исполнителя должно быть уникальным.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "artist"
                ],
                "summary": "Добавить нового исполнителя.",
                "parameters": [
                    {
                        "description": "Данные нового исполнителя",
                        "name": "artist",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/artistrest.CreateArtistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/artistrest.CreateArtistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/artists/{artist-id}": {
            "get": {
                "description": "Получить данные определенного исполнителя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "artist"
                ],
                "summary": "Получить данные определенного исполнителяяяя.",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "artist-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/artistrest.GetArtistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удалить данные исполнителя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "artist"
                ],
                "summary": "Удалить данные исполнителя.",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "artist-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/artistrest.RemoveArtistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Изменить данные исполнителя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "artist"
                ],
                "summary": "Изменить данные исполнителя.",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "artist-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новые данные исполнителя",
                        "name": "artist",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/artistrest.ChangeArtistRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/artistrest.ChangeArtistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/songs/": {
            "get": {
                "description": "Поиск определенной песни по всем атрибутам.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Поиск определенной песни.",
                "parameters": [
                    {
                        "type": "string",
                        "name": "artistName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "link",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/songrest.SearchSongsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавление новой песни. Для разделения куплетов необходимо использовать '\\n\\n'.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Добавить новую песню.",
                "parameters": [
                    {
                        "description": "Данные новой песни",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/songrest.CreateSongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/songrest.CreateSongResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/songs/{song-id}": {
            "delete": {
                "description": "Удалить данные песни.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Удалить данные песни.",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "song-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/songrest.RemoveSongResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Изменить данные песни. Для разделения куплетов необходимо использовать '\\n\\n'.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Изменить данные песни.",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "song-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новые данные песни",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/songrest.ChangeSongRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/songrest.ChangeSongResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/songs/{song-id}/couplets": {
            "get": {
                "description": "Получить данные определенной песни с пагинацией по куплетами.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Получить данные определенной песни.",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "song-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "name": "pageNumber",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 10,
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/songrest.GetSongResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/mwerror.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "artistrest.ChangeArtistRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 130
                }
            }
        },
        "artistrest.ChangeArtistResponse": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                }
            }
        },
        "artistrest.CreateArtistRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 130
                }
            }
        },
        "artistrest.CreateArtistResponse": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                }
            }
        },
        "artistrest.GetArtistResponse": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                }
            }
        },
        "artistrest.RemoveArtistResponse": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.ArtistAPI": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                }
            }
        },
        "models.ArtistIDAPI": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.Pagination": {
            "type": "object",
            "properties": {
                "pageNumber": {
                    "type": "integer",
                    "minimum": 1
                },
                "pageSize": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 10
                }
            }
        },
        "models.PaginationMetadataAPI": {
            "type": "object",
            "properties": {
                "currentPageNumber": {
                    "type": "integer"
                },
                "pageCount": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "recordCount": {
                    "type": "integer"
                }
            }
        },
        "models.SongAPI": {
            "type": "object",
            "required": [
                "id",
                "link",
                "name",
                "releaseDate",
                "text"
            ],
            "properties": {
                "artist": {
                    "$ref": "#/definitions/models.ArtistAPI"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "mwerror.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "songrest.ChangeSongRequestBody": {
            "type": "object",
            "properties": {
                "artist": {
                    "$ref": "#/definitions/models.ArtistIDAPI"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "songrest.ChangeSongResponse": {
            "type": "object",
            "required": [
                "id",
                "link",
                "name",
                "releaseDate",
                "text"
            ],
            "properties": {
                "artist": {
                    "$ref": "#/definitions/models.ArtistAPI"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "songrest.CreateSongRequest": {
            "type": "object",
            "required": [
                "link",
                "name",
                "releaseDate",
                "text"
            ],
            "properties": {
                "artist": {
                    "$ref": "#/definitions/models.ArtistIDAPI"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "songrest.CreateSongResponse": {
            "type": "object",
            "required": [
                "id",
                "link",
                "name",
                "releaseDate",
                "text"
            ],
            "properties": {
                "artist": {
                    "$ref": "#/definitions/models.ArtistAPI"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 130
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "songrest.GetSongResponse": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/models.PaginationMetadataAPI"
                },
                "song": {
                    "$ref": "#/definitions/models.SongAPI"
                }
            }
        },
        "songrest.RemoveSongResponse": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "songrest.SearchSongsResponse": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/models.PaginationMetadataAPI"
                },
                "songs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SongAPI"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Song-library-service",
	Description:      "Микросервис библиотеки песен.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
