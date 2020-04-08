package seed

import (
	"log"

	"github.com/eencogneeto/backend/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Name:     "jack",
		Email:    "jack@gmail.com",
		Password: "jacky",
	},
	models.User{
		Name:     "ben",
		Email:    "ben@gmail.com",
		Password: "beny",
	},
}

var nouns = []models.Noun{
	models.Noun{
		Noun:    "Handy",
		Gender:  "n",
		Meaning: "handphone",
		Plural:  "Handys",
	},
	models.Noun{
		Noun:    "Kind",
		Gender:  "n",
		Meaning: "kid",
		Plural:  "Kinder",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Noun{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Noun{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		err = db.Debug().Model(&models.Noun{}).Create(&nouns[i]).Error
		if err != nil {
			log.Fatalf("cannot seed nouns table: %v", err)
		}
	}
}
