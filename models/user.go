package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		Id      string `json:"_id" bson:"_id"`
		Name    string `json:"name" bson:"name"`
		Email   string `json:"email" bson:"email"`
		Gender  string `json:"gender" bson:"gender"`
		Image   string `json:"image" bson:"image"`
		Pass    string `json:"passHash" bson:"passHash"`
		PrivK   string `json:"privKey" bson:"privKey"`
		PubK    string `json:"pubKey" bson:"pubKey"`
		AuthC   string `json:"authCode" bson:"authCode"`
		Data    string `json:"data" bson:"data"`
		Submit  bool   `json:"submitted" bson:"submitted"`
		Matches string `json:"matches" bson:"matches"`
		State   int32  `json:"state" bson:"state"`
	}
)

// ----------------------------------------
type TypeUserNew struct {
	Id       string `json:"roll"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Image    string `json:"image"`
	PassHash string `json:"passHash"`
}

func NewUser(info *TypeUserNew) User {
	return User{
		Id:      info.Id,
		Name:    info.Name,
		Email:   info.Email,
		Gender:  info.Gender,
		Image:   info.Image,
		Pass:    info.PassHash,
		PrivK:   "",
		PubK:    "",
		AuthC:   "assdasdas",
		Data:    "",
		Submit:  false,
		Matches: "",
		State:   0,
	}
}

// ----------------------------------------
type TypeUserFirst struct {
	Id       string `json:"roll"`
	AuthCode string `json:"authCode"`
	PassHash string `json:"passHash"`
	PubKey   string `json:"pubKey"`
	PrivKey  string `json:"privKey"`
	Data     string `json:"data"`
}

func (u User) FirstLogin(info *TypeUserFirst) mgo.Change {
	return mgo.Change{
		Update: bson.M{"$set": bson.M{
			"authCode": "",
			"passHash": info.PassHash,
			"pubKey":   info.PubKey,
			"privKey":  info.PrivKey,
			"data":     info.Data,
			"matches":  "",
		}},
		ReturnNew: true,
	}
}

// ----------------------------------------
func (u User) ValidPass(pass string) bool {
	return u.Pass == pass
}

// ----------------------------------------
func (u User) HasSubmitted() mgo.Change {
	return mgo.Change{
		Update: bson.M{"$set": bson.M{
			"submitted": true,
		}},
		ReturnNew: true,
	}
}

// ----------------------------------------
type TypeUserUpdateData struct {
	Data string `json:"data"`
}

func (u User) UpdateData(info *TypeUserUpdateData) mgo.Change {
	return mgo.Change{
		Update: bson.M{"$set": bson.M{
			"data": info.Data,
		}},
		ReturnNew: true,
	}
}
