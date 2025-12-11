## 1. Dependencia de Testing a instalar

```sh
go get github.com/stretchr/testify
```

## 2. Tidy

```sh
go mod tidy
```


## 2. Comandos para correr
```sh

go test -coverprofile=coverage.out ./...
 
go tool cover -func=coverage
 
go tool cover -html=coverage -o coverage.html
 
start .\coverage.html
 
```