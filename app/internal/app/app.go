package app

import (
	"rss/internal/services/fetchPosts/constructor"

	"github.com/sirupsen/logrus"
)

func NewApp(logger *logrus.Logger) {
	constructor.FetchPostsController.Fetch()
}
