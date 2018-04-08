FROM golang:1.10-alpine AS builder
RUN apk add --no-cache bash make
RUN apk --update add git
WORKDIR /go/src/github.com/jimareed/metrics-tutorial/
COPY . /go/src/github.com/jimareed/metrics-tutorial/
RUN make build

FROM alpine:3.6

ARG CREATED
ARG VERSION
ARG REVISION

LABEL org.opencontainers.image.created=$CREATED
LABEL org.opencontainers.image.url="metrics-tutorial"
LABEL org.opencontainers.image.source="https://github.com/jimareed/metrics-tutorial"
LABEL org.opencontainers.image.version=$VERSION
LABEL org.opencontainers.image.revision=$REVISION

EXPOSE 8080
COPY --from=builder /go/src/github.com/jimareed/metrics-tutorial/metrics-tutorial /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/metrics-tutorial && chmod +x /usr/local/bin/metrics-tutorial
USER nobody
CMD metrics-tutorial
