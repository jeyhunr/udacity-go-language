# udacity-go-language

Udacity Go Language (Golang)

| Route            | HTTP Request |                                               Body                                                | Description                    |
| :--------------- | :----------: | :-----------------------------------------------------------------------------------------------: | :----------------------------- |
| `/customers`     |    `GET`     |                                                N/A                                                | Get all customers              |
| `/customer`      |    `POST`    | `{ "id": "abc", "name": "abc", "role": "abc", "email": "abc", "phone": "abc", "contacted": true}` | Register a new customer        |
| `/customer`      |    `PUT`     | `{ "id": "abc", "name": "abc", "role": "abc", "email": "abc", "phone": "abc", "contacted": true}` | Confirm customer registration  |
| `/customer/{ID}` |    `GET`     |                                                N/A                                                | Get customer information by ID |
| `/customer/{ID}` |   `DELETE`   |                                                N/A                                                | Delete customer by ID          |
