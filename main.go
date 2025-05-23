package main

import (
	"github.com/hairyhenderson/gomplate/v4"
	"net/url"
)

func main() {
	m1 := New("m1")
	u, _ := url.Parse("./example/simple/greeting.yml")
	m1.Src = gomplate.DataSource{URL: u}

	m1.dst = "out.yml"

	u2, _ := url.Parse("https://ipinfo.io")
	c := gomplate.DataSource{URL: u2}
	cs := make(map[string]gomplate.DataSource)
	cs["."] = c

	m1.g.Context = cs

	err := m1.Fetch()
	if err != nil {
		m1.l.Error(err.Error())
	}

}
