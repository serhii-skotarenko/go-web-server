A simple golang web server.
It can be packaged into a docker image.

Build:
```shell
docker build -t go-web-server .
```

Run:
```shell
docker run -p 3333:8080 go-web-server
```
