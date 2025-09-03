package utils

func ConvertResults[T any](results []interface{}) []T{
	var typedResults []T

	for _, result := range results {
		if typedResult, ok := result.(T); ok {
			typedResults = append(typedResults, typedResult)
		}
	}

	return typedResults
}