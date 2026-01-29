## Niri IPC in Go

Features:
 - [X] map all IPC calls
 - [ ] convenient high-level API
 - [ ] example implementations

## Example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/frizinak/goniri/ipc"
	"github.com/frizinak/goniri/ipc/types"
)

func main() {
	c := ipc.New(os.Getenv("NIRI_SOCKET"))
	if err := c.Connect(); err != nil {
		panic(err)
	}
	defer c.Close()

	req := ipc.RequestFocusedWindow()
	if err := c.Do(req); err != nil {
		panic(err)
	}
	fmt.Println(req.Response())

	err = c.Events(func(e types.Event) error {
		fmt.Printf("ev %s %+v\n", time.Now().Format("15:04:05"), e)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

```

