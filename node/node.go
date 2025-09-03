package node


type Node struct {
	Name string
	Ip string
	Memory int
	Cores int
	MemoryAllocated int
	Disk int
	DiskAllocated int
	Role string
	taskCount int
}