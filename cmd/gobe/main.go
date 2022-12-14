package main

import (
	"fmt"
	"go/types"
	"os"
	"sort"

	"github.com/lemon-mint/gobe/build"
	"github.com/lemon-mint/gobe/build/golang"
	"github.com/lemon-mint/gobe/typeconv"
	"github.com/lemon-mint/gobe/typeinfo"
	"golang.org/x/tools/go/loader"
)

func main() {
	var conf loader.Config
	conf.Import(os.Args[1])
	lp, err := conf.Load()
	if err != nil {
		panic(err)
	}
	var typeinfoMap = make(map[string]typeinfo.Type)
	var ctx = &typeconv.TypeConvCtx{}
	for _, pkg := range lp.InitialPackages() {
		for _, objs := range pkg.Defs {
			if tn, ok := objs.(*types.TypeName); ok {
				switch t := tn.Type().(type) {
				case *types.Named:
					if t.TypeParams() != nil &&
						t.TypeParams().Len() > 0 &&
						!(t.TypeArgs() != nil && t.TypeArgs().Len() == t.TypeParams().Len()) {
						continue
					}
					key := t.Obj().Name()
					if _, ok := typeinfoMap[key]; !ok {
						v := typeconv.TypeConv(ctx, t)
						if v != typeinfo.InvalidType {
							typeinfoMap[key] = v
						}
					} else {
						panic(fmt.Sprint("duplicate type name:", key))
					}
				}
			}
		}
	}

	/*
		for k, v := range typeinfoMap {
				fmt.Printf("==== %s ====\n", k)
				fmt.Printf("kind: %s\n", typeinfo.TypeKind(v).String())
				fmt.Printf("string: %s\n", v.String())
				switch v := v.(type) {
				case *typeinfo.StructType:
					fmt.Println("FIELDS:")
					for _, f := range v.Fields {
						fmt.Printf("  %s: %s\n", f.Name, f.Type.String())
					}
				}
				fmt.Println(strings.Repeat("=", len(k)+10))
			}
	*/
	keys := make([]string, 0, len(typeinfoMap))
	for k := range typeinfoMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	nis := make([]build.NamedInfo, len(typeinfoMap))
	for i, k := range keys {
		nis[i] = build.NamedInfo{
			Name: k,
			Data: build.Conv(typeinfoMap[k], nil),
		}
	}

	for i := range nis {
		fmt.Printf("%s\n", golang.Generate(nis[i]))
	}
}
