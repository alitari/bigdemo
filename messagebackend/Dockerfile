FROM golang:1.11.1 as dep-builder
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go get -d -v ./...

FROM golang:1.11.1 as exe-builder
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY --from=dep-builder /go /go
COPY main.go .
COPY main_test.go .
RUN go test -run Mock -v
RUN go build -o /app main.go

FROM golang:1.11.1
CMD ["./app"]
COPY --from=exe-builder /app .
