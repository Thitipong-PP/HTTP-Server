# HTTP Server and API (Basic)
[![GoDoc](https://godoc.org/github.com/google/trillian?status.svg)](https://godoc.org/github.com/google/trillian)

The simulation of HTTP Server for website and api

> [!NOTE] <br>
> Server start at port 8080 <br>
> The server api don't have database we use slice to mock database now <br>
> Goroutine don't have in this project.

## Getting the Source Code
Using `git clone` allows you to work in whatever directory you want. You will
still need to set GOPATH in order to build some apps (recommended to put this in
a cache dir). E.g.:

```
$ cd ${WORKDIR}
$ git clone https://github.com/Thitipong-PP/HTTP-Server.git
$ cd HTTP-Server
```

And try to run
```
$ go run main.go
```

## Routing of server api
GET Method
```
localhost:8080/server
localhost:8080/server/<id>
```
POST Method
```
localhost:8080/server
```
DELETE Method
```
localhost:8080/server/<id>
```
PUT Method
```
localhost:8080/server/<id>
```

## Server Struct
This is structure of server struct in project

```
type Server struct {
	ID int `json:"id"`
	Name string `json:"name"`
	IP string `json:"ip"`
	Status bool `json:"status"`
}
```

A example json data (You don't need to put ID when post or put because ID will auto increment)
```
{
    "name": "TP_Server",
    "ip": "123.456.789",
    "status": true
}
```