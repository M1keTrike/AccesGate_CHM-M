
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://accessgate.com/terms/",
        "contact": {
            "name": "Equipo de Desarrollo AccessGate",
            "email": "contacto@accessgate.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {

        "/events": {

        "/api/event-attendees": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Registers a user to attend a specific event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Attendees"
                ],
                "summary": "Register a user for an event",
                "parameters": [
                    {
                        "description": "Attendee Registration Data",
                        "name": "attendee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.EventAttendee"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.EventAttendee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/api/event-attendees/event/{eventId}/attendees": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves all users registered to attend a specific event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Attendees"
                ],
                "summary": "Get all attendees for an event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/api/event-attendees/event/{eventId}/user/{userId}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Removes a user's registration from a specific event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Attendees"
                ],
                "summary": "Remove an attendee from an event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/api/event-attendees/event/{eventId}/user/{userId}/check": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Verifies if a specific user is registered to attend an event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Attendees"
                ],
                "summary": "Check if a user is registered for an event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/api/event-attendees/user/{userId}/events": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],


                "description": "Retrieves all events that a user is registered to attend",

                "produces": [
                    "application/json"
                ],
                "tags": [

                    "Events"
                ],
                "summary": "Lista todos los eventos",

                    "Event Attendees"
                ],
                "summary": "Get all events for a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],

                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Event"
                            }
                        }


                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/api/nfc-assignments": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves all NFC card assignments in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NFC Card Assignments"
                ],
                "summary": "Gets all NFC card assignments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.NfcCardAssignment"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],


                "description": "Assigns an NFC card to a user",

                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [

                    "Events"
                ],
                "summary": "Crea un nuevo evento",
                "parameters": [
                    {
                        "description": "Evento a crear",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Event"

                    "NFC Card Assignments"
                ],
                "summary": "Creates a new NFC card assignment",
                "parameters": [
                    {
                        "description": "Assignment Data",
                        "name": "assignment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.NfcCardAssignment"

                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {

                            "$ref": "#/definitions/entities.Event"
                            "$ref": "#/definitions/entities.NfcCardAssignment"

                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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

        "/events/creator/{user_id}": {

        "/api/nfc-assignments/card/{cardUid}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves the active assignment for a specific NFC card",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NFC Card Assignments"
                ],
                "summary": "Gets an NFC card assignment by card UID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card UID",
                        "name": "cardUid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.NfcCardAssignment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
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
        "/api/nfc-assignments/user/{userId}": {

            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],


                "description": "Retrieves all NFC card assignments associated with a specific user",

                "produces": [
                    "application/json"
                ],
                "tags": [

                    "Events"
                ],
                "summary": "Lista eventos por creador",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del creador",
                        "name": "user_id",

                    "NFC Card Assignments"
                ],
                "summary": "Gets all NFC card assignments for a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",

                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {

                                "$ref": "#/definitions/entities.Event"

                                "$ref": "#/definitions/entities.NfcCardAssignment"

                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
        "/events/{id}": {

                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/api/nfc-assignments/{id}": {

            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],


                "description": "Retrieves a specific NFC card assignment using its ID",

                "produces": [
                    "application/json"
                ],
                "tags": [

                    "Events"
                ],
                "summary": "Obtiene un evento por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del evento",

                    "NFC Card Assignments"
                ],
                "summary": "Gets an NFC card assignment by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Assignment ID",

                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {

                            "$ref": "#/definitions/entities.Event"

                            "$ref": "#/definitions/entities.NfcCardAssignment"

                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],


                "description": "Updates the details of an existing NFC card assignment",

                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [

                    "Events"
                ],
                "summary": "Actualiza un evento",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del evento",

                    "NFC Card Assignments"
                ],
                "summary": "Updates an NFC card assignment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Assignment ID",

                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {

                        "description": "Datos del evento",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {

                        "description": "Updated Assignment Data",
                        "name": "assignment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.NfcCardAssignment"

                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }

            },
            "delete": {

            }
        },
        "/api/nfc-assignments/{id}/deactivate": {
            "put": {

                "security": [
                    {
                        "BearerAuth": []
                    }
                ],


                "description": "Sets an NFC card assignment as inactive",

                "produces": [
                    "application/json"
                ],
                "tags": [

                    "Events"
                ],
                "summary": "Elimina un evento",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del evento",

                    "NFC Card Assignments"
                ],
                "summary": "Deactivates an NFC card assignment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Assignment ID",

                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Obtiene todos los usuarios",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Registra un usuario con nombre, email, contraseña y rol",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Crea un nuevo usuario",
                "parameters": [
                    {
                        "description": "Nuevo usuario",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
        "/users/email": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Obtiene un usuario por email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email del usuario",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
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
        "/users/login": {
            "post": {
                "description": "Valida credenciales y retorna un token JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Inicia sesión de usuario",
                "parameters": [
                    {
                        "description": "Credenciales",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
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
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Obtiene un usuario por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Actualiza un usuario existente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Usuario actualizado",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Borra un usuario existente utilizando su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Elimina un usuario por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateUserRequest": {
            "type": "object",
            "properties": {
                "biometric_auth": {
                    "type": "boolean",
                    "example": false
                },
                "email": {
                    "type": "string",
                    "example": "admin@example.com"
                },
                "fingerprint_id": {
                    "type": "integer",
                    "example": 0
                },
                "name": {
                    "type": "string",
                    "example": "Admin"
                },
                "password_hash": {
                    "type": "string",
                    "example": "123456"
                },
                "role": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "controllers.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "admin@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "entities.Event": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },


        "entities.EventAttendee": {
            "type": "object",
            "properties": {
                "event_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "registered_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "entities.NfcCardAssignment": {
            "type": "object",
            "properties": {
                "assigned_at": {
                    "type": "string"
                },
                "card_uid": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },

        "entities.User": {
            "type": "object",
            "properties": {
                "biometric_auth": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fingerprint_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8084",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "AccessGate API",
	Description:      "API REST con arquitectura hexagonal para gestión de usuarios, tarjetas NFC y clientes.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
