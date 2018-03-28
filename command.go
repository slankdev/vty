package vty

var showCmdMap = map[string]func(string) *string{
	"vty_show": vtyShow,
}

var execCmdMap = map[string]func(string) *string{
	"vty_exec": vtyExec,
}

func vtyShow(line string) *string {
	return kamueeVtysh(line)
}

func vtyExec(line string) *string {
	return kamueeVtysh(line)
}

const showCmdSpec = `
[
    {
        "name": "vty_show",
        "line": "show author",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Author information"
        ]
    },
    {
        "name": "vty_show",
        "line": "show running-config",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "running configuration"
        ]
    },
    {
        "name": "vty_show",
        "line": "show history",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Display the session command history"
        ]
    },
    {
        "name": "vty_show",
        "line": "show logging",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Show current logging configuration"
        ]
    },
    {
        "name": "vty_show",
        "line": "show memory",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics"
        ]
    },
    {
        "name": "vty_show",
        "line": "show memory all",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "All memory statistics"
        ]
    },
    {
        "name": "vty_show",
        "line": "show startup-config",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Contentes of startup configuration"
        ]
    },
    {
        "name": "vty_show",
        "line": "show thread cpu",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Thread information",
            "Thread CPU usage"
        ]
    },
    {
        "name": "vty_show",
        "line": "show version",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Displays zebra version"
        ]
    },
    {
        "name": "vty_show",
        "line": "show work-queues",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Work Queue information"
        ]
    },
]
`

const execCmdSpec = `
[
    {
        "name": "vty_exec",
        "line": "clear bgp * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 view WORD * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp view WORD * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp view WORD * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp * in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 * in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp view WORD * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp * out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 * out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp view WORD * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp <1-4294967295> in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 <1-4294967295> in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp <1-4294967295> soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 <1-4294967295> soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp <1-4294967295> in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp <1-4294967295> soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 <1-4294967295> in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 <1-4294967295> soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp <1-4294967295> out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp <1-4294967295> soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 <1-4294967295> out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 <1-4294967295> soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp external in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 external in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp external soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 external soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp external in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp external soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 external WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 external soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp external out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp external soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 external WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 external soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 peer-group WORD in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp peer-group WORD in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 peer-group WORD soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp peer-group WORD soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 peer-group WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 peer-group WORD soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp peer-group WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp peer-group WORD soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 peer-group WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 peer-group WORD soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp peer-group WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp peer-group WORD soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 view WORD (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "BGP view",
            "view name",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp view WORD (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp view WORD *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * vpnv4 unicast soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * vpnv4 unicast in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * vpnv4 unicast soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * vpnv4 unicast out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp * vpnv4 unicast soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp dampening",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp dampening A.B.C.D",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information",
            "Network to clear damping information"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp dampening A.B.C.D A.B.C.D",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information",
            "Network to clear damping information",
            "Network mask"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp dampening A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp external",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 external",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp external soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family modifier",
            "Address Family modifier"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp ipv6 peer-group WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear bgp peer-group WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp peer-group WORD soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp view WORD (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip prefix-list",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "Build a prefix list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ip prefix-list WORD A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ipv6 prefix-list",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IPv6 information",
            "Build a prefix list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ipv6 prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear ipv6 prefix-list WORD X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "vty_exec",
        "line": "clear thread cpu [FILTER]",
        "mode": "exec",
        "helps": [
            "Clear stored data",
            "Thread information",
            "Thread CPU usage",
            "Display filter (rwtexb)"
        ]
    }
]
`
