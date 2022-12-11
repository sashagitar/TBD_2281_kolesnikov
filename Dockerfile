FROM golang:1.19.3-alpine as gouni

WORKDIR /go/gouni
COPY . .
RUN go build -o buy_list ./

FROM alpine:3.17 as release
WORKDIR /app

COPY --from=gouni /go/gouni/buy_list ./

CMD ["/app/buy_list"]