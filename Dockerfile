FROM vidsyhq/go-base:latest
LABEL maintainer="Nick Pocock"

RUN go build -o recs-api

ADD recs-api /
ADD config /config

ENTRYPOINT ["/recs-api"]