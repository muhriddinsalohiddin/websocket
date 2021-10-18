package utils

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DB() (*gorm.DB, error) {

    PG_HOST := "localhost"
    PG_USER := "muhriddin"
    PG_PASSWORD := "1"
    PG_PORT := "5432"
    PG_DBNAME := "users"
    PG_SSLMODE := "disable"
    PG_TIMEZONE := "Asia/Tashkent"

	PGConnection := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		PG_HOST,
		PG_USER,
		PG_PASSWORD,
		PG_PORT,
		PG_DBNAME,
		PG_SSLMODE,
		PG_TIMEZONE,
	)

	return gorm.Open(postgres.Open(PGConnection), &gorm.Config{})
}

