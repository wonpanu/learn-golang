package repo

import (
	"github.com/kamva/mgm/v3"
	"github.com/wonpanu/learn-golang/pkg/entity"
	"gopkg.in/mgo.v2/bson"
)

type BlogRepo struct {
	coll *mgm.Collection
}

func (b *BlogRepo) GetAll() ([]entity.Blog, error) {
	blogs := []entity.Blog{}
	err := b.coll.SimpleFind(&blogs, bson.M{})
	return blogs, err
}

func NewBlogRepo(coll *mgm.Collection) BlogRepo {
	return BlogRepo{
		coll: coll,
	}
}
