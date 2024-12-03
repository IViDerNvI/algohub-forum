FROM golang:1.23 AS builder

WORKDIR /algohub
COPY ./* ./

RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o build/algohub .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /algohub/build/algohub .

EXPOSE 80
EXPOSE 443

CMD ["./alogohub"]