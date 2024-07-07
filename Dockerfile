#Build stage
FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .

RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz

#Run stage
FROM alpine
WORKDIR /app 

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
RUN chmod +x start.sh

COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]