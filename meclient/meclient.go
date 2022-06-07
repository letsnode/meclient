package meclient

import (
	"demo/src/meclient/methods"
	"time"
)

type MeClient struct {
	*methods.Methods
}

var DefaultMeClient = &MeClient{methods.NewMethods(time.Second, 2)}

func NewMeClient(rateTime time.Duration, burst int) *MeClient {
	return &MeClient{methods.NewMethods(rateTime, burst)}
}
