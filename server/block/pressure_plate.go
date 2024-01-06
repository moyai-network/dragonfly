package block

import (
	"math/rand"
	"time"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// Lever is a non-solid block that can provide switchable redstone power.
type PressurePlate struct {
	empty
	transparent
	flowingWaterDisplacer

	Wooden bool
	Wood   WoodType
	// Powered is if the PressurePlate is switched on.
	Powered bool
}

// FaceSolid ...
func (u PressurePlate) FaceSolid(cube.Pos, cube.Face, *world.World) bool {
	return true
}

// Source ...
func (l PressurePlate) Source() bool {
	return true
}

// WeakPower ...
func (l PressurePlate) WeakPower(cube.Pos, cube.Face, *world.World, bool) int {
	if l.Powered {
		return 15
	}
	return 0
}

// StrongPower ...
func (l PressurePlate) StrongPower(_ cube.Pos, face cube.Face, _ *world.World, _ bool) int {
	if l.Powered {
		return 15
	}
	return 0
}

// NeighbourUpdateTick ...
func (l PressurePlate) NeighbourUpdateTick(pos, _ cube.Pos, w *world.World) {
	if !l.Powered {
		return
	}
	w.ScheduleBlockUpdate(pos, 0)
}

// UseOnBlock ...
func (l PressurePlate) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) bool {
	return true
}

// Activate ...
func (l PressurePlate) Activate(pos cube.Pos, _ cube.Face, w *world.World, _ item.User, _ *item.UseContext) bool {
	return false
}

func (l PressurePlate) EntityInside(pos cube.Pos, w *world.World, e world.Entity) {
	l.Powered = true
	w.SetBlock(pos, l, nil)
	w.ScheduleBlockUpdate(pos, time.Millisecond*200)
	updateAroundRedstone(pos, w)
}

func (l PressurePlate) ScheduledTick(pos cube.Pos, w *world.World, _ *rand.Rand) {
	if !l.Powered {
		return
	}
	l.Powered = false
	w.SetBlock(pos, l, nil)
	//updateAroundRedstone(pos, w)
}

// BreakInfo ...
func (l PressurePlate) BreakInfo() BreakInfo {
	return newBreakInfo(0.8, pickaxeHarvestable, pickaxeEffective, nil)

}

// EncodeItem ...
func (l PressurePlate) EncodeItem() (name string, meta int16) {
	if !l.Wooden {
		return "minecraft:stone_pressure_plate", 0
	} else {
		w := l.Wood.String()
		if w == "oak" {
			w = "wooden"
		}
		return "minecraft:" + w + "_pressure_plate", 0
	}
}

// EncodeBlock ...
func (l PressurePlate) EncodeBlock() (string, map[string]any) {
	if !l.Wooden {
		return "minecraft:stone_pressure_plate", map[string]any{
			"redstone_signal": int32(boolByte(l.Powered)),
		}
	} else {
		w := l.Wood.String()
		if w == "oak" {
			w = "wooden"
		}
		return "minecraft:" + w + "_pressure_plate", map[string]any{
			"redstone_signal": int32(boolByte(l.Powered)),
		}
	}
}

// allPlanks returns all planks types.
func allPressurePlates() (plates []world.Block) {
	for _, w := range WoodTypes() {
		for _, b := range []bool{true, false} {
			plates = append(plates, PressurePlate{Wooden: true, Wood: w, Powered: b})
		}
	}
	for _, b := range []bool{true, false} {
		plates = append(plates, PressurePlate{Powered: b})
	}
	return
}
