FROM golang:1.17-alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /anukrati

CMD ["go","run","main.go","cache_random.go"]