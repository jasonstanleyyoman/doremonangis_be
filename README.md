# Backend repository for Doremonangis

## Author

Jason Stanley Yoman (NIM : 13519019)

## How to run

Requirement:

- `docker`

Execute this command after requirement is installed properly

```sh
docker-compuse up -d
```

### API Endpoint List

- `GET /v1/store/all` (Get all store available)
- `POST /v1/store/new` (Create new store)

Request Body:

name: string

address: string

kecamatan: string

province: string

- `POST /v1/store/dorayaki/add` (Add dorayaki stock to a store)

Request Body:

store_id: uint

dorayaki_id: uint

amount: int

- `POST /v1/store/dorayaki/remove` (Remove dorayaki stock to a store)

Request Body:

Same as add stock to dorayaki above

- `POST /v1/store/dorayaki/move` (Move dorayaki to other store)

Request Body:

src: uint

dest: uint

dorayaki_id: uint

- `DELETE /v1/store/:id` (Remove a store based on id)

- `GET /v1/store/:id` (Get store by id)

- `GET /v1/dorayaki/all` (Get all dorayaki available)

- `POST /v1/dorayaki/new` (Add new dorayaki)

Request Body;

flavor: string

description: string

image_path: string

- `DELETE /v1/dorayaki/:id` (Delete dorayaki by id)

- `GET /v1/dorayaki/:id` (Get dorayaki by id)
