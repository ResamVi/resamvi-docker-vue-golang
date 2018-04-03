FROM golang

COPY *.go /go/src/server/
COPY ./gaestebuch.html /go/src/server
COPY ./entry.html /go/src/server
COPY ./style.css /go/src/server

WORKDIR /go/src/server

RUN go get -d -v ./
RUN go install -v ./

CMD ["server"]