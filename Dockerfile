FROM golang:latest

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY simple-server.go .
COPY simple-server_test.go .
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
