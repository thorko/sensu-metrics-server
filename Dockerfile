From alpine:latest

RUN apk update && apk add git go libc-dev
RUN mkdir /tmp/go && export GOPATH=/tmp/go
COPY cmd/getmetrics.go /tmp/go
RUN cd /tmp/go && go get || : 
RUN cd /tmp/go && env GOOS=linux GARCH=amd64 go build -o /tmp/getmetrics getmetrics.go
CMD ["/bin/sleep", "300"]
