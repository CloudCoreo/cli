// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"fmt"
)

// Link struct
type Link struct {
	Ref    string `json:"ref"`
	Method string `json:"method"`
	Href   string `json:"href"`
}

//GetLinkByRef get link object for given property
func GetLinkByRef(links []Link, ref string) (Link, error) {
	link := Link{}

	for _, l := range links {
		if l.Ref == ref {
			link = l
			break
		}
	}

	if link.Href == "" {
		return link, fmt.Errorf("resource for given ID not found")
	}
	return link, nil
}
