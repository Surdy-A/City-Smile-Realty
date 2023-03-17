package main

import (
	"hollister/handlers"
	"hollister/models"
	"hollister/repo"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		t := time.Now()
		day := t.Day()
		month := t.Month()
		year := t.Year()
		weekday := t.Weekday()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Main website",
			"day":     day,
			"month":   month,
			"year":    year,
			"weekday": weekday,
		})
	})

	router.GET("/rent_property", handlers.GetProperties)

	router.POST("/rent_property", handlers.FilterProperty)

	router.GET("/rent_property_:id", func(c *gin.Context) {

		id := c.Param("id")
		var h models.House
		property, err := repo.GetPropertyById(id, h)
		if err != nil {
			log.Fatal("Unable to retreive property: ", err)
		}

		c.HTML(http.StatusOK, "rent_property_detail.tmpl", gin.H{
			"title":    "hh",
			"id":       id,
			"property": property,
		})
	})

	router.POST("/submit_rent_property", handlers.CreateProperty)

	router.GET("/faq", handlers.GetFAQ)
	router.GET("/contact", handlers.ContactUs)

	//buy handlers
	router.GET("/buy_property", handlers.GetPropertiesForSale)

	router.POST("/buy_property", func(c *gin.Context) {
		var hs models.House

		name := c.PostForm("name")
		city := c.PostForm("city")
		property_id := c.PostForm("property_id")
		address := c.PostForm("address")
		bedroom := c.PostForm("bedroom")
		washroom := c.PostForm("washroom")
		carspace := c.PostForm("carspace")
		price := c.PostForm("price")

		user_bedroom, _ := strconv.ParseInt(bedroom, 10, 64)
		user_washroom, _ := strconv.ParseInt(washroom, 10, 64)
		user_carspace, _ := strconv.ParseInt(carspace, 10, 64)
		user_price, _ := strconv.ParseFloat(price, 64)

		var houses []models.House
		var h models.House

		var errorMessage string
		properties, err := repo.GetProperties(h)

		if err != nil {
			log.Fatal("Unable to retreive property: ", err)
		}

		for _, property := range properties {
			if property.PropertyName == name || property.City == city || property.ID == property_id ||
				property.Address == address || property.Bedroom == user_bedroom || property.Washroom == user_washroom ||
				property.CarSpace == user_carspace || property.Price == user_price {

				houses = properties
			}
		}

		if hs.PropertyName == "" {
			errorMessage = "No Record Found"

		}

		c.HTML(http.StatusOK, "rent_property.tmpl", gin.H{
			"title":      "Main website",
			"name":       name,
			"properties": houses,
			"message":    errorMessage,
		})

		c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
	})

	router.GET("/buy_property_:id", func(c *gin.Context) {

		id := c.Param("id")
		var h models.House
		property, err := repo.GetPropertyById(id, h)
		if err != nil {
			log.Fatal("Unable to retreive property: ", err)
		}

		c.HTML(http.StatusOK, "rent_property_detail.tmpl", gin.H{
			"title":    "hh",
			"id":       id,
			"property": property,
		})
	})

	router.GET("/submit_sell_property", handlers.AddPropertyForSale)
	router.GET("/submit_rent_property", handlers.AddPropertyForRent)
	router.GET("/about", handlers.AboutUs)
	// router.POST("/update_property_:id", handlers.EditPropertyForSale)
	// router.GET("/update_property_:id", handlers.GetEditPropertyForSale)
	router.DELETE("/delete_property_:id", handlers.DeleteProperty)
	router.GET("/delete_property_:id", handlers.DeleteProperty)

	defer repo.ConnectToDB()
	router.Run()
}
