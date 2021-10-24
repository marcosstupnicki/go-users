# go-users


## Api Operations

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
