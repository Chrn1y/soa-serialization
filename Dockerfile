FROM golang:latest

WORKDIR soa-serialization
COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]
