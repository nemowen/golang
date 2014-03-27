package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
	"path"
)

func main() {
	paths, _ := os.Getwd()
	tos := path.Join(paths, "/workspace/.metadata/.plugins/org.eclipse.core.runtime/.settings/com.comresource.eshoes.pos.client.prefs")
	fmt.Println(tos)
	c, err := goconfig.LoadConfigFile(tos)
	if err != nil {
		fmt.Println(err)
	} else {
		v, err := c.GetValue("", "cr.update")
		if err != nil {
			ok := c.SetValue("", "cr.update", "http://192.168.2.19:2000/posUpdate")
			fmt.Println(ok)
			goconfig.SaveConfigFile(c, tos)
			v2, _ := c.GetValue("", "cr.update")
			fmt.Println("...", v2)
		} else {
			fmt.Println(v)
		}
	}

}
