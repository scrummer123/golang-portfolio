package models

import (
	"context"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/scrummer123/golang-portfolio/database"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

// UserPost => title: title from user post, content: content from user post, userid: user id from post
type UserPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// User => Username: litteraly means how its called, Password: password. Dont show the password when getting the user data
type User struct {
	Username string    `json:"username"`
	Password []byte    `json:"-"`
	Post     *UserPost `json:"Post"`
}

var users map[string]User = make(map[string]User)

// GetAll returns all posts and refreshes the local userposts variable
func (User) GetAll() map[string]User {
	log.SetPrefix("[models.GetAll()] :: ")
	db := database.GetFirestoreClient()

	i := db.Collection("users").Documents(context.Background())
	for {
		doc, err := i.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error occured: %v", err)
		}

		data := doc.Data()

		var User User
		err = mapstructure.Decode(data, &User)
		if err != nil {
			log.Fatal(err)
		}

		users[doc.Ref.ID] = User
	}

	return users
}

// GetByID returns user by UserID
func (User) GetByID(UserID string) (User, bool) {
	User{}.GetAll()

	User, UserIsset := users[UserID]
	return User, UserIsset
}

// Create makes a new document in the database
// return true if successful, false if not successful
func (User) Create(u User) (User, error) {
	db := database.GetFirestoreClient()

	pass, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = pass

	doc := db.Collection("users").NewDoc()
	_, docErr := doc.Set(context.Background(), u)

	if docErr != nil {
		log.Fatalf("%v", err)
		return User{}, err
	}

	users[doc.ID] = u

	return u, nil
}