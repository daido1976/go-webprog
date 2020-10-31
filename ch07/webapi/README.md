```sh
# Start-up postgres with Docker
$ docker-compose -f ./ch07/webapi/docker-compose.yml up -d
# Start-up HTTP Server
$ go run ./ch07/webapi/server.go ./ch07/webapi/data.go
# See curl-script and run the curl command
```
