FROM golang:alpine3.16
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0
WORKDIR /app

RUN apk add curl

ADD go.mod .
ADD go.sum .
RUN go mod download
RUN go mod verify
ADD . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000
EXPOSE 27017

ENTRYPOINT CompileDaemon --build="go build -o main app/main.go" --command=./main