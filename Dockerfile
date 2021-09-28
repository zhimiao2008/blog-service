FROM alpine

WORKDIR /blog

COPY configs configs
COPY blog-service ./

EXPOSE 8000

ENTRYPOINT [ "/blog/blog-service" ]
