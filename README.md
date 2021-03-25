# Launch
```sh
go run main.go
```

port 8000

# Usage example

CREATE:
```sh
curl --location --request POST 'localhost:8000/autos' \
--header 'Content-Type: text/plain' \
--data-raw '    {
        "brand": "Lada",
        "model": "Priora",
        "price": 600000,
        "status": 1,
        "mileage": 2220
    }
```

READ by id:
```sh
curl --location --request GET 'localhost:8000/autos/2'
```

READ all:
```sh
curl --location --request GET 'localhost:8000/autos'
```

UPDATE by id:
```sh
curl --location --request PUT 'localhost:8000/autos/2' \
--header 'Content-Type: text/plain' \
--data-raw '    {
        "id": 2,
        "brand": "Mazda",
        "model": "RX-8",
        "price": 1740000,
        "status": 2,
        "mileage": 15000
    }'
```

DELETE by id:
```sh
curl --location --request DELETE 'localhost:8000/autos/2'
```




