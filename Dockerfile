# kayac-listen80-builderステージ
FROM golang:1.23 AS kayac-listen80-builder
WORKDIR /app
VOLUME ./

ARG COMPETITION='kayac-listen80'
ENV PORT=${PORT}

# Makefileの実行を実行して、ベンチマーカーのビルドと必要なデータの取得
RUN cd modules && git clone https://github.com/kayac/kayac-isucon-2022.git  && make init
RUN mkdir -p /app/


# メインステージ
FROM golang:1.23 AS main-builder
WORKDIR /app
COPY . /app/

RUN go mod download -x
RUN go build -o /app/bench-runner


# FROM debian:11.11-slim AS runtime
# WORKDIR /app

# ARG PORT=8080
# ENV PORT=${PORT}

# COPY --from=main-builder /app/bench-runner /app/bench-runner
# RUN if [ "$STAGE" = "isucon12" ]; then \
#         cp /app/isucon12/bench /bin/bench; \
#     elif [ "$STAGE" = "isucon13" ]; then \
#         cp /app/isucon13/bench /bin/bench; \
#     else \
#         echo "Unknown stage: $STAGE" && exit 1; \
#     fi
