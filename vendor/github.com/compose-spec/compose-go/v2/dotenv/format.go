/*
   Copyright 2020 The Compose Specification Authors.

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

package dotenv

import (
	"fmt"
	"io"
)

var formats = map[string]Parser{}

type Parser func(r io.Reader, filename string, lookup func(key string) (string, bool)) (map[string]string, error)

func RegisterFormat(format string, p Parser) {
	formats[format] = p
}

func ParseWithFormat(r io.Reader, filename string, resolve LookupFn, format string) (map[string]string, error) {
	parser, ok := formats[format]
	if !ok {
		return nil, fmt.Errorf("unsupported env_file format %q", format)
	}
	return parser(r, filename, resolve)
}