FROM golang

ARG service=scooters-api
ARG org=fguy
ARG pkg=github.com/${org}/${service}

ADD . $GOPATH/src/${pkg}

WORKDIR $GOPATH/src/${pkg}

RUN dep ensure

EXPOSE 8080

RUN make test && make
CMD ["./server"]
