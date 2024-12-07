FROM golang:1.23-bullseye AS bench-builder
WORKDIR /app
COPY . /app/
ARG ISUCON_NAME='kayac-listen80'

RUN bash bench-build.sh

# ---------------------------------------------------------------------
# メインステージ
FROM golang:1.23-bullseye AS main-builder
WORKDIR /app
COPY . /app/

RUN go mod download -x
RUN go build -o /app/kinben /app/cmd/kinben/main.go


# ---------------------------------------------------------------------
FROM debian:11.11-slim AS runtime
WORKDIR /app

ARG PORT=8080
ENV PORT=${PORT}
ARG ISUCON_NAME='kayac-listen80'
ENV ISUCON_NAME=${ISUCON_NAME}

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates openssl && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=bench-builder /app/bench /app/bench
COPY --from=bench-builder /app/data /app/data
COPY --from=main-builder /app/kinben /app/kinben

CMD [ "/app/kinben" ]
