#stageI (Build Binary)
FROM golang:1.18.3-alpine3.16 AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main

#stageII
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8000
CMD [ "./main" ]
