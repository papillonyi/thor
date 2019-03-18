FROM busybox
ADD ./thor /bin/thor
Add ./conf ./
RUN chmod +x /bin/thor

EXPOSE 8080

ENTRYPOINT ["thor"]