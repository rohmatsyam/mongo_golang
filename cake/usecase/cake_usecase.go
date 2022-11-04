package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/rohmatsyam/mongo_golang/domain"
)

type cakeUseCase struct {
	cakeRepo domain.CakeRepository
}

func NewCakeUseCase(repo domain.CakeRepository) domain.CakeUseCase {
	return cakeUseCase{
		cakeRepo: repo,
	}
}

func (c cakeUseCase) GetCakesUC() ([]*domain.Cake, error) {
	return c.cakeRepo.GetCakesRepository()
}

func (c cakeUseCase) GetCakeIdUC(ctx *gin.Context) (*domain.Cake, error) {
	return c.cakeRepo.GetCakeIdRepository(ctx)
}

func (c cakeUseCase) CreateCakeUC(ctx *gin.Context) (*domain.Cake, error) {
	return c.cakeRepo.CreateCakeRepository(ctx)
}

func (c cakeUseCase) UpdateCakeUC(ctx *gin.Context) (*domain.Cake, error) {
	return c.cakeRepo.UpdateCakeRepository(ctx)
}

func (c cakeUseCase) DeleteCakeUC(ctx *gin.Context) error {
	return c.cakeRepo.DeleteCakeRepository(ctx)
}
