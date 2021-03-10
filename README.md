# A rest based microservice for banking

This project is based off of [Ashish Juyal](https://github.com/ashishjuyal) [REST based banking microservices application](https://github.com/ashishjuyal/banking). The projects [Software Development Life Cycle](SDLC.md).

The purpose of this repo, is to gain experience with:
- creating a rest based microservice apis
- continuing the project structure and updating
    - essential routing endpoints to make the service useable
    - configuration file for env variables
    - documentation and tests
- creating a [Software Development Life Cycle](SDLC.md)

### The banking databases

This banking service has 4 different databases. Only the Users db is stored in [banking_auth repo](https://github.com/JonathanWamsley/banking_auth), while all the other are stored in this banking repo.
- Users 
    - stores login credentials
    - used for logging in users and generating a jwt token
- Customers
    - stores locality information
    - used to associate a customer id to different accounts
- Accounts
    - stores banking amount in either a checking or savings account
    - used to get account details and make transactions
- Transactions 
    - stores withdrawal and deposit transactions
    - used to update account balances

### API table
| Method | Route                                         | Name            | Action                                     | Access Level |
|--------|-----------------------------------------------|-----------------|--------------------------------------------|--------------|
| GET    | /customers                                    | GetAllCustomers | returns all customers                      | admin        |
| POST   | /customers                                    | CreateCustomers | creates a new customer                     | admin        |
| GET    | /customers/{customer_id}                      | GetCustomer     | returns a customer by id                   | user / admin |
| DELETE | /customers/{customer_id}                      | DeleteCustomer  | deletes a custmer by id                    | admin        |
| GET    | /customers/{customer_id}/account              | GetAccount      | returns customer's accounts                | user / admin |
| POST   | /customers/{customer_id}/account              | CreateAccount   | creates a new account                      | admin        |
| DELETE | /customers/{customer_id}/account              | DeleteAccount   | deletes an account type                    | admin        |
| POST   | /customers/{customer_id}/account/{account_id} | MakeTransaction | creates a new transaction, updates account | user / admin |
| GET    | /users                                        | GetUsers        | returns all users                          | N/A          |
| POST   | /users                                        | CreateUser      | creates a user                             | N/A          |
| POST   | /admins                                       | CreateAdmin     | creates a admin                            | N/A          |
| POST   | /auth/login                                   | Login           | create and returns a jwt token             | N/A          |
| GET    | /auth/verify                                  | Verify          | returns a bool if jwt is valid             | N/A          |
### Example usage

Both the banking and the banking_auth service need to be running at the same time. I use port 8080 for the banking and 8181 for the banking_auth service.

#### First, create a user/admin account
I do admin since they have full access

- Request: Create admin account
    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"username": "admin1002", "password": "password1002"}' http://localhost:8181/admins
    ```

- Response: Created admin response
    ```yml
        {"Username":"admin1002","role":"admin","created_on":"2006-01-02 15:04:05"}
    ```
<hr>

#### Next, generate a token by logging in
- Request: Login as admin
    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"username": "admin1002", "password": "password1002"}' http://localhost:8181/auth/login
    ```

- Response: Access Token
    ```sh
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUzODM5MjgsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiIn0.SbE9lAFR7jrkyJK-U1kbs26NL-5IosLeaqHsEvSalLQ
    ```
<hr>

#### Next, get the customer accounts

- Request: Get customers using admin role
    ```sh
    curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUzODM5MjgsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiIn0.SbE9lAFR7jrkyJK-U1kbs26NL-5IosLeaqHsEvSalLQ"  http://localhost:8080/customers
    ```
- Response: Limiting to first 2 responses
    ```yml
    [
        {
            "customer_id":"2000",
            "full_name":"Steve",
            "city":"Delhi",
            "zipcode":"110075",
            "date_of_birth":"1978-12-15",
            "status":"active"
        }
    ,{"customer_id":"2001","full_name":"Arian","city":"Newburgh, NY","zipcode":"12550","date_of_birth":"1988-05-21","status":"active"}]
    ```
<hr>

#### Get Account information

- Request: Get account information for customer 2001
    ```sh
    curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUzODM5MjgsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiIn0.SbE9lAFR7jrkyJK-U1kbs26NL-5IosLeaqHsEvSalLQ"  http://localhost:8080/customers/2001/account
    ```
- Response: banking accounts for 2001
    ```yml
    [{"account_id":"95472","customer_id":"2001","opening_date":"2020-08-09 10:35:22","account_type":"saving","amount":7000},
    {"account_id":"95481","customer_id":"2001","opening_date":"2006-01-02 15:04:05","account_type":"checking","amount":5000}]
    ```

<hr>

#### Make a transaction

- Request: add 10,000 deposit to account 95472
    ```sh
    curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUzODM5MjgsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiIn0.SbE9lAFR7jrkyJK-U1kbs26NL-5IosLeaqHsEvSalLQ" -d '{"transaction_type":"deposit", "amount": 10000}' http://localhost:8080/customers/2001/account/95472
    ```

- Response: Get the transaction results
    ```yml
        {
            "transaction_id": "6",
            "account_id": "95472",
            "new_balance": 17000,
            "transaction_type": "deposit",
            "transaction_date": "2021-03-10 09:02:44"
        }
    ```
