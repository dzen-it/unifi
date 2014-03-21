package main

import (
	"flag"
	"fmt"
	"github.com/dim13/unifi"
)

func main() {
	user := flag.String("user", "admin", "User")
	pass := flag.String("pass", "unifi", "Password")
	url := flag.String("url", "unifi", "URL")
	flag.Parse()

	u := new(unifi.Unifi)
	u.Login(*user, *pass, *url)
	defer u.Logout()

	aps := u.GetAps()
	apmap := make(map[string]string)
	for _, ap := range aps {
		apmap[ap.Mac] = ap.Name
	}

	sta := u.GetClients()
	for _, v := range sta {
		fmt.Printf("%s at %s/%d\n", v.GetName(), apmap[v.Ap_mac], v.Channel)
	}
}
