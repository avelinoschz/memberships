# Memberships API

Micro-service for managing members and payments. It follows [jsonapi](https://jsonapi.org/) standard.

URL: http://localhost.com/memberships/api/v1/

## Request/Response examples

Simple examples of body request and the responses.

### /alive GET

Health checking endpoint. Returns the server's current time and the service uptime.

```shell
curl --location --request GET 'localhost:8000/memberships/api/v1/alive'
```

```json
{
    "data": {
        "type": "",
        "attributes": {
            "currentTime": 1593461161,
            "uptime": 1.77
        }
    }
}
```

### /members POST

Registers new member to the service.

```shell
curl --location --request POST 'localhost:8000/memberships/api/v1/members' \
--header 'Content-Type: application/vnd.api+json' \
--data-raw '{
    "data": {
        "type": "members",
        "attributes": {
            "name": "Avelino Sanchez",
            "email": "avelino@mail.com",
            "phone": "+5215568040528",
            "password": "changeit"
        }
    }
}'
```

```json
{
    "data": {
        "type": "members",
        "id": "1",
        "attributes": {
            "email": "avelino@mailcom",
            "name": "Avelino Sanchez",
            "password": "changeit",
            "phone": "+5215568040528"
        }
    }
}
```

### /members/{id}/payments/ GET

Retrieves payments' information related to one specific member.

```shell
curl --location --request GET 'localhost:8000/memberships/api/v1/members/1/payments'
```

```json
{
    "data": {
        "type": "payments",
        "id": "1",
        "attributes": {
            "displayName": "A A Sanchez Alvarez",
            "lastFour": "1234",
            "expirationDate": "20-08-23T18:25:43.511Z"
        },
        "relationships": {
            "member": {
                "data": {
                    "type": "members",
                    "id": "1"
                }
            },
            "card": {
                "data": {
                    "type": "cards",
                    "id": "1"
                }
            }
        }
    }
}
```

### /members/{id}/payments/ POST

Registers a new transaction the service

```shell
curl --location --request POST 'localhost:8000/memberships/api/v1/members/1/payments' \
--header 'Content-Type: application/vnd.api+json' \
--data-raw '{
    "data": {
        "type": "payments",
        "attributes": {
            "displayName": "A A Sanchez Alvarez",
            "lastFour": "1234",
            "expirationDate": "20-08-23T18:25:43.511Z"
        },
        "relationships": {
            "member": {
                "data": {
                    "type": "members",
                    "id": "1"
                }
            }
        },
        "card": {
            "data": {
                "type": "cards",
                "id": "1"
            }
        }
    }
}'
```

```json
{
    "data": {
        "type": "payments",
        "attributes": {
            "displayName": "A A Sanchez Alvarez",
            "lastFour": "1234",
            "expirationDate": "20-08-23T18:25:43.511Z"
        },
        "relationships": {
            "member": {
                "data": {
                    "type": "members",
                    "id": "1"
                }
            }
        },
        "card": {
            "data": {
                "type": "cards",
                "id": "1"
            }
        }
    }
}
```
