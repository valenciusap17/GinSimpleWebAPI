## Creating docker for psql database
```docker run --name {containername} -e POSTGRES_PASSWORD={password} -e POSTGRES_USER={host_user} -e POSTGRES_DB={host_db} -p {port}:5432 -d postgres```

notes: {} means place it with the yours
