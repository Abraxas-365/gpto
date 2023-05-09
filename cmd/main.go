package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Abraxas-365/gpto/internal/application"
	"github.com/Abraxas-365/gpto/internal/infrastructure/controllers"
	"github.com/Abraxas-365/gpto/internal/infrastructure/database"
	"github.com/Abraxas-365/gpto/internal/infrastructure/repository"
	"github.com/Abraxas-365/gpto/pkg/summary"
	"github.com/Abraxas-365/gpto/pkg/utils"
	"github.com/Abraxas-365/gpto/pkg/visitor"
	"github.com/Abraxas-365/gpto/pkg/visitor/govisitor"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize the Golang parser

	ctx := context.Background()
	conn, err := database.NewConnection("localhost", 5432, "myuser", "mypassword", "mydb")
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.RunMigrations(ctx); err != nil {
		fmt.Println(err)
	}
	goVisitor := &govisitor.GoVisitor{}
	fnVisitor := visitor.New(goVisitor)
	nodes, err := utils.NewIndexer("test", *fnVisitor, summary.Create)
	if err != nil {
		fmt.Println("Error while walking the directory:", err)
	}
	repo := repository.NewMetadataRepository(conn.Pool)
	gpto, err := application.NewApp(repo)
	if err != nil {
		fmt.Println(err)
	}

	if err := gpto.SaveMetadataConcurrently(nodes); err != nil {
		fmt.Println(err)
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	controllers.ControllerFactory(app, *gpto)
	app.Listen(":8000")

}
