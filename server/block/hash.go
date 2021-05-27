// Code generated by cmd/blockhash; DO NOT EDIT.

package block

const hashPlanks = 0
const hashShroomlight = 1
const hashLapisBlock = 2
const hashEndBrickStairs = 3
const hashStone = 4
const hashNetherGoldOre = 5
const hashCarrot = 6
const hashStainedTerracotta = 7
const hashStainedGlass = 8
const hashFarmland = 9
const hashChest = 10
const hashTerracotta = 11
const hashConcrete = 12
const hashCoalOre = 13
const hashStainedGlassPane = 14
const hashMelon = 15
const hashCocoaBean = 16
const hashNetherBrickFence = 17
const hashNoteBlock = 18
const hashEndStone = 19
const hashBasalt = 20
const hashWoodStairs = 21
const hashWheatSeeds = 22
const hashFire = 23
const hashBedrock = 24
const hashQuartz = 25
const hashQuartzBricks = 26
const hashBlueIce = 27
const hashWoodSlab = 28
const hashEndBricks = 29
const hashWoodTrapdoor = 30
const hashDirt = 31
const hashCarpet = 32
const hashGrassPlant = 33
const hashLight = 34
const hashTorch = 35
const hashSeaLantern = 36
const hashEmeraldBlock = 37
const hashPumpkinSeeds = 38
const hashWoodFence = 39
const hashQuartzPillar = 40
const hashChiseledQuartz = 41
const hashLeaves = 42
const hashAndesite = 43
const hashAncientDebris = 44
const hashSoulSoil = 45
const hashGildedBlackstone = 46
const hashClay = 47
const hashBeetrootSeeds = 48
const hashSand = 49
const hashDirtPath = 50
const hashLapisOre = 51
const hashCoalBlock = 52
const hashPumpkin = 53
const hashNetherrack = 54
const hashGlassPane = 55
const hashDiorite = 56
const hashConcretePowder = 57
const hashIronOre = 58
const hashSoulSand = 59
const hashIronBlock = 60
const hashWool = 61
const hashCoralBlock = 62
const hashGlazedTerracotta = 63
const hashInvisibleBedrock = 64
const hashWater = 65
const hashBricks = 66
const hashCobblestone = 67
const hashGlass = 68
const hashGoldOre = 69
const hashNetherQuartzOre = 70
const hashCoral = 71
const hashLog = 72
const hashEmeraldOre = 73
const hashWoodDoor = 74
const hashIronBars = 75
const hashLitPumpkin = 76
const hashGranite = 77
const hashDragonEgg = 78
const hashNetheriteBlock = 79
const hashGoldBlock = 80
const hashNetherWart = 81
const hashGravel = 82
const hashGlowstone = 83
const hashKelp = 84
const hashDiamondBlock = 85
const hashLava = 86
const hashSponge = 87
const hashObsidian = 88
const hashBoneBlock = 89
const hashDiamondOre = 90
const hashGrass = 91
const hashMelonSeeds = 92
const hashBeacon = 93
const hashCake = 94
const hashBarrier = 95
const hashAir = 96
const hashWoodFenceGate = 97
const hashLantern = 98
const hashPotato = 99

func (Barrier) Hash() uint64 {
	return hashBarrier
}

func (Air) Hash() uint64 {
	return hashAir
}

func (f WoodFenceGate) Hash() uint64 {
	return hashWoodFenceGate | uint64(f.Wood.Uint8())<<7 | uint64(f.Facing)<<11 | uint64(boolByte(f.Open))<<13 | uint64(boolByte(f.Lowered))<<14
}

func (DiamondOre) Hash() uint64 {
	return hashDiamondOre
}

func (Grass) Hash() uint64 {
	return hashGrass
}

func (m MelonSeeds) Hash() uint64 {
	return hashMelonSeeds | uint64(m.Growth)<<7 | uint64(m.Direction)<<15
}

func (Beacon) Hash() uint64 {
	return hashBeacon
}

func (c Cake) Hash() uint64 {
	return hashCake | uint64(c.Bites)<<7
}

func (p Potato) Hash() uint64 {
	return hashPotato | uint64(p.Growth)<<7
}

func (l Lantern) Hash() uint64 {
	return hashLantern | uint64(boolByte(l.Hanging))<<7 | uint64(l.Type.Uint8())<<8
}

func (NetherGoldOre) Hash() uint64 {
	return hashNetherGoldOre
}

func (c Carrot) Hash() uint64 {
	return hashCarrot | uint64(c.Growth)<<7
}

func (t StainedTerracotta) Hash() uint64 {
	return hashStainedTerracotta | uint64(t.Colour.Uint8())<<7
}

func (p Planks) Hash() uint64 {
	return hashPlanks | uint64(p.Wood.Uint8())<<7
}

func (Shroomlight) Hash() uint64 {
	return hashShroomlight
}

func (LapisBlock) Hash() uint64 {
	return hashLapisBlock
}

