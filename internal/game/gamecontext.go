package game

import (
	"context"
	. "detective/internal/locations"
	"fmt"
	"log"
)

const TIMEUNIT = 5

type GameContext struct {
	gm     *GameMap
	events []EventTriggeredByTic
	context.Context
}

type GameDate int

type EventTriggeredByTic func(x any)

func (d *GameDate) tic() {
	*d = *d + TIMEUNIT
}

func (d *GameDate) skip(amount int) bool {
	if amount%5 == 0 {
		rp := amount / 5
		for i := 0; i < rp; i++ {
			d.tic()
		}
		return true
	}
	return false
}

func (d *GameDate) GetFormatDate() string {
	return fmt.Sprintf("Day: %d, %02d:%02d", *d/1440, *d%1440/60, *d%1440%60)
}

func (this *GameContext) SetMap(gm *GameMap) {
	this.gm = gm
}

func (this *GameContext) GetMap() *GameMap {
	return this.gm
}

type GameContextOptions func(this *GameContext)

func Init(opts ...GameContextOptions) (gCtx *GameContext) {
	gCtx = &GameContext{}

	for _, opt := range opts {
		opt(gCtx)
	}

	return gCtx
}

func WithMap(path string) GameContextOptions {
	return func(this *GameContext) {
		var err error
		this.gm, err = NewGameMap(path)
		if err != nil {
			log.Panicf("Cannot load map from %s path", path)
		}
	}
}
