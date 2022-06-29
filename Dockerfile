FROM golang:latest

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o app ./cmd/app/main.go

EXPOSE 4000

CMD [ "./app" ]

