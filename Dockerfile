FROM golang:latest

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o main simple-server.go

EXPOSE 8080

CMD ["./main"]
