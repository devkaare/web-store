# Kaare's Webstore

> WIP

## Usage

### Env

```
PORT=3000
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=database
DB_USERNAME=admin
DB_PASSWORD=password
DB_SCHEMA=public
API_KEY=
STRIPE_PUBLIC_KEY=
STRIPE_SECRET_KEY=
```

Here are A FEW handy commands for using the API

```
curl localhost:3000/products
```

```
curl -X POST http://localhost:3000/products/ \
                                                    -F "product_name=Test Name" \
                                                    -F "category_name=Test Category"
\
                                                    -F "description=This is a test product." \
                                                    -F "price=10" \
                                                    -F "image=@test.png"
```

```
curl localhost:3000/products/1
```

```
curl -X DELETE localhost:3000/products?product_id=1
```
