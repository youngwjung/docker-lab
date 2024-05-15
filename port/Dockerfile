FROM golang:1.20-alpine as builder
WORKDIR /app
COPY *.go ./
RUN go mod init main
RUN go build -o /main

FROM alpine
WORKDIR /
COPY --from=builder /main /main
EXPOSE 0-65535
CMD ["/main"]