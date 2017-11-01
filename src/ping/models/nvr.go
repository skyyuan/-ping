package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

//相机表
type Nvr struct {
	Id_         bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	IP          string        `bson:"ip"`
	SdkPort     string        `bson:"sdk_port"`
	RtspPort    string        `bson:"rtsp_port"`
	UserName    string        `bson:"user_name"`
	Password    string        `bson:"password"`
	Status      string        `bson:"status"`
	CommonModel               `bson:",inline"`
}

func GetNvrs(db *mgo.Database) (nvrs []Nvr, err error) {
	collection := db.C("nvrs")
	err = collection.Find(nil).All(&nvrs)
	return
}


func (self *Nvr)UpdateByParams(db *mgo.Database, status string)(err error){
	userCollection := db.C("nvrs")
	err = userCollection.Update(bson.M{"_id": self.Id_},bson.M{"$set": bson.M{"status": status, "updated_at": bson.Now()}})
	return
}