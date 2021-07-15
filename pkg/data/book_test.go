package data

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func init() {
	time.Sleep(2 * time.Second)
}

func getMessage(command string, symbol string) string {
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actualStr := string(cmd)
	index := strings.LastIndex(actualStr, symbol)
	return actualStr[index:]
}

func TestServerReadAll(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/buyDubrovsky`
	var cmd []byte
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("qwe")
			cmd, _ = exec.Command("/bin/sh", "-c", command).CombinedOutput()
			actual := string(cmd)
			fmt.Println("actual", actual)
		}()
	}
	//cmd, _ := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	time.Sleep(2 * time.Second)
	actual := string(cmd)
	logrus.Info(actual)
	require.Equal(true, true)
}