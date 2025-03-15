FROM gcr.io/distroless/static-debian11
COPY --from=busybox:1.36.0-uclibc /bin/cp /bin/cp

WORKDIR /
COPY v1alpha .

ENTRYPOINT ["/app"]
