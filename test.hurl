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


## Update user
PATCH http://localhost:8080/api/user
Authorization: Bearer {{token}}
{
    "name": "Другое Фамилия Имя Отчество",
    "email": "new{{email}}",
	"password": "notdeadinsideanymore69"
}

HTTP 200

GET http://localhost:8080/api/user/{{user_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$['name']" == "Другое Фамилия Имя Отчество"
jsonpath "$['email']" == "new{{email}}"

POST http://localhost:8080/api/user/login
{
	"email": "new{{email}}",
	"password": "notdeadinsideanymore69"
}

HTTP 200
[Asserts]
jsonpath "$['id']" == {{user_id}}
jsonpath "$['token']" != null


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
[Captures]
option_id: jsonpath "$['options'][0]['id']"


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
pool_id_temp: jsonpath "$['id']"

DELETE http://localhost:8080/api/pool/{{pool_id_temp}}
Authorization: Bearer {{token}}

HTTP 200


# -- VOTE --

## Create vote
POST http://localhost:8080/api/pool/{{pool_id}}/vote
Authorization: Bearer {{token}}
{
	"id": {{option_id}}
}

HTTP 200

GET http://localhost:8080/api/pool/{{pool_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$['options'][0]['count']" == 1
jsonpath "$['vote']" == {{option_id}}


## Delete vote
DELETE http://localhost:8080/api/pool/{{pool_id}}/vote
Authorization: Bearer {{token}}

HTTP 200

GET http://localhost:8080/api/pool/{{pool_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$['options'][0]['count']" == 0
jsonpath "$['vote']" == 0


# -- GROUP --

## Get all groups
GET http://localhost:8080/api/group
Authorization: Bearer {{token}}

HTTP 200
[Captures]
groups_count: jsonpath "$" count


## Create group
POST http://localhost:8080/api/group
Authorization: Bearer {{token}}
{
	"name": "My awesome group"
}

HTTP 200
[Captures]
group_id: jsonpath "$['id']"

GET http://localhost:8080/api/group
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$" count > {{groups_count}}


## Update group
PATCH http://localhost:8080/api/group/{{group_id}}
Authorization: Bearer {{token}}
{
	"name": "My awesomest group"
}

HTTP 200


## Delete group
POST http://localhost:8080/api/group
Authorization: Bearer {{token}}
{
	"name": "My temp group"
}

HTTP 200
[Captures]
group_id_temp: jsonpath "$['id']"

GET http://localhost:8080/api/group
Authorization: Bearer {{token}}

HTTP 200
[Captures]
groups_count: jsonpath "$" count

DELETE http://localhost:8080/api/group/{{group_id_temp}}
Authorization: Bearer {{token}}

HTTP 200

GET http://localhost:8080/api/group
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$" count < {{groups_count}}


## Get users in group
GET http://localhost:8080/api/group/{{group_id}}/user
Authorization: Bearer {{token}}

HTTP 200
[Captures]
users_count: jsonpath "$" count


## Add user to group
POST http://localhost:8080/api/group/{{group_id}}/user
Authorization: Bearer {{token}}
{
	"id": {{user_id}}
}

HTTP 200

GET http://localhost:8080/api/group/{{group_id}}/user
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$" count > {{users_count}}


## Remove user from group

DELETE http://localhost:8080/api/group/{{group_id}}/user/{{user_id}}
Authorization: Bearer {{token}}

HTTP 200

GET http://localhost:8080/api/group/{{group_id}}/user
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$" count == {{users_count}}


# -- CHAT --

## Write message
POST http://localhost:8080/api/group/{{group_id}}/user
Authorization: Bearer {{token}}
{
	"id": {{user_id}}
}

HTTP 200

POST http://localhost:8080/api/chat/{{group_id}}
Authorization: Bearer {{token}}
{
	"text": "Testing Testing"
}

HTTP 200


## Get messages
GET http://localhost:8080/api/chat/{{group_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$[0]['text']" == "Testing Testing"
