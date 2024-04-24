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
