package main

import (
	tokenHttp "github.com/sirajudheenam/GoRepo/openstack/token_http"
	token "github.com/sirajudheenam/GoRepo/openstack/tokens"
)

func main() {
	// Get token from raw HTTP method
	tokenHttp.GetOSTokenHttp()
	// Get token from token package of gophercloud
	token.GetOSToken()
}
