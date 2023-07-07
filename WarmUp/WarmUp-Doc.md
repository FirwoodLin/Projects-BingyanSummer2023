---
title: WarmUp-Doc v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# WarmUp-Doc

> v1.0.0

Base URLs:

* <a href="http://localhost:8080">正式环境: http://localhost:8080</a>

# WarmUp-23Summer

## POST 注册/register

POST /register

> Body Parameters

```json
{
  "name": "Carol Hall",
  "password": "pswd1234",
  "email": "123@qq.com",
  "tel": "+8618159300662",
  "nickname": "王秀兰"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» password|body|string| yes |none|
|» email|body|string| yes |none|
|» tel|body|string| yes |none|
|» nickname|body|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

## POST 登陆/signin

POST /signin

> Body Parameters

```json
{
  "email": "123@qq.com",
  "password": "pswd1234"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» email|body|string| yes |none|
|» password|body|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

## PATCH 修改个人信息/user/info

PATCH /user

> Body Parameters

```json
{
  "nickname": "曾明",
  "password": "culpa occaecat sint",
  "email": "q.bbwegii@qq.com",
  "tel": "18185189808",
  "name": "证议党头极安"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|SESSIONID|cookie|string| no |none|
|body|body|object| no |none|
|» nickname|body|string| yes |none|
|» password|body|string| yes |none|
|» email|body|string| yes |none|
|» tel|body|string| yes |none|
|» name|body|string| yes |none|

> Response Examples

> 成功

```json
{
  "status": "successfully updated"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» status|string|true|none||none|

## GET *获取所有人信息/user

GET /user

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|SESSIONID|cookie|string| no |none|

> Response Examples

> 成功

```json
[
  {
    "name": "admin",
    "nickname": "admin",
    "email": "admin@qq.com",
    "tel": "+8618888888888",
    "is_admin": true,
    "id": 12
  },
  {
    "name": "Carol Hall",
    "nickname": "王秀兰",
    "email": "123@qq.com",
    "tel": "+8618159300662",
    "is_admin": false,
    "id": 13
  }
]
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» name|string|true|none||none|
|» nickname|string|true|none||none|
|» email|string|true|none||none|
|» tel|string|true|none||none|
|» is_admin|boolean|true|none||none|
|» id|integer|true|none||none|

## DELETE *删除用户/user/:id

DELETE /user/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|SESSIONID|cookie|string| no |none|
|id|path|string| yes |none|

> Response Examples

> 成功

```json
{
  "msg": "删除成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» msg|string|true|none||none|

## GET 获取个人信息/user/:id

GET /user/{id}

用户本人/管理员拥有查询用户信息的权利

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|SESSIONID|cookie|string| no |none|
|id|path|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

# Data Schema

