package repositories

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/kiketordera/basic-gin-examples/domain"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionUsers              string = "users"
	collectionProperties         string = "properties"
	collectionPropertiesFeatured string = "properties-featured"
	projectID                           = "inmo-b3610"
)

// BoltRepository implements models.Repository with Bolt DataBse
type FirestoreRepository struct {
	// DataBase *bolthold.Store
}

// // InitDatabase creates the instance of the BoltHold Database
// func InitFirestoreDatabase() _domain.IRepository {

// 	return FirestoreRepository{
// 		// DataBase: db,
// 	}
// }

func Init() (*firestore.Client, context.Context) {
	// Init of the firestore
	ctx := context.Background()
	sa := option.WithCredentialsFile("/Users/enriquetorderaramos/go/src/github.com/kiketordera/basic-gin-examples/repositories/firebase-services.json")
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("antes")
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("despues")
	return client, ctx
}

func Init2() (*firestore.Client, context.Context) {
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("antes")
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("despues")
	return client, ctx
}

// GetUserByID finds and returns the user by his unique ID
func GetUserByID(id bson.ObjectId) (domain.User, error) {
	client, ctx := Init()
	defer client.Close()

	// The actual operation we want to make
	dsnap, err := client.Collection(collectionUsers).Doc(id.Hex()).Get(ctx)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return domain.User{}, err
	}
	var u domain.User
	dsnap.DataTo(&u)

	return u, nil
}

func GetUserByMail(email string) (domain.User, bool) {
	client, ctx := Init()
	defer client.Close()

	var user domain.User
	var users []domain.User
	query := client.Collection(collectionUsers).Where("Email", "==", email).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return domain.User{}, false
		}

		if err := doc.DataTo(&user); err != nil {
			// Handle error, possibly by returning the error
			// to the caller. Continue the loop,
			// break the loop or return.
			fmt.Println("ERROR: ", err)
			return domain.User{}, false
		}
		users = append(users, user)
	}
	return users[0], true
}

func GetAllCustomers() []domain.User {
	client, ctx := Init()
	defer client.Close()

	var user domain.User
	var users []domain.User
	query := client.Collection(collectionUsers).Where("Role", "==", domain.Customer).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.User{}
		}

		if err := doc.DataTo(&user); err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.User{}
		}
		users = append(users, user)
	}
	return users
}

func GetAllUsersNotCurtomer() []domain.User {
	client, ctx := Init()
	defer client.Close()

	var user domain.User
	var users []domain.User
	query := client.Collection(collectionUsers).Where("Role", "!=", domain.Customer).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.User{}
		}

		if err := doc.DataTo(&user); err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.User{}
		}
		users = append(users, user)
	}
	return users
}

func GetSellPropertiesLanding() []domain.Property {
	client, ctx := Init()
	defer client.Close()

	var property domain.Property
	var properties []domain.Property
	query := client.Collection(collectionProperties).Limit(20).Where("IsSale", "==", true).Where("Public", "==", true).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}

		if err := doc.DataTo(&property); err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}
		properties = append(properties, property)
	}
	return properties
}

func GetRentPropertiesLanding() []domain.Property {
	client, ctx := Init()
	defer client.Close()

	var property domain.Property
	var properties []domain.Property
	query := client.Collection(collectionProperties).Limit(20).Where("IsSale", "==", false).Where("Public", "==", true).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}

		if err := doc.DataTo(&property); err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}
		properties = append(properties, property)
	}
	return properties
}

func GetPropertiesActive(id bson.ObjectId) []domain.Property {
	client, ctx := Init()
	defer client.Close()

	var property domain.Property
	var properties []domain.Property
	query := client.Collection(collectionProperties).Limit(20).Where("Public", "==", true).Where("UserID", "==", id.Hex()).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}

		if err := doc.DataTo(&property); err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}
		properties = append(properties, property)
	}
	return properties
}
func GetPropertiesInactive(id bson.ObjectId) []domain.Property {
	client, ctx := Init()
	defer client.Close()

	var property domain.Property
	var properties []domain.Property
	query := client.Collection(collectionProperties).Limit(20).Where("Public", "==", false).Where("UserID", "==", id.Hex()).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}

		if err := doc.DataTo(&property); err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}
		properties = append(properties, property)
	}
	return properties
}
func GetPropertiesFromUser(id bson.ObjectId) []domain.Property {
	client, ctx := Init()
	defer client.Close()

	var property domain.Property
	var properties []domain.Property
	query := client.Collection(collectionProperties).Limit(20).Where("UserID", "==", id.Hex()).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}

		if err := doc.DataTo(&property); err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.Property{}
		}
		properties = append(properties, property)
	}
	return properties
}

func GetAllUsers() []domain.User {
	client, ctx := Init()
	defer client.Close()

	var user domain.User
	var users []domain.User
	iter := client.Collection(collectionUsers).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return []domain.User{}
		}
		if err := doc.DataTo(&user); err != nil {
			// Handle error, possibly by returning the error
			// to the caller. Continue the loop,
			// break the loop or return.
			fmt.Println("ERROR: ", err)
			return []domain.User{}
		}
		users = append(users, user)
	}
	return users
}

