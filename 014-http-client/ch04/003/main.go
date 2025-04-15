package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	protocol := parsedUrl.Scheme
	username := parsedUrl.User.Username()
	password, _ := parsedUrl.User.Password()
	hostname := parsedUrl.Hostname()
	port := parsedUrl.Port()
	pathname := parsedUrl.Path
	search := parsedUrl.RawQuery
	hash := parsedUrl.Fragment

	return ParsedURL{
		protocol: protocol,
		username: username,
		password: password,
		hostname: hostname,
		port:     port,
		pathname: pathname,
		search:   search,
		hash:     hash,
	}
}
