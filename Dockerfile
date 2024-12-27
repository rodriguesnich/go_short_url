# Build Stage
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src .

RUN go build -v -o server ./cmd/main.go

# Final Stage
FROM scratch

COPY --from=builder /app/server /server

COPY --from=builder /app/public /public

ENTRYPOINT ["./server"]
