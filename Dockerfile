FROM golang:1.20-alpine

WORKDIR /usr/src/app

COPY . .

RUN apk add --no-cache git

RUN go mod tidy

CMD ["go", "run", "./cmd/main.go"]
