package service

import pkg "clean-arch-template/pkg/logger"

type (
	ServiceDependencies struct {
		Logger pkg.Logger
	}

	templateService struct {
		log pkg.Logger
	}
)

func NewService(dep ServiceDependencies) Service {
	return &templateService{
		log: dep.Logger,
	}
}
