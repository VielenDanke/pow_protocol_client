FROM golang:latest AS builder
WORKDIR /go/src/github.com/vielendanke/pow_protocol_server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
ENV CLIENT_ADDRESS=':8090'
ENV CLIENT_USERNAME='user'
ENV CLIENT_PASSWORD='password'
WORKDIR /root/
COPY --from=builder /go/src/github.com/vielendanke/pow_protocol_server/app ./
CMD ["./app"]