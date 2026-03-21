# 1. Goのバージョンを指定。alpineは、サイズが小さく、無駄なものが入っていない環境
FROM golang:1.25-alpine AS builder

# コンテナ内の作業ディレクトリを作成
WORKDIR /app

# 先に依存関係ファイルをコピー（ビルドの高速化のため）
COPY go.mod ./

# 現状、外部ライブラリなどは入れてないが、一応
RUN go mod download

# プロジェクトの全ファイルをコピー（viewsやrouteも含まれる）
COPY . .

# アプリをビルド（main という実行ファイルを作成）
RUN go build -o ./bin/main .

# 2. 実行環境 (本番でも使える軽量イメージ)
FROM scratch
WORKDIR /app

# ビルドしたバイナリだけをコピー
COPY --from=builder /app/bin/main ./bin/main
# HTMLやCSSなどの静的ファイルが入ったディレクトリもコピー
COPY --from=builder /app/views ./views

# アプリを起動
CMD ["./bin/main"]
