package vty

import (
	"fmt"
	"github.com/coreswitch/cmd"
)

var (
	configParser *cmd.Node
)

func kamueeInterfacesInterfaceMirror(Cmd int, Args cmd.Args) int {
  // interfaces interface WORD mirror WORD
  switch Cmd {
  case cmd.Set:
		kamueeVtysh(
      fmt.Sprintf("set port %s mirror %s",
      Args[0], Args[1]))
	case cmd.Delete:
		kamueeVtysh(
      fmt.Sprintf("set port %s no mirror %s",
      Args[0], Args[1]))
  }
  return cmd.Success
}

func initConfig() {
	configParser = cmd.NewParser()
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "mirror", "WORD"},
		kamueeInterfacesInterfaceMirror)
}

