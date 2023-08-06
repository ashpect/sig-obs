package controllers

import (
	requests "github.com/ashpect/obs-api/obs_requests"
	"github.com/gofiber/fiber/v2"
)

func CallBasicApi(c *fiber.Ctx) error {
	return requests.BasicApiHit(c)
}

func CredsHitApi(c *fiber.Ctx) error {
	return requests.CredsHitApi(c)
}
