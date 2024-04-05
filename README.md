## Creating docker for psql database
docker run --name got -e POSTGRES_PASSWORD=primayudha -e POSTGRES_USER=valencius -e POSTGRES_DB=valencius -p 6543:5432 -d postgres