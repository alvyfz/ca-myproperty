package controller

import (
	openapi "ca-myproperty/thirdparties/ipapi"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckLocationByIP(c echo.Context) error {

	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Your location",
			"data":    err.Error(),
		})
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := ipapiClient.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Your location",
			"data":    err.Error(),
		})
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Your location",
			"data":    err.Error(),
		})
	}
	var res openapi.Ipapi
	err = json.Unmarshal(body, &res)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Your location",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your location",
		"data":    res,
	})
}
