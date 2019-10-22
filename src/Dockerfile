FROM golang

WORKDIR /golang-csv-postgres/src
ADD . .

RUN go get
RUN go build -o main .
RUN go get "github.com/Nhanderu/brdoc"

CMD ["go", "run", "main.go"]