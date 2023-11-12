# 2023年3月時点の最新版goの軽量版「alpine」
FROM golang:1.20.2-alpine

# ホットリロード用にAirをインストール
RUN go install github.com/cosmtrek/air@latest

# マイグレーション用にglang-migrate（MySQL）をインストール
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            git \
            gcc \
            musl-dev \
            bash

# 作業ディレクトリの設定
WORKDIR /backend/src