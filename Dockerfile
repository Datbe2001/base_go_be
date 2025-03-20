FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

WORKDIR /build

RUN go build -o base_go ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/base_go /

ENTRYPOINT ["/base_go", "config/local.yaml"]
