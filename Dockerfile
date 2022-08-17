FROM golang

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN go install github.com/ribeirosaimon/go_flight_api@latest

CMD ["tail", "-f", "/dev/null"]