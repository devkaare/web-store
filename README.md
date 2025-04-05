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

Add product:

```
curl -X POST localhost:3000/products\?api_key=81566e986cf8cc685a05ac5b634af7f8 -d 'name=shirt&price=10&sizes=["small","medium","large","extra large"]&image_path=../assets/product-imgs/placeholder1.png'
```

View products:

```
curl localhost:3000/products
```

```
curl "localhost:3000/products/?page=1"
```

```
curl -X POST http://localhost:3000/products/ \
                                                    -F "product_name=test" \
                                                    -F "category_name=testCategory" \

                                                    -F "description=this is a test product." \
                                                    -F "price=99" \
                                                    -F "image=@test.png"

```
