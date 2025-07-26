RUNSC_VOLUME_NAME="runsc-bin" docker compose run --rm runsc-installer
docker run --rm --runtime runsc ubuntu uname -a

docker inspect temporal-trino-1 | grep Runtime
docker inspect trino2 | grep Runtime
