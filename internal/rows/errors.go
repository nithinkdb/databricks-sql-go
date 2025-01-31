package rows

import "fmt"

var errRowsNoClient = "databricks: instance of Rows missing client"
var errRowsNilRows = "databricks: nil Rows instance"
var errRowsUnknowRowType = "databricks: unknown rows representation"
var errRowsCloseFailed = "databricks: Rows instance Close operation failed"
var errRowsMetadataFetchFailed = "databricks: Rows instance failed to retrieve result set metadata"

func errRowsInvalidColumnIndex(index int) string {
	return fmt.Sprintf("databricks: invalid column index: %d", index)
}
