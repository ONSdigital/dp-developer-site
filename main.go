package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/ONSdigital/dp-developer-site/renderer"
	"github.com/ONSdigital/dp-developer-site/spec"

	openAPI "github.com/go-openapi/spec"
)

type site map[string]Page

type Page struct {
	Data         interface{}
	Nav          interface{}
	Title        string
	templateName string
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
	Spec    *spec.API
	Path    string
	Methods []PathMethod
}

type PathMethod struct {
	openAPI.OperationProps
	Method           string
	OrderedResponses []MethodResponse
}

type MethodResponse struct {
	Status int
	openAPI.ResponseProps
}

func main() {
	sources := spec.APIs{
		{"dataset-api", "https://raw.githubusercontent.com/ONSdigital/dp-dataset-api/cmd-develop/swagger.yaml", nil, nil},
		{"import-api", "https://raw.githubusercontent.com/ONSdigital/dp-import-api/cmd-develop/swagger.yaml", nil, nil},
	}

	if err := sources.Load(); err != nil {
		log.Fatal(err)
	}

	siteModel := generateModel(sources)
	fmt.Println("Creating files...")
	for key, value := range siteModel {
		if err := os.MkdirAll("assets/"+key, 0755); err != nil {
			log.Fatal(err)
		}

		file, err := os.Create("assets/" + key + "/index.html")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if err = renderer.Render(file, value.templateName, value); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Created files")
	fmt.Println("Finished!")
}

func generateModel(APIs spec.APIs) site {
	var siteModel = make(site)

	for _, api := range APIs {
		var orderedPaths []APIPath
		apiDir := strings.TrimSuffix(api.ID, "-api")

		for key, path := range api.Spec.Paths.Paths {
			pathDir := strings.Replace(strings.TrimPrefix(key, "/"), "/", "-", -1)
			orderedPaths = append(orderedPaths, APIPath{
				APIURL:        key,
				SiteURL:       pathDir + "/index.html",
				PathItemProps: path.PathItemProps,
			})

			siteModel[apiDir+"/"+pathDir] = Page{
				templateName: "path",
				Title:        api.Spec.Info.Title,
				Data: PathPage{
					Spec:    api,
					Path:    key,
					Methods: generateMethods(path),
				},
			}
		}

		sort.Slice(orderedPaths, func(i, j int) bool {
			return orderedPaths[i].APIURL < orderedPaths[j].APIURL
		})

		siteModel[apiDir] = Page{
			templateName: "api",
			Title:        api.Spec.Info.Title,
			Data: APIPage{
				Spec:         api,
				OrderedPaths: orderedPaths,
			},
			Nav: nil,
		}
	}

	return siteModel
}

func generateMethods(path openAPI.PathItem) (methods []PathMethod) {
	if path.Get != nil {
		methods = append(methods, PathMethod{
			Method:           "GET",
			OperationProps:   path.Get.OperationProps,
			OrderedResponses: generateResponses(path.Get.Responses),
		})
	}
	if path.Put != nil {
		methods = append(methods, PathMethod{
			Method:           "PUT",
			OperationProps:   path.Put.OperationProps,
			OrderedResponses: generateResponses(path.Put.Responses),
		})
	}
	if path.Post != nil {
		methods = append(methods, PathMethod{
			Method:           "POST",
			OperationProps:   path.Post.OperationProps,
			OrderedResponses: generateResponses(path.Post.Responses),
		})
	}

	return
}

func generateResponses(responses *openAPI.Responses) (orderedResponses []MethodResponse) {
	for status, response := range responses.StatusCodeResponses {
		orderedResponses = append(orderedResponses, MethodResponse{
			Status:        status,
			ResponseProps: response.ResponseProps,
		})
	}

	sort.Slice(orderedResponses, func(i, j int) bool {
		return orderedResponses[i].Status < orderedResponses[j].Status
	})

	return
}
