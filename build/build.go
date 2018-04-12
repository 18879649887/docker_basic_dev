package build

import (
	"../flag"
	"../container"
	"../network"
)

func Build(flagstring *flag.SFlag) {
	name, port1, port2 := flag.CheckParams(flagstring);
	
	switch(flagstring.Command) {
		case "network": network.CreateNetwork(name, port1, port2); break;
		case "run": container.CreateContainer(name); break;
		default: break;
	}
}
