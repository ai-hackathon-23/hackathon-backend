FROM golang:1.16-alpine

RUN mkdir /app
WORKDIR /app

COPY ./app .

RUN go mod tidy
RUN go mod download

RUN go build -o main .

CMD ["/app/main"]