/*
Package log  library is an abstract over the zap library.
*/
package golog

import (
        "github.com/kesavand/golog/zaplog"
        "github.com/kesavand/golog/api"
)


/*
NewLogger Creates new logger based on the logger type
*/
func NewLogger(loglib string) (api.Logger, error) {

        /*call the logger based on the loglib*/
        return zaplog.NewZapLogger()

}

