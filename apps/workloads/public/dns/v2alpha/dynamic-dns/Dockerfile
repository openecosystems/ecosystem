FROM gcr.io/distroless/static-debian11
COPY --from=busybox:1.36.0-uclibc /bin/cp /bin/cp

WORKDIR /
COPY . .


ENTRYPOINT ["/nebula"]
# Allow users to override the args passed to nebula
CMD ["-config", "/config.yml"]
