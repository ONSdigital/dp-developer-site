package site

import (
	"strings"
)

type Nav []NavItem

type NavItem struct {
	Name    string
	SiteURL string
	JSOnly  bool
}

func (n *Nav) AppendNavItem(title string, url string, requiresJS bool) {
	*n = append(*n, NavItem{
		Name:    title,
		SiteURL: url,
		JSOnly:  requiresJS,
	})
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
