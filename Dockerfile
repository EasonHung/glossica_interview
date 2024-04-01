FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# build go app
RUN go mod download
RUN go build ./main.go

CMD ["./main"]