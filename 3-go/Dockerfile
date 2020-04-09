FROM golang:1.10-alpine3.7 as img

ADD ./main.go /go/src
WORKDIR /go/src
RUN go build main.go

FROM golang:1.10-alpine3.7
COPY --from=img /go/src/main /main

ENTRYPOINT [ "/main" ]
EXPOSE 8000