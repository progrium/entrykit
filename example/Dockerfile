FROM gliderlabs/alpine:latest
EXPOSE 80

RUN apk-install nginx
RUN mkdir -p /demo/includes
RUN echo "Hello world" > /demo/index.html
RUN echo "" > /demo/empty.html

COPY ./nginx.conf.tmpl /demo/nginx.conf.tmpl
COPY ./reloader /bin
COPY ./entrykit /bin
RUN entrykit --symlink

ENTRYPOINT [ \
  "switch", \
    "shell=/bin/sh", \
    "version=nginx -v", "--", \
  "render", "/demo/nginx.conf", "--", \
  "prehook", "nginx -V", "--", \
  "codep", \
    "/bin/reloader 3", \
    "/usr/sbin/nginx -c /demo/nginx.conf" ]
