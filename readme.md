<h1>meclient</h1>

`meclient` is a library for interacting with MagicEden.io API. You can find the API specification at the link: https://api.magiceden.dev/

<h3>Installation</h3>

`go get github.com/letsnode/meclient`

<h3>Usage example</h3>

```
package main

import (
	"context"
	"fmt"
	
	"github.com/letsnode/meclient"
	"github.com/letsnode/meclient/methods"
)

func main()  {
	client := meclient.NewDefaultMeClient()
	lp, err := client.LaunchpadCollections(context.Background(), &methods.LimitOffsetOpts{Limit: 5})
	if err != nil {
		fmt.Println(err)
		return
	}
	...
}
```
