FROM alpine:3.10

COPY build/justproxy_linux_amd64 /bin/justproxy

CMD [ "/bin/justproxy","/var/justproxy/conf.json" ]

