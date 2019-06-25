FROM golang:1.12.6-alpine3.9

WORKDIR /usr/src/app
COPY . .

RUN go build -o app .
CMD ["app"]
