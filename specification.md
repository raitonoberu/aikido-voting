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
	"groups": [
		{
			"id": 1,
			"name": "Администраторы",
		},
		// ...
	],
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
	"password": "newPassword",
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

## Group

-- **GET** `/group` - получить все существующие группы.

Response body:
```json
[
	{
		"id": 1,
		"name": "Администраторы",
	},
	{
		"id": 2,
		"name": "Все пользователи",
	},
	{
		"id": 3,
		"name": "Группа 1",
	},
	// ...
]
```

-- **POST** `/group` - создать новую группу. Доступно только администраторам.

Request body:
```json
{
	"name": "Группа 2",
}
```

Response body:
```json
{
	"id": 4,
}
```

-- **DELETE** `/group/{id}` - удалить группу. Доступно только администраторам.

-- **GET** `/group/{id}/user` - получить всех пользователей в группе.

Response body:
```json
[
	{
		"id": 1,
		"name": "Фамилия Имя Отчество",
		"email": "lussypicker@mail.ru",
	},
	// ...
]
```

-- **POST** `/group/{id}/user` - добавить пользователя в группу. Доступно только администраторам.

Request body:
```json
{
	"id": 1,
}
```

-- **DELETE** `/group/{id}/user/{id}` - удалить пользователя из группы. Доступно только администраторам.

## Chat

- **POST** `/chat/{id}` - отправить сообщение группе.

Request body:
```json
{
	"text": "Текст сообщения",
}
```

- **GET** `/chat/{id}?count=25&offset=0` - получить последние сообщения группы.

Request body:
```json
[
	{
		"text": "Текст сообщения",
		"user": {
			"id": 1,
			"email": "lussypicker@mail.ru",
			"name": "Фамилия Имя Отчество",
		},
		"created_at": "2023-11-23T13:19:46Z",
	},
	// ...
]
```

- **WS** `/chat/{id}/ws` - подписаться на сообщения группы.

Message body:
```json
{
	"text": "Текст сообщения",
	"user": {
		"id": 1,
		"email": "lussypicker@mail.ru",
		"name": "Фамилия Имя Отчество",
	},
	"created_at": "2023-11-23T13:19:46Z",
}
```