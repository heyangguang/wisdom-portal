# VERSION 1.0
# Author: heyang

FROM 172.16.140.21/heyang/busybox:1.28.4-glibc

MAINTAINER heyang <13833232533@163.com>

COPY wisdoms-ctl /bin/wisdoms-ctl

RUN chmod +x /bin/wisdoms-ctl

ARG DB

ENV db_env_var=$DB

RUN mkdir -p /opt/wisdom

WORKDIR /opt/wisdom

RUN mkdir logs && mkdir static

CMD /bin/wisdoms-ctl run --db "$db_env_var"