package main

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + a.username + ":" + a.password
}
