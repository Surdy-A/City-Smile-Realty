package handlers

import (
	"hollister/models"
	"hollister/repo"
	"hollister/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProperty(c *gin.Context) {
	property_name := c.PostForm("property_name")
	city := c.PostForm("city")
	address := c.PostForm("address")
	bedroom := c.PostForm("bedroom")
	bathroom := c.PostForm("bathroom")
	washroom := c.PostForm("washroom")
	carspace := c.PostForm("car_space")
	room_area := c.PostForm("room_area")
	price := c.PostForm("price")
	availability := c.PostForm("availability")
	status := c.PostForm("status")
	//images := c.PostForm("images")
	description := c.PostForm("description")
	agent := c.PostForm("agent")
	video := c.PostForm("video")
	category := c.PostForm("category")
	property_type := c.PostForm("property_type")
	state := c.PostForm("state")
	postal_code := c.PostForm("postal_code")

	var houseFeature models.Features
	c.ShouldBind(&houseFeature)

	uploadImages, err := utils.UploadImages(c)
	if err != nil {
		log.Fatal("Unable to Upload Image", err)
	}

	bedroomInt, err := strconv.ParseInt(bedroom, 10, 64)
	if err != nil {
		log.Println("Conversion Error 1", err)
	}

	bathroomInt, err := strconv.ParseInt(bathroom, 10, 64)
	if err != nil {
		log.Println("Conversion Error 2", err)
	}

	washroomInt, err := strconv.ParseInt(washroom, 10, 64)
	if err != nil {
		log.Println("Conversion Error 3", err)
	}

	carspaceInt, err := strconv.ParseInt(carspace, 10, 64)
	if err != nil {
		log.Println("Conversion Error 4", err)
	}

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Conversion Error 5", err)
	}

	h := models.House{
		PropertyName: property_name,
		City:         city,
		Address:      address,
		Bedroom:      bedroomInt,
		Bathroom:     bathroomInt,
		Washroom:     washroomInt,
		CarSpace:     carspaceInt,
		RoomArea:     room_area,
		Price:        priceFloat,
		Availability: availability,
		Status:       status,
		Images:       uploadImages,
		Description:  description,
		Agent:        agent,
		Video:        video,
		Category:     category,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now(),
		PropertyType: property_type,
		State:        state,
		PostalCode:   postal_code,
		Features:     houseFeature,
	}

	//var db *sql.DB
	err = repo.CreateProperty(h)
	if err != nil {
		log.Fatal("Unable to create property:", err)
	}

	c.HTML(http.StatusOK, "add_rent_property.tmpl", gin.H{
		"title": "hh",
	})

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
}

func GetProperties(c *gin.Context) {
	var hs models.House
	properties, err := repo.GetProperties(hs)

	if err != nil {
		log.Fatal("Unable to retreive property: ", err)
	}

	c.HTML(http.StatusOK, "rent_property.tmpl", gin.H{
		"properties": properties,
	})
}

func FilterProperty(c *gin.Context) {
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
}

func GetFAQ(c *gin.Context) {
	c.HTML(http.StatusOK, "faq.tmpl", gin.H{})
}

func ContactUs(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.tmpl", gin.H{})
}

func AboutUs(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{})
}

// Buy Handlers
func GetPropertiesForSale(c *gin.Context) {
	var hs models.House
	properties, err := repo.GetProperties(hs)

	if err != nil {
		log.Fatal("Unable to retreive property: ", err)
	}

	c.HTML(http.StatusOK, "buy_property.tmpl", gin.H{
		"properties": properties,
	})
}

func AddPropertyForRent(c *gin.Context) {

	c.HTML(http.StatusOK, "add_rent_property.tmpl", gin.H{})

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
}

func AddPropertyForSale(c *gin.Context) {

	c.HTML(http.StatusOK, "add_sell_property.tmpl", gin.H{})

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
}

