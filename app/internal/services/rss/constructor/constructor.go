package constructor

import (
	"rss/internal/services/rss/controller"
	"rss/internal/services/rss/service"
	"rss/internal/services/rss/storage"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	RssRepository storage.RssStorage
	RssService    service.RssService
	RssController controller.RssController
)

func RssRequirementsCreator(client *gorm.DB, logger *logrus.Logger) {
	RssRepository = storage.NewRssStorage(client, logger)
	RssService = service.NewRssService(RssRepository, logger)
	RssController = controller.NewRssController(RssService, logger)
}
