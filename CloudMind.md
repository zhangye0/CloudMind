---
title: CloudMind v1.0.0
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

# CloudMind

> v1.0.0

Base URLs:

# 登陆服务

## POST 账号登录

POST /usercenter/v1/login/EmailLogin

> Body 请求参数

```json
{
  "email": "e.fwhcuio@1",
  "passWord": "q7myH&(tL@DQ"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» email|body|string| 是 | 邮箱号|none|
|» passWord|body|string| 是 | 密码|none|

> 返回示例

> 200 Response

```json
{
  "accessToken": "string",
  "accessExpire": 0,
  "refreshAfter": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|登录成功|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|账号不存在|Inline|
|202|[Accepted](https://tools.ietf.org/html/rfc7231#section-6.3.3)|密码错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» accessToken|string|false|none|令牌内容|none|
|» accessExpire|integer|false|none|令牌有效期|none|
|» refreshAfter|integer|true|none|令牌刷新周期|令牌刷新期一定小于等于令牌有效期|

## POST 账号注册

POST /usercenter/v1/login/Register

> Body 请求参数

```json
{
  "code": "732613",
  "email": "s.xulom@txdozfcwg.gp",
  "nickName": "秦",
  "passWord": "bxA[^bxA["
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» nickName|body|string| 是 | 用户名|none|
|» passWord|body|string| 是 | 密码|none|
|» email|body|string| 是 | 邮箱号|none|
|» code|body|string| 是 | 邮箱验证码|none|

> 返回示例

> 注册成功

```json
{
  "accessToken": "tgK6l",
  "accessExpire": 378,
  "refreshAfter": 22
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|注册成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» accessToken|string|true|none|令牌内容|none|
|» accessExpire|integer|true|none|令牌有效期|none|
|» refreshAfter|integer|true|none|令牌刷新周期|令牌刷新周期要小于等于令牌有效期|

## POST 发送邮件

POST /usercenter/v1/tools/SendEmail

> Body 请求参数

```json
{
  "email": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» email|body|string| 是 | 邮箱号|none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|发送成功|Inline|

### 返回数据结构

## POST QQ登录

POST /usercenter/v1/login/QqLogin

> 返回示例

> 登录成功

```json
{
  "accessToken": "DlB",
  "accessExpire": 902,
  "refreshAfter": 812
}
```

> 201 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|登录成功|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|登录失败|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» accessToken|string|true|none|令牌内容|none|
|» accessExpire|integer|true|none|令牌有效期|none|
|» refreshAfter|integer|true|none|令牌刷新周期|令牌刷新周期要小于令牌有效期|

## POST 微信登录

POST /usercenter/v1/login/WxLogin

> 返回示例

> 登录成功

```json
{
  "accessToken": "DlB",
  "accessExpire": 902,
  "refreshAfter": 812
}
```

> 201 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|登录成功|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|登录失败|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» accessToken|string|true|none|令牌内容|none|
|» accessExpire|integer|true|none|令牌有效期|none|
|» refreshAfter|integer|true|none|令牌刷新周期|令牌刷新周期要小于令牌有效期|

# 数据模型

