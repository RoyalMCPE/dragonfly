package block

import (
	"github.com/dragonfly-tech/dragonfly/dragonfly/block/material"
	"github.com/dragonfly-tech/dragonfly/dragonfly/item/inventory"
)

// Log is a naturally occurring block found in trees, primarily used to create planks. It comes in six
// species: oak, spruce, birch, jungle, acacia, and dark oak.
// Stripped log is a variant obtained by using an axe on a log.
type Log struct {
	// Wood is the type of wood of the log. This field must have one of the values found in the material
	// package. Using Log without a Wood type will panic.
	Wood material.Wood
	// Stripped specifies if the log is stripped or not.
	Stripped bool

	// Axis is the axis which the log block faces.
	Axis Axis
}

// Name returns the name of the log, including the wood type and whether it is stripped or not.
func (l Log) Name() (name string) {
	if l.Wood == nil {
		panic("log has no wood type")
	}
	if l.Stripped {
		return "Stripped " + l.Wood.Name() + " Log"
	}
	return l.Wood.Name() + " Log"
}

// Drops returns the drops of the log, which is always the block itself excluding the rotation.
func (l Log) Drops() []inventory.Item {
	return []inventory.Item{Log{
		Wood:     l.Wood,
		Stripped: l.Stripped,
	}}
}

func (l Log) Minecraft() (name string, properties map[string]interface{}) {
	switch l.Wood {
	case nil:
		panic("log has no wood type")
	case material.OakWood(), material.SpruceWood(), material.BirchWood(), material.JungleWood():
		return "minecraft:log", map[string]interface{}{"pillar_axis": l.Axis.String(), "old_log_type": l.Wood.Minecraft()}
	case material.AcaciaWood(), material.DarkOakWood():
		return "minecraft:log2", map[string]interface{}{"pillar_axis": l.Axis.String(), "new_log_type": l.Wood.Minecraft()}
	default:
		panic("invalid wood type")
	}
}

// allLogs returns a list of all possible log states.
func allLogs() (logs []Block) {
	f := func(axis Axis, stripped bool) {
		logs = append(logs, Log{Axis: axis, Stripped: stripped, Wood: material.OakWood()})
		logs = append(logs, Log{Axis: axis, Stripped: stripped, Wood: material.SpruceWood()})
		logs = append(logs, Log{Axis: axis, Stripped: stripped, Wood: material.BirchWood()})
		logs = append(logs, Log{Axis: axis, Stripped: stripped, Wood: material.JungleWood()})
		logs = append(logs, Log{Axis: axis, Stripped: stripped, Wood: material.AcaciaWood()})
		logs = append(logs, Log{Axis: axis, Stripped: stripped, Wood: material.DarkOakWood()})
	}
	for axis := Axis(0); axis < 3; axis++ {
		f(axis, true)
		f(axis, false)
	}
	return
}
