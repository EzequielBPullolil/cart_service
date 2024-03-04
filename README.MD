# cart_service

`cart_service` is a shopping cart service developed in Golang

## Description

The `cart_service` aims to provide the client with the abstraction of shopping cart logic by offering a well-defined and easy-to-integrate REST API into any system. This system DOES NOT HAVE THE PRODUCT ENTITY defined, meaning the cart can be used for anything; it just needs to receive the UUID of the product or a reference, and optionally, additional data that can be provided using a string.

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


## Endpoint examples

### Create Cart
To create a cart, you need to send the user data along with the currency that the cart accepts.

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

---
### Add item to cart
```bash
curl --location '<api_url>/carts/<cart_id>/items' \
--header 'Content-Type: application/json' \
--data '{
    "id": <item_id>,
    "name": <item_name>,
    "price": <item_price>,
    "currency": <item_currency>,
    "quantity": <item_cuantity>
}'
```
To add an item to the cart, you should make a request to the `carts/:cart_id/item` endpoint, sending the [Item fields](#Item-fields)


None of the fields can be empty. In addition, the quantity and price fields must be greater than zero. The price can be a float value (e.g., 5.10).

---
### Modify item from cart
To modify a product in a cart, you should send the fields to be updated in the body to the endpoint with the cart_id and item_id.
```bash
curl --location --request PATCH '<api_url>/carts/<cart_id>/items/<item_id>' \
--header 'Content-Type: application/json' \
--data '{
    "price": 900.5
}'
```

### Delete item from cart

To delete a product from a cart, you should send a request to the endpoint passing the `cart_id and the `id` of the item to be deleted.
```bash
curl --location --request DELETE '<api_url>/carts/<cart_id>/items/<item_id>'
```


### Get Cart 
To retrieve the cart, you should make a GET request to the `carts/<cart_id>` endpoint. This will return the cart data and all items associated with the cart.
```bash
curl --location --request GET '<api_url>/carts/<cart_id>/'
```


## Item fields 
The fields of the Item must be as follows and comply with the following rules:

| Field | Type | Rule |
|-------|------|------|
| id | String | It is a non-empty string |
| name | String | It is a non-empty string |
| price| float | is greater than zero |
| currency| string | It is the same currency as the car |
| quantity| int | is greater than zero |
