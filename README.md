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
API_KEY=81566e986cf8cc685a05ac5b634af7f8
STRIPE_KEY=sk_test_4S68v29DeKcE4RxJcrJnUn5s
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
