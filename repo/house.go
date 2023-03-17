package repo

import (
	"database/sql"
	"fmt"
	"hollister/models"
	"log"

	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

var db *sql.DB
var err error

func ConnectToDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	var psqlInfo = user + "://" + user + ":" + password + "@" + host + "/" + dbname + "?sslmode=disable"
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("succesfuly connected to database")
}

func CreateProperty(h models.House) error {
	ConnectToDB()
	insertStmt := `INSERT INTO house_rent(propertyname, city, address, bedroom, bathroom, washroom,
				   carspace, price, availability, status, images, description, agent, video, category,
				   createddate, updateddate, propertytype, state, postalcode, features) 
				   VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)`

	_, err := db.Exec(insertStmt, h.PropertyName, h.City, h.Address, h.Bedroom, h.Bathroom, h.Washroom,
		h.CarSpace, h.Price, h.Availability, h.Status, pq.Array(h.Images), h.Description, h.Agent, h.Video,
		h.Category, h.CreatedDate, h.UpdatedDate, h.PropertyType, h.State, h.PostalCode, pq.Array(h.Features.Features))

	if err != nil {
		return err
	}

	return nil
}

func GetProperties(h models.House) ([]models.House, error) {
	ConnectToDB()

	var properties []models.House
	selectStmt := `SELECT id, propertyname, city, address, bedroom, bathroom, washroom,
	carspace, price, availability, status, images, description, agent, video, category,
	createddate, updateddate, propertytype, state, postalcode, features
	FROM house_rent`

	rows, err := db.Query(selectStmt)
	if err != nil {
		log.Fatal("Unable to retreive property: ", err)
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&h.ID, &h.PropertyName, &h.City, &h.Address, &h.Bedroom, &h.Bathroom, &h.Washroom,
			&h.CarSpace, &h.Price, &h.Availability, &h.Status, pq.Array(&h.Images), &h.Description, &h.Agent, &h.Video,
			&h.Category, &h.CreatedDate, &h.UpdatedDate, &h.PropertyType, &h.State, &h.PostalCode, pq.Array(&h.Features.Features))

		if err != nil {
			fmt.Println("Unable to retreive property: ", err)
		}

		properties = append(properties, h)
	}

	return properties, nil
}

func GetProperty(id string, h models.House) (models.House, error) {
	ConnectToDB()

	var property models.House
	selectStmt := `SELECT id, propertyname, city, address, bedroom, bathroom, washroom,
	carspace, price, availability, status, images, description, agent, video, category,
	createddate, updateddate, propertytype, state, postalcode, features
	FROM house_rent WHERE id =$1;`

	row, err := db.Query(selectStmt, 7)
	if err != nil {
		log.Fatal("Unable to retreive property: ", err)
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&h.ID, &h.PropertyName, &h.City, &h.Address, &h.Bedroom, &h.Bathroom, &h.Washroom,
			&h.CarSpace, &h.Price, &h.Availability, &h.Status, pq.Array(&h.Images), &h.Description, &h.Agent, &h.Video,
			&h.Category, &h.CreatedDate, &h.UpdatedDate, &h.PropertyType, &h.State, &h.PostalCode,
			pq.Array(&h.Features))

		if err != nil {
			log.Fatal(err)
		}
		property = h
	}

	return property, nil
}

func GetPropertyById(id string, h models.House) (models.House, error) {
	ConnectToDB()

	var property models.House
	selectStmt := `SELECT id, propertyname, city, address, bedroom, bathroom, washroom,
	carspace, price, availability, status, images, description, agent, video, category,
	createddate, updateddate, propertytype, state, postalcode, features
	FROM house_rent WHERE id =$1;`

	err = db.QueryRow(selectStmt, id).Scan(&h.ID, &h.PropertyName, &h.City, &h.Address, &h.Bedroom, &h.Bathroom, &h.Washroom,
		&h.CarSpace, &h.Price, &h.Availability, &h.Status, pq.Array(&h.Images), &h.Description, &h.Agent, &h.Video,
		&h.Category, &h.CreatedDate, &h.UpdatedDate, &h.PropertyType, &h.State, &h.PostalCode, pq.Array(&h.Features.Features))

	if err != nil {
		log.Fatal("Unable to retreive property: ", err)
	}

	property = h
	return property, nil
}

// func EditPropertyForSale(id string, h models.House) (models.House, error) {
// 	ConnectToDB()

// 	// var property models.House
// 	// updateStmt := `UPDATE house_rent set propertyname = '$1', city = '$2', address = '$3',  price = '$8', availability = '$9', status = '$10',
// 	// images = '$11', description = '$12', agent = '$13', video = '$14', category = '$15', updateddate = '$16',
// 	// propertytype = '$17', state = '$18', postalcode = '$19', features = '$20', category = '$21', status = '$22'
// 	// WHERE id = 1`

// 	// updatedProperty, err := db.Exec(updateStmt, id, &h.PropertyName, &h.City, &h.Address, &h.Bedroom, &h.Bathroom, &h.Washroom,
// 	// 	&h.CarSpace, &h.Price, &h.Availability, &h.Status, pq.Array(&h.Images), &h.Description, &h.Agent, &h.Video,
// 	// 	&h.Category, &h.CreatedDate, &h.UpdatedDate, &h.PropertyType, &h.State, &h.PostalCode, pq.Array(&h.Features.Features))

// 	var property models.House
// 	updateStmt := `UPDATE house_rent set propertyname = '$1' WHERE id = 1`

// 	updatedProperty, err := db.Exec(updateStmt, id, &h.PropertyName)

// 	if err != nil {
// 		log.Fatal("Unable to retreive property: ", err)
// 	}

// 	fmt.Println(updatedProperty)
// 	property = h
// 	return property, nil
// }

func DeleteProperty(id string, h models.House) (models.House, error) {
	ConnectToDB()

	var property models.House
	deleteStmt := `DELETE FROM house_rent WHERE id =$1;`

	_, err = db.Exec(deleteStmt, id)

	if err != nil {
		log.Fatal("Unable to retreive property: nnnn", err)
	}

	property = h
	return property, nil
}
