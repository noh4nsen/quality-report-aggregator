FROM scratch

COPY ./build/quality-report-aggregator /usr/bin/quality-report-aggregator

ENTRYPOINT [ "quality-report-aggregator" ]