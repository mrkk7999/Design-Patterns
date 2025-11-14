package main

import proxyrepo "proxy/proxy-repo"

func main() {
	dbRepo := proxyrepo.Proxyrepo{}
	dbRepo.GetByID()
}
