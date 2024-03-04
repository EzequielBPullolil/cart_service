# cart_service

`cart_service` is a shopping cart service developed in Golang

## Description

The service aims to offer the user the abstraction of a simple shopping cart using the RESTful API architecture.


## Run service
1. First, you need to download the repository.
    ```bash
    git clone https://github.com/EzequielBPullolil/cart_service && cd cart_service/
    ```
2. With the MongoDB service started, we configure the environment variables DB_URI and DB_NAME so that the service connects to the desired database.
    ```bash
    export DB_URI=<our_mongodb_uri> && export DB_NAME=<our_db_name>
    ```
3. Finally, we compile the binary and run it, passing the port and host of the service as flags.
    ```bash
    go build -o ./cart_service ./src/cmd 
    ./cart_service -host <our_host> -port <our_port>
    ```

## Api Endpoints

The service provides the following endpoints, following the REST-API architecture.

| HTTP Verbs | Endpoints                   | Action                                          |
| :--------: | --------------------------- | ----------------------------------------------- |
|  **POST**  | /carts/                     | To create a cart                                |
|  **GET**   | /carts/:id                  | To retrieve the cart data associated with an ID |
| **DELETE** | /carts/:id                  | To delete the cart with the passed ID           |
|  **POST**  | /carts/:id/items            | To add content to the cart                      |
| **DELETE** | /carts/:id/items/:item_uuid | To delete a specific item from the cart         |
| **PATCH**  | /carts/:id/items/:item_uuid | To modify the quantity of a specific item       |


## Create Cart
To create a cart, you need to send the user data along with the currency that the cart accepts.
### request
```bash
curl --location '<api_url>/carts' \
--header 'Content-Type: application/json' \
--data '{
    "user": {
        "id": <user_id>,
        <more_data>
    },
    "currency": <currency_accepted>
}'
```
The user field must have at least the ID field. Whenever the cart data is requested, the user data will be returned as well.

### Response
    {
        
        "cart": {
            "id": <cart_uuid>,
            "amount": 0,
            "currency": <currency>,
            "items": [],
            "user": {
                "id": "<user_id>"
                <more_data>
            }
        },
        "status": "cart created"
    }

## Add item to cart

To add an item to the cart, you should make a request to the `carts/:cart_id/item` endpoint, sending the [Item fields](#Item-fields)


None of the fields can be empty. In addition, the quantity and price fields must be greater than zero. The price can be a float value (e.g., 5.10).
### Request
```bash
curl --location '<api_url>/carts/<cart_id>/items' \
--header 'Content-Type: application/json' \
--data '{
    "id": "dotella",
    "name": "boorsdd",
    "price": 500,
    "currency": "ARS",
    "quantity": 1
}'
```

### Response
    {
        "data": {
            "id": <cart:id>,
            "amount": 0,
            "currency": "ARS",
            "items": [
                {
                    "id": "dotella",
                    "name": "boorsdd",
                    "price": 500,
                    "currency": "ARS",
                    "quantity": 1
                }
            ],
            "user": {
                "id": "<user_id>"
                <more_data>
            }
        },
        "status": "item added to cart"
    }


## Modify item from cart
To modify a product in a cart, you should send the fields to be updated in the body to the endpoint with the cart_id and item_id.
### Request
```bash
curl --location --request PATCH '<api_url>/carts/<cart_id>/items/<item_id>' \
--header 'Content-Type: application/json' \
--data '{
    "price": 900.5
}'
```
### Response 
    {
        "cart": {
            "id": <cart_id>,
            "amount": 0,
            "currency": "ARS",
            "items": [
                {
                    "id": <item_id>,
                    "name": <item_name>,
                    "price": 900.5,
                    "currency": "ARS",
                    "quantity": 1
                },
            ],
            "user": {}
        },
        "status": "cart item updated"
    }


Any of the fields can be modified except the ID
## Delete item from cart

To delete a product from a cart, you should send a request to the endpoint passing the `cart_id and the `id` of the item to be deleted.
### Request
```bash
curl --location --request DELETE '<api_url>/carts/<cart_id>/items/<item_id>'
```

### Response
    {
        
        "cart": {
            "id": <cart_uuid>,
            "amount": 0,
            "currency": <currency>,
            "items": [],
            "user": {
                "id": "<user_id>"
                <more_data>
            }
        },
        "status": "cart item deleted"
    }

## Get Cart 
To retrieve the cart, you should make a GET request to the `carts/<cart_id>` endpoint. This will return the cart data and all items associated with the cart.

### Request
```bash
curl --location --request GET '<api_url>/carts/<cart_id>/'
```


### Response
    {
    "cart": {
        "id": "86beabc5-caf9-4694-ac6f-a8d0c09709b3",
        "amount": 0,
        "currency": "ARS",
        "items": [],
        "user": []
    }
    }

## Item fields 
The fields of the Item must be as follows and comply with the following rules:

| Field | Type | Rule |
|-------|------|------|
| id | String | It is a non-empty string |
| name | String | It is a non-empty string |
| price| float | is greater than zero |
| currency| string | It is the same currency as the car |
| quantity| int | is greater than zero |


## Todo

- [x] MVP
- [ ] Possibility of using several Currencies in the same cart
- [ ] Add query options to endpoints
- [ ] Add app tokens to restrict API usage
- [ ] 