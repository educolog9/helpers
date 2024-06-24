package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   primitive.ObjectID `bson:"_id"`
	Email                string             `bson:"email"`
	Name                 string             `bson:"name"`
	LastName             string             `bson:"lastName"`
	Password             string             `bson:"password"`
	Roles                []string           `bson:"roles"`
	IsConfirmed          bool               `bson:"isConfirmed"`
	IsBlocked            bool               `bson:"isBlocked"`
	ProfilePicture       string             `bson:"profilePicture"`
	Department           string             `bson:"department"`
	Organization         primitive.ObjectID `bson:"organization"`
	Group                primitive.ObjectID `bson:"group"`
	Version              int                `bson:"version"`
	UpdatedAt            time.Time          `bson:"updatedAt"`
	LessonEngagementTime int                `bson:"lessonEngagementTime"`
	MicroLessons         int                `bson:"microLessons"`
	GroupName            string             `bson:"groupName"`
	CreatedAt            time.Time          `bson:"createdAt"`
}
