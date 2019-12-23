FROM golang:1.12.10-alpine as builder

WORKDIR /go/src/github.com/mpolden/wakeup
RUN apk --no-cache add bash make gcc libc-dev git

RUN go get github.com/jessevdk/go-flags
COPY . /go/src/github.com/mpolden/wakeup

RUN make install

FROM alpine:3.8

COPY --from=builder /go/src/github.com/mpolden/wakeup/static /opt/wakeup/
COPY --from=builder /go/bin /opt/wakeup

EXPOSE 8080

CMD ["-c" ,"/opt/wakeup/wakeup-cache","-s","/opt/wakeup/static"]
ENTRYPOINT [ "/opt/wakeup/wakeup" ]
