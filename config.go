package vty

import (
	"fmt"
	"github.com/coreswitch/cmd"
)

var (
	configParser *cmd.Node
)

/*
	tag:
	type: u32
	help: MD5 key id

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "ID must be between (1-255)"
	val_help: u32:1-255; MD5 key id

	commit:expression: $VAR(md5-key/) != ""; \
	       "Must add the md5-key for key-id $VAR(@)"

*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyId(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication md5 key-id WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: MD5 key
	syntax:expression: pattern $VAR(@) "^[^[:space:]]{1,16}$"; "MD5 key must be 16 characters or less"
	val_help: MD5 Key (16 characters or less)

	# If this node is created
	create:
		vtysh -c "configure terminal" -c "interface $VAR(../../../../../../@)" \
	             -c "ip ospf message-digest-key $VAR(../@) md5 $VAR(@)"

	# If the value of this node is changed
	update:
		vtysh -c "configure terminal" -c "interface $VAR(../../../../../../@)" \
	             -c "no ip ospf message-digest-key $VAR(../@)" \
	             -c "ip ospf message-digest-key $VAR(../@) md5 $VAR(@)"

	# If this node is deleted
	delete:
	        vtysh -c "configure terminal" -c "interface $VAR(../../../../../../@)" \
	             -c "no ip ospf message-digest-key $VAR(../@)"
*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyIdMd5Key(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication md5 key-id WORD md5-key WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf message-digest-key ", Args[1], " md5 ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("no ip ospf message-digest-key ", Args[1]))
	}
	return cmd.Success
}

/*
	help: MD5 parameters

	create: vtysh -c "configure terminal" \
		-c "interface $VAR(../../../../@)" \
		-c "no ip ospf authentication" \
		-c "ip ospf authentication message-digest"

	delete: vtysh -c "configure terminal" \
		-c "interface $VAR(../../../../@)" \
	        -c "no ip ospf authentication"
*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication md5
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication",
			"ip ospf authentication message-digest")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication")
	}
	return cmd.Success
}

/*
	help: OSPF interface authentication

*/
func quaggaInterfacesInterfaceIpv4OspfAuthentication(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Plain text password
	syntax:expression: pattern $VAR(@) "^[^[:space:]]{1,8}$" ; "Password must be 8 characters or less"
	val_help: Plain text password (8 characters or less)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../../@)" \
		 -c "no ip ospf authentication " -c "ip ospf authentication " \
		 -c "ip ospf authentication-key $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../../@)" \
		 -c "no ip ospf authentication " -c "no ip ospf authentication-key"
*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationPlaintextPassword(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication plaintext-password WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication",
			"ip ospf authentication",
			fmt.Sprint("ip ospf authentication-key ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication",
			"no ip ospf authentication-key")
	}
	return cmd.Success
}

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
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("bandwidth ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no bandwidth ")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interface cost
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; OSPF interface cost

	update:vtysh -c "configure terminal" \
		-c "interface $VAR(../../../@)" \
		-c "ip ospf cost $VAR(@)"
	delete:vtysh -c "configure terminal" \
		-c "interface $VAR(../../../@)" \
		-c "no ip ospf cost"
*/
func quaggaInterfacesInterfaceIpv4OspfCost(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf cost WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf cost ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf cost")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval after which neighbor is dead
	default: 40
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; OSPF dead interval in seconds (default 40)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf dead-interval $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf dead-interval"
*/
func quaggaInterfacesInterfaceIpv4OspfDeadInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf dead-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf dead-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf dead-interval")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between hello packets
	default: 10
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Interval between OSPF hello packets in seconds (default 10)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf hello-interval $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf hello-interval"
*/
func quaggaInterfacesInterfaceIpv4OspfHelloInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf hello-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf hello-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf hello-interval")
	}
	return cmd.Success
}

/*
	help: Disable Maximum Transmission Unit (MTU) mismatch detection
	create:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf mtu-ignore"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf mtu-ignore"
*/
func quaggaInterfacesInterfaceIpv4OspfMtuIgnore(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf mtu-ignore
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ip ospf mtu-ignore")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf mtu-ignore")
	}
	return cmd.Success
}

/*
	type: txt
	help: Network type
	syntax:expression: $VAR(@) in "broadcast", "non-broadcast", "point-to-multipoint", "point-to-point"; \
	       "Must be (broadcast|non-broadcast|point-to-multipoint|point-to-point)"
	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf network $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf network"

	val_help: broadcast; Broadcast network type
	val_help: non-broadcast; Non-broadcast network type
	val_help: point-to-multipoint; Point-to-multipoint network type
	val_help: point-to-point; Point-to-point network type
*/
func quaggaInterfacesInterfaceIpv4OspfNetwork(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf network WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf network ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf network")
	}
	return cmd.Success
}

