package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"github.com/streadway/amqp"
	"github.com/wonpanu/learn-golang/service/pkg/entity"
	"github.com/wonpanu/learn-golang/service/pkg/util"
	"gopkg.in/mgo.v2/bson"
)

type BlogAdapter struct {
	coll          *mgm.Collection
	amqpCh        *amqp.Channel
	blogQueueName map[string]string
}

type BlogRepo interface {
	GetAll() ([]entity.Blog, error)
	CreateBlog(blog entity.Blog) (entity.Blog, error)
	UpdateBlog(id string, blog entity.Blog) (entity.Blog, error)
	DeleteBlog(id string) (entity.Blog, error)
	Publish(routingKey string, payload []byte) error
}

func (b BlogAdapter) GetAll() ([]entity.Blog, error) {
	blogs := []entity.Blog{}
	err := b.coll.SimpleFind(&blogs, bson.M{
		"is_show": bson.M{
			operator.Ne: false,
		},
	})
	return blogs, err
}

func (b BlogAdapter) CreateBlog(blog entity.Blog) (entity.Blog, error) {
	seed := time.Now().Unix()
	blog.ID = util.Hash(fmt.Sprint(seed))
	blog.IsShow = true
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	_, err := b.coll.InsertOne(ctx, blog)
	if err != nil {
		return entity.Blog{}, err
	}
	return blog, err
}

func (b BlogAdapter) UpdateBlog(id string, blog entity.Blog) (entity.Blog, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": id}
	blog.ID = id
	blog.UpdatedAt = time.Now()
	update := bson.M{
		"$set": blog,
	}
	_, err := b.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (b BlogAdapter) DeleteBlog(id string) (entity.Blog, error) {
	var blog entity.Blog
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": id}
	err := b.coll.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		return blog, err
	}

	blog.ID = id
	blog.UpdatedAt = time.Now()
	blog.IsShow = false
	update := bson.M{
		"$set": blog,
	}
	_, err = b.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func (b BlogAdapter) Publish(routingKey string, payload []byte) error {
	return b.amqpCh.Publish(
		"",                          // exchange
		b.blogQueueName[routingKey], // routing key
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         payload,
		},
	)
}

func NewBlogRepo(coll *mgm.Collection, amqpCh *amqp.Channel, blogQueueName map[string]string) BlogRepo {
	return BlogAdapter{
		coll:          coll,
		amqpCh:        amqpCh,
		blogQueueName: blogQueueName,
	}
}
