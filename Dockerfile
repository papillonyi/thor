FROM busybox

ADD ./thor /bin/thor
RUN mkdir /conf
ADD ./conf/app.ini /conf/app.ini
RUN chmod +x /bin/thor

EXPOSE 8080

ENTRYPOINT ["thor"]