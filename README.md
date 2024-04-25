## Welcome Note

```
Everyone is welcome for code contributions or code enhancemnets.
please give a star, fork the repo and create PR for your code contributions
```
## db commands

```
create database userdb;
use userdb;
create table users(username varchar(255), password varchar(255));
```
## run the application with air live reloading

```
air
```
## run the application with go run

```
go run main.go
```
## signup request

```
curl --location 'localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data '{
    "username": "user1",
    "password": "password1"
    
}'
```
## signin request

```
curl --location 'localhost:8080/signin' \
--header 'Content-Type: application/json' \
--data '{
    "username": "user1",
    "password": "password1"
    
}'
```
