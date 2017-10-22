package set

import (
	"reflect"
	"testing"
)

func Test_Union(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	x := New("5", "6", "7")

	u := Union(s, r, x)
	if settype := reflect.TypeOf(u).String(); settype != "*set.Set" {
		t.Error("Union should derive its set type from the first passed set, got", settype)
	}
	if u.Size() != 7 {
		t.Error("Union: the merged set doesn't have all items in it.")
	}

	if !u.Has("1", "2", "3", "4", "5", "6", "7") {
		t.Error("Union: merged items are not availabile in the set.")
	}

	z := Union(x, r)
	if z.Size() != 5 {
		t.Error("Union: Union of 2 sets doesn't have the proper number of items.")
	}
	if settype := reflect.TypeOf(z).String(); settype != "*set.Set" {
		t.Error("Union should derive its set type from the first passed set, got", settype)
	}

}

func Test_Difference(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	x := New("5", "6", "7")
	u := Difference(s, r, x)

	if u.Size() != 2 {
		t.Error("Difference: the set doesn't have all items in it.")
	}

	if !u.Has("1", "2") {
		t.Error("Difference: items are not availabile in the set.")
	}

	y := Difference(r, r)
	if y.Size() != 0 {
		t.Error("Difference: size should be zero")
	}

}

func Test_Intersection(t *testing.T) {
	s1 := New("1", "3", "4", "5")
	s2 := New("2", "3", "5", "6")
	s3 := New("4", "5", "6", "7")
	u := Intersection(s1, s2, s3)

	if u.Size() != 1 {
		t.Error("Intersection: the set doesn't have all items in it.")
	}

	if !u.Has("5") {
		t.Error("Intersection: items after intersection are not availabile in the set.")
	}
}

func benchmarkUnion(b *testing.B, numberOfItems int) {
	s := New()
	u := New()

	for i := 0; i < numberOfItems/2; i++ {
		s.Add(i)
	}

	for i := 0; i < numberOfItems; i++ {
		u.Add(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Union(s, u)
	}
}

func benchmarkDifference(b *testing.B, numberOfItems int) {
	s := New()
	u := New()

	for i := 0; i < numberOfItems; i++ {
		s.Add(i)
	}

	for i := 0; i < numberOfItems/2; i++ {
		u.Add(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Difference(s, u)
	}
}

func benchmarkIntersection(b *testing.B, numberOfItems int) {
	s1 := New()
	s2 := New()

	for i := 0; i < b.N/2; i++ {
		s1.Add(i)
	}

	for i := 0; i < b.N; i++ {
		s2.Add(i)
	}

	b.ResetTimer()

	for i:= 0; i < b.N; i++{
		Intersection(s1, s2)
	}
}

func BenchmarkUnion2(b *testing.B) {
	benchmarkUnion(b, 2)
}

func BenchmarkUnion10(b *testing.B) {
	benchmarkUnion(b, 10)
}

func BenchmarkUnion100(b *testing.B) {
	benchmarkUnion(b, 100)
}

func BenchmarkUnion1000(b *testing.B) {
	benchmarkUnion(b, 1000)
}

func BenchmarkUnion100000(b *testing.B) {
	benchmarkUnion(b, 100000)
}

func BenchmarkDifference2(b *testing.B) {
	benchmarkDifference(b, 2)
}

func BenchmarkDifference10(b *testing.B) {
	benchmarkDifference(b, 10)
}

func BenchmarkDifference100(b *testing.B) {
	benchmarkDifference(b, 100)
}

func BenchmarkDifference1000(b *testing.B) {
	benchmarkDifference(b, 1000)
}

func BenchmarkDifference100000(b *testing.B) {
	benchmarkDifference(b, 100000)
}

func BenchmarkIntersection2(b *testing.B) {
	benchmarkIntersection(b, 2)
}

func BenchmarkIntersection10(b *testing.B) {
	benchmarkIntersection(b, 10)
}

func BenchmarkIntersection100(b *testing.B) {
	benchmarkIntersection(b, 100)
}

func BenchmarkIntersection1000(b *testing.B) {
	benchmarkIntersection(b, 1000)
}

func BenchmarkIntersection100000(b *testing.B) {
	benchmarkIntersection(b, 100000)
}