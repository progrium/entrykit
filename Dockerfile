# This is for dev/debugging. No intention of an entrykit base image...
FROM alpine:3.2
ADD ./build/Linux/entrykit /bin/entrykit
RUN /bin/entrykit --symlink
RUN echo -e "#!/bin/sh\nwhile true; do sleep 1; echo \$RANDOM \$1; done" > /bin/numbers \
  && chmod +x /bin/numbers

ENTRYPOINT ["codep", "-p", "alpha=/bin/numbers foo", "beta=/bin/numbers bar"]
