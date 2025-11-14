package proxyrepo

import (
	"proxy/repository"
)

type Proxyrepo struct {
	repo *repository.Repository
}

func (p *Proxyrepo) GetByID() interface{} {
	return p.repo.GetByID()
}
