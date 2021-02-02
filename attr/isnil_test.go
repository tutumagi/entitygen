package attr

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

type People struct{}

func Test_IsNil(t *testing.T) {
	{
		var a interface{} = (*People)(nil)
		Equal(t, a, nil)
		Equal(t, a == nil, false)
		Equal(t, isNil(a), true)
	}

	{
		var a interface{} = nil
		Equal(t, a, nil)
		Equal(t, a == nil, true)
		Equal(t, isNil(a), true)
	}

	{
		var a interface{} = ([]string)(nil)
		Equal(t, a, nil)
		Equal(t, a == nil, false)
		Equal(t, isNil(a), true)
	}

	{
		var a interface{} = (map[int32]string)(nil)
		Equal(t, a, nil)
		Equal(t, a == nil, false)
		Equal(t, isNil(a), true)
	}

	{
		var a interface{} = int32(0)
		NotEqual(t, a, nil)
		Equal(t, a == nil, false)
		Equal(t, isNil(a), false)
	}

	{
		var a interface{} = string("")
		NotEqual(t, a, nil)
		Equal(t, a == nil, false)
		Equal(t, isNil(a), false)
	}
}

func Benchmark_IsNil(b *testing.B) {
	var data []interface{} = []interface{}{
		(map[int32]string)(nil),
		nil,
		([]string)(nil),
		(*People)(nil),
		"",
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		float32(0),
		float64(0),
		complex(10, 100),
		complex(200, 300),
	}
	count := len(data)
	for i := 0; i < b.N; i++ {
		isNil(data[i%count])
	}

}

func Benchmark_IsNilPure(b *testing.B) {
	var data []interface{} = []interface{}{
		(map[int32]string)(nil),
		nil,
		([]string)(nil),
		(*People)(nil),
		"",
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		float32(0),
		float64(0),
		complex(10, 100),
		complex(200, 300),
	}
	count := len(data)
	for i := 0; i < b.N; i++ {
		is := data[i%count] == nil
		_ = is
	}

}
