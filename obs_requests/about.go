package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func BasicApiHit(c *fiber.Ctx) error {
	apiURL := "https://api.opensuse.org/about"

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return c.Status(500).SendString("Error making API request")
	}

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
