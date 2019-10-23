FROM golang

ENV GOPATH /htdocs/golang-csv-postgres
ENV PATH $GOPATH:/usr/local/go/bin:$PATH

#RUN mkdir -p "$GOPATH/src" "$GOPATH/src" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src

#WORKDIR /htdocs/golang-csv-postgres/src

ADD . .

#RUN go get
RUN go get "github.com/Nhanderu/brdoc"
RUN go build -o main .

CMD ["go", "run", "main.go"]