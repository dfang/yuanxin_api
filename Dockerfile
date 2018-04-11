FROM alpine:3.4

EXPOSE 9090

ADD news /bin/news

CMD ["news"]