// // Buy Handlers
// func EditPropertyForSale(c *gin.Context) {
// 	id := c.Param("id")
// 	property_name := c.PostForm("property_name")
// 	city := c.PostForm("city")
// 	address := c.PostForm("address")
// 	bedroom := c.PostForm("bedroom")
// 	bathroom := c.PostForm("bathroom")
// 	washroom := c.PostForm("washroom")
// 	carspace := c.PostForm("car_space")
// 	room_area := c.PostForm("room_area")
// 	price := c.PostForm("price")
// 	availability := c.PostForm("availability")
// 	status := c.PostForm("status")
// 	images := c.PostForm("images")
// 	description := c.PostForm("description")
// 	agent := c.PostForm("agent")
// 	video := c.PostForm("video")
// 	category := c.PostForm("category")
// 	property_type := c.PostForm("property_type")
// 	state := c.PostForm("state")
// 	postal_code := c.PostForm("postal_code")
// 	fmt.Println("------------")
// 	fmt.Println(property_name, city, address, bedroom, bathroom, washroom, carspace, room_area, price, availability,
// 		status, images, description, agent, video, category, property_type, state, postal_code)

// 	var houseFeature models.Features
// 	c.ShouldBind(&houseFeature)
// 	fmt.Println("_______________________")
// 	fmt.Println(houseFeature.Features)
// 	fmt.Println(houseFeature)

// 	uploadImages, err := utils.UploadImages(c)
// 	if err != nil {
// 		log.Fatal("Unable to Upload Image", err)
// 	}

// 	// bedroomInt, err := strconv.ParseInt(bedroom, 10, 64)
// 	// if err != nil {
// 	// 	log.Println("Conversion Error 1", err)
// 	// }

// 	// bathroomInt, err := strconv.ParseInt(bathroom, 10, 64)
// 	// if err != nil {
// 	// 	log.Println("Conversion Error 2", err)
// 	// }

// 	// washroomInt, err := strconv.ParseInt(washroom, 10, 64)
// 	// if err != nil {
// 	// 	log.Println("Conversion Error 3", err)
// 	// }

// 	// carspaceInt, err := strconv.ParseInt(carspace, 10, 64)
// 	// if err != nil {
// 	// 	log.Println("Conversion Error 4", err)
// 	// }

// 	// priceFloat, err := strconv.ParseFloat(price, 64)
// 	// if err != nil {
// 	// 	log.Println("Conversion Error 5", err)
// 	// }

// 	h := models.House{
// 		PropertyName: property_name,
// 		City:         city,
// 		Address:      address,
// 		// Bedroom:      bedroomInt,
// 		// Bathroom:     bathroomInt,
// 		// Washroom:     washroomInt,
// 		// CarSpace:     carspaceInt,
// 		// RoomArea:     room_area,
// 		// Price:        priceFloat,
// 		Availability: availability,
// 		Status:       status,
// 		Images:       uploadImages,
// 		Description:  description,
// 		Agent:        agent,
// 		Video:        video,
// 		Category:     category,
// 		CreatedDate:  time.Now(),
// 		UpdatedDate:  time.Now(),
// 		PropertyType: property_type,
// 		State:        state,
// 		PostalCode:   postal_code,
// 		Features:     houseFeature,
// 	}
// 	properties, err := repo.EditPropertyForSale(id, h)

// 	if err != nil {
// 		log.Fatal("Unable to retreive property: ", err)
// 	}

// 	c.HTML(http.StatusOK, "edit_property.tmpl", gin.H{
// 		"properties": properties,
// 	})
// }

func GetEditPropertyForSale(c *gin.Context) {

	c.HTML(http.StatusOK, "update_rent_property.tmpl", gin.H{})

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
}

func DeleteProperty(c *gin.Context) {
	id := c.Param("id")
	var h models.House
	_, err := repo.DeleteProperty(id, h)

	if err != nil {
		log.Fatal("Unable to retreive property: ", err)
	}

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Host)
}

func DeleteVuyProperty(c *gin.Context) {
	id := c.Param("id")
	var h models.House
	_, err := repo.DeleteProperty(id, h)

	if err != nil {
		log.Fatal("Unable to retreive property: ", err)
	}

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Host)
}
