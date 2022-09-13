package constructor

import (
	"rss/internal/services/rssLinks/controller"
	"rss/internal/services/rssLinks/service"
	"rss/internal/services/rssLinks/storage"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	RssLinksRepository storage.RssLinksStorage
	RssLinksService    service.RssLinksService
	RssLinksController controller.RssLinksController
)

func RssLinksRequirementsCreator(client *gorm.DB, logger *logrus.Logger) {
	RssLinksRepository = storage.NewRssLinksStorage(client, logger)
	RssLinksService = service.NewRssLinksService(RssLinksRepository, logger)
	RssLinksController = controller.NewRssLinksController(RssLinksService, logger)
}
