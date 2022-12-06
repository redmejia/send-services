FROM alpine:latest

RUN mkdir /app

COPY /dist/wallet /app

CMD [ "/app/wallet" ]