#!/bin/sh
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

  \connect $POSTGRES_DB $POSTGRES_USER
  BEGIN;
  
    create table shorts (
        id serial primary key not null,
        link text not null,
        short text not null
    );
    
  COMMIT;
EOSQL