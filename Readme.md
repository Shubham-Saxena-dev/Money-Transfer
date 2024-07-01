**This is a HTTP service that exposes an endpoint "/transfer". This endpoint receives "POST" request
and return a status string as a response**

To start the application, run the following command in your terminal:

```bash 
make up
```

You can also start the application by directly running the following command in your terminal:

```bash
docker-compose up 
```
This will build the application and start the server. The server will be available at http://localhost:8080

Requests can be made to the server using the following curl command:

```copy
curl --location 'http://localhost:8080/transfer' \
--header 'Content-Type: application/json' \
--data '{
  "organization_name": "My corp",
  "organization_bic": "BIC1234",
  "organization_iban": "ABCD1234XX5678",
  "credit_transfers": [
    {
      "amount": "14.5",
      "counterparty_name": "Bip Bip",
      "counterparty_bic": "CRLYFRPPTOU",
      "counterparty_iban": "EE383680981021245685",
      "description": "Wonderland/4410"
    }
  ]
}
'
```

You can also check the api doc after running the swagger on :

```copy
http://localhost:8080/swagger/index.html
```


To stop the application, run the following command in your terminal:

```bash
make stop
```
