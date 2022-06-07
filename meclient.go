package meclient

import (
	"time"

	"github.com/letsnode/meclient/methods"
)

type MeClient struct {
	*methods.Methods
}

func NewMeClient(rateTime time.Duration, burst int) *MeClient {
	return &MeClient{
		Methods: methods.NewMethods(rateTime, burst),
	}
}

func NewDefaultMeClient() *MeClient {
	return &MeClient{
		Methods: methods.NewMethods(time.Second, 2),
	}
}
