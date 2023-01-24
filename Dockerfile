FROM golang:1.19-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o todo-app ./main.go

CMD ["./todo-app"]