FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .


RUN go build

ENV PORT ":8000"
CMD ["./randapi.com"]

