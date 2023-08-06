package requests

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashpect/obs-api/helper"
	"github.com/gofiber/fiber/v2"
)

func CredsHitApi(c *fiber.Ctx) error {
	apiURL := "https://api.opensuse.org/architectures"

	config, err := helper.ReadConfig("config.toml")
	if err != nil {
		fmt.Println("Error reading config:", err)
		return c.SendString("Error getting obs credentials")
	}
	fmt.Println("Username:", config.ObsCreds.Username)
	fmt.Println("Password:", config.ObsCreds.Password)

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return c.Status(500).SendString("Error making API request")
	}

	auth := config.ObsCreds.Username + ":" + config.ObsCreds.Password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)
	req.Header.Set("Accept", "application/xml")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return c.Status(500).SendString("Error making API request")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return c.Status(500).SendString("Error reading API response")
	}

	fmt.Println("API Response:", string(body))
	return c.SendString("Request sent to API. The response is : " + string(body) + "\n")

}
