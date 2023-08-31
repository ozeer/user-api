package global

import (
	"user-api/config"

	"go.uber.org/zap"
)

var (
	Conf *config.Conf = &config.Conf{}
	Log  *zap.SugaredLogger

	Page     uint32 = 1
	PageSize uint32 = 15
)
