package setop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet_Includes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name        string
		setElements []string
		str         string
		exp         bool
	}{
		{
			name:        "should include",
			setElements: []string{"cat", "dog", "mouse"},
			str:         "dog",
			exp:         true,
		},
		{
			name:        "should not include",
			setElements: []string{"cat", "dog", "mouse"},
			str:         "parrot",
			exp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.setElements...)

			if got := set.Includes(tt.str); got != tt.exp {
				t.Errorf("Set.Includes() = %v, want %v", got, tt.exp)
			}
		})
	}
}

func TestSet_IncludesSet(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		setElements  []string
		setElements2 []string
		exp          bool
	}{
		{
			name:         "equal - should include",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          true,
		},
		{
			name:         "smaller - should include",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog"},
			exp:          true,
		},
		{
			name:         "one - should include",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat"},
			exp:          true,
		},
		{
			name:         "bigger - should not include",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "mouse", "parrot"},
			exp:          false,
		},
		{
			name:         "wrong - should not include",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "parrot"},
			exp:          false,
		},
		{
			name:         "one wrong - should not include",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"parrot"},
			exp:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.setElements...)
			set2 := NewSet(tt.setElements2...)

			if got := set.IncludesSet(set2); got != tt.exp {
				t.Errorf("Set.Includes() = %v, want %v", got, tt.exp)
			}
		})
	}
}

func TestSet_EqualTo(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		setElements  []string
		setElements2 []string
		exp          bool
	}{
		{
			name:         "equal - should be equal",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          true,
		},
		{
			name:         "equal shuffled - should be equal",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"dog", "cat", "mouse"},
			exp:          true,
		},
		{
			name:         "smaller - should not be equal",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog"},
			exp:          false,
		},
		{
			name:         "bigger - should not be equal",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "mouse", "parrot"},
			exp:          false,
		},
		{
			name:         "wrong - should not be equal",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "parrot"},
			exp:          false,
		},
		{
			name:         "one wrong - should not be equal",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"parrot"},
			exp:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.setElements...)
			set2 := NewSet(tt.setElements2...)

			if got := set.EqualTo(set2); got != tt.exp {
				t.Errorf("Set.EqualTo() = %v, want %v", got, tt.exp)
			}
		})
	}
}

func TestSet_LargerThan(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		setElements  []string
		setElements2 []string
		exp          bool
	}{
		{
			name:         "bigger - should be larger",
			setElements:  []string{"cat", "dog", "mouse", "parrot"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          true,
		},
		{
			name:         "equal - should not be larger",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          false,
		},
		{
			name:         "smaller - should not be larger",
			setElements:  []string{"cat", "dog"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          false,
		},
		{
			name:         "bigger and wrong - should not be larger",
			setElements:  []string{"cat", "dog", "parrot", "monkey"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.setElements...)
			set2 := NewSet(tt.setElements2...)

			if got := set.LargerThan(set2); got != tt.exp {
				t.Errorf("Set.LargerThan() = %v, want %v", got, tt.exp)
			}
		})
	}
}

func TestSet_SmallerThan(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		setElements  []string
		setElements2 []string
		exp          bool
	}{
		{
			name:         "bigger - should not be smaller",
			setElements:  []string{"cat", "dog", "mouse", "parrot"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          false,
		},
		{
			name:         "equal - should not be smaller",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          false,
		},
		{
			name:         "smaller - should be smaller",
			setElements:  []string{"cat", "dog"},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          true,
		},
		{
			name:         "smaller and wrong - should not be smaller",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "parrot", "monkey"},
			exp:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.setElements...)
			set2 := NewSet(tt.setElements2...)

			if got := set.SmallerThan(set2); got != tt.exp {
				t.Errorf("Set.SmallerThan() = %v, want %v", got, tt.exp)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	tests := []struct {
		name         string
		setElements  []string
		setElements2 []string
		exp          []string
	}{
		{
			name:         "should add to set",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"parrot"},
			exp:          []string{"cat", "dog", "mouse", "parrot"},
		},
		{
			name:         "should add nothing",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{},
			exp:          []string{"cat", "dog", "mouse"},
		},
		{
			name:         "should add to empty",
			setElements:  []string{},
			setElements2: []string{"cat", "dog", "mouse"},
			exp:          []string{"cat", "mouse", "dog"},
		},
		{
			name:         "should add with repeated elements",
			setElements:  []string{"cat", "dog", "mouse"},
			setElements2: []string{"cat", "dog", "mouse", "parrot"},
			exp:          []string{"cat", "parrot", "mouse", "dog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.setElements...)
			expSet := NewSet(tt.exp...)

			set.Add(tt.setElements2...)
			require.Equal(t, expSet, set)
		})
	}
}
