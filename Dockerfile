FROM ubuntu

COPY gushkin /home/gushkin

RUN chmod +x /home/gushkin

WORKDIR /home

EXPOSE 8080

ENTRYPOINT ./gushkin
