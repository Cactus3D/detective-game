package main

import (
	"detective/internal/game"
	"fmt"
)

func main() {
	gCtx := game.Init(game.WithMap("configs/gamemap.yaml"))
	gm := gCtx.GetMap()

	v := gm.GetItem(5)
	t := gm.GetItem(5)
	nwp, _ := gm.NextWaypoint(v, t)
	if nwp != nil {
		fmt.Println(nwp.GetInfo())
	} else {
		fmt.Println("Not Found")
	}

}
