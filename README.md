# Introduction

Boilerplate for an HTTP API service, in Golang.

# Features

- ORM with sqlite3 backend.
- CRUD APIs.
- Swagger, auto-generated from declarative comments.
- Documentation by `godoc`.
- Unit tests.
- Logs to stdout (12-factor), collected to Sumo via K8s annotations.
- Prometheus metrics exported to `/metrics`, collected to Dash via K8s annotations.
- Dockerfile.
- Kubernetes deployment.

# Getting started

Simply clone the repo and update functionalities in `model` package.

## Run Service

```shell
GIN_MODE=release go run main.go
```

## Regenerate swagger

```shell
swag init
```

## Generate documentation

```shell
godoc -http=:6060
```

Go to http://localhost:6060/pkg/github.com/huyanhvn/quick-rest-go/

# Unit tests

Boilerplate unit tests are stubbed for `model` and `controller`.

```shell
huy.nguyen-> GIN_MODE=release go test -v ./...
?       github.com/huyanhvn/quick-rest-go [no test files]
=== RUN   TestController_GetAllItems
--- PASS: TestController_GetAllItems (0.00s)
=== RUN   TestController_GetItemByID
2022/01/19 13:27:01 Setup for unit tests.
=== RUN   TestController_GetItemByID/non-existent_item

2022/01/19 13:27:01 /Users/huy.nguyen/repos/quick-rest-go/model/items.go:73 record not found
[1.054ms] [rows:0] SELECT * FROM `items` WHERE `items`.`id` = 1 ORDER BY `items`.`id` LIMIT 1
2022/01/19 13:27:01 Tear down setup.
--- PASS: TestController_GetItemByID (0.01s)
    --- PASS: TestController_GetItemByID/non-existent_item (0.00s)
=== RUN   TestController_CreateItem
--- PASS: TestController_CreateItem (0.00s)
=== RUN   TestController_UpdateItem
--- PASS: TestController_UpdateItem (0.00s)
=== RUN   TestController_DeleteItem
--- PASS: TestController_DeleteItem (0.00s)
PASS
ok      github.com/huyanhvn/quick-rest-go/controller      (cached)
?       github.com/huyanhvn/quick-rest-go/docs    [no test files]
=== RUN   TestGetAllItems
--- PASS: TestGetAllItems (0.00s)
PASS
ok      github.com/huyanhvn/quick-rest-go/model   (cached)
?       github.com/huyanhvn/quick-rest-go/utils   [no test files]
```

# Deployment

## Docker image

```shell
docker build -t quick-rest:1.0.0 .
```

## Kubernetes

```shell
kubectl apply -f kubernetes.yaml
```
