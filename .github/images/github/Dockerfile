# Brew Installs
FROM --platform=linux/amd64 cimg/go:1.20.11-node as brew-installs
RUN sudo apt-get update && sudo apt-get install build-essential procps curl file git && \
    NONINTERACTIVE=1 /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)" && \
    echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> /home/circleci/.profile
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install protobuf@3
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install protoc-gen-grpc-web
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install swift-protobuf
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install grpc-swift
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install hugo
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install fastly/tap/fastly
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install fastly/tap/fastly
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install golangci-lint
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install pnpm

# Python Installs
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && brew install poetry flake8

# Go Installs
FROM --platform=linux/amd64 cimg/go:1.20.11 AS go-installs
RUN go install github.com/envoyproxy/protoc-gen-validate@v0.9.0
RUN go install github.com/lyft/protoc-gen-star/protoc-gen-debug@v0.6.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
RUN go install github.com/gogo/protobuf/protoc-gen-gogo@v1.3.2
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.14.0
RUN go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.8
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.14.0
RUN go install mvdan.cc/gofumpt@v0.7.0

# Buf Installs
FROM --platform=linux/amd64 bufbuild/buf:1.13.1 AS buf

# NX Environment
FROM --platform=linux/amd64 cimg/go:1.20.11-node
WORKDIR /home/circleci

# Bashrc
COPY --from=brew-installs --chown=circleci /home/linuxbrew/.linuxbrew /home/linuxbrew/.linuxbrew
RUN echo 'export PATH=/home/linuxbrew/.linuxbrew/bin:$PATH' >> ~/.bashrc
RUN echo 'export PATH=/home/linuxbrew/.linuxbrew/sbin:$PATH' >> ~/.bashrc
RUN /home/linuxbrew/.linuxbrew/bin/brew link --overwrite protobuf@3

# Dotnet
RUN WGET_DOTNET="https://packages.microsoft.com/config/ubuntu/$(lsb_release -sr)/packages-microsoft-prod.deb" && \
    wget "$WGET_DOTNET" -O packages-microsoft-prod.deb && \
    sudo dpkg -i packages-microsoft-prod.deb && \
    rm packages-microsoft-prod.deb
RUN sudo touch /etc/apt/preferences.d/99microsoft-dotnet.pref && \
    echo $'Package: *\nPin: origin "packages.microsoft.com"\nPin-Priority: 1001' | sudo tee /etc/apt/preferences.d/99microsoft-dotnet.pref

# apt java, dotnet
RUN sudo apt-get update && sudo apt-get install -y openjdk-17-jdk aspnetcore-runtime-6.0 dotnet-sdk-7.0
RUN curl -o go/bin/protoc-gen-grpc-java -fsSL https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.54.0/protoc-gen-grpc-java-1.54.0-linux-x86_64.exe
RUN chmod +x go/bin/protoc-gen-grpc-java

# Golang
COPY --from=go-installs /home/circleci/go/bin /home/circleci/go/bin

# Buf
COPY --from=buf /usr/local/bin/buf /home/circleci/go/bin

# Node
RUN echo 'export PATH=/home/circleci/project/node_modules/.bin:$PATH' >> ~/.bashrc

# Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --default-toolchain stable
RUN echo 'export PATH=/home/circleci/.cargo/bin:$PATH' >> ~/.bashrc
RUN /home/circleci/.cargo/bin/rustup target add wasm32-wasi --toolchain stable
RUN /home/circleci/.cargo/bin/rustup toolchain add stable
RUN /home/circleci/.cargo/bin/rustup target add wasm32-wasi --toolchain stable
