package handle

import (
	"github.com/atompi/autom/cmd/autom-apiserver/app/options"
	"go.uber.org/zap"
)

func Handle(opts options.Options) {
	zap.L().Sugar().Infof("options: ", opts)
}
