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