/*
	help: Open Shortest Path First (OSPF) parameters
*/
func quaggaInterfacesInterfaceIpv4Ospf(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Router priority
	default: 1
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Must be between 0-255"
	val_help: u32:0-255; Priority (default 1)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf priority $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf priority"
*/
func quaggaInterfacesInterfaceIpv4OspfPriority(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf priority WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf priority ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf priority")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between retransmitting lost link state advertisements
	default: 5
	syntax:expression: $VAR(@) >= 3 && $VAR(@) <= 65535; "Must be between 3-65535"
	val_help: u32: 3-65535; Retransmit interval in seconds (default 5)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
		-c "ip ospf retransmit-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
		-c "no ip ospf retransmit-interval"
*/
func quaggaInterfacesInterfaceIpv4OspfRetransmitInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf retransmit-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf retransmit-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf retransmit-interval")
	}
	return cmd.Success
}

/*
	type: u32
	help: Link state transmit delay
	default: 1
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Transmit delay in seconds (default 1)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf transmit-delay $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf transmit-delay"
*/
func quaggaInterfacesInterfaceIpv4OspfTransmitDelay(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf transmit-delay WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf transmit-delay ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf transmit-delay")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interface cost
	default: 1
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; OSPFv3 cost

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 cost $VAR(@)"

*/
func quaggaInterfacesInterfaceIpv6Ospfv3Cost(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 cost WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 cost ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 cost")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval after which neighbor is declared dead
	default: 40
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Neighbor dead interval in seconds (default 40)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 dead-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 dead-interval 40"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3DeadInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 dead-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 dead-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 dead-interval 40")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between hello packets
	default: 10
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Interval between OSPFv3 hello packets in seconds (default 10)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 hello-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 hello-interval 10"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3HelloInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 hello-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 hello-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 hello-interval 10")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interface MTU
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Interface MTU

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 ifmtu $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ipv6 ospf6 ifmtu"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3Ifmtu(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 ifmtu WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 ifmtu ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 ifmtu")
	}
	return cmd.Success
}

/*
	type: u32
	help: Instance-id
	default: 0
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Must be between 0-255"
	val_help: u32:0-255; Instance Id (default 0)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 instance-id $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 instance-id 0"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3InstanceId(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 instance-id WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 instance-id ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 instance-id 0")
	}
	return cmd.Success
}

/*
	help: Disable Maximum Transmission Unit mismatch detection
	create:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 mtu-ignore"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ipv6 ospf6 mtu-ignore"

*/
func quaggaInterfacesInterfaceIpv6Ospfv3MtuIgnore(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 mtu-ignore
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 mtu-ignore")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 mtu-ignore")
	}
	return cmd.Success
}

/*
	help: IPv6 Open Shortest Path First (OSPFv3)
*/
func quaggaInterfacesInterfaceIpv6Ospfv3(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable forming of adjacency
	create: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 passive"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ipv6 ospf6 passive"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3Passive(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 passive
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 passive")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 passive")
	}
	return cmd.Success
}

/*
	type: u32
	help: Router priority
	default: 1
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Must be between 0-255"
	val_help: u32:0-255; Priority (default 1)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 priority $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 priority 1"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3Priority(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 priority WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 priority ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 priority 1")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between retransmitting lost link state advertisements
	default: 5
	syntax:expression: $VAR(@) >= 3 && $VAR(@) <= 65535; "Must be between 3-65535"
	val_help: u32:3-65535; Retransmit interval in seconds (default 5)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 retransmit-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 retransmit-interval 5"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3RetransmitInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 retransmit-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 retransmit-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 retransmit-interval 5")
	}
	return cmd.Success
}

/*
	type: u32
	help: Link state transmit delay
	default: 1
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Link state transmit delay (default 1)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 transmit-delay $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 transmit-delay 1"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3TransmitDelay(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 transmit-delay WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 transmit-delay ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 transmit-delay 1")
	}
	return cmd.Success
}

func initConfig() {
	configParser = cmd.NewParser()
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "md5", "key-id", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyId)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "md5", "key-id", "WORD", "md5-key", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyIdMd5Key)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "md5"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication"},
		quaggaInterfacesInterfaceIpv4OspfAuthentication)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "plaintext-password", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationPlaintextPassword)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "bandwidth", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfBandwidth)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "cost", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfCost)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "dead-interval", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfDeadInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "hello-interval", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfHelloInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "mtu-ignore"},
		quaggaInterfacesInterfaceIpv4OspfMtuIgnore)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "network", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfNetwork)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf"},
		quaggaInterfacesInterfaceIpv4Ospf)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "priority", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfPriority)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "retransmit-interval", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfRetransmitInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "transmit-delay", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfTransmitDelay)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "cost", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3Cost)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "dead-interval", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3DeadInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "hello-interval", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3HelloInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "ifmtu", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3Ifmtu)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "instance-id", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3InstanceId)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "mtu-ignore"},
		quaggaInterfacesInterfaceIpv6Ospfv3MtuIgnore)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3"},
		quaggaInterfacesInterfaceIpv6Ospfv3)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "passive"},
		quaggaInterfacesInterfaceIpv6Ospfv3Passive)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "priority", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3Priority)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "retransmit-interval", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3RetransmitInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "transmit-delay", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3TransmitDelay)
}
