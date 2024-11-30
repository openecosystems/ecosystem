FROM --platform="linux/amd64" rust:1.70-slim-bullseye AS buildbase
WORKDIR /opt

RUN apt-get update && \
    apt-get install -y git clang wget && \
    rustup target add wasm32-wasi

ARG FASTLY_CLI_VERSION="10.5.0"
ARG FASTLY_ARCH="amd64"

RUN wget https://github.com/fastly/cli/releases/download/v${FASTLY_CLI_VERSION}/fastly_${FASTLY_CLI_VERSION}_linux_${FASTLY_ARCH}.deb
COPY . .
RUN apt-get -y install ./fastly_${FASTLY_CLI_VERSION}_linux_${FASTLY_ARCH}.deb
RUN tar -xvzf pkg/edge-router.tar.gz

RUN chmod -R -x pkg
ENTRYPOINT [ "fastly", "compute", "serve", "--file", "/opt/edge-router/bin/main.wasm", "--addr", "0.0.0.0:7676", "--env", "container", "--verbose", "--skip-build"]
