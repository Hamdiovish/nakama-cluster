FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build samples/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/app/main .

EXPOSE 10000-20000 
EXPOSE 30000-40000

CMD ["main"]
