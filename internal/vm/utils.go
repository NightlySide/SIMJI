package vm

import (
	"bufio"
	"fmt"
	"os"
	"simji/internal/log"
)

func getNumberInput() int {
	stdin := bufio.NewReader(os.Stdin)

    var value int

    for {
		fmt.Print("[SCALL 0] Enter : R1 <= ")
        _, err := fmt.Fscan(stdin, &value)
        if err == nil {
            break
        }

        stdin.ReadString('\n')
        log.GetLogger().Error("Input not valid. Please enter a NUMBER.")
    }

    return value
}