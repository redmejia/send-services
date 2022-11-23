# build a tiny docker image
FROM alpine:latest

RUN mkdir /app


COPY /dist/users_acc_service /app


CMD [ "/app/users_acc_service" ]

