FROM vidsyhq/go-base:latest
LABEL maintainer="Nick Pocock"

ARG VERSION
LABEL version=$VERSION

ADD recs-api /
ADD config /config

ENTRYPOINT ["/recs-api"]