package handle

import (
	"github.com/atompi/autom/cmd/options"
	"go.uber.org/zap"
)

func Handle(opts options.Options) {
	zap.L().Sugar().Infof("options: ", opts)
}
