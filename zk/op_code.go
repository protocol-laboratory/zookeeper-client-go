package zk

type OpCode int32

const (
	OpCloseSession         OpCode = -11
	OpCreateSession        OpCode = -10
	OpError                OpCode = -1
	OpNotification         OpCode = 0
	OpCreate               OpCode = 1
	OpDelete               OpCode = 2
	OpExists               OpCode = 3
	OpGetData              OpCode = 4
	OpSetData              OpCode = 5
	OpGetAcl               OpCode = 6
	OpSetAcl               OpCode = 7
	OpGetChildren          OpCode = 8
	OpSync                 OpCode = 9
	OpDummy10              OpCode = 10
	OpPing                 OpCode = 11
	OpGetChildren2         OpCode = 12
	OpCheck                OpCode = 13
	OpMulti                OpCode = 14
	OpCreate2              OpCode = 15
	OpReconfig             OpCode = 16
	OpCheckWatches         OpCode = 17
	OpRemoveWatches        OpCode = 18
	OpCreateContainer      OpCode = 19
	OpDeleteContainer      OpCode = 20
	OpCreateTtl            OpCode = 21
	OpMultiRead            OpCode = 22
	OpAuth                 OpCode = 100
	OpSetWatches           OpCode = 101
	OpSasl                 OpCode = 102
	OpGetEphemerals        OpCode = 103
	OpGetAllChildrenNumber OpCode = 104
	OpSetWatches2          OpCode = 105
	OpAddWatch             OpCode = 106
	OpWhoAmI               OpCode = 107
)
