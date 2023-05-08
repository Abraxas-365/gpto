package rest

import (
	"github/Abraxas-365/akeneo-connector/pkg/application"
	"github/Abraxas-365/akeneo-connector/pkg/core/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ControllerFactory(fiberApp *fiber.App, app application.Application) {
	r := fiberApp.Group("/api")

	r.Get("/attributes", func(c *fiber.Ctx) error {

		resp, err := app.SyncAttributes()
		if err != nil {
			return c.Status(500).JSON(err)
		}

		return c.Status(200).JSON(resp.ToAkeneoStruct())
	})

	r.Get("/family", func(c *fiber.Ctx) error {

		resp, err := app.SyncFamilies()
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(resp)
	})

	r.Get("/category/isNew=:isNew<bool>", func(c *fiber.Ctx) error {
		isNew := false
		isNewParam := c.Params("isNew", "false")
		if isNewParam == "true" {
			isNew = true
		}

		err := app.SyncCategories(isNew)

		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.SendStatus(200)
	})

	r.Get("/options", func(c *fiber.Ctx) error {

		options, err := app.SyncOptions()

		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(options)
	})

	r.Get("/groups", func(c *fiber.Ctx) error {

		resp, err := app.SyncAttributeGroups()
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(resp)
	})

	r.Get("/family/id=:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id", "0")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.Status(500).JSON(err)
		}

		resp, err := app.GetFamilyNameWithId(id)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(resp)
	})

	r.Get("/product/id=:id/:locale?/:scope?", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		log.Println(id)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		localParam := c.Params("locale")
		scopeParam := c.Params("scope")
		var local *string
		var scope *string
		local = &localParam
		scope = &scopeParam
		if localParam == "" {
			local = nil
		}

		if scopeParam == "" {
			scope = nil
		}

		resp, err := app.SyncProductById(models.NewProductAkeneoBuilder().AddLocale(local).AddScope(scope), id)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(resp)
	})

	r.Get("/product/page=:page/:locale?/:scope?", func(c *fiber.Ctx) error {
		pageParam := c.Params("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		localParam := c.Params("locale")
		scopeParam := c.Params("scope")
		var local *string
		var scope *string
		local = &localParam
		scope = &scopeParam
		if localParam == "" {
			local = nil
		}

		if scopeParam == "" {
			scope = nil
		}

		resp, err := app.SyncProducts(models.NewProductAkeneoBuilder().AddLocale(local).AddScope(scope), page)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(resp)
	})

	r.Get("/product/code=:code", func(c *fiber.Ctx) error {
		codeParam := c.Params("code")
		code, err := strconv.Atoi(codeParam)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		product, err := app.ConectorProduct(code)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(product)
	})
}
