package site

import (
	"github.com/ONSdigital/dp-developer-site/spec"

	openAPI "github.com/go-openapi/spec"
)

type Page struct {
	Data         interface{}
	nav          *Nav
	Title        string
	Path         string
	TemplateName string
}

// Nav dereferencing the pointer so that the template can iterate over it
func (p Page) Nav() Nav {
	// We use a pointer in `nav` so that we can build the nav through the loop that generates data for the templates
	// and not have to do a second loop after we've gone through the API specs (because the nav is built from the specs)
	return *p.nav
}

type APIPage struct {
	Spec         *spec.API
	OrderedPaths []APIPath
}

type APIPath struct {
	APIURL        string
	SiteURL       string
	PathItemProps openAPI.PathItemProps
}

type PathPage struct {
	Spec       *spec.API
	Path       string
	Methods    []PathMethod
	Parameters []openAPI.Parameter
	APITitle   string
}

type PathMethod struct {
	openAPI.OperationProps
	Method           string
	OrderedResponses []MethodResponse
}

type MethodResponse struct {
	Status int
	openAPI.ResponseProps
	ExampleResponse string
}

type tags struct {
	Public      string
	Private     string
	PrivateUser string
}

var Tags = tags{
	Public:      "Public",
	Private:     "Private",
	PrivateUser: "Private user",
}
