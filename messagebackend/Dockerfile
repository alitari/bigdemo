FROM golang:1.10.1 as builder
WORKDIR /go/src/app
COPY src/main/main.go .
RUN go get -d -v ./...
RUN go build -o /app main.go

FROM golang:1.10.1
CMD ["./app"]
COPY --from=builder /app .
