FROM alpine

ENV PORT 80
EXPOSE 80

RUN mkdir -p /public
ADD bin/server /server
ADD public /public

CMD ["/server"]
