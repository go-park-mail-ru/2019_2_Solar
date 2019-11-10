FROM golang:latest

WORKDIR /app
ENV GO111MODULE "on"
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o main simple-server.go

EXPOSE 8080

CMD ["./main"]
