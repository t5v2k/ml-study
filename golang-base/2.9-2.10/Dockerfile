FROM golang:1.21.1-bullseye AS builder
WORKDIR /workspace
COPY main/go.mod go.mod
RUN go mod download
COPY main/main.go main.go
RUN go build -a -o example main.go

FROM gcr.io/distroless/base-debian11:latest
WORKDIR /
COPY --from=builder /workspace/example .
ENTRYPOINT ["/example"]
EXPOSE 8080
EXPOSE 8089
