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
        "line": "show port",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Port information"
        ]
    },
    {
        "name": "vty_show",
        "line": "show port all",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Port information",
            "All port information"
        ]
    },
    {
        "name": "vty_show",
        "line": "show port WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Port information",
            "Specify port number"
        ]
    },
    {
        "name": "vty_show",
        "line": "show running status",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "running information",
            "running status"
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
        "line": "show memory mempool",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory information",
            "Mempool information"
        ]
    },
    {
        "name": "vty_show",
        "line": "show memory mempool all",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory information",
            "Mempool information",
            "All mempool information"
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
