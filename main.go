package main

import (
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"

	"agent301/core"

)

func main() {
	fmt.Println(`

#!/bin/bash
merah="\e[31m"
kuning="\e[33m"
hijau="\e[32m"
biru="\e[34m"
UL="\e[4m"
bold="\e[1m"
italic="\e[3m"
reset="\e[m"
echo -e "$bold$italic$biru"
 echo " █████╗ ██╗     ███████╗ ██████╗ ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗  "
echo " ██╔══██╗██║     ██╔════╝██╔═══██╗████╗  ██║██╔═══██╗██║   ██║██╔══██╗ "
echo " ███████║██║     █████╗  ██║   ██║██╔██╗ ██║██║   ██║██║   ██║███████║ "
echo " ██╔══██║██║     ██╔══╝  ██║   ██║██║╚██╗██║██║   ██║╚██╗ ██╔╝██╔══██║ "
echo " ██║  ██║███████╗██║     ╚██████╔╝██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║ "
echo " ╚═╝  ╚═╝╚══════╝╚═╝      ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝ "
echo -e "$reset$bold$merah====================>>($hijau https://github.com/Alfonova-Node $merah)<<===================$reset\n"
																	
																	
																	
	
	`)
	fmt.Println(`ρσωєяє∂ ву : ALFONODES`)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yml")
	if err != nil {
		panic(err)
	}

	core.ProcessBot(config.Default())
}
