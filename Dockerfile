#FROM golang
#
#WORKDIR /go/src
#ENV PATH="/go/bin:${PATH}"
#ENV CGO_ENABLED=0
#
#RUN go install github.com/ribeirosaimon/go_flight_api@latest
#
#CMD ["tail", "-f", "/dev/null"]


FROM golang as builder

WORKDIR /build/api
COPY go.mod ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o api
# post build stage
FROM alpine
WORKDIR /root
COPY --from=builder /build/api/api .
EXPOSE 3000
CMD ["./api"]