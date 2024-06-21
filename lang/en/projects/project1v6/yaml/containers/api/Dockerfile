FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o api .
EXPOSE 1234
CMD ["./api"]
