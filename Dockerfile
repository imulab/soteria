FROM golang:1.11
RUN mkdir -p /go/src/github.com/imulab/soteria
ADD . /go/src/github.com/imulab/soteria
WORKDIR /go/src/github.com/imulab/soteria
RUN CGO_ENABLED=0 GOOS=linux go build -o soteria .

FROM alpine:3.9.2
RUN apk --no-cache add ca-certificates
WORKDIR /bin
COPY --from=0 /go/src/github.com/imulab/soteria/soteria .

CMD ["soteria"]