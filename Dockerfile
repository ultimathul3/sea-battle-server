FROM golang:1.20

WORKDIR /usr/src/sea-server

COPY . .

RUN go build -o sea-server cmd/app/main.go

EXPOSE 8080

CMD ["./sea-server"]
