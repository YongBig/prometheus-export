FROM golang:1.17 as builder
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -o deploy-worker main.go && chmod +x deploy-worker

FROM alpine:3.17.0
WORKDIR /opt/www
COPY --from=builder /workspace/deploy-worker .

CMD ./deploy-worker
EXPOSE 9100