# helpet-api-rest

##  Docker Instructions
Download MariaDb image

```
docker pull mariadb
```

Run Mariadb container with environment variables

```
docker run --name [nameContainer] -e MARIADB_USER=[user] -e MARIADB_PASSWORD=[password] -e MARIADB_ROOT_PASSWORD=[passwordRoot] -e MARIADB_DATABASE=[nameDB] -p 3306:3306 -d mariadb
```
Access the container

```
docker exec -it [nameContainer] bash
```

Access the Mariadb in Container

```
mariadb -u [user] -p
psql -U [user] --password --db [nameDB]
```
