# API Specification

Для всех эндпоинтов кроме регистрации и входа нужен токен.
## User

- **POST** `/user` - добавить пользователя.

Request body:
```json
{
	"email": "lussypicker@mail.ru",
	"name": "Фамилия Имя Отчество",
	"password": "deadinside69",
}
```
Response body:
```json
{
	"id": 1,
	"token": "qwertyqwertyqwertyqwerty",
}
```

- **GET** `/user/{id}` - получить информацию о пользователе.

Response body:
```json
{
	"id": 1,
	"name": "Фамилия Имя Отчество",
	"email": "lussypicker@mail.ru",
}
```

- **POST** `/user/login` - войти в аккаунт.

Request body:
```json
{
	"email": "lussypicker@mail.ru",
	"password": "deadinside69",
}
```

Response body:
```json
{
	"id": 1,
	"token": "qwertyqwertyqwertyqwerty",
}
```

- **PATCH** `/user` - обновить информацию о пользователе.

Request body:
```json
{
	"name": "Фамилия Имя Отчество",
	"email": "lussypicker@mail.ru",
}
```

- **DELETE** `/user` - удалить пользователя.

## Pool

- **POST** `/pool` - создать новый опрос.

Request body:
```json
{
	"name": "Название опроса",
	"description": "Описание опроса",
	"group_id": 1,
	"is_anonymous": false,
	"options": [
		"Вариант 1",
		"Вариант 2",
		// ...
	],
	"open_for": 24,
}
```

Response body:
```json
{
	"id": 1,
}
```

- **GET** `/pool` - получить доступные голосования.

Response body:
```json
[
	{
		"id": 1,
		"name": "Название опроса",
		"description": "Описание опроса",
		"user": {
			"id": 1,
			"email": "lussypicker@mail.ru",
			"name": "Фамилия Имя Отчество",
		},
		"is_anonymous": false,
		"options": [
			{"id": 1, "text": "Вариант 1", "count": 0},
			{"id": 2, "text": "Вариант 2", "count": 0},
		],
		"created_at": "2023-11-23T13:19:46Z",
		"expires_at": "2023-11-23T13:19:46Z",
	},
	// ...
]
```

- **GET** `/pool/{id}` - получить информацию о голосовании.

Response body:
```json
{
	"id": 1,
	"name": "Название опроса",
	"description": "Описание опроса",
	"user": {
		"id": 1,
		"email": "lussypicker@mail.ru",
		"name": "Фамилия Имя Отчество",
	},
	"is_anonymous": false,
	"options": [
		{"id": 1, "text": "Вариант 1", "count": 0},
		{"id": 2, "text": "Вариант 2", "count": 0},
		// ...
	],
	"created_at": "2023-11-23T13:19:46Z",
	"expires_at": "2023-11-23T13:19:46Z",
}
```

- **DELETE** `/pool/{id}` - удалить голосование.

- **POST** `/pool/{id}/vote` - проголосовать.

Request body:
```json
{
	"id": 1,
}
```

- **DELETE** `/pool/{id}/vote` - отменить голос.

- **POST** `/pool/{id}/comment` - добавить комментарий.

Request body:
```json
{
	"text": "Текст комментария",
}
```

- **GET** `/pool/{id}/comment` - получить комментарии.

Response body:
```json
[
	{
		"id": 1,
		"user": {
			"id": 1,
			"email": "lussypicker@mail.ru",
			"name": "Фамилия Имя Отчество",
		},
		"text": "Текст комментария",
		"created_at": "2023-11-23T13:19:46Z",
	},
	// ...
]
```

- **DELETE** `/pool/{id}/comment/{id}` - удалить комментарий.

# Group

`TODO`

# Role

`TODO`
