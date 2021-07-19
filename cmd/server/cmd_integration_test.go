package main

import (
	"github.com/stretchr/testify/require"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func init() {
	time.Sleep(2 * time.Second)
}

func containsBook(str string) bool {
	return strings.Contains(str, "You buy")
}

func containsSerialize(str string) bool {
	return strings.Contains(str, "serialize access due")
}

func TestServerBuyNoEqual(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/buy/Dubrovsky`
	expected := "falsefalse"
	str := ""
	for i := 0; i < 5; i++ {
		go func() {
			strCmd := runCommand(command)
			if containsBook(strCmd) || containsSerialize(strCmd) {
				str += "true"
			} else {
				str += "false"
			}
		}()
	}
	time.Sleep(1 * time.Second)
	require.NotEqual(expected, str)
}

func TestServerBuyEqual(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/buy/Dubrovsky`
	expected := "truetruetruetruetrue"
	str := ""
	for i := 0; i < 5; i++ {
		go func() {
			strCmd := runCommand(command)
			if containsBook(strCmd) || containsSerialize(strCmd) {
				str += "true"
			} else {
				str += "false"
			}
		}()
	}
	time.Sleep(1 * time.Second)
	require.Equal(expected, str)
}

func TestServerBuyContainsTrue(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/buy/Dubrovsky`
	str := ""
	for i := 0; i < 5; i++ {
		go func() {
			strCmd := runCommand(command)
			if containsBook(strCmd) || containsSerialize(strCmd) {
				str += "true"
			} else {
				str += "false"
			}
		}()
	}
	time.Sleep(1 * time.Second)
	require.Contains(str, "true")
}

func TestServerBuyNoContainsFalse(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/buy/Dubrovsky`
	str := ""

	for i := 0; i < 5; i++ {
		go func() {
			strCmd := runCommand(command)
			if containsBook(strCmd) || containsSerialize(strCmd) {
				str += "true"
			} else {
				str += "false"
			}
		}()
	}
	time.Sleep(1 * time.Second)
	require.NotContains(str, "false")
}

func runCommand(command string) string {
	cmd, _ := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	strCmd := string(cmd)
	return strCmd
}
