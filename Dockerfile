FROM golang:1.21.1-alpine as builder

RUN apk update && apk add --no-cache git build-base

WORKDIR /app
COPY . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD CompileDaemon --build="go build cmd/app/main.go" --command="./main" --color