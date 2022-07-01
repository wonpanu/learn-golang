package usecase

import (
	"encoding/json"

	"github.com/wonpanu/learn-golang/service/pkg/entity"
	"github.com/wonpanu/learn-golang/service/pkg/repo"
)

type BlogUsecase struct {
	blogRepo repo.BlogRepo
}

func (b *BlogUsecase) GetAll() (blogs []entity.Blog, err error) {
	blogs, err = b.blogRepo.GetAll()
	if err != nil {
		return blogs, err
	}

	bulkLogJSON, _ := json.Marshal(map[string]string{
		"msg": "User get all blog data",
	})
	err = b.blogRepo.Publish("blog", bulkLogJSON)
	return blogs, err
}

func NewBlogUsecase(blogRepo repo.BlogRepo) BlogUsecase {
	return BlogUsecase{
		blogRepo: blogRepo,
	}
}
