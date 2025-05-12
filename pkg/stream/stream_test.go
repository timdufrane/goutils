package stream

import (
	"maps"
	"slices"
	"testing"
)

var elementsIntSlice = []int{
	1, 3, 5, 8, 10, 2, 3, 5,
}

// TestFindElement tests whether the Find function returns the expected
// array
func TestFindElement(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	res := stream.Find(func(ii int) bool {
		return ii == 3
	})

	if !slices.Equal(res, []int{3, 3}) {
		t.Errorf("Expected equal slice to %v, got %v", []any{3, 3}, res)
	}
}

func TestFilter(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	stream.Filter(func(ii int) bool {
		return ii%2 == 0
	})

	if !slices.Equal(stream.Slice(), []int{8, 10, 2}) {
		t.Errorf("Expected equal slice to %v, got %v", []any{3, 3}, stream.Slice())
	}

	stream.Slice()[0] = 100

	if stream.Slice()[0] == 100 {
		t.Errorf("Expected index 0 of stream slice to be 1, got %v", stream.Slice()[0])
	}
}

func TestFindFirst(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	res, err := stream.FindFirst(func(ii int) bool {
		return ii%2 == 0
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if res != 8 {
		t.Errorf("Expected value 8, got %v", res)
	}
}

func TestLength(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	if stream.Length() != 8 {
		t.Errorf("Expected stream contents length of 8, got %v", stream.Length())
	}

	// Now let's filter and see if it changes
	stream.Filter(func(in int) bool {
		return in%3 == 0
	})

	if stream.Length() != 2 {
		t.Errorf("Expected filtered stream contents length of 2, got %v", stream.Length())
	}
}

func TestMap(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	stream.Map(func(in int) int {
		return in * 3
	})

	expectedResult := []int{3, 9, 15, 24, 30, 6, 9, 15}

	if !slices.Equal(stream.Slice(), expectedResult) {
		t.Errorf("Expected slice contents of %v, got %v", expectedResult, stream.Slice())
	}
}

func TestForEach(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	compareMap := map[int]int{
		0: 1, 1: 3, 2: 5, 3: 8, 4: 10, 5: 2, 6: 3, 7: 5,
	}

	outMap := make(map[int]int, 0)

	stream.Foreach(func(ii int, ele int) {
		outMap[ii] = ele
	})

	if !maps.Equal(outMap, compareMap) {
		t.Errorf("Expected map contents of %v, got %v", compareMap, outMap)
	}
}

func TestChaining(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	stream.Filter(func(a int) bool { return a%2 == 0 || a%3 == 0 }).Map(func(a int) int { return a + 3 }).Foreach(func(ii, a int) {
		var out string
		if a%2 == 0 {
			out += "Fizz"
		}

		if a%3 == 0 {
			out += "Buzz"
		}
	})

	compareSlice := []int{6, 11, 13, 5, 6}

	if !slices.Equal(stream.Slice(), compareSlice) {
		t.Errorf("Expected slice contents of %v, got %v", compareSlice, stream.Slice())
	}
}

func TestReducer(t *testing.T) {
	stream := NewStream(elementsIntSlice)
	res := stream.Reduce(func(previous, current int) int {
		return previous + current
	})

	if res != 37 {
		t.Errorf("Expected result of 37, got %v", res)
	}
}
