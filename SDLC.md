
# Software Development Life Cycle (SDLC)

#### SDLC Index
1. [Planning](#1-header)
2. [Requirement Analysis]("#2-header)
3. [Design]("#3-header)
4. [Implementation/ Development]("#4-header)
5. [Testing]("#5-header)
6. [Deployment]("#6-header)
7. [Maintenance and improvement]("#7-header)
8. [Time Log(personal use)]("#8-header)
<hr></hr>

A software development life cycle helps produce high-quality software quickly by following a procedure targeting key stages.

<hr></hr>

<h3 id="1-header"> 1. Planning procedure</h3>

> To conduct research on the product you are planning to develop
- The product is a banking application API with endpoints for the customers and admins to make transaction on their accounts securely. Right now, several essential api calls are missing, like creating new users.

> identify the pros and cons of the current software method you are using
- Golang is being used to create this project. Golang offers high performance and is paired with microframeworks to build SOA and microservice applications. It does not come with as many features as Ruby on Rails or Laravel, which means  3rd party packages are only added when required. This results in a vary lean code base.


<h3 id="2-header"> 2. Requirement Analysis</h3>

> The Software Requirements Specification(SRS) describes what the software will do and how it will be expected to perform. The functionality and features needed will be described.

> Purpose

The purpose for this application is to create an internal REST api endpoints for a banking business. This api can be connected to a frontend application or other intermediary apis within the business. This api therefor inherits some of the expected user use cases as well as being easy to use for internal developers. The users of the api are Admins and customers with bank accounts that will require some expected functionalities.

> Project structure

This banking appliaction builds off of [Ashish Juyal](https://github.com/ashishjuyal) [REST based banking microservices application](https://github.com/ashishjuyal/banking). This repo uses golang to build an application using a hexagonal architecture pattern. The hexagonal architecture (also referred to Ports and Adapters), structures projects in such a way to separate business logic from other services, ports(ui, cli, logging) and adapters(database, email). It helps clean code organization, allows for faster testing and can easily integrate with external services. One of the main drawbacks is working in several folders to get a simple feature working and requiring lots of testing using mocks.

> expected functionalities

Bellow are features that the banking application currently has and features that will be built to provide the essential banking services. Here account refers to a banking account like saving or checking account. A customer refers to the owners (customer/ admin) personal information like location status. A user identifies access credentials to login.

- Customer needs
    - [ ] Create a customer
    - [ ] Update their customer info
    - [ ] Delete their customer info
    - [ ] Create a user
    - [ ] Update a user
    - [ ] Delete a user
    - [x] Login as a user
    - [x] Get balance from their account
    - [x] Create a new transaction for their account
- Admins needs
    - [ ] Update a customer
    - [ ] Delete a customer
    - [x] Create an account
    - [ ] Delete an account
    - [x] Login as an admin
    - [x] Get a customer
    - [x] Get customers
    - [x] Get balance from any account
    - [x] Create new transactions

		"admin": {"GetCustomer", "GetCustomers", "CreateCustomer", "CreateCustomer", "DeleteCustomer", "CreateAccount", "GetAccount", "DeleteAccount", "NewTransaction"},

    Update: Some of the fields the customer should have to go through an admin to make certain calls. The new permissions have been changed to:

- Customer needs
    - [x] Login as a user
    - [x] Get info from their customer account
    - [x] Get balance from their banking account
    - [x] Create a new transaction for their account
- Admins needs
    - [x] Login as an admin
    - [x] Get a customer
    - [x] Get customers
    - [x] Create a customer
    - [x] Delete a customer
    - [x] Create an account
    - [~] Delete an account
    - [x] Get balance from any account
    - [x] Create new transactions

> security requirements

- [x] The banking application uses a banking authorization api to create jwt. A 6 step process is used to allow a customer or admin to access the banking api. All of this has already been implemented.

1. A customer/admin request an jwt by logging in
2. The banking auth api response with a jwt if given valid credentials
3. The customer/admin request to the banking api using jwt
4. The banking api sends this jwt to be verified to the banking auth api
5. The banking auth api verifies the jwt and authorizes the customer/admin
6. The banking api processes the users request and responds

- [ ] The users login credentials are not secure. A bcrypt hashing algorithm will be used replace raw passwords and help secure sensitive information.

> software requirements

mysql  Ver 8.0.23-0ubuntu0.20.04.1 for Linux on x86_64 
go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.1
	github.com/golang/mock v1.5.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
)

> testing requirements

- [ ] The majority of the code base is not tested. Many unit tests and mocks will need to be created.

testing will give the owners and developers interacting with the api some assurance that the code is working as expected. This will allow for more reliable code and will force decoupling of any messy code.

> Documentation

- [ ] All public types should have a comment about what it is and its purpose.


> Validation

- [ ] A customer should be limited to only one checking and saving account.
- [ ] A password should be a minimum length of 8 characters with a capital


<h3 id="3-header"> 3. Design</h3>

Design Document Specification

> User Interface
- N/a

> Functionality

The project will build on top of existing code and the functionality changes are outlined in 'expected functionalities'

> Milestones

- create git profile with SDLC
- create the customers api
    - add Create, Update, Delete api endpoints
    - add testing and documentation
- create the account api
    - add delete account
    - add limited account verification
    - add testing and documentation
- create the transaction api
    - add testing
- create the auth banking api
    - add create, update, delete user login account
    - add testing and documentation
    - add password verification
- finishing up lose ends and explore deployment options

> Time frame

The banking app course was already completed prior. Much of the code will remain the same with added features discussed in the 'expected functionalities' section.

- git interactions (2 hours)
    - creating issues
    - submitting a pr using a feature branch
    - approval/ update until approval
    - merge into master and delete the feature branch

- banking tutorial already completed prior
- reviewing course content (4 hours)
- outside documentation (6 hours)
- Coding (8 hours)
- code documentation (1 hour)
- testing (3 hours)

- total approximation: 24 hours

> budget

- some potential options like using digital ocean or google app engine to host application will be explored

<h3 id="4-header"> 4. Implementation / Development</h3>

This banking project has already been started using golang. A few api endpoints, testing, validation and documentation will be added. The code will be be redone by typing and submitting through the 'git interaction'.


<h3 id="5-header"> 5. Testing</h3>

The majority of the project should be tested using unit test and mocks. Any specific names should be refactored into constants and tested so they do not accidentally get changed.


<h3 id="6-header"> 6. Deployment</h3>

Researching deployment options like digital ocean or google app engine to host application will be explored
- dockerizing container
- kubernetes orchestration

Creating a guide how to build and deploy application for new users will be documented

<h3 id="7-header"> 7. Maintenance and Improvement</h3>

Future improvements that can be made will be mentioned, mainly focusing on deployment and api end points.


<h3 id="8-header"> 8. Time Log</h3>

- 18 hours going through Banking tutorial
- 4 hours reading about SDLC and documenting

