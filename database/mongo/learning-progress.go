package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LearningProgress struct {
	ID             primitive.ObjectID   `bson:"_id"`
	Progress       float64              `bson:"progress"`
	StudentID      primitive.ObjectID   `bson:"studentId"`
	OrganizationID primitive.ObjectID   `bson:"organizationId"`
	CourseID       primitive.ObjectID   `bson:"courseId"`
	CreatedAt      time.Time            `bson:"createdAt"`
	Lessons        []primitive.ObjectID `bson:"lessons"`
	Status         string               `bson:"status"`
	Version        int                  `bson:"version"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
}
