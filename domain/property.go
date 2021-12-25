package domain

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Property Represents the properties in the DataBase
type Property struct {
	ID               bson.ObjectId `json:"id"`
	UserID           bson.ObjectId `json:"userid"`
	DateInscription  time.Time     `json:"date"`
	Description      string        `json:"description" form:"flongdesc" san:"max=200,trim,cap,xss"`
	ShortDescription string        `json:"shortdescription" form:"fshortdesc" san:"max=1000,trim,cap,xss"`
	Price            int           `json:"price" form:"fprice" san:"trim,xss"`
	Country          string        `json:"country" form:"fcountry" san:"max=500,trim,xss"`
	Region           string        `json:"region" form:"fregion" san:"max=50,trim,title,xss"`
	State            string        `json:"state" form:"fstate" san:"max=500,trim,title,xss"`
	Street           string        `json:"street" form:"fstreet" san:"max=500,trim,title,xss"`
	PostalCode       int           `json:"postal" form:"fpostal" san:"trim,xss"`
	City             string        `json:"city" form:"fcity" san:"max=50,trim,title,xss"`
	Type             PropertyType  `json:"type" form:"ftype" san:"max=15,trim,xss"`
	MetersBuilt      int           `json:"meters-built" form:"fbulit" san:"trim,xss"`
	MetersAvailable  int           `json:"meters-available" form:"futil" san:"trim,cap,xss"`
	HasEvelator      bool          `json:"elevator" form:"felevator[]"`
	Public           bool          `json:"public" form:"fpublic[]"`
	IsSale           bool          `json:"sale" form:"fsale[]"`
	Rooms            int           `json:"rooms" form:"frooms" san:"trim,xss"`
	Toilets          int           `json:"toilets" form:"ftoilet" san:"trim,xss"`
	Floor            int           `json:"floor" form:"ffloor" san:"max=5,trim,xss"`
	Garage           int           `json:"garage" form:"fgarage"  san:"trim,xss"`
	StorageRoom      int           `json:"storage-room" form:"fstorage" san:"trim,xss"`
	Photos           []string      `json:"photos"`
	Documents        []string      `json:"documents"`
	// User only for interface, no real data
	IsFavoriteForCurrentUser bool `json:"favourite"`
}

// PropertyType is the close list of different types of property that the users can upload
type PropertyType string

// The options for PropertyType
const (
	Chalet PropertyType = "Chalet"
	Flat   PropertyType = "Flat"
	House  PropertyType = "House"
	Garage PropertyType = "Garage"
)

// PropertyUsecase are the different Usecases that we can have regarding Property
type PropertyUsecase interface {
	CreateNewProperty()
	EditProperty()
	DeleteProperty()
}

type FeaturedProperties struct {
	ID       bson.ObjectId   `json:"id"`
	Featured []bson.ObjectId `json:"featured"`
}
