FROM golang:latest AS build

WORKDIR /app

RUN go mod download

COPY . ./

RUN go build -o /binary

## Deploy
FROM golang:latest

WORKDIR /

COPY --from=build /binary /binary

EXPOSE 8080

ENTRYPOINT ["/binary", "serve"]
