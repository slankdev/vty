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
    }
]
`

const execCmdSpec = `
[
    {
        "name": "vty_exec",
        "line": "clear thread cpu WORD",
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
