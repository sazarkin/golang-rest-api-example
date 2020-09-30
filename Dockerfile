FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git
WORKDIR /app/golang-rest-api-example
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go test -v ./...
RUN go build -o app.exe ./cmd/app

FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=build_base /app/golang-rest-api-example/app.exe /usr/bin
EXPOSE 8080
ENTRYPOINT ["app.exe"]