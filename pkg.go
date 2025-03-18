package mypackage

type Result struct {
	Value int
}

func ProcessData(data []byte) (*Result, error) {
	result := &Result{
		Value: len(data),
	}
	return result, nil
}

func Validate(result *Result) bool {
	return result.Value > 0
}
