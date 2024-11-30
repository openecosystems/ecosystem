# Setup RUST

Use Rustrover and install rust plugin

Install Fastly CLI

```shell
brew install fastly/tap/fastly
```

Create a Fastly CLI token

```
https://manage.fastly.com/account/personal/tokens
```

Configure Fastly CLI with token

```shell
fastly profile create --automation-token
```

Download and install rustup and the Rust stable toolchain:

```shell
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --default-toolchain stable
```

Update your .zshrc file to always source rust binaries

```shell
vi ~/.zshrc

# Add Rust toolchain
export PATH="$HOME/.cargo/bin:$PATH"
```

Install the wasm32-wasi target for the stable toolchain:

```shell
rustup target add wasm32-wasi --toolchain stable
```

Install Rust:

```shell
rustup toolchain add stable
```

Install the wasm32-wasi target for the stable toolchain:

```shell
rustup target add wasm32-wasi --toolchain stable
```

Update docker desktop to support running wasm workloads
https://docs.docker.com/desktop/wasm/

```browser
Open the Docker Desktop Settings.
Go to the Features in development tab.
Check the following checkboxes:
Use containerd for storing and pulling images

Restart then go back to settings and
Enable Wasm
Select Apply & restart to save the settings.
In the confirmation dialog, select Install to install the Wasm runtimes.
```

docker buildx build --platform wasi/wasm -t workloads/internal/edge/v2alpha/edge-router:latest .
docker buildx build --platform wasi/wasm --provenance=false -t workloads/internal/edge/v2alpha/edge-router:latest .
docker run -d --runtime=io.containerd.wasmedge.v1 --platform=wasi/wasm --name edge-router -p 7676:7676 workloads/internal/edge/v2alpha/edge-router:latest

docker build --platform linux/amd64 -t workloads/internal/edge/v2alpha/edge-router:latest .
docker run -d --platform=linux/amd64 --name edge-router -p 7676:7676 workloads/internal/edge/v2alpha/edge-router:latest


# To compile ring, we will not be able to use Apple's version of LLVM https://github.com/briansmith/ring/issues/1824
brew install gcc
brew install llvm

## Verify configuration
llvm-config --version

echo 'export PATH="/usr/local/opt/llvm/bin:$PATH"' >> ~/.zshrc
echo 'export LDFLAGS="-L/usr/local/opt/llvm/lib"' >> ~/.zshrc
echo 'export CPPFLAGS="-I/usr/local/opt/llvm/include"' >> ~/.zshrc

To test locally

```shell
fastly compute serve
```

To publish package:

```shell
fastly compute publish
fastly compute publish --env communication
```

To tail logs from running service

```shell
fastly log-tail
```

export FASTLY_HOSTNAME=localhost

curl \
--header "Content-Type: application/json" \
--header 'x-spec-apikey: 12345678' \
--header "x-timezone:America/New_York" \
--header "x-spec-debug: true" \
--data '{"email": "dimy@jeannot.company"}' \
http://127.0.0.1:7676/platform.communication.v1alpha.MarketingEmailService/Subscribe

curl -iv --http1.1 -X OPTIONS \
--header "Content-Type: application/json" \
-H "origin: http://localhost:8080" \
--header "x-spec-debug: true" \
-H 'access-control-request-headers: GET' \
http://127.0.0.1:7676/platform.communication.v1alpha.MarketingEmailService/Subscribe

curl 'http://localhost:7676/platform.communication.v1alpha.MarketingEmailService/Subscribe' \
-X 'OPTIONS' \
-H 'Accept: */*' \
-H 'Accept-Language: en-US' \
-H 'Access-Control-Request-Headers: content-type' \
-H 'Access-Control-Request-Method: GET' \
-H 'Connection: keep-alive' \
-H 'Origin: https://project-r6mnllyjzrbssnlskcjd.framercanvas.com' \
-H 'Sec-Fetch-Dest: empty' \
-H 'Sec-Fetch-Mode: cors' \
-H 'Sec-Fetch-Site: cross-site' \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) FramerElectron/2024.34.1 Chrome/122.0.6261.156 Electron/29.3.0 Safari/537.36'

curl -X GET \
--header "Content-Type: application/json" \
--header 'x-spec-apikey: 12345678' \
--header "x-spec-debug: true" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--header "x-spec-fieldmask: spec_context.organization_slug,configuration.id" \
http://localhost:7676/v2/configurations/123