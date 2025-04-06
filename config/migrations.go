package config

import (
	"github.com/TheSaifZaman/GoFast/migrations"
)

type MigrationListStruct struct {
	Name      string
	Structure any
}

var LandlordMigrationList = []MigrationListStruct{
	{
		Name:      "users",
		Structure: &migrations.User{},
	},
	{
		Name:      "user_profiles",
		Structure: &migrations.UserProfile{},
	},
}

var TenantMigrationList = []MigrationListStruct{
	{
		//
	},
}
