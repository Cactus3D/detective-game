package character

import (
	. "detective/internal/game"
	. "detective/internal/locations"
	"log"
)

type Character struct {
	name             string
	currentLocation  *Vertex
	waypointLocation *Vertex
	targetLocation   *Vertex
	trevelCooldown   int
}

type schedule struct {
}

func (this *Character) Travel(gCtx *GameContext) {

	gameMap := gCtx.GetMap()

	if this.targetLocation == nil || this.targetLocation == this.currentLocation {
		return
	}

	if this.trevelCooldown > 0 {
		this.trevelCooldown = this.trevelCooldown - TIMEUNIT
		return
	}

	this.currentLocation = this.waypointLocation

	if this.currentLocation == this.targetLocation {
		this.waypointLocation = nil
		this.targetLocation = nil
		return
	}

	waypoint, err := gameMap.NextWaypoint(this.currentLocation, this.targetLocation)

	if err != nil {
		this.waypointLocation = nil
		this.targetLocation = nil
		log.Println("WARNING: cannot get next waypoint, character:", this.name)
		return
	}

	this.waypointLocation = waypoint
	this.trevelCooldown = this.currentLocation.GetDurationToTarget(this.waypointLocation)
}

// SetTargetLocation func sets new target to character's travel, returning ok = true if set was succesfull, otherwise return false and error.
// Updates waypoint and target locations. Sets travelCooldown, if next waypoint was changed.
func (this *Character) SetTargetLocation(target *Vertex, gCtx *GameContext) (ok bool, err error) {

	if this.targetLocation == target {
		return true, nil
	}

	if this.currentLocation == target {
		this.waypointLocation = nil
		this.targetLocation = nil
		this.trevelCooldown = 0
		return true, nil
	}

	gameMap := gCtx.GetMap()

	waypoint, err := gameMap.NextWaypoint(this.currentLocation, target)

	if err != nil {
		return false, err
	}

	this.targetLocation = target

	if this.waypointLocation != waypoint {
		this.waypointLocation = waypoint
		this.trevelCooldown = this.currentLocation.GetDurationToTarget(this.waypointLocation)
	}

	return true, nil
}
