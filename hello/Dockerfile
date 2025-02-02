FROM golang:1.23.5-alpine AS builder

WORKDIR /app

# ルートの go.mod/go.sum とプロト用ファイル、ソースをコピー
COPY go.mod go.sum ./
COPY proto/ ./proto/
COPY hello/ ./hello/

WORKDIR /app/hello

RUN go mod download

# バイナリのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o hello .

# 実行ステージ
FROM scratch

COPY --from=builder /app/hello/hello /hello

EXPOSE 50051

ENTRYPOINT ["/hello"]
