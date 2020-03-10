# The API For Go language used embed all type media.

## Suport Website

1. Twiiter
2. Youtube
3. Facebook
4. Instagram

## Install

`go get github.com/didikprabowo/go-embed/embed`

## Example main

```
// example main
package main

import (
	"fmt"
	"github.com/didikprabowo/go-embed/embed"
)

// example main
func main() {
	g := embed.InitEmbed()

	g = embed.NewFacebook("https://www.facebook.com/ceritakatajawalucu/videos/1352107424999312")
	s, _ := g.Get()
	fmt.Println(s)

	g = embed.NewYoutube("https://www.youtube.com/watch?v=iEkETgm-fFo")
	j, _ := g.Get()
	fmt.Println(j)

	g = embed.NewTwitter("https://twitter.com/ReynaOlivia10/status/1237079194497994758")
	t, _ := g.Get()
	fmt.Println(t)

	g = embed.NewInstagram("http://instagram.com/p/V8UMy0LjpX/")
	ig, _ := g.Get()
	fmt.Println(ig)
}

```
