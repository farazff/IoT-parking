FROM golang:1.19 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /binary

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /binary /binary

EXPOSE 8080

ENTRYPOINT ["/binary", "serve"]
