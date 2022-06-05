package memcachemock

const (
	OperationFlushAll       Operation = "FlushAll"
	OperationGet            Operation = "Get"
	OperationGetMulti       Operation = "GetMulti"
	OperationSet            Operation = "Set"
	OperationAdd            Operation = "Add"
	OperationReplace        Operation = "Replace"
	OperationCompareAndSwap Operation = "CompareAndSwap"
	OperationDelete         Operation = "Delete"
	OperationIncrement      Operation = "Increment"
	OperationDecrement      Operation = "Decrement"
	OperationExists         Operation = "Exists"
	OperationTouch          Operation = "Touch"
	OperationDeleteAll      Operation = "DeleteAll"
	OperationPing           Operation = "Ping"
)

type Args []interface{}

type Return interface{}

type Operation string

type Mock struct {
	Operation Operation
	Args      Args

	Return Return
	Error  error
}
