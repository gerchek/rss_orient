package app

import (
	"rss/internal/services/rss/constructor"

	"github.com/sirupsen/logrus"
)

func NewApp(logger *logrus.Logger) {
	constructor.RssController.Fetch()
}
