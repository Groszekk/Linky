FROM postgres

ENV POSTGRES_DB=linky
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

COPY setup.sh /docker-entrypoint-initdb.d/setup.sh
WORKDIR /