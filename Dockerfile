FROM golang:1.13-alpine

WORKDIR /app/goblog
COPY . .

RUN apk add tar wget

RUN go get -d -v ./...
RUN go build .
RUN ./download_fonts.sh

CMD ["./go-blog"]
