FROM alpine:latest

RUN mkdir /app

COPY /dist/fake_bank /app

CMD [ "/app/fake_bank" ]