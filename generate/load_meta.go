package main

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/hsiafan/cocoa/generate/data"
)

//go:embed meta
var metaFS embed.FS

type ModuleTypes struct {
	Name  string          // module name
	Types []data.TypeInfo // types in this module
}

// load all module meta datas
func loadAllMeta() []ModuleTypes {
	entries, err := metaFS.ReadDir("meta")
	if err != nil {
		panic(err)
	}
	var moduleTypeses []ModuleTypes
	for _, de := range entries {
		mts := loadModule(de.Name())
		moduleTypeses = append(moduleTypeses, mts)
	}
	return moduleTypeses
}

func loadModule(name string) ModuleTypes {
	log.Println("load module:", name)
	entries, err := metaFS.ReadDir("meta/" + name)
	if err != nil {
		panic(err)
	}
	var typeInfos []data.TypeInfo
	for _, e := range entries {
		ti := loadOneTypeInfo(name, strings.TrimSuffix(e.Name(), ".json"))
		typeInfos = append(typeInfos, ti)
	}
	return ModuleTypes{
		Name:  name,
		Types: typeInfos,
	}
}

var typeInfoCache map[string]data.TypeInfo = map[string]data.TypeInfo{}

func loadOneTypeInfo(module string, name string) data.TypeInfo {
	key := module + "." + name
	ti, ok := typeInfoCache[key]
	if ok {
		return ti
	}
	ti = loadOneTypeInner(module, name)
	typeInfoCache[key] = ti
	return ti
}

func loadOneTypeInner(module string, name string) data.TypeInfo {
	log.Println("load type meta:", name)
	f, err := metaFS.Open("meta/" + module + "/" + name + ".json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	d, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var kt data.KindType
	err = json.Unmarshal(d, &kt)
	if err != nil {
		panic(err)
	}
	switch kt.Kind {
	case data.KindClass:
		var c data.Class
		err = json.Unmarshal(d, &c)
		if err != nil {
			panic(err)
		}
		return &c
	case data.KindProtocol:
		var p data.Protocol
		err = json.Unmarshal(d, &p)
		if err != nil {
			panic(err)
		}
		return &p
	case data.KindAlias:
		var a data.Alias
		err = json.Unmarshal(d, &a)
		if err != nil {
			panic(err)
		}
		return &a
	case data.KindStruct:
		var s data.Struct
		err = json.Unmarshal(d, &s)
		if err != nil {
			panic(err)
		}
		return &s
	default:
		panic("unkown type: " + string(kt.Kind))
	}
}
