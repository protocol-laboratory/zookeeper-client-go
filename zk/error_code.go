package zk

type ErrorCode int32

const (
	EcOk                      ErrorCode = 0
	EcSystemError             ErrorCode = -1
	EcRuntimeInconsistency    ErrorCode = -2
	EcDataInconsistency       ErrorCode = -3
	EcConnectionLoss          ErrorCode = -4
	EcMarshallingError        ErrorCode = -5
	EcUnimplemented           ErrorCode = -6
	EcOperationTimeout        ErrorCode = -7
	EcBadArguments            ErrorCode = -8
	EcUnknownSession          ErrorCode = -12
	EcNewConfigNoQuorum       ErrorCode = -13
	EcReconfigInProgress      ErrorCode = -14
	EcAPIError                ErrorCode = -100
	EcNoNode                  ErrorCode = -101
	EcNoAuth                  ErrorCode = -102
	EcBadVersion              ErrorCode = -103
	EcNoChildrenForEphemerals ErrorCode = -108
	EcNodeExists              ErrorCode = -110
	EcNotEmpty                ErrorCode = -111
	EcSessionExpired          ErrorCode = -112
	EcInvalidCallback         ErrorCode = -113
	EcInvalidACL              ErrorCode = -114
	EcAuthFailed              ErrorCode = -115
	EcSessionMoved            ErrorCode = -118
	EcEphemeralOnLocalSession ErrorCode = -120
)
