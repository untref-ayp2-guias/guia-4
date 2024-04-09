package guia4

import (
	"testing"

	"github.com/untref-ayp2/data-structures/set"
)

func TestUnion(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(1, 3)
	result := Union(set1, set2)
	if result.Size() != 3 {
		t.Errorf("La union de set1 y set2 deberia tener 3 elementos")
	}
	esperado := []int{3, 1, 2}
	for _, v := range esperado {
		if !result.Contains(v) {
			t.Errorf("La union de set1 y set2 deberia dar [1,2,3]")
			break
		}
	}
}
func TestIntersection(t *testing.T) {
	set1 := set.NewSetList(1, 2, 5)
	set2 := set.NewSetList(1, 3, 5)
	result := Intersection(set1, set2)
	if result.Size() != 2 {
		t.Errorf("La intersección de set1 y set2 deberia tener 3 elementos")
	}
	esperado := []int{1, 5}
	for _, v := range esperado {
		if !result.Contains(v) {
			t.Errorf("La intersección de set1 y set2 deberia dar [1,5]")
			break
		}
	}
}

func TestDifference(t *testing.T) {
	set1 := set.NewSetList(1, 2, 4, 7)
	set2 := set.NewSetList(1, 3, 7)
	result := Difference(set1, set2)
	if result.Size() != 2 {
		t.Errorf("La differencia de set1 y set2 deberia tener 2 elementos")
	}
	esperado := []int{2, 4}
	for _, v := range esperado {
		if !result.Contains(v) {
			t.Errorf("La differencia de set1 y set2 deberia dar [2,4]")
			break
		}
	}
}

func TestSubset(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(1)
	if !Subset(set1, set2) {
		t.Errorf("set2 deberia ser un subconjunto de set1")
	}
}

func TestEqual(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(2, 1)
	if !Equal(set1, set2) {
		t.Errorf("Los conjuntos deberian ser iguales")
	}
}
