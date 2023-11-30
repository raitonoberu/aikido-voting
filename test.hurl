#  -- USER --

## Registration
POST http://localhost:8080/api/user
{
	"email": "{{email}}",
	"name": "Фамилия Имя Отчество",
	"password": "deadinside69"
}

HTTP 200
[Captures]
user_id: jsonpath "$['id']"
token: jsonpath "$['token']"


## Login
POST http://localhost:8080/api/user/login
{
    "email": "{{email}}",
    "password": "deadinside69"
}

HTTP 200
[Asserts]
jsonpath "$['id']" == {{user_id}}
jsonpath "$['token']" != null


## Get user
GET http://localhost:8080/api/user/{{user_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$['id']" == {{user_id}}
jsonpath "$['name']" == "Фамилия Имя Отчество"
jsonpath "$['email']" == "{{email}}"


## Update user (TODO)
# PATCH http://localhost:8080/api/user/{{user_id}}
# Authorization: Bearer {{token}}
# {
#     "name": "Другое Фамилия Имя Отчество",
#     "email": "sickducker@mail.ru"
# }

# HTTP 200
# [Asserts]
# jsonpath "$['id']" == {{user_id}}
# jsonpath "$['name']" == "Другое Фамилия Имя Отчество"
# jsonpath "$['email']" == "sickducker@mail.ru"


## Delete user
POST http://localhost:8080/api/user
{
	"email": "testing@testing.ru",
	"name": "name",
	"password": "deadinside69"
}

HTTP 200
[Captures]
token_temp: jsonpath "$['token']"

DELETE http://localhost:8080/api/user
Authorization: Bearer {{token_temp}}

HTTP 200


# -- POOL --

## Create pool
POST http://localhost:8080/api/pool
Authorization: Bearer {{token}}
{
	"name": "Название опроса",
	"description": "Описание опроса",
	"group_id": 1,
	"is_anonymous": true,
	"options": [
		"Вариант 1",
		"Вариант 2"
	],
	"open_for": 24
}

HTTP 200
[Captures]
pool_id: jsonpath "$['id']"


## Get pool
GET http://localhost:8080/api/pool/{{pool_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$['id']" == {{pool_id}}
jsonpath "$['name']" == "Название опроса"
jsonpath "$['description']" == "Описание опроса"
jsonpath "$['user']['id']" == {{user_id}}
jsonpath "$['is_anonymous']" == true
jsonpath "$['options']" count == 2
jsonpath "$['options'][0]['text']" == "Вариант 1"
jsonpath "$['options'][0]['id']" != null
jsonpath "$['options'][0]['count']" == 0
jsonpath "$['options'][1]['text']" == "Вариант 2"
jsonpath "$['options'][1]['id']" != null
jsonpath "$['options'][1]['count']" == 0
jsonpath "$['created_at']" != null
jsonpath "$['expires_at']" != null


## Delete pool
POST http://localhost:8080/api/pool
Authorization: Bearer {{token}}
{
	"name": "temp",
	"description": "temp",
	"group_id": 1,
	"is_anonymous": true,
	"options": [
		"Вариант 1",
		"Вариант 2"
	],
	"open_for": 24
}

HTTP 200
[Captures]
pool_id: jsonpath "$['id']"

DELETE http://localhost:8080/api/pool/{{pool_id}}
Authorization: Bearer {{token}}

HTTP 200