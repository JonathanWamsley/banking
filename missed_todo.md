# TODO list 

contains ToDo's of things that could not be implemented right away, forgotten, mistyped:
I do not want to forget them, but want to maintain focus on feature, so this is the missed_todo's dump

- [ ] to get a single customer, the endpoint /customer should be /customers/{customer_id:[0-9]+}
- [ ] customer_handler_test.go should be customer_handlers_test.
- [ ] writeResponse(w, http.StatusBadRequest, "invalid json") -> writeResponse(w, http.StatusBadRequest, err.Error())
- [ ] CreateAccount should validate that an account of the same type does not exist. Need to first implement GetAccount

### other notes

In real life, if a customer deletes their account, they will no longer have access to that account through the business. This means to close the account, their should be no balance, outstanding fees.

How do I deal with an api endpoint that needs to contact multiple controllers. I think one solution would be to create a new service that includes both endpoints. This means I would need to create a new handler to take that duel service. Maybe I could design a way to include a generic service, but that is too complicated for now.
