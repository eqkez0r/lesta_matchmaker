FROM golang:1.22.1-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git go-task gcc gettext musl-dev

# dependecies
COPY ../go.mod go.sum ./
RUN go mod download

COPY .. ./
RUN go build -o ./build/server ./cmd/server/main.go
RUN go build -o ./build/agent ./cmd/agent/main.go

FROM alpine AS server_runner

COPY --from=builder /usr/local/src/build/server /

CMD ["./server"]

FROM alpine AS agent_runner

COPY --from=builder /usr/local/src/build/agent /

CMD ["./agent"]