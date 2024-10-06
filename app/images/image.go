package images

import (
	handler "github.com/hammer-code/lms-be/app/images/delivery/http"
	repo "github.com/hammer-code/lms-be/app/images/repository"
	usecase "github.com/hammer-code/lms-be/app/images/usecase"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/db"
)

func InitRepository(db db.DatabaseTransaction) domain.ImageRepository {
	return repo.NewRepository(db)
}

func InitUsecase(repository domain.ImageRepository, dbTX db.DatabaseTransaction) domain.ImageUsecase {
	return usecase.NewUsecase(repository, dbTX)
}

func InitHandler(usecase domain.ImageUsecase) domain.ImageHandler {
	return handler.NewHandler(usecase)
}
