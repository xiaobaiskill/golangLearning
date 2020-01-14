FROM golang:1.13.4 as build

ENV GOPATH /go/cache
ENV GOPROXY https://goproxy.cn
ENV GOBIN ${GOPATH}/bin

WORKDIR /go/cache/src
ADD . .

RUN go get 2>&1

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags "-w -s" -installsuffix cgo -o app .


FROM alpine as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/cache/src /home

WORKDIR /home

EXPOSE 8082

CMD ["./app"]