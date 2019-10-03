FROM golang:latest

RUN mkdir /app
WORKDIR /app
ENV GO111MODULE=on
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
