package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStreak struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"userId"`
	LessonID  primitive.ObjectID `bson:"lessonId"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Version   int                `bson:"version"`
}
