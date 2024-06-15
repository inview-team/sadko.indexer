FROM golang:1.21 AS builder
ENV PROJECT_PATH=/app/vservice
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}
RUN go build cmd/vservice/main.go

FROM golang:alpine
WORKDIR /app/cmd/vservice
COPY --from=builder /app/vservice/main .
EXPOSE 30001
CMD ["./main"]