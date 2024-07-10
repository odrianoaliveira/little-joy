package pair

type Pair[T any] struct {
	Key   string `json:"key"`
	Value T      `json:"value"`
}
