FROM golang:1.11.1 as dep-builder
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY src/go.mod .
COPY src/go.sum .
RUN go get -d -v ./...

FROM golang:1.11.1 as exe-builder
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY --from=dep-builder /go /go
COPY src/main.go .
COPY src/main_test.go .
RUN go test -run Mock -v
RUN go build -o /app main.go

FROM golang:1.11.1
CMD ["./app"]
COPY --from=exe-builder /app .
