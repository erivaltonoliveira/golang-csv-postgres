FROM golang

ENV GOPATH $PWD
ENV PATH $GOPATH:/usr/local/go/bin:$PATH

#RUN mkdir -p "$GOPATH/src" "$GOPATH/src" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src

#WORKDIR /htdocs/golang-csv-postgres/src

ADD . .

#RUN go get
RUN go get "github.com/Nhanderu/brdoc"
RUN go get "github.com/lib/pq"
RUN go build -o main .

CMD ["go", "run", "main.go"]