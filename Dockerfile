FROM golang:1.17-alpine as builder

WORKDIR /go/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app ./cmd/web

FROM gcr.io/distroless/base
COPY --from=builder /go/app/ .

CMD ["/app"]
