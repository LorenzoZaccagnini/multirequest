# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /multirequest

EXPOSE 8080

CMD [ "/multirequest" ]