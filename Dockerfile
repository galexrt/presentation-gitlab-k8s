FROM busybox:1.38.0-glibc

COPY app /bin/app

RUN chmod +x /bin/app

CMD ["/bin/app"]
