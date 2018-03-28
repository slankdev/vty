package vty

import (
	"fmt"
	"github.com/coreswitch/cmd"
)

var (
	configParser *cmd.Node
)

/*
 * Sample
 */
/*
	type: u32
	help: Bandwidth of interface (kilobits/sec)
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 10000000; "Must be between 1-10000000"
	val_help: u32:1-10000000; Bandwidth in kilobits/sec (for calculating OSPF cost)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "bandwidth $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no bandwidth"
*/
func quaggaInterfacesInterfaceIpv4OspfBandwidth(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf bandwidth WORD
	switch Cmd {
	case cmd.Set:
		fmt.Print("set")
		kamueeVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("bandwidth ", Args[1]))
	case cmd.Delete:
		fmt.Print("delete")
		kamueeVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no bandwidth ")
	}
	return cmd.Success
}

func initConfig() {
	configParser = cmd.NewParser()
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "bandwidth", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfBandwidth)
}

