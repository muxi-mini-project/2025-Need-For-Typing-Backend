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
        "/api/assets": {
            "get": {
                "description": "查询所有素材的信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "素材"
                ],
                "summary": "获取所有素材",
                "responses": {
                    "200": {
                        "description": "返回所有素材信息",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "查询失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "上传素材文件到七牛云，需提供 asset_id 与文件内容",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "素材"
                ],
                "summary": "上传素材",
                "parameters": [
                    {
                        "type": "string",
                        "description": "素材ID",
                        "name": "asset_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "上传的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回素材ID和文件URL",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "缺少 asset_id 或文件上传失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "上传到七牛云失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/essay": {
            "get": {
                "description": "根据传入的 topic 参数，通过 gRPC 流式调用生成文章，并使用 SSE 向前端实时推送生成结果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/event-stream"
                ],
                "tags": [
                    "文章",
                    "SSE",
                    "gRPC"
                ],
                "summary": "获取生成的文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章主题",
                        "name": "topic",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "SSE stream of generated essay",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部服务器错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/score": {
            "post": {
                "description": "接收 JSON 格式的总分数据并上传到服务器",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分数"
                ],
                "summary": "上传总分",
                "parameters": [
                    {
                        "description": "上传的总分信息",
                        "name": "totalScore",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TotalScore"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回上传成功的分数ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/scores": {
            "get": {
                "description": "根据传入的 song_id 查询该歌曲所有的总分记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分数"
                ],
                "summary": "获取歌曲所有总分信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "歌曲ID",
                        "name": "song_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回总分信息",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/song": {
            "get": {
                "description": "根据传入的 song_id 查询歌曲信息并返回文件ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "歌曲"
                ],
                "summary": "获取歌曲信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "歌曲ID",
                        "name": "song_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回歌曲的 file_id",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "缺少 song_id 参数",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "歌曲未找到",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "处理歌曲上传请求，将歌曲文件存储到七牛云并创建数据库记录",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "歌曲"
                ],
                "summary": "上传歌曲",
                "parameters": [
                    {
                        "type": "string",
                        "description": "歌曲ID",
                        "name": "song_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "上传的歌曲文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回歌曲ID和文件URL",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "请求参数错误或文件上传失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "上传到七牛云失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "查询所有歌曲信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "歌曲"
                ],
                "summary": "获取所有歌曲",
                "responses": {
                    "200": {
                        "description": "返回所有歌曲的列表",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user/forget_password": {
            "get": {
                "description": "向后端发起忘记密码请求，通过邮箱发送重置密码的链接",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "忘记密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户邮箱地址",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "密码重置链接已发送",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "缺少邮箱参数",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录，验证用户名和密码，生成 JWT Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "用户登录信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回登录成功消息及 JWT Token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "用户名或密码错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "403": {
                        "description": "邮箱未验证",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "用户不存在",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册接口，接收用户信息创建新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户注册信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "用户创建成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "409": {
                        "description": "用户已存在",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user/reset_password": {
            "get": {
                "description": "提供重置密码的 token、邮箱和新密码，完成密码重置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "重置密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "重置密码的 token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新的密码",
                        "name": "new_password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "密码重置成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "缺少参数或无效的 token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "密码重置失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user/send_code": {
            "get": {
                "description": "生成并向用户邮箱发送验证码",
                "tags": [
                    "Verification"
                ],
                "summary": "发送邮箱验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱地址",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "验证码发送成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "邮箱为空",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "发送失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user/verify_code": {
            "post": {
                "description": "验证用户提交的验证码是否有效",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Verification"
                ],
                "summary": "验证验证码",
                "parameters": [
                    {
                        "description": "邮箱和验证码",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.VerifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "验证成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "请求体无效",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "验证码无效或过期",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user_scores": {
            "get": {
                "description": "根据传入的 user_id 查询该用户所有的最佳分数记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分数"
                ],
                "summary": "获取用户所有最佳成绩",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回用户最佳分数记录",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "用户未找到",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/ws": {
            "get": {
                "description": "建立 WebSocket 连接并处理房间逻辑",
                "tags": [
                    "WebSocket"
                ],
                "summary": "WebSocket 连接处理",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "房间 ID",
                        "name": "room_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回房间加入信息",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "未授权",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.VerifyRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "models.Game": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键，自动递增",
                    "type": "integer"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "score": {
                    "description": "关联 TotalScore",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TotalScore"
                    }
                },
                "song": {
                    "description": "关联 Song",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Song"
                        }
                    ]
                },
                "song_id": {
                    "description": "关联 Song 的 ID",
                    "type": "integer"
                },
                "time": {
                    "type": "string"
                },
                "user_id": {
                    "description": "外键，指向 User 表的 ID",
                    "type": "integer"
                }
            }
        },
        "models.Song": {
            "type": "object",
            "properties": {
                "file_id": {
                    "description": "文件标识符",
                    "type": "string"
                },
                "id": {
                    "description": "主键",
                    "type": "integer"
                },
                "title": {
                    "description": "歌曲标题",
                    "type": "string"
                }
            }
        },
        "models.TotalScore": {
            "type": "object",
            "properties": {
                "game": {
                    "$ref": "#/definitions/models.Game"
                },
                "game_id": {
                    "description": "外键，指向 Game 表的 ID",
                    "type": "integer"
                },
                "id": {
                    "description": "主键",
                    "type": "integer"
                },
                "total_score": {
                    "description": "分值",
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "email_verified": {
                    "description": "邮箱是否验证",
                    "type": "boolean"
                },
                "games": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Game"
                    }
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "reset_token": {
                    "type": "string"
                },
                "token_expires_at": {
                    "type": "string"
                },
                "user_id": {
                    "description": "id",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "NeedForTyping",
	Description:      "一个打字游戏",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
