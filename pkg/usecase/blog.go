package usecase

import (
	"github.com/wonpanu/learn-golang/pkg/entity"
	"github.com/wonpanu/learn-golang/pkg/repo"
)

type BlogUsecase struct {
	blogRepo repo.BlogRepo
}

func (b *BlogUsecase) GetAll() ([]entity.Blog, error) {
	return b.blogRepo.GetAll()
}

func NewBlogUsecase(blogRepo repo.BlogRepo) BlogUsecase {
	return BlogUsecase{
		blogRepo: blogRepo,
	}
}
