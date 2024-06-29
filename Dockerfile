#Creates binary stage
FROM golang:1.22.4-alpine3.20 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags "-w -s" -o main

#Deploy
FROM alpine:latest as final
RUN apk update
COPY --from=builder /app/main .
CMD ["./main"]

