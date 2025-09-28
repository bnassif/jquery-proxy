# Final Stage -- Copy the binary from the build stage to a new container
FROM alpine:3

LABEL Name=jquery-proxy
# Version label supplied by build process; passed with `--label "Version=$(VERSION)"`

# Copy artifact
COPY ./dist/jquery-proxy /jquery-proxy

RUN addgroup -S app && adduser -S app -G app
RUN chown app:app /jquery-proxy

USER app

EXPOSE 8080

ENTRYPOINT ["/jquery-proxy"]