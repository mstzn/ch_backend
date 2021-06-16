FROM golang:rc-alpine3.13
RUN apk add build-base

RUN mkdir -p /app
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

ADD ./ /app

RUN go build ./main.go

CMD ["./main"]