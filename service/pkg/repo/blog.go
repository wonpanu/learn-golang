package repo

import (
	"github.com/kamva/mgm/v3"
	"github.com/wonpanu/learn-golang/service/pkg/entity"
	"gopkg.in/mgo.v2/bson"
)

type BlogRepo interface {
	GetAll() ([]entity.Blog, error)
}

type MongoDB struct {
	coll *mgm.Collection
}

func (b MongoDB) GetAll() ([]entity.Blog, error) {
	blogs := []entity.Blog{}
	err := b.coll.SimpleFind(&blogs, bson.M{})
	return blogs, err
}

func NewBlogRepo(coll *mgm.Collection) BlogRepo {
	return MongoDB{
		coll: coll,
	}
}
