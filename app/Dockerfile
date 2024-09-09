FROM alpine:3.19.0

ARG BINARY_PATH=runlocal/main

WORKDIR /app

COPY ${BINARY_PATH} .

ENTRYPOINT ["/app/main"]
