package constructor

import (
	rssConstructor "rss/internal/services/rss/constructor"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetConstructor(client *gorm.DB, logger *logrus.Logger) {
	rssConstructor.RssRequirementsCreator(client, logger)
}
