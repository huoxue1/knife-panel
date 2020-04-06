FROM alpine:3.11.2
LABEL maintainer=liuxiaodong2017@gmail.com
ADD linux_knife_panel /KnifePanel
RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash
EXPOSE 10088
CMD ["/KnifePanel"]