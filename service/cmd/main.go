package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	rest "github.com/wonpanu/learn-golang/service/pkg/handler"
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

	blogCollection := mgm.CollectionByName("blogs")
	blogRepo := repo.NewBlogRepo(blogCollection)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogHandler := rest.NewBlogHandler(blogUsecase)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(rest.Response{
			Status: "ok",
			Data:   "Welcome to Blog API!",
		})
	})
	app.Get("/blogs", blogHandler.GetAll)

	err = app.Listen(fmt.Sprintf(":%d", apiPort))
	if err != nil {
		log.Fatal(err)
	}
}
