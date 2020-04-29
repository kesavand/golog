/*
sample file to test the logger
*/
package main

import (
	"github.com/kesavand/golog"
	"fmt"
)


func main() {
	fmt.Println("The example is used to verify the logging library")
	logger,err := golog.NewLogger("zap");
	if err != nil {
		fmt.Println("Unable to create logger",err)
		return
	}

	logger.Debug("The log library is success")
	fmt.Println("Logger created Successfully")

}
