FROM alpine:latest

RUN mkdir /app

COPY /dist/controller /app


CMD [ "/app/controller" ]