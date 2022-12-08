package ledgerclient

import (
	"fmt"
	_neturl "net/url"
)

func addMetadataParams(queryParams _neturl.Values, metadata map[string]interface{}) {
	for k, v := range metadata {
		key := fmt.Sprintf("metadata[%s]", k)
		queryParams.Add(key, parameterToString(v, ""))
	}
}
