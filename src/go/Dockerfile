FROM golang

COPY *.go /go/src/server/

WORKDIR /go/src/server

RUN go get -d -v ./
RUN go install -v ./

CMD ["server"]