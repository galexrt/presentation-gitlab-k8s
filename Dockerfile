FROM busybox:1.37.0-glibc

COPY app /bin/app

RUN chmod +x /bin/app

CMD ["/bin/app"]
