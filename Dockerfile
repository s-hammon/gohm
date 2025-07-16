FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY bin/gohm /bin/gohm

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT [ "/bin/gohm" ]
