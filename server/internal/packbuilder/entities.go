package packbuilder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/df-mc/dragonfly/server/world"
)

func buildEntities(dir string) (count int, lang []string) {
	if err := os.Mkdir(filepath.Join(dir, "entity"), os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "textures/entity"), os.ModePerm); err != nil {
		panic(err)
	}

	for name, entity := range world.CustomEntities() {
		lang = append(lang, fmt.Sprintf("entity.%s.name=%s", name, entity.Name()))
	}
	return
}

func buildEntity(dir, identifier, name string, entity world.CustomEntity) {
	entityData, err := json.Marshal(map[string]any{
		"format_version": formatVersion,
		"minecraft:client_entity": map[string]any{
			"identifier": identifier,
		},
	})
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "entity", fmt.Sprint("%s.json", name)), entityData, 0666); err != nil {
		panic(err)
	}
}
