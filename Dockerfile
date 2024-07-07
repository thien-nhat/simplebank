#Build stage
FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .

RUN go build -o main main.go

#Run stage
FROM alpine
WORKDIR /app 

COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
RUN chmod +x start.sh

COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]