FROM busybox:1.30.1-glibc

COPY app /bin/app

RUN chmod +x /bin/app

CMD ["/bin/app"]
