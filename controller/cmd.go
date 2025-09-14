package controller

import (
	"errors"

	"github.com/EmreZURNACI/url-shortener/app/shortener"
	"github.com/EmreZURNACI/url-shortener/domain"
	"github.com/EmreZURNACI/url-shortener/infra"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func (h *Handler) Shortener(c *fiber.Ctx) error {

	var address domain.Address

	if err := c.BodyParser(&address); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": ErrInvalidRequestBody.Error(),
		})
	}

	if err := validate.Struct(&address); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	GetURL := shortener.NewGetURLHandler(h.Repository)
	res, err := GetURL.Handle(c.UserContext(), &shortener.GetURLRequest{
		Address: address,
	})
	if err != nil {
		if errors.Is(err, infra.ErrRecordNotFound) {
			createURLHandler := shortener.NewCreateURLHandler(h.Repository)
			res2, err := createURLHandler.Handle(c.UserContext(), &shortener.CreateURLRequest{
				Address: address,
			})
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			return c.Status(200).JSON(fiber.Map{
				"link": res2.ShortURL,
			})
		}
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"link": res.Address.ShortURL,
	})

}

func (h *Handler) Redirect(c *fiber.Ctx) error {

	var link string = c.Params("link")

	getShortURLHandler := shortener.NewGetShortURLHandler(h.Repository)
	res, err := getShortURLHandler.Handle(c.UserContext(), &shortener.GetShortURLRequest{
		Address: domain.Address{
			ShortURL: link,
		},
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(301).Redirect(res.Address.URL)

}
