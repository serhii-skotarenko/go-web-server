FROM golang:1.20.1-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-web-server

EXPOSE 8080

CMD [ "/go-web-server" ]
