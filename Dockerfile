# VERSION 1.0
# Author: heyang

FROM busybox:1.28.4-glibc

MAINTAINER heyang <13833232533@163.com>

COPY app /bin/app

RUN chmod +x /bin/app

CMD ["/bin/app"]