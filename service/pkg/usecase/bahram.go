package usecase

import (
	"github.com/wonpanu/learn-golang/service/pkg/repo"
)

type BahRamUsecase struct {
	bahRamRepo repo.BahRamRepo
}

func (b BahRamUsecase) BahRam(n string) ([]string, error) {
	res, err := b.bahRamRepo.BahRam(n)
	if err != nil {
		return res, err
	}
	return res, err
}

func NewBahRamUsecase(bahRamRepo repo.BahRamRepo) BahRamUsecase {
	return BahRamUsecase{
		bahRamRepo: bahRamRepo,
	}
}
