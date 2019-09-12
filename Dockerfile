FROM alpine:3.10

LABEL "com.github.actions.name"="Hello world action"
LABEL "com.github.actions.icon"="shield"
LABEL "com.github.actions.color"="green"

WORKDIR /app
COPY script.sh script.sh
CMD ["/app/script.sh"]
