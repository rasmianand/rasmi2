FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go test --cover ./...

RUN go build -o main .

CMD ["./main"]
