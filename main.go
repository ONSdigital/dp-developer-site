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

func main() {
	sources := spec.APIs{
		{"dataset-api", "https://raw.githubusercontent.com/ONSdigital/dp-dataset-api/cmd-develop/swagger.yaml", nil, nil},
		{"filter-api", "https://raw.githubusercontent.com/ONSdigital/dp-filter-api/cmd-develop/swagger.yaml", nil, nil},
		{"code-list-api", "https://raw.githubusercontent.com/ONSdigital/dp-code-list-api/cmd-develop/swagger.yaml", nil, nil},
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

	fmt.Println("Files created")
	fmt.Println("Finished!")
}

func generateModel(APIs spec.APIs) site {
	var siteModel = make(site)

	for _, api := range APIs {
		var orderedPaths []APIPath
		apiDir := strings.TrimSuffix(api.ID, "-api")

		for key, path := range api.Spec.Paths.Paths {
			pathDir := strings.Replace(strings.TrimPrefix(key, "/"), "/", "-", -1)
			pathMethods := generateMethods(path)

			// generateMethods() only includes public methods so checking the length
			// so we don't add a path if none of it's methods are public
			if len(pathMethods) == 0 {
				continue
			}

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
	if path.Get != nil && contains(path.Get.Tags, Tags.Public) {
		methods = append(methods, PathMethod{
			Method:           "GET",
			OperationProps:   path.Get.OperationProps,
			OrderedResponses: generateResponses(path.Get.Responses),
		})
	}
	if path.Head != nil && contains(path.Head.Tags, Tags.Public) {
		methods = append(methods, PathMethod{
			Method:           "HEAD",
			OperationProps:   path.Head.OperationProps,
			OrderedResponses: generateResponses(path.Head.Responses),
		})
	}
	if path.Post != nil && contains(path.Post.Tags, Tags.Public) {
		methods = append(methods, PathMethod{
			Method:           "POST",
			OperationProps:   path.Post.OperationProps,
			OrderedResponses: generateResponses(path.Post.Responses),
		})
	}
	if path.Put != nil && contains(path.Put.Tags, Tags.Public) {
		methods = append(methods, PathMethod{
			Method:           "PUT",
			OperationProps:   path.Put.OperationProps,
			OrderedResponses: generateResponses(path.Put.Responses),
		})
	}
	if path.Delete != nil && contains(path.Delete.Tags, Tags.Public) {
		methods = append(methods, PathMethod{
			Method:           "DELETE",
			OperationProps:   path.Delete.OperationProps,
			OrderedResponses: generateResponses(path.Delete.Responses),
		})
	}
	if path.Options != nil && contains(path.Options.Tags, Tags.Public) {
		methods = append(methods, PathMethod{
			Method:           "OPTIONS",
			OperationProps:   path.Options.OperationProps,
			OrderedResponses: generateResponses(path.Options.Responses),
		})
	}
	if path.Patch != nil && contains(path.Patch.Tags, Tags.Public) {
		methods = append(methods, PathMethod{
			Method:           "PATCH",
			OperationProps:   path.Patch.OperationProps,
			OrderedResponses: generateResponses(path.Patch.Responses),
		})
	}

	return
}

func contains(sl []string, s string) (b bool) {
	b = false
	for i := range sl {
		if sl[i] == s {
			b = true
			break
		}
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
