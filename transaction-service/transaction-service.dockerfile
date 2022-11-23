FROM alpine:latest


RUN mkdir /app

COPY /dist/tx_service /app

CMD [ "/app/tx_service" ]