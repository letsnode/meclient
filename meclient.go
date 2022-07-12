package meclient

import (
	"time"

	"github.com/letsnode/meclient/consts"
	"github.com/letsnode/meclient/methods"
)

type MeClient struct {
	*methods.Executor
}

func NewDefaultClient() *MeClient {
	return &MeClient{
		Executor: methods.NewMethodExecutor(consts.Host, time.Second, 2),
	}
}
