FROM golang:latest

RUN mkdir /app
WORKDIR /app
ENV GO111MODULE=on
COPY . .

RUN go get -d -v ./...
RUN go build -o main .

CMD ["./main"]
