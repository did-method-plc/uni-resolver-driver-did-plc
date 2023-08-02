
### Compile stage
FROM golang:1.20-alpine3.17 AS build-env

ADD . /dockerbuild
WORKDIR /dockerbuild

# timezone data for alpine builds
RUN go build -tags timetzdata -o /uni-resolver-driver-did-plc main.go

### Run stage
FROM alpine:3.17

RUN apk add --no-cache --update dumb-init ca-certificates
ENTRYPOINT ["dumb-init", "--"]

WORKDIR /
RUN mkdir -p data/uni-resolver-driver-did-plc
COPY --from=build-env /uni-resolver-driver-did-plc /

# small things to make golang binaries work well under alpine
ENV GODEBUG=netdns=go
ENV TZ=Etc/UTC

CMD ["/uni-resolver-driver-did-plc"]
EXPOSE 8000

LABEL org.opencontainers.image.source=https://github.com/bnewbold/uni-resolver-driver-did-plc
LABEL org.opencontainers.image.description="did:plc Universal Resolver Driver"
LABEL org.opencontainers.image.licenses=MIT
