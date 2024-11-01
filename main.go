package main

import (
	"context"
	"os"

	"github.com/ONSdigital/dp-developer-site/renderer"
	"github.com/ONSdigital/dp-developer-site/site"
	"github.com/ONSdigital/dp-developer-site/spec"
	"github.com/ONSdigital/log.go/v2/log"
)

const (
	serviceName   = "dp-developer-site"
	staticRootDir = "static"
)

func main() {
	log.Namespace = serviceName
	ctx := context.Background()

	// Load API specs
	sources := spec.APIs{
		{"dataset-api", "https://raw.githubusercontent.com/ONSdigital/dp-dataset-api/master/swagger.yaml", nil, nil},
		{"population-types-api", "https://raw.githubusercontent.com/ONSdigital/dp-population-types-api/master/swagger.yaml", nil, nil},
		{"filter-api", "https://raw.githubusercontent.com/ONSdigital/dp-filter-api/master/swagger.yaml", nil, nil},
		{"cantabular-filter-flex-api", "https://raw.githubusercontent.com/ONSdigital/dp-cantabular-filter-flex-api/master/swagger.yaml", nil, nil},
		{"code-list-api", "https://raw.githubusercontent.com/ONSdigital/dp-code-list-api/master/swagger.yaml", nil, nil},
		{"hierarchy-api", "https://raw.githubusercontent.com/ONSdigital/dp-hierarchy-api/master/swagger.yaml", nil, nil},
		{"dimension-search-api", "https://raw.githubusercontent.com/ONSdigital/dp-dimension-search-api/master/swagger.yaml", nil, nil},
		{"search-api", "https://raw.githubusercontent.com/ONSdigital/dp-search-api/master/swagger.yaml", nil, nil},
		{"topic-api", "https://raw.githubusercontent.com/ONSdigital/dp-topic-api/master/swagger.yaml", nil, nil},
	}

	if err := sources.Load(); err != nil {
		log.Error(ctx, "failed to load sources", err)
	}

	// Populate site model
	orderedNav := createNav()
	siteModel := site.NewSite(orderedNav, sources, staticRootDir)

	// Generate static HTML site assets
	log.Info(ctx, "creating assets directories and HTML files")
	for dir, page := range *siteModel {
		if err := os.MkdirAll("assets/"+dir, 0755); err != nil {
			log.Error(ctx, "failed to create directories", err)
		}

		file, err := os.Create("assets/" + dir + "/index.html")
		if err != nil {
			log.Error(ctx, "failed to create HTML files", err)
		}
		defer file.Close()

		if err = renderer.Render(file, page.TemplateName, page); err != nil {
			log.Error(ctx, "failed to render templates", err)
		}
	}

	log.Info(ctx, "files created")
}

func createNav() *site.Nav {
	var orderedNav = &site.Nav{}
	orderedNav.AppendNavItem("Introduction", "", false)
	// FIXME need to handle static content
	orderedNav.AppendNavItem("Guide to rate limiting and bot development", "bots", false)
	orderedNav.AppendNavItem("Take a tour of the API", "tour/getting-started", true)
	orderedNav.AppendNavItem("Create a custom dataset - Census 2021", "createyourowndataset", false)
	orderedNav.AppendNavItem("Guide to requesting Census 2021 observations", "censusobservations", false)
	orderedNav.AppendNavItem("Guide to requesting CMD observations", "cmdobservations", false)
	orderedNav.AppendNavItem("Guide to filtering a CMD dataset", "filters", false)
	orderedNav.AppendNavItem("Guide to filtering a Census 2021 dataset", "censusfilters", false)
	orderedNav.AppendNavItem("Guide to retirement of API endpoints", "retirement", false)

	return orderedNav
}
