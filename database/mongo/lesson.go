package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lesson struct {
	ID                primitive.ObjectID `bson:"_id"`
	Title             string             `bson:"title"`
	TitleTranslations map[string]string  `bson:"titleTranslations"`
	LockedBy          primitive.ObjectID `bson:"lockedBy,omitempty"`
	IsOptional        bool               `bson:"isOptional"`
	Step              int                `bson:"step"`
	Duration          int                `bson:"duration"`
	Version           int                `bson:"version"`
	UpdatedAt         time.Time          `bson:"updatedAt"`
	CourseID          primitive.ObjectID `bson:"courseId"`
	CreatedAt         time.Time          `bson:"createdAt"`
}
