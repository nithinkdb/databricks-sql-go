// Code generated by "stringer -type=clientMethod -trimprefix=clientMethod"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[clientMethodUnknown-0]
	_ = x[clientMethodOpenSession-1]
	_ = x[clientMethodCloseSession-2]
	_ = x[clientMethodFetchResults-3]
	_ = x[clientMethodGetResultSetMetadata-4]
	_ = x[clientMethodExecuteStatement-5]
	_ = x[clientMethodGetOperationStatus-6]
	_ = x[clientMethodCloseOperation-7]
	_ = x[clientMethodCancelOperation-8]
}

const _clientMethod_name = "UnknownOpenSessionCloseSessionFetchResultsGetResultSetMetadataExecuteStatementGetOperationStatusCloseOperationCancelOperation"

var _clientMethod_index = [...]uint8{0, 7, 18, 30, 42, 62, 78, 96, 110, 125}

func (i clientMethod) String() string {
	if i < 0 || i >= clientMethod(len(_clientMethod_index)-1) {
		return "clientMethod(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _clientMethod_name[_clientMethod_index[i]:_clientMethod_index[i+1]]
}
