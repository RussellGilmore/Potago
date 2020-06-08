# Potago
A basic go RESTful api designed to put together a multiple go practices that I have been learning. Seeds a Postgres database with a few entries for test purposes and contains a postman request collection for ease of use.

# Prerequisites
### Tools
* Docker
* Git
### Setup Database
```sh
docker run --rm   --name pg-docker -e \
POSTGRES_PASSWORD=docker -d -p 5432:5432 \
-v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
```
### Deploy
To start the api server and observe the following to ensure it is working. 
```sh
~ go run main.go
We are getting the env values
We are connected to the postgres database
(/Users/gilmorera/Documents/playground/potago/api/seed/seeder.go:23)
[2020-02-11 16:08:45]  [68.85ms]  DROP TABLE "potatos"
[0 rows affected or returned ]

(/Users/gilmorera/Documents/playground/potago/api/seed/seeder.go:27)
[2020-02-11 16:08:45]  [83.33ms]  CREATE TABLE "potatos" ("id" bigserial,"name" varchar(255) NOT NULL UNIQUE,"description" varchar(255) NOT NULL,"created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,"updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP
 , PRIMARY KEY ("id"))
[0 rows affected or returned ]

(/Users/gilmorera/Documents/playground/potago/api/seed/seeder.go:33)
[2020-02-11 16:08:45]  [12.06ms]  INSERT INTO "potatos" ("name","description","created_at","updated_at") VALUES ('Red Potato','A good red potato.','2020-02-11 16:08:45','2020-02-11 16:08:45') RETURNING "potatos"."id"
[1 rows affected or returned ]

(/Users/gilmorera/Documents/playground/potago/api/seed/seeder.go:33)
[2020-02-11 16:08:45]  [3.78ms]  INSERT INTO "potatos" ("name","description","created_at","updated_at") VALUES ('Yellow Potato','A good yellow potato.','2020-02-11 16:08:45','2020-02-11 16:08:45') RETURNING "potatos"."id"
[1 rows affected or returned ]
Listening to port 8080
```
# Curl Commands
### Get All Potato
```sh
curl --location --request GET 'localhost:8080/potato'
```
### Get Potato By ID
```sh
curl --location --request GET 'localhost:8080/potato/1'
```
### Post Potato
```sh
curl --location --request POST 'localhost:8080/potato' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Orange Potato",
    "description": "A good orange potato."
}'
```
### Update Potato
```sh
curl --location --request PUT 'localhost:8080/potato/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Updated Potato",
    "description": "A good updated potato."
}'
```
### Delete Potato
```sh
curl --location --request DELETE 'localhost:8080/potato/1'
```