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
	nav          *Nav
	Title        string
	Path         string
	templateName string
}

// Nav dereferencing the pointer so that the template can iterate over it
func (p Page) Nav() Nav {
	// We use a pointer in `nav` so that we can build the nav through the loop that generates data for the templates
	// and not have to do a second loop after we've gone through the API specs (because the nav is built from the specs)
	return *p.nav
}

type Nav []NavItem

type NavItem struct {
	Name    string
	SiteURL string
}

func (n NavItem) IsActive(currentPath string) bool {
	return strings.HasPrefix(currentPath, n.SiteURL)
}

func (n NavItem) GetRelativePath(currentPath string) string {
	c := strings.Count(currentPath, "/")
	return strings.Repeat("../", c) + n.SiteURL
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
	Spec     *spec.API
	Path     string
	Methods  []PathMethod
	APITitle string
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
		{"hierarchy-api", "https://raw.githubusercontent.com/ONSdigital/dp-hierarchy-api/cmd-develop/swagger.yaml", nil, nil},
		{"search-api", "https://raw.githubusercontent.com/ONSdigital/dp-search-api/cmd-develop/swagger.yaml", nil, nil},
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
	var orderedNav = &Nav{}

	for _, api := range APIs {
		var orderedPaths []APIPath
		apiDir := strings.TrimSuffix(api.ID, "-api")

		orderedNav.appendNavItem(api.Spec.Info.Title, apiDir)

		for key, path := range api.Spec.Paths.Paths {
			// generateMethods() only includes public methods so checking the length
			// so we don't add a path if none of it's methods are public
			pathMethods := generateMethods(path)
			if len(pathMethods) == 0 {
				continue
			}

			pathDir := strings.Replace(strings.TrimPrefix(strings.TrimSuffix(key, "index.html"), "/"), "/", "-", -1)
			fmt.Printf("pathDir: %+v\n", pathDir)
			orderedPaths = append(orderedPaths, APIPath{
				APIURL:        key,
				SiteURL:       pathDir,
				PathItemProps: path.PathItemProps,
			})

			siteModel[apiDir+"/"+pathDir] = Page{
				templateName: "path",
				Title:        api.Spec.Info.Title,
				Path:         apiDir + "/" + pathDir + "/",
				Data: PathPage{
					Spec:     api,
					Path:     key,
					Methods:  pathMethods,
					APITitle: api.Spec.Info.Title,
				},
				nav: orderedNav,
			}
		}

		sort.Slice(orderedPaths, func(i, j int) bool {
			return orderedPaths[i].APIURL < orderedPaths[j].APIURL
		})

		siteModel[apiDir] = Page{
			templateName: "api",
			Title:        api.Spec.Info.Title,
			Path:         apiDir + "/",
			Data: APIPage{
				Spec:         api,
				OrderedPaths: orderedPaths,
			},
			nav: orderedNav,
		}
	}

	return siteModel
}

func (n *Nav) appendNavItem(title string, url string) {
	*n = append(*n, NavItem{
		Name:    title,
		SiteURL: url,
	})
}

func generateMethods(path openAPI.PathItem) (methods []PathMethod) {
	//FIXME We're checking for the lack of 'Private' or 'Private user' tag on a method at the point
	// it'd be safer to check for 'Public' but the hierarchy API is currently missing that tag,
	// so this fixes that until the APIs spec is updated.
	if path.Get != nil && !contains(path.Get.Tags, Tags.Private) && !contains(path.Get.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "GET",
			OperationProps:   path.Get.OperationProps,
			OrderedResponses: generateResponses(path.Get.Responses),
		})
	}
	if path.Head != nil && !contains(path.Head.Tags, Tags.Private) && !contains(path.Head.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "HEAD",
			OperationProps:   path.Head.OperationProps,
			OrderedResponses: generateResponses(path.Head.Responses),
		})
	}
	if path.Post != nil && !contains(path.Post.Tags, Tags.Private) && !contains(path.Post.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "POST",
			OperationProps:   path.Post.OperationProps,
			OrderedResponses: generateResponses(path.Post.Responses),
		})
	}
	if path.Put != nil && !contains(path.Put.Tags, Tags.Private) && !contains(path.Put.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "PUT",
			OperationProps:   path.Put.OperationProps,
			OrderedResponses: generateResponses(path.Put.Responses),
		})
	}
	if path.Delete != nil && !contains(path.Delete.Tags, Tags.Private) && !contains(path.Delete.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "DELETE",
			OperationProps:   path.Delete.OperationProps,
			OrderedResponses: generateResponses(path.Delete.Responses),
		})
	}
	if path.Options != nil && !contains(path.Options.Tags, Tags.Private) && !contains(path.Options.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "OPTIONS",
			OperationProps:   path.Options.OperationProps,
			OrderedResponses: generateResponses(path.Options.Responses),
		})
	}
	if path.Patch != nil && !contains(path.Patch.Tags, Tags.Private) && !contains(path.Patch.Tags, Tags.PrivateUser) {
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
