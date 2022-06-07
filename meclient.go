package meclient

import (
	"meclient/methods"
	"time"
)

type MeClient struct {
	*methods.Methods
}

var DefaultMeClient = &MeClient{
	Methods: methods.NewMethods(time.Second, 2),
}

func NewMeClient(rateTime time.Duration, burst int) *MeClient {
	return &MeClient{
		Methods: methods.NewMethods(rateTime, burst),
	}
}
