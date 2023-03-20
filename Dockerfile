FROM cr.loongnix.cn/library/alpine:3.11

LABEL maintainer "znley <shanjiantao@loonson.cn>"

ARG BIN
ADD $BIN /server

ENTRYPOINT ["/server"]
