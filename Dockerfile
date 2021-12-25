FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY . .

RUN go mod tidy -v

CMD ["go", "run" ,"cmd/main.go"]