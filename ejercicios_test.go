package guia4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/untref-ayp2/data-structures/set"
)

func TestUnion(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(1, 3)
	result := Union(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 3, result.Size(), "La union de set1 y set2 deberia tener 3 elementos")
	esperado := []int{3, 1, 2}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La union de set1 y set2 deberia dar [1,2,3]")
	}
}

func TestIntersection(t *testing.T) {
	set1 := set.NewSetList(1, 2, 5)
	set2 := set.NewSetList(1, 3, 5)
	result := Intersection(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.Size(), "La intersección de set1 y set2 deberia tener 2 elementos")
	esperado := []int{1, 5}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La intersección de set1 y set2 deberia dar [1,5]")
	}
}

func TestDifference(t *testing.T) {
	set1 := set.NewSetList(1, 2, 4, 7)
	set2 := set.NewSetList(1, 3, 7)
	result := Difference(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.Size(), "La differencia de set1 y set2 deberia tener 2 elementos")
	esperado := []int{2, 4}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La differencia de set1 y set2 deberia dar [2,4]")
	}
}

func TestSubset(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(1)
	assert.True(t, Subset(set1, set2), "set2 deberia ser un subconjunto de set1")
}

func TestEqual(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(2, 1)
	assert.True(t, Equal(set1, set2), "Los conjuntos deberian ser iguales")
}

func TestNotEqual(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(2, 3)
	assert.False(t, Equal(set1, set2), "Los conjuntos deberian ser distintos")
}
