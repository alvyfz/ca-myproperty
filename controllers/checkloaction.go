package controller

import (
	openapi "ca-myproperty/thirdparties/ipapi"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckLocationByIP(c echo.Context) error {
	ip:=getIP(c.Request())
	fmt.Println(ip)
	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/json", ip), nil)
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


func getIP(req *http.Request) string{
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        panic(err)
    }

    userIP := net.ParseIP(ip)
    if userIP == nil {
        panic(err)
    }
  return ip
}
