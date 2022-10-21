FROM golang:1.18 as builder

WORKDIR /workspace

COPY . .

RUN go build -o  bin/pubsub ./main

FROM gcr.io/distroless/base:latest
COPY --from=builder /workspace/bin/pubsub pubsub
ENTRYPOINT ["./pubsub"]