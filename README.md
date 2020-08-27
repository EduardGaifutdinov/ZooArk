### Create .env file from .env.example
```
cp .env.example .env
```
___
### If you cloned this project somewhere outside of $GOPATH/src/ directory
### please be sure to update godotenv.Load() path in ./src/config/.env.go file
___
### Install all Go dependencies
```
go get -v ./...
```
### Generate and update swagger docs
```
go get -u github.com/swaggo/swag/cmd/swag
swag init
```
### [How to describe swagger routes](https://github.com/swaggo/swag/blob/master/README.md)
#### [Examples](https://github.com/swaggo/swag/blob/master/example/celler/controller/examples.go)
___
### How to 
##### Run migrations and seeds
```
go run db/migrate.go
```

use flag -count=1 to clear cache
##### Run linter
```
go fmt ./...
```
___

## Without live reload
```
go run main.go
```
