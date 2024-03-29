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
    "/v1/schema/{id}": {
      "get": {
        "summary": "Get Schema",
        "operationId": "GetSchema",
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
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Backend"
        ]
      }
    },
    "/v1/schema/{projectID}": {
      "post": {
        "summary": "Create Schema",
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
            "name": "projectID",
            "in": "path",
            "required": true,
            "type": "string"
          },
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
    "protoSchema": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "properties": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/protobufStruct"
          }
        },
        "required": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        },
        "projectID": {
          "type": "string"
        },
        "version": {
          "type": "string"
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
    },
    "protobufListValue": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufValue"
          },
          "description": "Repeated field of dynamically typed values."
        }
      },
      "description": ""ListValue" is a wrapper around a repeated field of values.\n\nThe JSON representation for "ListValue" is JSON array."
    },
    "protobufNullValue": {
      "type": "string",
      "enum": [
        "NULL_VALUE"
      ],
      "default": "NULL_VALUE",
      "description": ""NullValue" is a singleton enumeration to represent the null value for the\n"Value" type union.\n\n The JSON representation for "NullValue" is JSON "null".\n\n - NULL_VALUE: Null value."
    },
    "protobufStruct": {
      "type": "object",
      "properties": {
        "fields": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/protobufValue"
          },
          "description": "Unordered map of dynamically typed values."
        }
      },
      "description": ""Struct" represents a structured data value, consisting of fields\nwhich map to dynamically typed values. In some languages, "Struct"\nmight be supported by a native representation. For example, in\nscripting languages like JS a struct is represented as an\nobject. The details of that representation are described together\nwith the proto support for the language.\n\nThe JSON representation for "Struct" is JSON object."
    },
    "protobufValue": {
      "type": "object",
      "properties": {
        "null_value": {
          "$ref": "#/definitions/protobufNullValue",
          "description": "Represents a null value."
        },
        "number_value": {
          "type": "number",
          "format": "double",
          "description": "Represents a double value."
        },
        "string_value": {
          "type": "string",
          "description": "Represents a string value."
        },
        "bool_value": {
          "type": "boolean",
          "format": "boolean",
          "description": "Represents a boolean value."
        },
        "struct_value": {
          "$ref": "#/definitions/protobufStruct",
          "description": "Represents a structured value."
        },
        "list_value": {
          "$ref": "#/definitions/protobufListValue",
          "description": "Represents a repeated "Value"."
        }
      },
      "description": ""Value" represents a dynamically typed value which can be either\nnull, a number, a string, a boolean, a recursive struct value, or a\nlist of values. A producer of value is expected to set one of that\nvariants, absence of any variant indicates an error.\n\nThe JSON representation for "Value" is JSON value."
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
