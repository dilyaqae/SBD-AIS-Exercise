  # todo
# builder stage
FROM golang:1.25 AS builder 

WORKDIR /app

COPY . .

RUN sed -i 's/\r$//' scripts/build-application.sh && \
    chmod +x scripts/build-application.sh && \
    ./scripts/build-application.sh

# run stage 
FROM alpine:latest AS run
WORKDIR /app

COPY --from=builder /app/ordersystem /app/ordersystem

EXPOSE 3000
ENTRYPOINT ["/app/ordersystem"]
