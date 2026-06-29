FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/raizes-api .

FROM alpine:3.22

WORKDIR /app

RUN adduser -D -g '' appuser

COPY --from=builder /app/raizes-api /app/raizes-api
COPY --from=builder /app/docs /app/docs

EXPOSE 8080

USER appuser

CMD ["/app/raizes-api"]
