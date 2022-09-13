package constructor

import (
	"rss/internal/services/fetchPosts/controller"
	"rss/internal/services/fetchPosts/service"
	"rss/internal/services/fetchPosts/storage"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	FetchPostsRepository storage.FetchPostsStorage
	FetchPostsService    service.FetchPostsService
	FetchPostsController controller.FetchPostsController
)

func FetchPostsRequirementsCreator(client *gorm.DB, logger *logrus.Logger) {
	FetchPostsRepository = storage.NewFetchPostsStorage(client, logger)
	FetchPostsService = service.NewFetchPostsService(FetchPostsRepository, logger)
	FetchPostsController = controller.NewFetchPostsController(FetchPostsService, logger)
}
