FROM golang:1.18-alpine AS build

WORKDIR /usr/src/app

COPY  go.mod .
COPY  go.sum .

RUN go mod download
COPY . .
RUN go build -o ./bin/server  ./cmd/server/main.go



FROM alpine
WORKDIR /usr/src/app
COPY --from=build /usr/src/app/bin/ ./bin/

