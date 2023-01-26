FROM golang:1.19-alpine

WORKDIR /app

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o my-app ./main.go

CMD ["./my-app"]