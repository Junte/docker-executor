FROM golang:1.15.2-alpine3.12 as builder

ENV CGO_ENABLED=0

WORKDIR /go/src/executor
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GO111MODULE=on go build -i -v -a -installsuffix cgo -o app executor/cmd/server

FROM docker:stable-dind
WORKDIR /root
COPY --from=builder /go/src/executor/app .

CMD ["./app"]