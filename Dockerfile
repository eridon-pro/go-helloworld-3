# FROM命令の末尾にAS builderとすることでステージ名を指定します
FROM golang:alpine AS builder

COPY . /app
WORKDIR /app

RUN go get

# ビルドのみに留め、CMDによるアプリケーションの起動は削除します
RUN go build -o server


# 2回目のFROM命令で、ビルドしたバイナリを実行するイメージを指定する
FROM alpine

WORKDIR /app
# builderステージのコンテナ内からビルド済みのバイナリのみコピーする
COPY --from=builder /app/server /app/

CMD ["/app/server"]
