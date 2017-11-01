package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

//相机表
type Camera struct {
	Id_         bson.ObjectId `bson:"_id"`
	CameraId    int           `bson:"camera_id"`
	DetectorId   int           `bson:"detector_id"`
	DeviceId    string        `bson:"device_id"`
	Location    string        `bson:"location"`
	Level       int           `bson:"level"`
	NvrId      int            `bson:"nvr_id"`
	ChannelId      int            `bson:"channel_id"`
	Ip    string        `bson:"ip"`
	Status    string        `bson:"status"`
	CommonModel `bson:",inline"`
}


func GetCameras(db *mgo.Database) (camera []Camera, err error) {
	collection := db.C("cameras")
	err = collection.Find(nil).All(&camera)
	return
}


func (self *Camera)UpdateByParams(db *mgo.Database, status string)(err error){
	userCollection := db.C("cameras")
	err = userCollection.Update(bson.M{"_id": self.Id_},bson.M{"$set": bson.M{"status": status, "updated_at": bson.Now()}})
	return
}