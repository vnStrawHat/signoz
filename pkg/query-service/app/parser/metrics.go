package parser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.signoz.io/signoz/pkg/query-service/model"
)

func ParseMetricAutocompleteTagParams(r *http.Request) (*model.MetricAutocompleteTagParams, *model.ApiError) {

	metricName := r.URL.Query().Get("metricName")
	if len(metricName) == 0 {
		err := fmt.Errorf("metricName not present in params")
		return nil, &model.ApiError{Typ: model.ErrorBadData, Err: err}
	}

	tagsStr := r.URL.Query().Get("tags")

	// parsing tags
	var tags map[string]string
	if tagsStr != "" && len(tagsStr) != 0 {

		err := json.Unmarshal([]byte(tagsStr), &tags)
		if err != nil {
			return nil, &model.ApiError{Typ: model.ErrorBadData, Err: fmt.Errorf("unable to parse tags in params: %v", err)}
		}
	}

	matchText := r.URL.Query().Get("match")

	tagKey := r.URL.Query().Get("tagKey")

	metricAutocompleteTagParams := &model.MetricAutocompleteTagParams{
		MetricName: metricName,
		MetricTags: tags,
		Match:      matchText,
		TagKey:     tagKey,
	}

	return metricAutocompleteTagParams, nil
}
