# 使用 validator 进行参数校验

gin 框架默认采用的是 go-playground/validator 包进行参数校验，因此，我们不用额外引入包，就可以使用 validator 进行参数校验。

体验：

```bash
go run validator.go
```

然后请求

```bash
curl --location '127.0.0.1:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "age": 18,
    "date": "2024-08-10",
    "email": "12345@qq.com",
    "name": "alex",
    "password": "123",
    "re_password": "123"
}'
```