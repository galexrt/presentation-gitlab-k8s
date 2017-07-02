FROM busybox:1.26.2-glibc

COPY app /bin/app

RUN chmod +x /bin/app

CMD ["/bin/app"]
