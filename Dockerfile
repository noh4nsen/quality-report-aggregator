FROM alpine:3.18.3

COPY ./build/quality-report-aggregator /usr/bin/quality-report-aggregator
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]