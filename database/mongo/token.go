package mongo

import "time"

type Token struct {
	Email     string    `bson:"email"`     // Email is the user's email address.
	Token     string    `bson:"token"`     // Token is the authentication token.
	CreatedAt time.Time `bson:"createdAt"` // CreatedAt is the timestamp when the token was created.
}
