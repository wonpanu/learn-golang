package repo

import (
	"github.com/kamva/mgm/v3"
	"github.com/streadway/amqp"
	"github.com/wonpanu/learn-golang/service/pkg/entity"
	"gopkg.in/mgo.v2/bson"
)

type BlogAdapter struct {
	coll          *mgm.Collection
	amqpCh        *amqp.Channel
	blogQueueName map[string]string
}

type BlogRepo interface {
	GetAll() ([]entity.Blog, error)
	Publish(routingKey string, payload []byte) error
}

func (b BlogAdapter) GetAll() ([]entity.Blog, error) {
	blogs := []entity.Blog{}
	err := b.coll.SimpleFind(&blogs, bson.M{})
	return blogs, err
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
