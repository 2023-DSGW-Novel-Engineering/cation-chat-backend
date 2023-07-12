package initializers

import (
	"github.com/2023-DSGW-Novel-Engineering/cation-chat-backend/models"
)

func SyncDatabase() {
	if err := DB.AutoMigrate(new(models.Chatory)); err != nil {
		panic(err)
	}
}
