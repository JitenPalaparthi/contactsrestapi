# How to run the code
- This project has been designed to use multiple databases like mysql, postgres and also filedb(file based) implementations.
- To use mysql or postgres make sure that mysql and | or postgres are running in your machine or as a docker containers.

### windows or any machine

- install docker. Make sure that docker is up and run. Check docker.com documentation for instalation.
- run the below command to start postgres docker container
- ```docker run -d -p 5432:5432 --name mypostgres -e POSTGRES_PASSWORD=jiten123 -e POSTGRES_USER=jiten -e POSTGRES_DB=contactsdb postgres```
- run the below command to start mysql docker container
- ```docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=jiten -e MYSQL_PASSWORD=jite123 --name mysqlcontactdb mysql```

### How to run the project

- after cloning the project to run project on postgres db
- ```go run main.go -stderrthreshold=INFO --mode=postgres --port=50001 --dsn="host=localhost user=jiten password=jiten123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"```
- to run project on mysql
- ```go run main.go -stderrthreshold=INFO --mode=mysql --port=50001 --dsn="jiten:jite123@tcp(127.0.0.1:3306)/contactsdb?charset=utf8mb4&parseTime=True&loc=Local"```
- to run the project as filedb
- ```go run main.go -stderrthreshold=INFO --mode=filedb --port=50001```

### To test use curl or postman.
- to insert data
```curl --location --request POST 'localhost:50090/v1/person' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Jiten",
    "address": "Bangalore",
    "email": "Jitenp@outlook.com",
    "mobile": "9618558500"
}'```

- to delete data
```curl --location --request DELETE 'localhost:50090/v1/person/10``` 10 is the id. Based on inserted id it should be given