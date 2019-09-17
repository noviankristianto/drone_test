FROM golang:1.13-alpine3.10 as builder
WORKDIR /src

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o drone_test


FROM alpine:3.10.2
COPY --from=builder /src/drone_test drone_test

CMD "./drone_test"
