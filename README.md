# Kaare's Webstore

> WIP

## Usage

Here are A FEW handy commands for using the API

Add product:

```
curl -X POST localhost:3000/products -d name="shirt"&price=10&sizes="[\"small\",\"medium\",\"large\",\"extra large\"]"&image_path=""
```

View products:

```
curl localhost:3000/products
```
