FROM frolvlad/alpine-glibc

RUN mkdir /app
COPY ShorterService /app
WORKDIR /app
ENV CORS-HOST=http://reverseproxy
ENV DB-HOST=database
ENV DB-NAME=linky
ENV DB-USER=postgres
ENV DB-PASS=postgres
CMD ./ShorterService
