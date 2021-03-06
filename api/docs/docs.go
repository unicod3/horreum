// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/articles/": {
            "get": {
                "description": "Get all articles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get all articles",
                "operationId": "list-articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/article.Article"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a article with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Create a article with given data",
                "operationId": "create-article",
                "parameters": [
                    {
                        "description": "Article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.ArticleRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/article.Article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/article.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/articles/{id}": {
            "get": {
                "description": "Get single article by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get single article by id",
                "operationId": "get-article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/article.Article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/article.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/article.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a article with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Update a article with given data",
                "operationId": "update-article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.ArticleRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/article.Article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/article.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/article.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a article by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Delete a article by id",
                "operationId": "delete-article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "NoContent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/article.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/article.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/orders/": {
            "get": {
                "description": "Get all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get all orders",
                "operationId": "list-orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/order.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a order with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a order with given data",
                "operationId": "create-order",
                "parameters": [
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/order.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Get single order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get single order by id",
                "operationId": "get-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/order.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/order.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a order with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update a order with given data",
                "operationId": "update-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/order.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/order.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete a order by id",
                "operationId": "delete-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "NoContent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/order.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/order.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/": {
            "get": {
                "description": "Get all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get all products",
                "operationId": "list-products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/product.ProductArticle"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a article with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create a article with given data",
                "operationId": "create-product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/product.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get single product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get single product by id",
                "operationId": "get-product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/product.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/product.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a product with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update a product with given data",
                "operationId": "update-product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/product.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/product.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete a product by id",
                "operationId": "delete-product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "NoContent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/product.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/product.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/warehouses/": {
            "get": {
                "description": "Get all warehouses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "warehouses"
                ],
                "summary": "Get all warehouses",
                "operationId": "list-warehouse",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/warehouse.Warehouse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a warehouse with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "warehouses"
                ],
                "summary": "Create a warehouse with given data",
                "operationId": "create-warehouse",
                "parameters": [
                    {
                        "description": "Warehouse",
                        "name": "warehouse",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/warehouse.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/warehouse.Warehouse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/warehouse.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/warehouses/{id}": {
            "get": {
                "description": "Get single warehouse by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "warehouses"
                ],
                "summary": "Get single warehouse by id",
                "operationId": "get-warehouse",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Warehouse ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/warehouse.Warehouse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/warehouse.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/warehouse.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a warehouse with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "warehouses"
                ],
                "summary": "Update a warehouse with given data",
                "operationId": "update-warehouse",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Warehouse ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Warehouse",
                        "name": "warehouse",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/warehouse.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/warehouse.Warehouse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/warehouse.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/warehouse.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a warehouse by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "warehouses"
                ],
                "summary": "Delete a warehouse by id",
                "operationId": "delete-warehouse",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Warehouse ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "NoContent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/warehouse.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/warehouse.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "article.Article": {
            "type": "object",
            "properties": {
                "amount_of": {
                    "type": "integer"
                },
                "available_inventory": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "article.ArticleRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "article.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "order.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "order.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customer": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.OrderLine"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "warehouse_id": {
                    "type": "integer"
                }
            }
        },
        "order.OrderLine": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "unit_cost": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "order.RequestBody": {
            "type": "object",
            "properties": {
                "customer": {
                    "type": "string"
                },
                "lines": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "product_id": {
                                "type": "integer"
                            },
                            "quantity": {
                                "type": "integer"
                            },
                            "unit_cost": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "warehouse_id": {
                    "type": "integer"
                }
            }
        },
        "product.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "product.Product": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/article.Article"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "sellable_inventory": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "product.ProductArticle": {
            "type": "object",
            "properties": {
                "amount_of": {
                    "type": "integer"
                },
                "available_inventory": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "productID": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "product.ProductRequestBody": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "amount_of": {
                                "type": "integer"
                            },
                            "id": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "warehouse.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "warehouse.RequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "warehouse.Warehouse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
