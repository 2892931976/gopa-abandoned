/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import "time"

type KV struct {
	Key   string   `json:"key,omitempty"`
	Value []string `storm:"inline" json:"value,omitempty"`
}

type LinkGroup struct {
	Internal []PageLink `json:"internal,omitempty"`
	External []PageLink `json:"external,omitempty"`
}

type Snapshot struct {
	ID      string `json:"id,omitempty" gorm:"not null;unique;primary_key"`
	Version int    `json:"version,omitempty"`
	Path    string `json:"path,omitempty"` //path of this file
	File    string `json:"file,omitempty"` //filename of this page

	StatusCode int    `json:"-"`
	Payload    []byte `json:"-"`
	Size       uint64 `json:"size,omitempty"`

	Headers    map[string][]string     `json:"-"`
	Metadata   *map[string]interface{} `json:"-"`
	Parameters []KV                    `json:"-"`

	Language string `json:"lang,omitempty"`

	Title       string `json:"title,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Text        string `json:"text,omitempty"`
	ContentType string `json:"content_type,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Links LinkGroup `json:"links,omitempty"`

	Images struct {
		Internal []PageLink `json:"internal,omitempty"`
		External []PageLink `json:"external,omitempty"`
	} `json:"images,omitempty"`

	H1     []string `json:"h1,omitempty"`
	H2     []string `json:"h2,omitempty"`
	H3     []string `json:"h3,omitempty"`
	H4     []string `json:"h4,omitempty"`
	H5     []string `json:"h5,omitempty"`
	Bold   []string `json:"bold,omitempty"`
	Italic []string `json:"italic,omitempty"`

	Classifications  []string                `json:"classifications,omitempty"`
	EnrichedFeatures *map[string]interface{} `json:"enriched_features,omitempty"`

	Hash    string `json:"hash,omitempty"`
	SimHash string `json:"sim_hash,omitempty"`

	CreateTime *time.Time `json:"created,omitempty"`
}

type PageLink struct {
	Url   string `json:"url"`
	Label string `json:"label"`
}
