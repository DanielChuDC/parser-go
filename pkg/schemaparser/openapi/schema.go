package openapi

func getSchema() []byte {
	return []byte(`
{
    "title": "OpenAPI/AsyncAPI schema.",
    "$schema": "http://json-schema.org/draft-04/schema#",
    "description": "A deterministic version of a JSON Schema object.",
    "type": "object",
    "additionalProperties": false,
    "patternProperties": {
        "^x-[\\w\\d\\.\\-\\_]+$": {
            "$ref": "#/definitions/specificationExtension"
        }
    },
    "properties": {
        "$ref": {
            "$ref": "#/definitions/ReferenceObject"
        },
        "format": {
            "type": "string"
        },
        "nullable": {
            "type": "boolean",
            "default": false
        },
        "title": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/title"
        },
        "description": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/description"
        },
        "default": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/default"
        },
        "multipleOf": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/multipleOf"
        },
        "maximum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/maximum"
        },
        "exclusiveMaximum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/exclusiveMaximum"
        },
        "minimum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/minimum"
        },
        "exclusiveMinimum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/exclusiveMinimum"
        },
        "maxLength": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
        },
        "minLength": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
        },
        "pattern": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/pattern"
        },
        "maxItems": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
        },
        "minItems": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
        },
        "uniqueItems": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/uniqueItems"
        },
        "maxProperties": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveInteger"
        },
        "minProperties": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0"
        },
        "required": {
            "$ref": "http://json-schema.org/draft-04/schema#/definitions/stringArray"
        },
        "enum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/enum"
        },
        "deprecated": {
            "type": "boolean",
            "default": false
        },
        "additionalProperties": {
            "anyOf": [
                {
                    "$ref": "#"
                },
                {
                    "type": "boolean"
                }
            ],
            "default": {}
        },
        "type": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/type"
        },
        "items": {
            "anyOf": [
                {
                    "$ref": "#"
                },
                {
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "$ref": "#"
                    }
                }
            ],
            "default": {}
        },
        "allOf": {
            "type": "array",
            "minItems": 1,
            "items": {
                "$ref": "#"
            }
        },
        "oneOf": {
            "type": "array",
            "minItems": 2,
            "items": {
                "$ref": "#"
            }
        },
        "anyOf": {
            "type": "array",
            "minItems": 2,
            "items": {
                "$ref": "#"
            }
        },
        "not": {
            "$ref": "#"
        },
        "properties": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#"
            },
            "default": {}
        },
        "discriminator": {
            "type": "string"
        },
        "readOnly": {
            "type": "boolean",
            "default": false
        },
        "xml": {
            "$ref": "#/definitions/xml"
        },
        "externalDocs": {
            "$ref": "#/definitions/externalDocs"
        },
        "example": {},
        "examples": {
            "type": "array",
            "items": {}
        }
    },
    "definitions": {
        "ReferenceObject": {
            "type": "string",
            "format": "uri"
        },
        "specificationExtension": {
            "description": "Any property starting with x- is valid.",
            "additionalProperties": true,
            "additionalItems": true
        },
        "xml": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "prefix": {
                    "type": "string"
                },
                "attribute": {
                    "type": "boolean",
                    "default": false
                },
                "wrapped": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "externalDocs": {
            "type": "object",
            "additionalProperties": false,
            "description": "information about external documentation",
            "required": [
                "url"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "url": {
                    "type": "string",
                    "format": "uri"
                }
            },
            "patternProperties": {
                "^x-[\\w\\d\\.\\-\\_]+$": {
                    "$ref": "#/definitions/specificationExtension"
                }
            }
        }
    }
}
	`)
}
