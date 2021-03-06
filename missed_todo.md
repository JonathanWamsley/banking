# TODO list 

contains ToDo's of things that could not be implemented right away, forgotten, mistyped:
I do not want to forget them, but want to maintain focus on feature, so this is the missed_todo's dump

- [ ] to get a single customer, the endpoint /customer should be /customers/{customer_id:[0-9]+}
- [ ] customer_handler_test.go should be customer_handlers_test.
- [ ] writeResponse(w, http.StatusBadRequest, "invalid json") -> writeResponse(w, http.StatusBadRequest, err.Error())
- [ ] CreateAccount should validate that an account of the same type does not exist. Need to first implement GetAccount