FROM debian:stretch-slim

LABEL maintainer="wushaobo.china@gmail.com"

EXPOSE 80
RUN mkdir -p /var/log/zipper

COPY docker/prod/run.sh /opt/zipper/run.sh
COPY bin/zipper /opt/zipper/zipper
