FROM golang:1.10.1-alpine3.7 as builder
COPY main.go .
RUN go build -o /hello-world main.go

FROM alpine:3.7  
CMD ["./hello-world"]
COPY --from=builder /hello-world .
