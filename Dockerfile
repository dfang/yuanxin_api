FROM golang:1.10 AS builder
WORKDIR $GOPATH/src/github.com/dfang/yuanxin_api
ADD . ./
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/yuanxin_api .

FROM alpine:3.4
LABEL maintainer="df1228@gmail.com"
EXPOSE 9090
COPY --from=builder /go/bin/* /usr/local/bin/
CMD [ "yuanxin_api" ]
