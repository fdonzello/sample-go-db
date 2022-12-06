FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/fdonzello/sample-go-db/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/fdonzello/sample-go-db/app .
CMD ["./app"]
