package client

import "fmt"

type Link struct {
	Ref string `json:"ref"`
	Method string `json:"method"`
	Href string `json:"href"`
}

func GetLinkByRef(links []Link, ref string) (Link, error) {
	link := Link{}

	for _, l := range links {
		if l.Ref == ref {
			link = l
			break
		}
	}

	if link.Href == "" {
		return link, fmt.Errorf("%s link is empty", ref)
	}

	return link, nil
}