// SaveObject saves the element in the DataBase
func SaveObject(object interface{}, id bson.ObjectId) error {
	client, ctx := Init()
	defer client.Close()

	// The actual operation we want to make
	switch v := object.(type) {
	case domain.User:
		u := object.(domain.User)
		_, err := client.Collection(collectionUsers).Doc(id.Hex()).Set(ctx, u)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return err
		}
	case domain.Property:
		p := object.(domain.Property)
		_, err := client.Collection(collectionProperties).Doc(id.Hex()).Set(ctx, p)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return err
		}
	case domain.FeaturedProperties:
		fp := object.(domain.FeaturedProperties)
		_, err := client.Collection(collectionPropertiesFeatured).Doc(id.Hex()).Set(ctx, fp)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return err
		}
	default:
		fmt.Println("Struct not handled")
		fmt.Println("This is v: ", v)
	}
	return nil
}

// GetUserByID finds and returns the user by his unique ID
func GetPropertyByID(id bson.ObjectId) (domain.Property, error) {
	client, ctx := Init()
	defer client.Close()

	// The actual operation we want to make
	dsnap, err := client.Collection(collectionProperties).Doc(id.Hex()).Get(ctx)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return domain.Property{}, err
	}
	var p domain.Property
	dsnap.DataTo(&p)

	return p, nil
}

// GetUserByID finds and returns the user by his unique ID
func GetUserFomCookie(c *gin.Context) domain.User {
	// We get here the username from the Cookie
	u, e := c.Get("username")
	if e {
		user, _ := GetUserByMail(u.(string))
		return user
	}
	return domain.User{}
}

func GetFeatured() domain.FeaturedProperties {
	client, ctx := Init()
	defer client.Close()

	var user domain.FeaturedProperties
	var users []domain.FeaturedProperties
	iter := client.Collection(collectionPropertiesFeatured).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err)
			return domain.FeaturedProperties{}
		}
		if err := doc.DataTo(&user); err != nil {
			// Handle error, possibly by returning the error
			// to the caller. Continue the loop,
			// break the loop or return.
			fmt.Println("ERROR: ", err)
			return domain.FeaturedProperties{}
		}
		users = append(users, user)
	}
	return users[0]
}
func GetFeaturedProperties() []domain.Property {
	feat := GetFeatured()
	return GetFavoritePropertiesFromSlice(feat.Featured)
}

// GetSellPropertiesLanding gets all the properties for renting for the landing page from the DataBase
func GetFavoritePropertiesFromSlice(ids []bson.ObjectId) []domain.Property {
	var properties []domain.Property
	var p domain.Property
	for _, id := range ids {
		p, _ = GetPropertyByID(id)
		properties = append(properties, p)
	}
	return properties
}

// Gets the User from the cookie
func AddFeatured(id bson.ObjectId) {
	feat := GetFeatured()
	feat.Featured = append(feat.Featured, id)
	SaveObject(feat, feat.ID)
}

//////////// ****************************************************************************************************************************************************************************************************

