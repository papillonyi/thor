FROM busybox

ADD ./app /bin/app
RUN mkdir /conf
ADD ./conf/app_docker.ini /conf/app.ini
RUN chmod +x /bin/app

ENTRYPOINT ["app"]