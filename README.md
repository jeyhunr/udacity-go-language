# udacity-go-language

# Udacity Go Language (Golang)

## Getting Started

### Traditional Way

#### Prerequisites

Before you start, make sure you have Go installed on your machine.

#### Installation

1. Clone the repository.
   ```sh
   git clone https://github.com/jeyhunr/udacity-go-language.git
   ```
2. Navigate to the project directory.
   ```sh
   cd udacity-go-language
   ```
3. Build the application.
   ```sh
   go build
   ```

#### Usage

1. Run the executable.
   ```sh
   ./udacity-go-language
   ```
2. Open your web browser and go to [http://localhost:8080](http://localhost:8080) to view the application.

### Docker

```sh
docker-compose up -d
```

## Routes

| Route            | HTTP Request |                                               Body                                                | Description                    |
| :--------------- | :----------: | :-----------------------------------------------------------------------------------------------: | :----------------------------- |
| `/customers`     |    `GET`     |                                                N/A                                                | Get all customers              |
| `/customer`      |    `POST`    | `{ "id": "abc", "name": "abc", "role": "abc", "email": "abc", "phone": "abc", "contacted": true}` | Register a new customer        |
| `/customer`      |    `PUT`     | `{ "id": "abc", "name": "abc", "role": "abc", "email": "abc", "phone": "abc", "contacted": true}` | Confirm customer registration  |
| `/customer/{ID}` |    `GET`     |                                                N/A                                                | Get customer information by ID |
| `/customer/{ID}` |   `DELETE`   |                                                N/A                                                | Delete customer by ID          |

Ready to see the magic in action? 🚀 Take a tour with our interactive demo or dive into the full experience at https://udacity-go.rahimli.net
