FROM alpine:latest

RUN mkdir /app

COPY basicCRUDApp /app

CMD ["/app/basicCRUDApp"]