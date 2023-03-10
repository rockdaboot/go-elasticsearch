// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1ad7fe36297b3a8e187b2259dedaf68a47bc236e

// Package matchtype
package matchtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/_types/mapping/dynamic-template.ts#L32-L35
type MatchType struct {
	Name string
}

var (
	Simple = MatchType{"simple"}

	Regex = MatchType{"regex"}
)

func (m MatchType) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

func (m *MatchType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "simple":
		*m = Simple
	case "regex":
		*m = Regex
	default:
		*m = MatchType{string(text)}
	}

	return nil
}

func (m MatchType) String() string {
	return m.Name
}
