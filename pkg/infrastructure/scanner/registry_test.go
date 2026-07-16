//     MIT License
//
//     Copyright (c) Microsoft Corporation.
//
//     Permission is hereby granted, free of charge, to any person obtaining a copy
//     of this software and associated documentation files (the "Software"), to deal
//     in the Software without restriction, including without limitation the rights
//     to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//     copies of the Software, and to permit persons to whom the Software is
//     furnished to do so, subject to the following conditions:
//
//     The above copyright notice and this permission notice shall be included in all
//     copies or substantial portions of the Software.
//
//     THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//     IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//     FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//     AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//     LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//     OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//     SOFTWARE

package scanner

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNextPageFromLink(t *testing.T) {
	base, err := url.Parse("https://mcr.microsoft.com/v2/repo/tags/list?n=100")
	require.NoError(t, err)

	tests := []struct {
		name   string
		header string
		want   string
	}{
		{"empty", "", ""},
		{"next relative", `</v2/repo/tags/list?n=100&last=abc>; rel="next"`, "https://mcr.microsoft.com/v2/repo/tags/list?n=100&last=abc"},
		{"next absolute", `<https://mcr.microsoft.com/v2/repo/tags/list?n=100&last=z>; rel="next"`, "https://mcr.microsoft.com/v2/repo/tags/list?n=100&last=z"},
		{"prev only", `</v2/repo/tags/list?n=100>; rel="prev"`, ""},
		{"mixed", `</v2/repo/tags/list?n=100&last=a>; rel="prev", </v2/repo/tags/list?n=100&last=b>; rel="next"`, "https://mcr.microsoft.com/v2/repo/tags/list?n=100&last=b"},
		{"rel without quotes", `</v2/repo/tags/list?n=100&last=c>; rel=next`, "https://mcr.microsoft.com/v2/repo/tags/list?n=100&last=c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nextPageFromLink(base, tt.header))
		})
	}
}

func TestGetTags_FollowsPagination(t *testing.T) {
	var hits int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		assert.Equal(t, "/v2/azurelinux/base/python/tags/list", r.URL.Path)

		switch r.URL.Query().Get("last") {
		case "":
			assert.Equal(t, "100", r.URL.Query().Get("n"))
			w.Header().Set("Link", `</v2/azurelinux/base/python/tags/list?n=100&last=3.11>; rel="next"`)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name": "azurelinux/base/python",
				"tags": []string{"3.12", "3.11"},
			})
		case "3.11":
			// Final page — no Link header
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name": "azurelinux/base/python",
				"tags": []string{"3.10", "3.9"},
			})
		default:
			t.Fatalf("unexpected last=%q", r.URL.Query().Get("last"))
		}
	}))
	defer srv.Close()

	s := &RegistryScanner{
		client:      srv.Client(),
		registryURL: srv.URL,
	}

	tags, err := s.GetTags("azurelinux/base/python")
	require.NoError(t, err)
	assert.Equal(t, 2, hits, "should request both pages")
	assert.Equal(t, []string{"3.12", "3.11", "3.10", "3.9"}, tags)
}

func TestGetTags_SinglePage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"name": "repo",
			"tags": []string{"1.0", "2.0"},
		})
	}))
	defer srv.Close()

	s := &RegistryScanner{client: srv.Client(), registryURL: srv.URL}
	tags, err := s.GetTags("repo")
	require.NoError(t, err)
	assert.Equal(t, []string{"1.0", "2.0"}, tags)
}

func TestGetTags_NonOKStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	s := &RegistryScanner{client: srv.Client(), registryURL: srv.URL}
	_, err := s.GetTags("missing")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}
