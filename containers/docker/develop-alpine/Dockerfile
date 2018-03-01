FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/Sberex/go-sberex && \
  (cd go-sberex && make geth) && \
  cp go-sberex/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-sberex && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/geth"]
