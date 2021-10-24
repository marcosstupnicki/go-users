# go-users

go-users is an api that implements the [go-webapp](https://github.com/marcosstupnicki/go-webapp) library. This is a basic CRUD example for handling users.

## Running

This example uses a mysql database. To generate this db, this api has a Docker recipe (docker-compose.yml), you can execute it:

```bash
$ docker-compose up
```

After creating the container, we need to migrate the user table:
```bash
$ go run cmd/tools/migrate/main.go
```

Finally, we run the app:
```bash
$ go run cmd/api/main.go
```

You can validate the operation of the application by pinging the app:
```
curl --location --request GET 'http://localhost:8080/ping'
```
Response:
```
pong
```

## Operations

### Create User
Request:
```
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
"email": "some@email.com",
"password": "12312312asdasdas"
}'
```

Response (status_code: 201):
```json
{
    "id": 7,
    "email": "some@email.com",
    "created_at": "2021-10-23T15:45:41.135-03:00",
    "updated_at": "2021-10-23T15:45:41.135-03:00"
}
```

### Get User

Request:
```
curl --location --request GET 'http://localhost:8080/users/7'
```

Response (status_code: 200):
```json
{
    "id": 7,
    "email": "otro@email.com",
    "created_at": "2021-10-23T15:45:41.135-03:00",
    "updated_at": "2021-10-23T15:46:10.847-03:00"
}
```

### Update User

Request:
```
curl --location --request PUT 'http://localhost:8080/users/7' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "otro@email.com",
    "password": "12312312asdasdas"
}'
```

Response (status_code: 200):
```json
{
    "id": 7,
    "email": "otro@email.com",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "2021-10-23T15:46:10.847-03:00"
}
```

### Delete User

Request:
```
curl --location --request DELETE 'http://localhost:8080/users/7'
```

Response (status_code: 204):
```json
```
