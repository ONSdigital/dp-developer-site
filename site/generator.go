package site

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

	"github.com/ONSdigital/dp-developer-site/spec"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/PuerkitoBio/goquery"
	blackfriday "github.com/russross/blackfriday/v2"

	openAPI "github.com/go-openapi/spec"
)

type Site map[string]Page

func NewSite(nav *Nav, APIs spec.APIs, staticRootDir string) *Site {
	siteModel := &Site{}
	siteModel.addAPISpecPages(nav, APIs)
	siteModel.addStaticPages(nav, staticRootDir)

	return siteModel
}

func (s Site) addAPISpecPages(orderedNav *Nav, a spec.APIs) {
	for _, api := range a {
		var orderedPaths []APIPath
		apiDir := strings.TrimSuffix(api.ID, "-api")

		orderedNav.AppendNavItem(api.Spec.Info.Title, apiDir, false)

		for key, path := range api.Spec.Paths.Paths {
			// generateMethods() only includes public methods so checking the length
			// so we don't add a path if none of it's methods are public
			pathMethods := getPublicMethodsList(path)
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
				TemplateName: "path",
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
			TemplateName: "api",
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

func getPublicMethodsList(path openAPI.PathItem) (methods []PathMethod) {
	//FIXME We're checking for the lack of 'Private' or 'Private user' tag on a method at the point
	// it'd be safer to check for 'Public' but the hierarchy API is currently missing that tag,
	// so this fixes that until the APIs spec is updated.
	if path.Get != nil && !contains(path.Get.Tags, Tags.Private) && !contains(path.Get.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "GET",
			OperationProps:   path.Get.OperationProps,
			OrderedResponses: getResponsesList(path.Get.Responses),
		})
	}
	if path.Head != nil && !contains(path.Head.Tags, Tags.Private) && !contains(path.Head.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "HEAD",
			OperationProps:   path.Head.OperationProps,
			OrderedResponses: getResponsesList(path.Head.Responses),
		})
	}
	if path.Post != nil && !contains(path.Post.Tags, Tags.Private) && !contains(path.Post.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "POST",
			OperationProps:   path.Post.OperationProps,
			OrderedResponses: getResponsesList(path.Post.Responses),
		})
	}
	if path.Put != nil && !contains(path.Put.Tags, Tags.Private) && !contains(path.Put.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "PUT",
			OperationProps:   path.Put.OperationProps,
			OrderedResponses: getResponsesList(path.Put.Responses),
		})
	}
	if path.Delete != nil && !contains(path.Delete.Tags, Tags.Private) && !contains(path.Delete.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "DELETE",
			OperationProps:   path.Delete.OperationProps,
			OrderedResponses: getResponsesList(path.Delete.Responses),
		})
	}
	if path.Options != nil && !contains(path.Options.Tags, Tags.Private) && !contains(path.Options.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "OPTIONS",
			OperationProps:   path.Options.OperationProps,
			OrderedResponses: getResponsesList(path.Options.Responses),
		})
	}
	if path.Patch != nil && !contains(path.Patch.Tags, Tags.Private) && !contains(path.Patch.Tags, Tags.PrivateUser) {
		methods = append(methods, PathMethod{
			Method:           "PATCH",
			OperationProps:   path.Patch.OperationProps,
			OrderedResponses: getResponsesList(path.Patch.Responses),
		})
	}

	return
}

func getResponsesList(responses *openAPI.Responses) (orderedResponses []MethodResponse) {
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

func (s Site) addStaticPages(orderedNav *Nav, rootDir string) {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
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

			templateBytes, metadata := getStaticMetadata(bytes)
			html := blackfriday.Run(templateBytes, blackfriday.WithExtensions(blackfriday.AutoHeadingIDs|blackfriday.FencedCode|blackfriday.Tables))
			styledHTML := formatHTMLCodeBlocks(html)
			fileDir := strings.TrimSuffix(strings.TrimPrefix(path, "static/"), "index.md")
			s[fileDir] = Page{
				Title:        metadata["title"],
				Path:         fileDir,
				Data:         template.HTML(styledHTML),
				nav:          orderedNav,
				TemplateName: "static",
			}
		}

		if strings.HasSuffix(path, "index.html") {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				log.Error(context.TODO(), "failed to read index.html", err)
			}

			templateBytes, metadata := getStaticMetadata(bytes)
			fileDir := strings.TrimSuffix(strings.TrimPrefix(path, "static/"), "index.html")
			s[fileDir] = Page{
				Title:        metadata["title"],
				Path:         fileDir,
				Data:         template.HTML(templateBytes),
				nav:          orderedNav,
				TemplateName: "html",
			}
		}

		return nil
	})
	if err != nil {
		log.Error(context.TODO(), "failed to generate static files", err)
	}
}

func getStaticMetadata(md []byte) (b []byte, metadata map[string]string) {
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

func formatHTMLCodeBlocks(html []byte) []byte {
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