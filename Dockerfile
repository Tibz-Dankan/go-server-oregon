FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o ./bin/myapp ./

EXPOSE 8000

CMD ["./bin/myapp"]
