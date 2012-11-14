package main

import ()

type SGApi struct {
	handlers []*SGOp
}

func (api *SGApi) Get(path string) *SGOp {
	return &SGOp{api: api, method: "GET"}
}
func (api *SGApi) Post(path string) *SGOp {
	return &SGOp{api: api, method: "POST"}
}
func (api *SGApi) Put(path string) *SGOp {
	return &SGOp{api: api, method: "PUT"}
}
func (api *SGApi) Delete(path string) *SGOp {
	return &SGOp{api: api, method: "DELETE"}
}

type SGHandler interface{}

type SGOp struct {
	api         *SGApi
	method      string
	handler     SGHandler
	description string
	params      map[string]string
	returns     string
	accepts     string
}

func (op *SGOp) DoOn(obj interface{}, funcName string) error {
	op.handler = func() {
		// call funcName on obj with passed in args
	}
	return nil
}
func (op *SGOp) Do(handler SGHandler) error {
	op.handler = handler
	return nil
}
func (op *SGOp) Description(description string) *SGOp {
	op.description = description
	return op
}
func (op *SGOp) Param(name string, description string) *SGOp {
	if op.params == nil {
		op.params = make(map[string]string)
	}
	op.params[name] = description
	return op
}
func (op *SGOp) Returns(typ string) *SGOp {
	op.returns = typ
	return op
}
func (op *SGOp) Accepts(typ string) *SGOp {
	op.accepts = typ
	return op
}

type Person struct {
}

type PersonResources struct{}

type SGApiResource interface {
	GetSwaggoApi() (*SGApi, error)
}

func (pr *PersonResources) GetMe(abc string, def string, resp interface{}) error {
	//resp.Println("Hello, world!")
	return nil
}

func (pr *PersonResources) GetSwaggoApi() (*SGApi, error) {

	api := new(SGApi)

	api.Get("/person/me").
		Description("A function that does stuff").
		Param("abc", "The value of awesomeness").
		Param("def", "Slightly more awesomeness").
		Returns("Person").
		DoOn(pr, "GetMe")

	api.Post("/person/me").
		Description("Create a person").
		Param("param1", "its description").
		Accepts("Person").
		Returns("Person").
		Do(func(param1 string, resp interface{}) {

	})

	return api, nil
}

func main() {

	pr := &PersonResources{}

	pr.GetSwaggoApi()

}
