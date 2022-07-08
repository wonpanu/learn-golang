package usecase

import (
	"encoding/json"

	"github.com/wonpanu/learn-golang/service/pkg/entity"
	"github.com/wonpanu/learn-golang/service/pkg/repo"
)

type BlogUsecase struct {
	blogRepo repo.BlogRepo
}

func (b BlogUsecase) GetAll() (blogs []entity.Blog, err error) {
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

func (b BlogUsecase) CreateBlog(blog entity.Blog) (entity.Blog, error) {
	res, err := b.blogRepo.CreateBlog(blog)
	if err != nil {
		return res, err
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		return res, err
	}
	bulkLogJSON, _ := json.Marshal(map[string]string{
		"msg":  "User create blog.",
		"data": string(resJSON),
	})
	err = b.blogRepo.Publish("blog", bulkLogJSON)
	return res, err
}

func (b BlogUsecase) UpdateBlog(id string, blog entity.Blog) (entity.Blog, error) {
	res, err := b.blogRepo.UpdateBlog(id, blog)
	if err != nil {
		return res, err
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		return res, err
	}
	bulkLogJSON, _ := json.Marshal(map[string]string{
		"msg":  "User update blog.",
		"data": string(resJSON),
	})
	err = b.blogRepo.Publish("blog", bulkLogJSON)
	return res, err
}

func (b BlogUsecase) DeleteBlog(id string) (entity.Blog, error) {
	res, err := b.blogRepo.DeleteBlog(id)
	if err != nil {
		return res, err
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		return res, err
	}
	bulkLogJSON, _ := json.Marshal(map[string]string{
		"msg":  "User delete blog.",
		"data": string(resJSON),
	})
	err = b.blogRepo.Publish("blog", bulkLogJSON)
	return res, err
}

func NewBlogUsecase(blogRepo repo.BlogRepo) BlogUsecase {
	return BlogUsecase{
		blogRepo: blogRepo,
	}
}
