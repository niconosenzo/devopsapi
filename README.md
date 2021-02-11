# go REST API

This piece of code has been created as a code exercise, it's not intended for production usage.


## Install & run the app

    make install
    ./api

## Build and Push the container to a given registry

    make build

## Run the unit tests

    make test

## Deploy app from image in a K8s cluster

    make deploy

# REST API

API endpoints.

## Get list of users (when deployed locally, otherwise replace localhost with the correct address)

### Request

`GET /users/`

    curl http://localhost:3000/users


## Create a new User

### Request

`POST /user/{id}`

     curl -XPOST http://localhost:3000/user -d '{"id": "5", "name": "Pp", "surname": "Mm"}'


## Get a specific User

### Request

`GET /user/{id}`

    curl http://localhost:3000/user/1


## Delete a User

### Request

`DELETE /user/{id}`

    curl -XDELETE  http://localhost:3000/user/1 
