package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sourcegraph/syntaxhighlight"

	"github.com/ONSdigital/dp-developer-site/renderer"
	"github.com/ONSdigital/dp-developer-site/spec"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/PuerkitoBio/goquery"
	blackfriday "github.com/russross/blackfriday/v2"

	openAPI "github.com/go-openapi/spec"
)

const (
	serviceName = "dp-developer-site"
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
	JSOnly  bool
}

func (n NavItem) IsActive(currentPath string) bool {
	// Stop false-positives when handling root path
	// (otherwise root path in navigation will always be active)
	if n.SiteURL == "" && currentPath != "" {
		return false
	}

	pathRoot := strings.Split(currentPath, "/")[0]
	linkRoot := strings.Split(n.SiteURL, "/")[0]

	return pathRoot == linkRoot
}

func (n NavItem) GetRelativePath(currentPath string) string {
	c := strings.Count(currentPath, "/")
	p := strings.Repeat("../", c)
	if len(n.SiteURL) > 0 {
		p += n.SiteURL + "/"
	}
	return p
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

func main() {
	log.Namespace = serviceName
	ctx := context.Background()

	sources := spec.APIs{
		{"dataset-api", "https://raw.githubusercontent.com/ONSdigital/dp-dataset-api/master/swagger.yaml", nil, nil, false},
		{"population-types-api", "https://raw.githubusercontent.com/ONSdigital/dp-population-types-api/master/swagger.yaml", nil, nil, false},
		{"filter-api", "https://raw.githubusercontent.com/ONSdigital/dp-filter-api/master/swagger.yaml", nil, nil, false},
		{"cantabular-filter-flex-api", "https://raw.githubusercontent.com/ONSdigital/dp-cantabular-filter-flex-api/master/swagger.yaml", nil, nil, false},
		{"code-list-api", "https://raw.githubusercontent.com/ONSdigital/dp-code-list-api/master/swagger.yaml", nil, nil, false},
		{"hierarchy-api", "https://raw.githubusercontent.com/ONSdigital/dp-hierarchy-api/master/swagger.yaml", nil, nil, false},
		{"dimension-search-api", "https://raw.githubusercontent.com/ONSdigital/dp-dimension-search-api/master/swagger.yaml", nil, nil, false},
		{"search-api", "/Users/rhysstromaine/src/github.com/ONSdigital/dp-search-api/swagger.yaml", nil, nil, true},
	}

	if err := sources.Load(); err != nil {
		log.Error(ctx, "failed to load sources", err)
	}

	siteModel := generateModel(sources)
	log.Info(ctx, "creating assets directories and HTML files")
	for key, value := range siteModel {
		if err := os.MkdirAll("assets/"+key, 0755); err != nil {
			log.Error(ctx, "failed to create directories", err)
		}

		file, err := os.Create("assets/" + key + "/index.html")
		if err != nil {
			log.Error(ctx, "failed to create HTML files", err)
		}
		defer file.Close()

		if err = renderer.Render(file, value.templateName, value); err != nil {
			log.Error(ctx, "failed to render templates", err)
		}
	}

	log.Info(ctx, "files created")
}

func generateModel(APIs spec.APIs) site {
	var siteModel = make(site)
	var orderedNav = &Nav{}
	orderedNav.appendNavItem("Introduction", "", false)
	// FIXME need to handle static content
	orderedNav.appendNavItem("Guide to rate limiting and bot development", "bots", false)
	orderedNav.appendNavItem("Take a tour of the API", "tour/getting-started", true)
	orderedNav.appendNavItem("Create a custom dataset - Census 2021", "createyourowndataset", false)
	orderedNav.appendNavItem("Guide to requesting Census 2021 observations", "censusobservations", false)
	orderedNav.appendNavItem("Guide to requesting CMD observations", "cmdobservations", false)
	orderedNav.appendNavItem("Guide to filtering a CMD dataset", "filters", false)
	orderedNav.appendNavItem("Guide to filtering a Census 2021 dataset", "censusfilters", false)
	orderedNav.appendNavItem("Guide to retirement of API endpoints", "retirement", false)
	siteModel.generateDynamicPages(APIs, orderedNav)
	siteModel.generateStaticPages(orderedNav)

	return siteModel
}

func (n *Nav) appendNavItem(title string, url string, requiresJS bool) {
	*n = append(*n, NavItem{
		Name:    title,
		SiteURL: url,
		JSOnly:  requiresJS,
	})
}

func (s site) generateDynamicPages(a spec.APIs, orderedNav *Nav) {
	for _, api := range a {
		var orderedPaths []APIPath
		apiDir := strings.TrimSuffix(api.ID, "-api")

		orderedNav.appendNavItem(api.Spec.Info.Title, apiDir, false)

		for key, path := range api.Spec.Paths.Paths {
			// generateMethods() only includes public methods so checking the length
			// so we don't add a path if none of it's methods are public
			pathMethods := generateMethods(path)
			if len(pathMethods) == 0 {
				continue
			}

			pathDir := strings.Replace(strings.TrimPrefix(strings.TrimSuffix(key, "index.html"), "/"), "/", "-", -1)

			// Remove instances of curly brackets from the pathDir to have a cleaner URL.
			urlSanitiser := strings.NewReplacer(
				"{", "",
				"}", "",
			)
			sanitisedPathDir := urlSanitiser.Replace(pathDir)

			orderedPaths = append(orderedPaths, APIPath{
				APIURL:        key,
				SiteURL:       sanitisedPathDir + "/",
				PathItemProps: path.PathItemProps,
			})

			s[apiDir+"/"+sanitisedPathDir] = Page{
				templateName: "path",
				Title:        key,
				Path:         apiDir + "/" + sanitisedPathDir + "/",
				Data: PathPage{
					Spec:       api,
					Path:       key,
					Methods:    pathMethods,
					APITitle:   api.Spec.Info.Title,
					Parameters: path.PathItemProps.Parameters,
				},
				nav: orderedNav,
			}
		}

		sort.Slice(orderedPaths, func(i, j int) bool {
			return orderedPaths[i].APIURL < orderedPaths[j].APIURL
		})

		s[apiDir] = Page{
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

func generateResponses(responses *openAPI.Responses) (orderedResponses []MethodResponse) {
	for status, response := range responses.StatusCodeResponses {

		json, err := json.MarshalIndent(response.ResponseProps.Schema, "", "  ")

		if err != nil {
			log.Error(context.TODO(), "creating assets directories and HTML files", err)
			json = []byte{}
		}

		orderedResponses = append(orderedResponses, MethodResponse{
			Status:          status,
			ResponseProps:   response.ResponseProps,
			ExampleResponse: string(json),
		})
	}

	sort.Slice(orderedResponses, func(i, j int) bool {
		return orderedResponses[i].Status < orderedResponses[j].Status
	})

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

func (s site) generateStaticPages(orderedNav *Nav) {
	err := filepath.Walk("static", func(path string, info os.FileInfo, err error) error {
		if err != nil {

			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", "static", err)
			return err
		}
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, "index.md") {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				log.Error(context.TODO(), "failed to read index.md file", err)
			}

			templateBytes, metadata := generateStaticMetadata(bytes)
			html := blackfriday.Run(templateBytes, blackfriday.WithExtensions(blackfriday.AutoHeadingIDs|blackfriday.FencedCode))
			styledHTML := generateStyledCodeHTML(html)
			fileDir := strings.TrimSuffix(strings.TrimPrefix(path, "static/"), "index.md")
			s[fileDir] = Page{
				Title:        metadata["title"],
				Path:         fileDir,
				Data:         template.HTML(styledHTML),
				nav:          orderedNav,
				templateName: "static",
			}
		}

		if strings.HasSuffix(path, "index.html") {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				log.Error(context.TODO(), "failed to read index.html", err)
			}

			templateBytes, metadata := generateStaticMetadata(bytes)
			fileDir := strings.TrimSuffix(strings.TrimPrefix(path, "static/"), "index.html")
			s[fileDir] = Page{
				Title:        metadata["title"],
				Path:         fileDir,
				Data:         template.HTML(templateBytes),
				nav:          orderedNav,
				templateName: "html",
			}
		}

		return nil
	})
	if err != nil {
		log.Error(context.TODO(), "failed to generate static files", err)
	}
}

func generateStaticMetadata(md []byte) (b []byte, metadata map[string]string) {
	metadata = make(map[string]string)
	s := string(md)
	lines := strings.Split(s, "\n")

	var body []string
	var isMetadata bool

	if len(lines) == 0 {
		return
	}

	if strings.TrimSpace(lines[0]) != "---" {
		return
	}

	for _, line := range lines {
		if !isMetadata && line == "---" {
			isMetadata = true
			continue
		}

		if isMetadata {
			if line == "---" {
				isMetadata = false
				continue
			}
			pair := strings.SplitN(line, ":", 2)
			if len(pair) == 2 {
				metadata[strings.TrimSpace(strings.ToLower(pair[0]))] = pair[1]
			}
			continue
		}

		body = append(body, line)
	}

	b = []byte(strings.Join(body, "\n"))

	return
}

func generateStyledCodeHTML(html []byte) []byte {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		log.Error(context.TODO(), "failed to read html file", err)
	}

	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("tabindex", "0")
	})

	doc.Find("code[class*=\"language-\"]").Each(func(i int, s *goquery.Selection) {
		formattedCode, err := syntaxhighlight.AsHTML([]byte(s.Text()))
		if err != nil {
			log.Error(context.TODO(), "failed to format HTML code blocks", err)
		}
		s.SetHtml(string(formattedCode))
	})

	formattedHTML, err := doc.Html()
	if err != nil {
		log.Error(context.TODO(), "failed to find formatted HTML", err)
	}

	formattedHTML = strings.Replace(formattedHTML, "<html><head></head><body>", "", 1)
	formattedHTML = strings.Replace(formattedHTML, "</body></html>", "", 1)

	return []byte(formattedHTML)
}
