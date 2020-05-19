Just a simple go CRUD app

Adapted from [this post](https://medium.com/@etiennerouzeaud/how-to-create-a-basic-restful-api-in-go-c8e032ba3181) with the goal of using as few outside packages as possible.

```
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/mattn/go-sqlite3
go run .
```

To run w/ autoreloads,

```
go get github.com/silenceper/gowatch
gowatch
```

`curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Test\", \"lastname\": \"User\" }" http://localhost:8080/user`

`curl http://localhost:8080/users`