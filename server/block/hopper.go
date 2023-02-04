package block

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/internal/nbtconv"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"strings"
	"sync"
)

// Hopper is a low-capacity storage block that can be used to collect item entities directly above it, as well as to
// transfer items into and out of other containers.
// TODO: Functionality!
type Hopper struct {
	transparent

	// Facing is the direction the hopper is facing.
	Facing cube.Face
	// Powered is whether the hopper is powered or not.
	Powered bool
	// CustomName is the custom name of the hopper. This name is displayed when the hopper is opened, and may include
	// colour codes.
	CustomName string

	inventory *inventory.Inventory
	viewerMu  *sync.RWMutex
	viewers   map[ContainerViewer]struct{}
}

// NewHopper creates a new initialised hopper. The inventory is properly initialised.
func NewHopper() Hopper {
	m := new(sync.RWMutex)
	v := make(map[ContainerViewer]struct{}, 1)
	return Hopper{
		inventory: inventory.New(5, func(slot int, _, item item.Stack) {
			m.RLock()
			defer m.RUnlock()
			for viewer := range v {
				viewer.ViewSlotChange(slot, item)
			}
		}),
		viewerMu: m,
		viewers:  v,
	}
}

// Model ...
func (Hopper) Model() world.BlockModel {
	// TODO: Implement me.
	return model.Solid{}
}

// BreakInfo ...
func (h Hopper) BreakInfo() BreakInfo {
	return newBreakInfo(3, pickaxeHarvestable, pickaxeEffective, oneOf(h))
}

// Inventory returns the inventory of the hopper.
func (h Hopper) Inventory() *inventory.Inventory {
	return h.inventory
}

// WithName returns the hopper after applying a specific name to the block.
func (h Hopper) WithName(a ...any) world.Item {
	h.CustomName = strings.TrimSuffix(fmt.Sprintln(a...), "\n")
	return h
}

// AddViewer adds a viewer to the hopper, so that it is updated whenever the inventory of the hopper is changed.
func (h Hopper) AddViewer(v ContainerViewer, _ *world.World, _ cube.Pos) {
	h.viewerMu.Lock()
	defer h.viewerMu.Unlock()
	h.viewers[v] = struct{}{}
}

// RemoveViewer removes a viewer from the hopper, so that slot updates in the inventory are no longer sent to it.
func (h Hopper) RemoveViewer(v ContainerViewer, _ *world.World, _ cube.Pos) {
	h.viewerMu.Lock()
	defer h.viewerMu.Unlock()
	delete(h.viewers, v)
}

// Activate ...
func (Hopper) Activate(pos cube.Pos, _ cube.Face, _ *world.World, u item.User, _ *item.UseContext) bool {
	if o, ok := u.(ContainerOpener); ok {
		o.OpenBlockContainer(pos)
		return true
	}
	return false
}

// UseOnBlock ...
func (h Hopper) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(w, pos, face, h)
	if !used {
		return false
	}

	//noinspection GoAssignmentToReceiver
	h = NewHopper()
	h.Facing = cube.FaceDown
	if h.Facing != face {
		h.Facing = face.Opposite()
	}

	place(w, pos, h, user, ctx)
	return placed(ctx)
}

// EncodeItem ...
func (h Hopper) EncodeItem() (name string, meta int16) {
	return "minecraft:hopper", 0
}

// EncodeBlock ...
func (h Hopper) EncodeBlock() (string, map[string]any) {
	return "minecraft:hopper", map[string]any{
		"facing_direction": int32(h.Facing),
		"toggle_bit":       h.Powered,
	}
}

// EncodeNBT ...
func (h Hopper) EncodeNBT() map[string]any {
	if h.inventory == nil {
		facing, powered, customName := h.Facing, h.Powered, h.CustomName
		//noinspection GoAssignmentToReceiver
		h = NewHopper()
		h.Facing, h.Powered, h.CustomName = facing, powered, customName
	}
	m := map[string]any{
		"Items": nbtconv.InvToNBT(h.inventory),
		"id":    "Hopper",
	}
	if h.CustomName != "" {
		m["CustomName"] = h.CustomName
	}
	return m
}

// DecodeNBT ...
func (h Hopper) DecodeNBT(data map[string]any) any {
	facing, powered := h.Facing, h.Powered
	//noinspection GoAssignmentToReceiver
	h = NewHopper()
	h.Facing = facing
	h.Powered = powered
	h.CustomName = nbtconv.String(data, "CustomName")
	nbtconv.InvFromNBT(h.inventory, nbtconv.Slice(data, "Items"))
	return h
}

// allHoppers ...
func allHoppers() (hoppers []world.Block) {
	for _, f := range cube.Faces() {
		for _, p := range []bool{false, true} {
			hoppers = append(hoppers, Hopper{Facing: f, Powered: p})
		}
	}
	return hoppers
}