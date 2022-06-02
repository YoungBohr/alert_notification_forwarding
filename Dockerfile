FROM  registry-vpc.cn-zhangjiakou.aliyuncs.com/data100/alpine:latest AS dev

ADD --chown=nobody:nobody $CI_PROJECT_DIR/anf /app/

RUN chmod +x /app/anf

COPY $CI_PROJECT_DIR/Shanghai /etc/localtime

WORKDIR /app

#ENV GIN_MODE=release

CMD ["/app/anf"]