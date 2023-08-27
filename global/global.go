package global

import "user-api/config"

var (
	Conf *config.Conf = &config.Conf{}

	Page     uint32 = 1
	PageSize uint32 = 15
)