func (s EndBrickStairs) Hash() uint64 {
	return hashEndBrickStairs | uint64(boolByte(s.UpsideDown))<<7 | uint64(s.Facing)<<8
}

func (s Stone) Hash() uint64 {
	return hashStone | uint64(boolByte(s.Smooth))<<7
}

func (c Concrete) Hash() uint64 {
	return hashConcrete | uint64(c.Colour.Uint8())<<7
}

func (g StainedGlass) Hash() uint64 {
	return hashStainedGlass | uint64(g.Colour.Uint8())<<7
}

func (f Farmland) Hash() uint64 {
	return hashFarmland | uint64(f.Hydration)<<7
}

func (c Chest) Hash() uint64 {
	return hashChest | uint64(c.Facing)<<7
}

func (Terracotta) Hash() uint64 {
	return hashTerracotta
}

func (c CocoaBean) Hash() uint64 {
	return hashCocoaBean | uint64(c.Facing)<<7 | uint64(c.Age)<<9
}

func (CoalOre) Hash() uint64 {
	return hashCoalOre
}

func (p StainedGlassPane) Hash() uint64 {
	return hashStainedGlassPane | uint64(p.Colour.Uint8())<<7
}

func (Melon) Hash() uint64 {
	return hashMelon
}

func (EndStone) Hash() uint64 {
	return hashEndStone
}

func (NetherBrickFence) Hash() uint64 {
	return hashNetherBrickFence
}

func (n NoteBlock) Hash() uint64 {
	return hashNoteBlock
}

func (b Bedrock) Hash() uint64 {
	return hashBedrock | uint64(boolByte(b.InfiniteBurning))<<7
}

func (b Basalt) Hash() uint64 {
	return hashBasalt | uint64(boolByte(b.Polished))<<7 | uint64(b.Axis)<<8
}

func (s WoodStairs) Hash() uint64 {
	return hashWoodStairs | uint64(s.Wood.Uint8())<<7 | uint64(boolByte(s.UpsideDown))<<11 | uint64(s.Facing)<<12
}

func (s WheatSeeds) Hash() uint64 {
	return hashWheatSeeds | uint64(s.Growth)<<7
}

func (f Fire) Hash() uint64 {
	return hashFire | uint64(f.Type.Uint8())<<7 | uint64(f.Age)<<11
}

func (QuartzBricks) Hash() uint64 {
	return hashQuartzBricks
}

func (q Quartz) Hash() uint64 {
	return hashQuartz | uint64(boolByte(q.Smooth))<<7
}

func (c Carpet) Hash() uint64 {
	return hashCarpet | uint64(c.Colour.Uint8())<<7
}

func (g GrassPlant) Hash() uint64 {
	return hashGrassPlant | uint64(boolByte(g.UpperPart))<<7 | uint64(g.Type.Uint8())<<8
}

func (l Light) Hash() uint64 {
	return hashLight | uint64(l.Level)<<7
}

func (BlueIce) Hash() uint64 {
	return hashBlueIce
}

func (s WoodSlab) Hash() uint64 {
	return hashWoodSlab | uint64(s.Wood.Uint8())<<7 | uint64(boolByte(s.Top))<<11 | uint64(boolByte(s.Double))<<12
}

func (EndBricks) Hash() uint64 {
	return hashEndBricks
}

func (t WoodTrapdoor) Hash() uint64 {
	return hashWoodTrapdoor | uint64(t.Wood.Uint8())<<7 | uint64(t.Facing)<<11 | uint64(boolByte(t.Open))<<13 | uint64(boolByte(t.Top))<<14
}

func (d Dirt) Hash() uint64 {
	return hashDirt | uint64(boolByte(d.Coarse))<<7
}

func (t Torch) Hash() uint64 {
	return hashTorch | uint64(t.Facing)<<7 | uint64(t.Type.Uint8())<<10
}

func (ChiseledQuartz) Hash() uint64 {
	return hashChiseledQuartz
}

func (l Leaves) Hash() uint64 {
	return hashLeaves | uint64(l.Wood.Uint8())<<7 | uint64(boolByte(l.Persistent))<<11 | uint64(boolByte(l.ShouldUpdate))<<12
}

func (a Andesite) Hash() uint64 {
	return hashAndesite | uint64(boolByte(a.Polished))<<7
}

func (SeaLantern) Hash() uint64 {
	return hashSeaLantern
}

func (EmeraldBlock) Hash() uint64 {
	return hashEmeraldBlock
}

func (p PumpkinSeeds) Hash() uint64 {
	return hashPumpkinSeeds | uint64(p.Growth)<<7 | uint64(p.Direction)<<15
}

func (w WoodFence) Hash() uint64 {
	return hashWoodFence | uint64(w.Wood.Uint8())<<7
}

func (q QuartzPillar) Hash() uint64 {
	return hashQuartzPillar | uint64(q.Axis)<<7
}

func (SoulSoil) Hash() uint64 {
	return hashSoulSoil
}

func (AncientDebris) Hash() uint64 {
	return hashAncientDebris
}

