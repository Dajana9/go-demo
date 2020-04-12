FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /app 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o webserver .

CMD ["./webserver"]