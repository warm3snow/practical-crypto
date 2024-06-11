FROM chyiyaqing/kingbase:v8r6 as builder


FROM centos:7

RUN groupadd kingbase && useradd -g kingbase -m -d /opt/kingbase -s /bin/bash kingbase

RUN usermod -u 1001 kingbase && groupmod -g 1001 kingbase

WORKDIR /opt/kingbase

COPY --from=builder /opt/kingbase ./

RUN chmod +x docker-entrypoint.sh

ADD ./license.dat ./

RUN chown -R kingbase:kingbase /opt/kingbase

ENV PATH /opt/kingbase/Server/bin:$PATH

ENV DB_VERSION V008R006C005B0014

USER kingbase

EXPOSE 54321

ENTRYPOINT ["/opt/kingbase/docker-entrypoint.sh"]