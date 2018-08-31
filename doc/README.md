# README

Source Code
https://github.com/dfang/yuanxin_api

Docker Image
https://dashboard.daocloud.io/

CI/CD
https://semaphoreci.com/dfang/yuanxin_api



由于chips/{:id}, buy_reqests/{:id}, help_reqests/{:id}, news/{:id} 的需求，在未登录的时候可以访问，
在登录的时候也能够访问，并且显示当前对其的收藏状态，而要获取用户身份需要用go-jwt-middleware中间处理请求，但是用了这个就得必须传Authorization Token，要满足这个需求需要对这个插件进行定制，修改CheckJWT 方法，这是个比较的特殊的需求没必要为了这个fork, 所以需要一个workdaround， 要么在客户端如果登录了的情况下，多发送一个查询收藏状态的请求，或者在发送/news/1请求的时候在header里带上一个Header X-UserID
