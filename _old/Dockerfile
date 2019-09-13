FROM golang:alpine3.10 as build

WORKDIR /build
COPY main.go main.go
RUN go build -o rultor .

FROM alpine:3.10

LABEL "com.github.actions.name"="Hello world action"
LABEL "com.github.actions.icon"="shield"
LABEL "com.github.actions.color"="green"

WORKDIR /app
COPY --from=build /build/rultor rultor
RUN apk --update add git
CMD ["/app/rultor"]
