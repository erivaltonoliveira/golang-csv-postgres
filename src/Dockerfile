FROM golang

ENV GOPATH $PWD
ENV PATH $GOPATH:/usr/local/go/bin:$PATH

WORKDIR $GOPATH/src

ADD . .

RUN go get "github.com/Nhanderu/brdoc"
RUN go get "github.com/lib/pq"
RUN go build -o main .

CMD ["go", "run", "main.go"]