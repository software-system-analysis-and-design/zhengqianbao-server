# zhengqianbao-server
挣钱宝 go后端源码，使用Iris实现，blotdb数据库支持

## Quick Start

安装依赖

```
go get -u github.com/kataras/iris
go get -u github.com/iris-contrib/middleware/jwt
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/boltdb/bolt
```

运行

```
go run main.go
```

## 使用说明

后台代码目前已经挂在到腾讯云服务器

## API设计文档

[Interface API design](https://software-system-analysis-and-design.github.io/Dashboard/docs/API.html)

## 数据库设计文档

- 7.2.1 [用户及权限系统数据库设计](https://software-system-analysis-and-design.github.io/Dashboard/docs/db_design.html)
- 7.2.2 [ER模型和关系模型](https://software-system-analysis-and-design.github.io/Dashboard/docs/db_er.html)
- 7.2.3 [第三方数据评审结果](https://github.com/software-system-analysis-and-design/Dashboard/issues/1)

