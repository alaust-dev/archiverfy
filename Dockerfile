FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY cmd/ cmd/

RUN go build -a -o archiverfy cmd/archiverfy/main.go


FROM alpine:latest

USER 1000
WORKDIR /app

COPY --from=build /app/archiverfy .

CMD [ "/app/archiverfy" ]