func (LapisOre) Hash() uint64 {
	return hashLapisOre
}

func (CoalBlock) Hash() uint64 {
	return hashCoalBlock
}

func (GildedBlackstone) Hash() uint64 {
	return hashGildedBlackstone
}

func (c Clay) Hash() uint64 {
	return hashClay
}

func (b BeetrootSeeds) Hash() uint64 {
	return hashBeetrootSeeds | uint64(b.Growth)<<7
}

func (s Sand) Hash() uint64 {
	return hashSand | uint64(boolByte(s.Red))<<7
}

func (DirtPath) Hash() uint64 {
	return hashDirtPath
}

func (d Diorite) Hash() uint64 {
	return hashDiorite | uint64(boolByte(d.Polished))<<7
}

func (p Pumpkin) Hash() uint64 {
	return hashPumpkin | uint64(boolByte(p.Carved))<<7 | uint64(p.Facing)<<8
}

func (Netherrack) Hash() uint64 {
	return hashNetherrack
}

func (GlassPane) Hash() uint64 {
	return hashGlassPane
}

func (c CoralBlock) Hash() uint64 {
	return hashCoralBlock | uint64(c.Type.Uint8())<<7 | uint64(boolByte(c.Dead))<<11
}

func (t GlazedTerracotta) Hash() uint64 {
	return hashGlazedTerracotta | uint64(t.Colour.Uint8())<<7 | uint64(t.Facing)<<11
}

func (InvisibleBedrock) Hash() uint64 {
	return hashInvisibleBedrock
}

func (c ConcretePowder) Hash() uint64 {
	return hashConcretePowder | uint64(c.Colour.Uint8())<<7
}

func (IronOre) Hash() uint64 {
	return hashIronOre
}

func (SoulSand) Hash() uint64 {
	return hashSoulSand
}

func (IronBlock) Hash() uint64 {
	return hashIronBlock
}

func (w Wool) Hash() uint64 {
	return hashWool | uint64(w.Colour.Uint8())<<7
}

func (w Water) Hash() uint64 {
	return hashWater | uint64(boolByte(w.Still))<<7 | uint64(w.Depth)<<8 | uint64(boolByte(w.Falling))<<16
}

func (GoldOre) Hash() uint64 {
	return hashGoldOre
}

func (Bricks) Hash() uint64 {
	return hashBricks
}

func (c Cobblestone) Hash() uint64 {
	return hashCobblestone | uint64(boolByte(c.Mossy))<<7
}

func (Glass) Hash() uint64 {
	return hashGlass
}

func (IronBars) Hash() uint64 {
	return hashIronBars
}

func (l LitPumpkin) Hash() uint64 {
	return hashLitPumpkin | uint64(l.Facing)<<7
}

func (NetherQuartzOre) Hash() uint64 {
	return hashNetherQuartzOre
}

func (c Coral) Hash() uint64 {
	return hashCoral | uint64(c.Type.Uint8())<<7 | uint64(boolByte(c.Dead))<<11
}

func (l Log) Hash() uint64 {
	return hashLog | uint64(l.Wood.Uint8())<<7 | uint64(boolByte(l.Stripped))<<11 | uint64(l.Axis)<<12
}

func (EmeraldOre) Hash() uint64 {
	return hashEmeraldOre
}

func (d WoodDoor) Hash() uint64 {
	return hashWoodDoor | uint64(d.Wood.Uint8())<<7 | uint64(d.Facing)<<11 | uint64(boolByte(d.Open))<<13 | uint64(boolByte(d.Top))<<14 | uint64(boolByte(d.Right))<<15
}

func (Gravel) Hash() uint64 {
	return hashGravel
}

func (g Granite) Hash() uint64 {
	return hashGranite | uint64(boolByte(g.Polished))<<7
}

func (DragonEgg) Hash() uint64 {
	return hashDragonEgg
}

func (NetheriteBlock) Hash() uint64 {
	return hashNetheriteBlock
}

func (GoldBlock) Hash() uint64 {
	return hashGoldBlock
}

func (n NetherWart) Hash() uint64 {
	return hashNetherWart | uint64(n.Age)<<7
}

func (o Obsidian) Hash() uint64 {
	return hashObsidian | uint64(boolByte(o.Crying))<<7
}

func (b BoneBlock) Hash() uint64 {
	return hashBoneBlock | uint64(b.Axis)<<7
}

func (Glowstone) Hash() uint64 {
	return hashGlowstone
}

func (k Kelp) Hash() uint64 {
	return hashKelp | uint64(k.Age)<<7
}

func (DiamondBlock) Hash() uint64 {
	return hashDiamondBlock
}

func (l Lava) Hash() uint64 {
	return hashLava | uint64(boolByte(l.Still))<<7 | uint64(l.Depth)<<8 | uint64(boolByte(l.Falling))<<16
}

func (s Sponge) Hash() uint64 {
	return hashSponge | uint64(boolByte(s.Wet))<<7
}
