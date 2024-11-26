FROM golang:1.23 AS builder
WORKDIR /app
COPY . /app/

RUN go mod download -x

RUN go build -o /app/bench-runner

ARG PORT=8080
ENV PORT=${PORT}

# Makefileの実行を実行して、ベンチマーカーのビルドと必要なデータの取得
RUN git submodule update --init && make init

CMD ["/app/bench-runner"]
