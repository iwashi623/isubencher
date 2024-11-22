FROM golang:1.23 AS builder
WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/,sharing=locked \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    go build -o /bin/bench-runner

# サブモジュールのアップデート
RUN --mount=type=cache,target=/go/pkg/mod/,sharing=locked \
    --mount=type=bind,target=. \
    git submodule update --init --recursive

# Makefileの実行を実行して、ベンチマーカーのビルドと必要なデータの取得
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    make init

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /bin/bench-runner /app/bench-runner
COPY --from=builder /bin/bench /app/bench
COPY ./data /app/

CMD ["/app/myapp"]
