FROM golang
MAINTAINER Gaurav Dhameeja gaurav.dhameeja@fiftyfive.io
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64
CMD mkdir /go/src/gdhameeja/work

WORKDIR /go/src/gdhameeja/work

# setup db
RUN apt-get update
RUN apt-get install sqlite3
COPY boot.sql /tmp/
RUN cat /tmp/boot.sql | sqlite3 books.db

# SETUP CODE
COPY main.go /go/src/gdhameeja/work
COPY go.mod /go/src/gdhameeja/work
COPY go.sum /go/src/gdhameeja/work
RUN go mod download

RUN go build main.go
EXPOSE 8888

ENTRYPOINT ["./main"]
