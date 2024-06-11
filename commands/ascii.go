package commands

import "lmdb-cli/core"

var ascii = []byte(`LMDB CLI`)

type Ascii struct {
}

func (cmd Ascii) Execute(context *core.Context, input []byte) (err error) {
	context.Output(ascii)
	return nil
}
