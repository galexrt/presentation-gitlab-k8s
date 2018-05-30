FROM busybox:1.28.4-glibc

COPY app /bin/app

RUN chmod +x /bin/app

CMD ["/bin/app"]
