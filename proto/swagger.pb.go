package proto

const (
	BackendSwagger = `{
  "swagger": "2.0",
  "info": {
    "title": "backend.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/schema": {
      "put": {
        "summary": "Update User Betting",
        "operationId": "CreateSchema",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/protoSchema"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/protoSchema"
            }
          }
        ],
        "tags": [
          "Backend"
        ]
      }
    },
    "/v1/version": {
      "get": {
        "summary": "Version",
        "operationId": "Version",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/protoVersionResponse"
            }
          }
        },
        "tags": [
          "Backend"
        ]
      }
    }
  },
  "definitions": {
    "protoColumn": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "minLength": {
          "type": "integer",
          "format": "int32"
        },
        "maxLength": {
          "type": "integer",
          "format": "int32"
        },
        "minimum": {
          "type": "string",
          "format": "int64"
        },
        "maximum": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "protoSchema": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "properties": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/protoColumn"
          }
        },
        "required": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "protoVersionResponse": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    }
  }
}
`
	ChemaSwagger = `{
  "swagger": "2.0",
  "info": {
    "title": "schema.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {},
  "definitions": {}
}
`
)
