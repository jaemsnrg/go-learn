
# Create and run postgres container
docker run --name postgres-24 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# Execute psql command line tool inside the container
# -it: Interactive terminal
# postgres-24: Container name
# psql: PostgreSQL interactive terminal
# -U root: Connect as user 'root'

# Enter psql
docker exec -it postgres-24 psql -U root

# Access shell
docker exec -it postgres-24 /bin/sh

# exit container
exit

# stop container
docker stop postgres-24

# logs
docker logs -f postgres-24

# start container
docker start postgres-24

# exec and create db (used in makefile)
docker exec -it postgres-24 createdb --username=root --owner=root simple_bank
docker exec -it postgres-24 psql -U root simple_bank