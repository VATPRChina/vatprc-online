package main

import (
	"VatprcOnline/model"
	"VatprcOnline/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getVatprcStatus() *model.VatprcStatus {
	var vatsimRawData []byte
	cached, err := service.GetDataFromCache("VatsimData")
	if (cached != nil) && (err == nil) {
		log.Println("Cached!")
		vatsimRawData = cached
	} else {
		log.Println("Cache Missed!")
		vatsimRawData, err = service.FetchOnlineDataFromVatsim()
		go func() {
			service.PutDataToCache("VatsimData", vatsimRawData, 30)
		}()

		if err != nil {
			// VATSIM Down
			return nil
		}
	}
	vatsimResponse, err := service.ParseVatsimResponse(vatsimRawData)
	if err != nil {
		// Data error
		return nil
	}
	return vatsimResponse.MakeVatprcStatus()
}

func getVatprcFutureAtc() *[]model.FutureController {
	var vatprcFutureAtcRaw []byte
	cached, err := service.GetDataFromCache("VatprcFutureAtc")
	if (cached != nil) && (err == nil) {
		log.Println("Cached!")
		vatprcFutureAtcRaw = cached
	} else {
		log.Println("Cache Missed!")
		vatprcFutureAtcRaw, err = service.FetchFutureAtcFromVatprc()
		go func() {
			service.PutDataToCache("VatprcFutureAtc", vatprcFutureAtcRaw, 600)
		}()

		if err != nil {
			// VATPRC Down
			return nil
		}
	}

	var vatprcFutureAtc []model.AtcCenterFutureController
	err = json.Unmarshal(vatprcFutureAtcRaw, &vatprcFutureAtc)
	if err != nil {
		// VATPRC Data error
		return nil
	}

	var result []model.FutureController
	for _, elem := range vatprcFutureAtc {
		futureController := elem.MakeFutureController()
		if futureController != nil {
			result = append(result, *futureController)
		}
	}
	return &result
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORS())
	r.LoadHTMLGlob("templates/*.tmpl")
	r.NoRoute(func(c *gin.Context) {
		requestType := c.Query("type")
		switch requestType {
		case "raw":
			vatprcStatus := getVatprcStatus()
			vatprcStatus.FutureControllers = *getVatprcFutureAtc()
			if vatprcStatus == nil {
				c.JSON(500, "Something went wrong")
			} else {
				c.JSON(200, vatprcStatus)
			}
		case "atc":
			vatprcStatus := getVatprcStatus()
			vatprcStatus.FutureControllers = *getVatprcFutureAtc()
			c.HTML(http.StatusOK, "atc.tmpl", gin.H{
				"Controllers": vatprcStatus.Controllers,
				"FutureControllers": vatprcStatus.FutureControllers,
			})
		case "pilot":
			vatprcStatus := getVatprcStatus()
			c.HTML(http.StatusOK, "pilot.tmpl", gin.H{
				"Pilots": vatprcStatus.Pilots,
			})
		default:
			c.JSON(404, "Not Found")
		}
	})

	r.Run("127.0.0.1:9000")
}
