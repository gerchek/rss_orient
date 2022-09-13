package constructor

import (
	fetchPostsConstructor "rss/internal/services/fetchPosts/constructor"
	rssLinksConstructor "rss/internal/services/rssLinks/constructor"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetConstructor(client *gorm.DB, logger *logrus.Logger) {
	fetchPostsConstructor.FetchPostsRequirementsCreator(client, logger)
	rssLinksConstructor.RssLinksRequirementsCreator(client, logger)
}
