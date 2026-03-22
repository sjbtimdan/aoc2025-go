package utils

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false
	}
	if uf.size[rootX] < uf.size[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}
	return true
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}

type PartitionInfo struct {
	Representative int
	Size           int
}

func (uf *UnionFind) PartitionInfos() []PartitionInfo {
	seen := make(map[int]bool)
	partitions := []PartitionInfo{}
	for i := range uf.parent {
		root := uf.Find(i)
		if !seen[root] {
			partitions = append(partitions, PartitionInfo{Representative: i, Size: uf.size[root]})
			seen[root] = true
		}
	}
	return partitions
}
