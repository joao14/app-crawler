FROM golang:alpine


WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build main.go

RUN find . -name "*.go" -type f -delete

EXPOSE 8010

CMD ["/app/main"]
