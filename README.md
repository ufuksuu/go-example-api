# Example Go RESTful API

## Run application

```
make dev
```

## Features

* API Framework: [Gin](https://gin-gonic.com/)
* Object Relational Mapping (ORM): [Gorm](https://gorm.io/index.html)
* Routing: [Gin](https://gin-gonic.com/)
* Package manager: go mod
* Base CRUD service
* Base CRUD controller

## Folder structure

* ```controllers/```: Controllers for web requests processing.
* ```models/```: Business logic.
* ```models/db```: DB connection.
* ```models/entity```: GORM entities.
* ```models/service```: Business logic.
* ```server```: Web requests routes and server settings.