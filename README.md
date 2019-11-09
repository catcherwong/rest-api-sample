# rest-api-sample

A simple sample that use gin to build rest api.


## Installation

```
git clone https://github.com/catcherwong/rest-api-sample.git
```

## How to run

```
$ cd rest-api-sample

$ docker build -t rest-api-sample:v1 .

$ docker run -p 9999:9999 rest-api-sample:v1
```

Swagger doc


![](./media/swagger.png)

## Features

- Gin
- Gorm
- Swagger
- Redis
- SQLite
- yaml
- prometheus