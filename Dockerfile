FROM golang:1.13-alpine

WORKDIR /app/goblog
COPY . .

RUN apk add tar wget

RUN go get -d -v ./...
RUN go build .
RUN ./configure.sh

CMD ["./wait-for", "db:5432", "--", "./go-blog"]
