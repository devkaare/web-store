# Kaare's Webstore

> WIP

## Usage

### Env

```
PORT=3000
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=dbwebstore
DB_USERNAME=kaare
DB_PASSWORD=password
DB_SCHEMA=public
```

Here are A FEW handy commands for using the API

Add product:

```
curl -X POST localhost:3000/products -d name="shirt"&price=10&sizes="[\"small\",\"medium\",\"large\",\"extra large\"]"&image_path=""
```

View products:

```
curl localhost:3000/products
```
