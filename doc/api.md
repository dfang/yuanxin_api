# API

### 新闻列表

<<<<<<< HEAD
GET /news
=======
GET /news  
>>>>>>> develop

参数   默认值
start  0
count  10
type   0


### 新闻详情

<<<<<<< HEAD
GET /news/:id
=======
GET /news/:id  
>>>>>>> develop


### 短信验证码

发送验证码
POST /captcha/send

参数
phone required


验证验证码
POST /captcha/validate

参数
phone required
captcha required

### Exists
验证（email, phone）是否存在

参数 phone 或 email


### 注册

POST /registrations


### 登录

POST /sessions


### 重置密码

PUT /passwords
<<<<<<< HEAD
=======




>>>>>>> develop
