package utils

import (
	"sort"
	"testing"

	"github.com/shoenig/test/must"
)

func TestUnionFind_InitialState(t *testing.T) {
	uf := NewUnionFind(5)

	for i := 0; i < 5; i++ {
		must.Eq(t, i, uf.Find(i))
		must.Eq(t, 1, uf.Size(i))
	}
}

func TestUnionFind_UnionAndFind(t *testing.T) {
	uf := NewUnionFind(6)

	must.Eq(t, true, uf.Union(0, 1))
	must.Eq(t, true, uf.Union(2, 3))
	must.Eq(t, true, uf.Union(1, 2))

	root0 := uf.Find(0)
	must.Eq(t, root0, uf.Find(1))
	must.Eq(t, root0, uf.Find(2))
	must.Eq(t, root0, uf.Find(3))

	must.Eq(t, 4, uf.Size(0))
	must.Eq(t, 4, uf.Size(3))
	must.Eq(t, 1, uf.Size(4))
	must.Eq(t, 1, uf.Size(5))
}

func TestUnionFind_UnionSameSetIsNoOp(t *testing.T) {
	uf := NewUnionFind(4)

	must.Eq(t, true, uf.Union(0, 1))
	sizeBefore := uf.Size(0)

	must.Eq(t, false, uf.Union(0, 1))
	must.Eq(t, false, uf.Union(1, 0))

	must.Eq(t, sizeBefore, uf.Size(0))
	must.Eq(t, uf.Find(0), uf.Find(1))
}

func TestUnionFind_PathCompression(t *testing.T) {
	uf := NewUnionFind(5)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)

	root := uf.Find(0)
	must.Eq(t, root, uf.Find(3))

	// Ensure path compression updated parent links after a find.
	must.Eq(t, root, uf.parent[3])
}

func TestUnionFind_PartitionInfos(t *testing.T) {
	uf := NewUnionFind(8)

	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(3, 4)
	uf.Union(6, 7)

	partitions := uf.PartitionInfos()
	got := make([]int, 0, len(partitions))
	for _, partition := range partitions {
		must.Eq(t, partition.Size, uf.Size(partition.Representative))
		got = append(got, partition.Size)
	}
	sort.Ints(got)

	must.Eq(t, []int{1, 2, 2, 3}, got)
}
