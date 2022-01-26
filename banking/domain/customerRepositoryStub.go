package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "112", Name: "John", City: "New York", Zipcode: "10001", DateofBirth: "01/01/1990", Status: "Active"},
		{Id: "113", Name: "Mary", City: "New York", Zipcode: "10001", DateofBirth: "01/01/1990", Status: "Active"},
		{Id: "114", Name: "Bob", City: "New York", Zipcode: "10001", DateofBirth: "01/01/1990", Status: "Active"},
	}
	return CustomerRepositoryStub{customers: customers}
}
