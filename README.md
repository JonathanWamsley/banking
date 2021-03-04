# A rest based microservice for banking

This project is based off of [Ashish Juyal](https://github.com/ashishjuyal) [REST based banking microservices application](https://github.com/ashishjuyal/banking). The projects [Software Development Life Cycle](SDLC.md).

The purpose of this repo, is to gain experience with:
- creating a rest based microservice apis
- continuing the project structure and updating
    - essential routing endpoints to make the service useable
    - password security
    - documentation and tests
- creating a [Software Development Life Cycle](SDLC.md)


### The banking databases
Only the Users db is stored in banking_auth repo, while all the other are stored in this banking repo.
- Users 
    - for logging in and getting a jwt token
- Transactions 
    - for processing withdrawal and deposit transactions
- Customers
    - the owners of the accounts
- Accounts
    - where money is stored in a checking or savings account
