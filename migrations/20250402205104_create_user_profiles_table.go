package migrations

type UserProfile struct {
	CustomIdColumn
	UserIdColumn
	CustomActorAndActingTimeColumns
}
