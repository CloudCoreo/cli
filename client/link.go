package client


type Link struct {
	Ref string `json:"ref"`
	Method string `json:"method"`
	Href string `json:"href"`
}

func GetLinkByRef(links []Link, ref string) (Link) {
	link := Link{}

	for _, l := range links {
		if l.Ref == ref {
			link = l
			break
		}
	}

	return link
}