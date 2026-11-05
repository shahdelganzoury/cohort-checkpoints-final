FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
