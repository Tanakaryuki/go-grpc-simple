FROM golang:1.23.5-alpine AS builder

WORKDIR /app

# ルートの go.mod/go.sum とプロト用ファイル、ソースをコピー
COPY go.mod go.sum ./
COPY proto/ ./proto/
COPY gateway/ ./gateway/

WORKDIR /app/gateway

RUN go mod download

# バイナリのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o gateway .

# 実行ステージ
FROM scratch

COPY --from=builder /app/gateway/gateway /gateway

EXPOSE 8080

ENTRYPOINT ["/gateway"]
