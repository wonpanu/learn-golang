package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"github.com/streadway/amqp"
	"github.com/wonpanu/learn-golang/amqputil"
	"github.com/wonpanu/learn-golang/service/pkg/handler"
	"github.com/wonpanu/learn-golang/service/pkg/repo"
	"github.com/wonpanu/learn-golang/service/pkg/usecase"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Welcome to Blog API!")
	godotenv.Load(".env")

	apiPortRaw := os.Getenv("API_PORT")
	apiPort, err := strconv.Atoi(apiPortRaw)
	if err != nil {
		log.Fatal(err)
	}

	mongoHost := os.Getenv("MONGO_HOST")
	mongoPortRaw := os.Getenv("MONGO_PORT")
	mongoDB := os.Getenv("MONGO_DB")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASS")
	mongoPort, _ := strconv.Atoi(mongoPortRaw)

	mongoURI := ""
	if mongoPort == 0 {
		mongoURI = fmt.Sprintf("mongodb+srv://%s:%s@%s", mongoUser, mongoPass, mongoHost)
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%d", mongoUser, mongoPass, mongoHost, mongoPort)
	}
	fmt.Println("mongoDB URI:", mongoURI)

	err = mgm.SetDefaultConfig(nil, mongoDB, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	amqpHost := os.Getenv("AMQP_HOST")
	amqpPortRaw := os.Getenv("AMQP_PORT")
	amqpUsername := os.Getenv("AMQP_USERNAME")
	amqpPassword := os.Getenv("AMQP_PASSWORD")
	blogQueueName := os.Getenv("LOG_API_QUEUE_NAME")
	amqpQueueName := map[string]string{
		"blog": blogQueueName,
	}

	amqpPort, err := strconv.Atoi(amqpPortRaw)
	if err != nil {
		log.Fatal(err)
	}

	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", amqpUsername, amqpPassword, amqpHost, amqpPort)
	fmt.Println("rabbitMQ URI:", amqpURI)

	amqpConn, amqpCh, _ := amqputil.CreatePublisherConnection(amqpURI, blogQueueName)
	amqpCloseNotify := amqpConn.NotifyClose(make(chan *amqp.Error))
	defer func() {
		amqpCh.Close()
		amqpConn.Close()
	}()

	go func() {
		for err := range amqpCloseNotify {
			log.Println("Rabbit MQ connection lost", err)
			os.Exit(1)
		}
	}()

	bahRamRepo := repo.NewBahRamRepo()
	bahRamUsecase := usecase.NewBahRamUsecase(bahRamRepo)
	bahRamHandler := handler.NewBahRamHandler(bahRamUsecase)

	blogCollection := mgm.CollectionByName("blogs")
	blogRepo := repo.NewBlogRepo(blogCollection, amqpCh, amqpQueueName)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogHandler := handler.NewBlogHandler(blogUsecase)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(handler.Response{
			Status: "ok",
			Data:   "Welcome to Blog API!",
		})
	})

	// Bah Ram
	app.Get("/bahram/:n", bahRamHandler.BahRam)

	// Blog API
	app.Get("/blogs", blogHandler.GetAll)
	app.Post("/create-blog", blogHandler.CreateBlog)
	app.Post("/update-blog/:id", blogHandler.UpdateBlog)
	app.Post("/delete-blogs/:id", blogHandler.DeleteBlog)

	err = app.Listen(fmt.Sprintf(":%d", apiPort))
	if err != nil {
		log.Fatal(err)
	}
}
