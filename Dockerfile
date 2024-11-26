# kayac-listen80-builderステージ
FROM golang:1.23 AS kayac-listen80-builder
WORKDIR /app
COPY . /app/

ARG COMPETITION='kayac-listen80'
ENV PORT=${PORT}

# Makefileの実行を実行して、ベンチマーカーのビルドと必要なデータの取得
RUN git clone https://github.com/kayac/kayac-isucon-2022.git  && make init


# メインステージ
FROM golang:1.23 AS main-builder
WORKDIR /app
COPY . /app/

RUN go mod download -x
RUN go build -o /app/bench-runner
CMD ["/app/bench-runner"]


FROM debial:11.11-slim AS runtime
WORKDIR /app

ARG PORT=8080
ENV PORT=${PORT}

COPY --from=main-builder /app/bench-runner /app/bench-runner
COPY --from=main-builder /app/bench /app/bench
