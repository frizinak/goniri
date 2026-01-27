## WIP

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/frizinak/goniri/niri"
	"github.com/frizinak/goniri/niri/types"
)

func main() {
	ipc := niri.New(os.Getenv("NIRI_SOCKET"))
	if err := ipc.Init(); err != nil {
		panic(err)
	}
	defer ipc.Close()

	req := niri.NewFocusedWindowRequest()
	err := ipc.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(req.Response())

	err = ipc.Events(func(e types.Event) error {
		fmt.Printf("ev %s %+v\n", time.Now().Format("15:04:05"), e)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

```
