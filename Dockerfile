FROM golang:1.20 AS builder

WORKDIR /algohub
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o build/algohub .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /algohub/build/algohub .

EXPOSE 80
EXPOSE 443

CMD ["./alogohub"]