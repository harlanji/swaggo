package main

import ()

type Swaggo struct {
	ApiVersion     string
	SwaggerVersion string
	BasePath       string
	ResourcePath   string

	Apis   []*SGApi
	Models []*SGModel
}

func NewSwaggo(resourcePath string, apiVersion string) *Swaggo {
	return &Swaggo{
		ApiVersion:     apiVersion,
		SwaggerVersion: "1.1",
		BasePath:       "http://localhost:9000/api", // FIXME calculate from http muxer
		ResourcePath:   resourcePath,
		Apis:           make([]*SGApi, 0),
		Models:         make([]*SGModel, 0),
	}
}

func (sg *Swaggo) Api(description string) (*SGApi, error) {
	api := &SGApi{
		Description: description,
		Operations:  make([]*SGOp, 4),
	}

	sg.Apis = append(sg.Apis, api)

	return api, nil
}

type SGModel struct {
}

type SGApi struct {
	Description string
	Operations  []*SGOp
}

func (api *SGApi) Get(path string) *SGOp {
	return api.newOp("GET", path)
}
func (api *SGApi) Post(path string) *SGOp {
	return api.newOp("POST", path)
}
func (api *SGApi) Put(path string) *SGOp {
	return api.newOp("PUT", path)
}
func (api *SGApi) Delete(path string) *SGOp {
	return api.newOp("DELETE", path)
}
func (api *SGApi) newOp(method string, path string) *SGOp {
	return &SGOp{
    api: api,
    method: method,
    params: make(map[string]string),
    errorResponses: make([]*SGErrorResp, 0),
  }
}

type SGHandler interface{}

type SGOp struct {
	api            *SGApi
	method         string
	handler        SGHandler
	summary        string
	notes          string
	params         map[string]string
	returns        string
	errorResponses []*SGErrorResp
}

type SGErrorResp struct{}

func (op *SGOp) MapOn(obj interface{}, funcName string) error {
	op.handler = func() {
		// call funcName on obj with passed in args
	}
	return nil
}
func (op *SGOp) Map(handler SGHandler) error {
	op.handler = handler
	return nil
}
func (op *SGOp) Summary(summary string) *SGOp {
	op.summary = summary
	return op
}
func (op *SGOp) Notes(notes string) *SGOp {
	op.notes = notes
	return op
}
func (op *SGOp) Param(name string, description string) *SGOp {
	op.params[name] = description
	return op
}
func (op *SGOp) ResponseClass(typ string) *SGOp {
	op.returns = typ
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

func (pr *PersonResources) SwaggoApi(sg *Swaggo) (*SGApi, error) {

	api, err := sg.Api("Person Resources")

	if err != nil {
		return nil, err
	}

	api.Get("/person/me").
		Summary("A function that does stuff").
		Param("abc", "The value of awesomeness").
		Param("def", "Slightly more awesomeness").
		ResponseClass("Person").
		MapOn(pr, "GetMe")

	api.Post("/person/me").
		Summary("Create a person").
		Param("param1", "its description").
		ResponseClass("Person").
		Map(func(param1 string, resp interface{}) {

	})

	return api, nil
}

func main() {
	sg := NewSwaggo("/{format}", "0.1")

	pr := &PersonResources{}

	pr.SwaggoApi(sg)

}
