FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]