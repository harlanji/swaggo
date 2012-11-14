package main

import (
)

type SwagGo struct{
  handlers []*SwagGoBody
}

func (sg *SwagGo) Get(path string) *SwagGoBody {
  return &SwagGoBody{ sg: sg, method: "GET" }
}
func (sg *SwagGo) Post(path string) *SwagGoBody {
  return &SwagGoBody{ sg: sg, method: "POST" }
}
func (sg *SwagGo) Put(path string) *SwagGoBody {
  return &SwagGoBody{ sg: sg, method: "PUT" }
}
func (sg *SwagGo) Delete(path string) *SwagGoBody {
  return &SwagGoBody{ sg: sg, method: "DELETE" }
}


type SwagGoHandler interface{}

type SwagGoBody struct{
  method string
  sg *SwagGo
  handler SwagGoHandler
  description string
  params map[string]string
  returns string
  accepts string
}

func (sgb *SwagGoBody) DoOn(obj interface{}, funcName string) error {
  sgb.handler = func() {
    // call funcName on obj with passed in args
  }
  return nil
}
func (sgb *SwagGoBody) Do(handler SwagGoHandler) error {
  sgb.handler = handler
  return nil
}
func (sgb *SwagGoBody) Description(description string) *SwagGoBody {
  sgb.description = description
  return sgb
}
func (sgb *SwagGoBody) Param(name string, description string) *SwagGoBody {
  sgb.params[name] = description
  return sgb
}
func (sgb *SwagGoBody) Returns(typ string) *SwagGoBody {
  sgb.returns = typ
  return sgb
}
func (sgb *SwagGoBody) Accepts(typ string) *SwagGoBody {
  sgb.accepts = typ
  return sgb
}

type Person struct {

}

type PersonResources struct {}


type SwagGoResources interface {
  SwagGoResources(sg *SwagGo) error
}


func (pr *PersonResources) GetMe(abc string, def string, resp interface{}) error {
  //resp.Println("Hello, world!")
  return nil
}

func (pr *PersonResources) SwagGoResources(sg *SwagGo) error {

  sg.Get("/person/me").
    Description("A function that does stuff").
    Param("abc", "The value of awesomeness").
    Param("def", "Slightly more awesomeness").
    Returns("Person").
    DoOn(pr, "GetMe")

  sg.Post("/person/me").
    Description("Create a person").
    Param("param1", "its description").
    Accepts("Person").
    Returns("Person").
    Do(func(param1 string, resp interface{}) {

    })

  return nil;
}

func main() {

  sg := new(SwagGo)


  pr := &PersonResources{}

  pr.SwagGoResources(sg)
}
