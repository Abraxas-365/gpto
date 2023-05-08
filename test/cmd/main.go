package main

import (
	"database/sql"
	"encoding/json"
	"github/Abraxas-365/akeneo-connector/pkg/application"
	"github/Abraxas-365/akeneo-connector/pkg/core/service"
	"github/Abraxas-365/akeneo-connector/pkg/infrastructure/akeneo"
	"github/Abraxas-365/akeneo-connector/pkg/infrastructure/magento"
	"github/Abraxas-365/akeneo-connector/pkg/infrastructure/rest"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Config struct {
	MaentoUrl       string `json:"magento_url"`
	MagentoDbUri    string `json:"magento_db_uri"`
	MagentoUser     string `json:"magento_user"`
	MagentoPassword string `json:"magento_password"`
	AkeneoSecret    string `json:"akeneo_secret"`
	AkeneoUrl       string `json:"akeneo_url"`
	AkeneoUser      string `json:"akeneo_user"`
	AkeneoPassword  string `json:"akeneo_password"`
	AkeneoDb        string `json:"akeneo_db_uri"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}

	if err := json.Unmarshal([]byte(file), &config); err != nil {
		log.Fatal(err)
	}

	appF := fiber.New()
	appF.Use(cors.New())
	appF.Use(logger.New())
	dbMagento, err := sql.Open("mysql", config.MagentoDbUri)
	if err != nil {
		log.Fatal(err)
	}
	defer dbMagento.Close()

	dbAkeneo, err := sql.Open("mysql", config.AkeneoDb)
	if err != nil {
		log.Fatal(err)
	}
	defer dbMagento.Close()

	magento := magento.MagentoFactory(dbMagento, config.MaentoUrl, config.MagentoUser, config.MagentoPassword)
	akeneo := akeneo.AkeneoFactory(dbAkeneo, config.AkeneoSecret, config.AkeneoUser, config.AkeneoPassword, config.AkeneoUrl)
	service := service.ServiceFactory(magento, akeneo)
	app := application.ApplicationFactory(magento, akeneo, service)
	rest.ControllerFactory(appF, app)

	appF.Listen(":3000")

}
