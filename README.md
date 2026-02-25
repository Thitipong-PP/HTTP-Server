# HTTP Server and API (Basic)

The simulation of HTTP Server for website and api

> [!NOTE]
> 1) Server start at port 8080 <br>
> 2) The server api don't have database we use slice to mock database now <br>
> 3) Goroutine don't have in this project.

## Getting the Source Code
Clone the repository to your local machine:

``` bash
$ cd ${WORKDIR}
$ git clone https://github.com/Thitipong-PP/HTTP-Server.git
$ cd HTTP-Server
```

And try to run
``` bash
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

``` go
type Server struct {
	ID int `json:"id"`
	Name string `json:"name"`
	IP string `json:"ip"`
	Status bool `json:"status"`
}
```

A example json data (You don't need to put ID when post or put because ID will auto increment)
``` json
{
    "name": "TP_Server",
    "ip": "123.456.789",
    "status": true
}
```
