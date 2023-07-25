FROM golang:1.19.0

WORKDIR /usr/src/app

COPY ./go.mod ./
COPY ./go.sum ./
COPY ./src ./src

RUN go mod download

RUN go build -o ./bin/app ./src/server.go
EXPOSE 8080
CMD [ "./bin/app" ]