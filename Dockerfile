FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd

EXPOSE 3000 2112

CMD ["./main"]