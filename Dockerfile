FROM busybox
ADD ./thor /bin/thor
RUN chmod +x /bin/thor

EXPOSE 8080

ENTRYPOINT ["thor"]