package errors

import (
	"testing"

	"go.uber.org/zap"
)

func TestErr(t *testing.T) {
	logger, _ := zap.NewProduction()
	err := New("第一次错误")
	err = Wrap(err, "2")
	err = Wrapf(err, "%v", "adadas")
	logger.Info("err", zap.Error(err))
}
