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

	"github.com/frizinak/goniri/niri"
	"github.com/frizinak/goniri/niri/types"
)

func main() {
	ipc := niri.New(os.Getenv("NIRI_SOCKET"))
	if err := ipc.Connect(); err != nil {
		panic(err)
	}
	defer ipc.Close()

	req := niri.RequestFocusedWindow()
	if err := ipc.Do(req); err != nil {
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

