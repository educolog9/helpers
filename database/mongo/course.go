package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID                      primitive.ObjectID `bson:"_id"`
	CountRating             int                `bson:"countRating"`
	TotalRating             int                `bson:"totalRating"`
	TotalStudents           int                `bson:"totalStudents"`
	TotalDuration           int                `bson:"totalDuration"`
	CreatedAt               time.Time          `bson:"createdAt"`
	Title                   string             `bson:"title"`
	TitleTranslations       map[string]string  `bson:"titleTranslations"`
	Description             string             `bson:"description"`
	DescriptionTranslations map[string]string  `bson:"descriptionTranslations"`
	MainPicture             string             `bson:"mainPicture"`
	Pictures                []string           `bson:"pictures"`
	AuthorID                primitive.ObjectID `bson:"authorId"`
	AuthorName              string             `bson:"authorName"`
	CategoryID              primitive.ObjectID `bson:"categoryId"`
	SupportLanguages        []LanguageEnum     `bson:"supportLanguages"`
	Status                  CourseStatus       `bson:"status"`
	CourseType              CourseType         `bson:"courseType"`
	Version                 int                `bson:"version"`
	UpdatedAt               time.Time          `bson:"updatedAt"`
}

type CourseStatus string

const (
	// Published represents a published course
	// @description Published course
	CoursePublished CourseStatus = "published"

	// Unpublished represents an unpublished course
	// @description Unpublished course
	CourseUnpublished CourseStatus = "unpublished"
)

type CourseType string

const (
	// CourseTypeFree represents a free course
	// @description Free course
	CourseTypeFree CourseType = "free"

	// CourseTypeSubscription represents a subscription course
	// @description Subscription course
	CourseTypeSubscription CourseType = "subscription"
)

// LanguageEnum represents the language of a course.
// @name LanguageEnum
type LanguageEnum string

const (

	// SPANISH represents the SPANISH language
	// @description SPANISH language
	SPANISH LanguageEnum = "spanish"

	// ENGLISH represents the ENGLISH language
	// @description ENGLISH language
	ENGLISH LanguageEnum = "english"

	// FRENCH represents the FRENCH language
	// @description FRENCH language
	FRENCH LanguageEnum = "french"
)
