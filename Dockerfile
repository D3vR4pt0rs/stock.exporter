FROM golang:latest AS builder
ENV PROJECT_PATH=/app/exporter
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}
RUN go build cmd/exporter/main.go

FROM golang:alpine
WORKDIR /app/cmd/exporter
COPY --from=builder /app/exporter/main .
EXPOSE 5000
CMD ["./main"]