// SaveObject saves the element in the DataBase
func ModifyObject(object interface{}, id bson.ObjectId) error {
	client, ctx := Init()
	defer client.Close()

	// The actual operation we want to make
	switch v := object.(type) {
	case domain.User:
		u := object.(domain.User)
		client.Collection(collectionUsers).Doc(id.Hex()).Update(ctx, []firestore.Update{
			{
				Path:  "ID",
				Value: u.ID,
			},
			{
				Path:  "Name",
				Value: u.Name,
			},
			{
				Path:  "Surname",
				Value: u.Surname,
			},
			{
				Path:  "IdentificationNumber",
				Value: u.IdentificationNumber,
			},
			{
				Path:  "Email",
				Value: u.Email,
			},
			{
				Path:  "Password",
				Value: u.Password,
			},
			{
				Path:  "Address",
				Value: u.Address,
			},
			{
				Path:  "Province",
				Value: u.Province,
			},
			{
				Path:  "Country",
				Value: u.Country,
			},
			{
				Path:  "State",
				Value: u.State,
			},
			{
				Path:  "Region",
				Value: u.Region,
			},
			{
				Path:  "CountryPhonePrefix",
				Value: u.CountryPhonePrefix,
			},
			{
				Path:  "Phone",
				Value: u.Phone,
			},
			{
				Path:  "PhotoUser",
				Value: u.PhotoUser,
			},
			{
				Path:  "PhotoIDFront",
				Value: u.PhotoIDFront,
			},
			{
				Path:  "PhotoIDBack",
				Value: u.PhotoIDBack,
			},
			{
				Path:  "PhotoWithID",
				Value: u.PhotoWithID,
			},
			{
				Path:  "IsValidated",
				Value: u.IsValidated,
			},
			{
				Path:  "IsEmailValidated",
				Value: u.IsEmailValidated,
			},
			{
				Path:  "Role",
				Value: u.Role,
			},
			{
				Path:  "DateInscription",
				Value: u.DateInscription,
			},
			{
				Path:  "FavouriteProperties",
				Value: u.FavouriteProperties,
			},
		})

	case domain.Property:
		p := object.(domain.Property)
		client.Collection(collectionUsers).Doc(id.Hex()).Update(ctx, []firestore.Update{
			{
				Path:  "ID",
				Value: p.ID,
			},
			{
				Path:  "UserID",
				Value: p.UserID,
			},
			{
				Path:  "DateInscription",
				Value: p.DateInscription,
			},
			{
				Path:  "Description",
				Value: p.Description,
			},
			{
				Path:  "ShortDescription",
				Value: p.ShortDescription,
			},
			{
				Path:  "Price",
				Value: p.Price,
			},
			{
				Path:  "Country",
				Value: p.Country,
			},
			{
				Path:  "Region",
				Value: p.Region,
			},
			{
				Path:  "State",
				Value: p.State,
			},
			{
				Path:  "PostalCode",
				Value: p.PostalCode,
			},
			{
				Path:  "City",
				Value: p.City,
			},
			{
				Path:  "Type",
				Value: p.Type,
			},
			{
				Path:  "MetersBuilt",
				Value: p.MetersBuilt,
			},
			{
				Path:  "MetersAvailable",
				Value: p.MetersAvailable,
			},
			{
				Path:  "HasEvelator",
				Value: p.HasEvelator,
			},
			{
				Path:  "Public",
				Value: p.Public,
			},
			{
				Path:  "IsSale",
				Value: p.IsSale,
			},
			{
				Path:  "Rooms",
				Value: p.Rooms,
			},
			{
				Path:  "Toilets",
				Value: p.Toilets,
			},
			{
				Path:  "Floor",
				Value: p.Floor,
			},
			{
				Path:  "Garage",
				Value: p.Garage,
			},
			{
				Path:  "StorageRoom",
				Value: p.StorageRoom,
			},
			{
				Path:  "Photos",
				Value: p.Photos,
			},
			{
				Path:  "Documents",
				Value: p.Documents,
			},
			{
				Path:  "IsFavoriteForCurrentUser",
				Value: p.IsFavoriteForCurrentUser,
			},
		})
	case domain.FeaturedProperties:
		fp := object.(domain.FeaturedProperties)
		client.Collection(collectionUsers).Doc(id.Hex()).Update(ctx, []firestore.Update{
			{
				Path:  "ID",
				Value: fp.ID,
			},
			{
				Path:  "Featured",
				Value: fp.Featured,
			},
		})

	default:
		fmt.Println("Struct not handled")
		fmt.Println("This is v: ", v)
	}
	return nil
}

// SaveObject saves the element in the DataBase
func SaveObject2(object interface{}, id bson.ObjectId) error {
	client, ctx := Init()
	defer client.Close()

	// The actual operation we want to make
	switch v := object.(type) {
	case domain.User:
		u := object.(domain.User)
		client.Collection(collectionUsers).Add(ctx, map[string]interface{}{
			"ID":                   u.ID,
			"Name":                 u.Name,
			"Surname":              u.Surname,
			"IdentificationNumber": u.IdentificationNumber,
			"Email":                u.Email,
			"Password":             u.Password,
			"Address":              u.Address,
			"Province":             u.Province,
			"Country":              u.Country,
			"State":                u.State,
			"Region":               u.Region,
			"CountryPhonePrefix":   u.CountryPhonePrefix,
			"Phone":                u.Phone,
			"PhotoUser":            u.PhotoUser,
			"PhotoIDFront":         u.PhotoIDFront,
			"PhotoIDBack":          u.PhotoIDBack,
			"PhotoWithID":          u.PhotoWithID,
			"IsValidated":          u.IsValidated,
			"IsEmailValidated":     u.IsEmailValidated,
			"Role":                 u.Role,
			"DateInscription":      u.DateInscription,
			"FavouriteProperties":  u.FavouriteProperties,
		})

	case domain.Property:
		p := object.(domain.Property)
		client.Collection(collectionUsers).Add(ctx, map[string]interface{}{
			"ID":                       p.ID.Hex(),
			"UserID":                   p.UserID,
			"DateInscription":          p.DateInscription,
			"Description":              p.Description,
			"ShortDescription":         p.ShortDescription,
			"Price":                    p.Price,
			"Country":                  p.Country,
			"Region":                   p.Region,
			"State":                    p.State,
			"Street":                   p.Street,
			"PostalCode":               p.PostalCode,
			"City":                     p.City,
			"Type":                     p.Type,
			"MetersBuilt":              p.MetersBuilt,
			"MetersAvailable":          p.MetersAvailable,
			"HasEvelator":              p.HasEvelator,
			"Public":                   p.Public,
			"IsSale":                   p.IsSale,
			"Rooms":                    p.Rooms,
			"Toilets":                  p.Toilets,
			"Floor":                    p.Floor,
			"Garage":                   p.Garage,
			"StorageRoom":              p.StorageRoom,
			"Photos":                   p.Photos,
			"Documents":                p.Documents,
			"IsFavoriteForCurrentUser": p.IsFavoriteForCurrentUser,
		})
	default:
		fmt.Println("Struct not handled")
		fmt.Println("This is v: ", v)
	}
	return nil
}
