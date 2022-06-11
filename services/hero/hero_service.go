package hero

import (
	"gin-bmg-restful/entities/web"
	"gin-bmg-restful/repositories/hero"
)

// Service main struct
type Service struct {
	heroRepo *hero.Repository
}

// NewService constructor
func NewService(heroRepo *hero.Repository) *Service {
	return &Service{
		heroRepo: heroRepo,
	}
}

// FindAll return heroes with search
func (service Service) FindAll(nameSearch string) ([]web.Hero, error) {
	heroes, err := service.heroRepo.FindAllHeroes(nameSearch)
	if err != nil {
		return []web.Hero{}, err
	}
	return heroes, nil
}