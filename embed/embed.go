package embed

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	Embed interface {
		Get() (map[string]string, error)
	}
	// Twiiter
	Youtube struct {
		URL          string
		AuthorName   string `json:"author_name"`
		ProviderURL  string `json:"provider_url"`
		HTML         string `json:"html"`
		ThumbnailURL string `json:"thumbnail_url"`
	}
	// Twiiter
	Facebook struct {
		URL         string `json:"url"`
		ProviderURL string `json:"provider_url"`
		AuthorName  string `json:"author_name"`
		HTML        string `json:"html"`
	}
	// Twiiter
	Twiiter struct {
		URL         string `json:"url"`
		AuthorName  string `json:"author_name"`
		ProviderURL string `json:"provider_url"`
		HTML        string `json:"html"`
	}
	// Instagram
	Instagram struct {
		URL          string `json:"url"`
		AuthorName   string `json:"author_name"`
		ProviderURL  string `json:"provider_url"`
		ThumbnailURL string `json:"thumbnail_url"`
		HTML         string `json:"html"`
	}
)

// const
const (
	FacebookURL  string = "https://www.facebook.com/plugins/video/oembed.json/?url="
	YoutubeURL   string = "https://www.youtube.com/oembed?url="
	TwitterURL   string = "https://publish.twitter.com/oembed?url="
	InstagramURL string = "https://api.instagram.com/oembed/?url="
)

// InitEmbed
func InitEmbed() Embed {
	var e Embed
	return e
}

// NewYoutube
func NewYoutube(URL string) *Youtube {
	return &Youtube{
		URL: URL,
	}
}

// NewFacebook
func NewFacebook(URL string) *Facebook {
	return &Facebook{
		URL: URL,
	}
}

// NewTwitter
func NewTwitter(URL string) *Twiiter {
	return &Twiiter{
		URL: URL,
	}
}

// NewInstagram
func NewInstagram(URL string) *Instagram {
	return &Instagram{
		URL: URL,
	}
}

// getHTML
func (f *Facebook) Get() (map[string]string, error) {
	url := fmt.Sprintf("%v%v", FacebookURL, f.URL)
	res, err := NewRequest(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&f)
	if err != nil {
		return nil, err
	}
	out := map[string]string{
		"url":         f.URL,
		"author_name": f.AuthorName,
		"html":        f.HTML,
		"provider":    f.ProviderURL,
	}
	return out, nil
}

// GetHTML
func (y *Youtube) Get() (map[string]string, error) {
	parseURL := url.QueryEscape(y.URL)
	url := fmt.Sprintf("%v%s", YoutubeURL, parseURL)
	res, err := NewRequest(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&y)
	if err != nil {
		return nil, err
	}

	out := map[string]string{
		"url":           y.URL,
		"author_name":   y.AuthorName,
		"html":          y.HTML,
		"provider":      y.ProviderURL,
		"thumbnail_url": y.ThumbnailURL,
	}
	return out, nil
}

// Get()
func (ig *Instagram) Get() (map[string]string, error) {
	parseURL := url.QueryEscape(ig.URL)
	url := fmt.Sprintf("%v%s", InstagramURL, parseURL)
	res, err := NewRequest(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ig)
	if err != nil {
		return nil, err
	}

	out := map[string]string{
		"url":           ig.URL,
		"author_name":   ig.AuthorName,
		"html":          ig.HTML,
		"provider":      ig.ProviderURL,
		"thumbnail_url": ig.ThumbnailURL,
	}

	return out, nil
}

// GetHTML
func (t *Twiiter) Get() (map[string]string, error) {
	parseURL := url.QueryEscape(t.URL)
	url := fmt.Sprintf("%v%s", TwitterURL, parseURL)
	res, err := NewRequest(url)
	if err != nil {
		return map[string]string{}, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		return map[string]string{}, err
	}

	out := map[string]string{
		"url":         t.URL,
		"author_name": t.AuthorName,
		"html":        t.HTML,
		"provider":    t.ProviderURL,
	}

	return out, nil
}

// NewRequest
func NewRequest(URL string) (*http.Response, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("Recovered from : %v", r))
		}
	}()

	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		panic(err)
	}

	ctx, canc := context.WithTimeout(req.Context(), 1*time.Microsecond)

	defer canc()

	req.WithContext(ctx)

	var client http.Client

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
