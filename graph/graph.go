package graph

import (
	"fmt"
	//"strconv"
)

type Graph struct {
	g      []gvector
	length int
}

type gvector []node

type node struct {
	next  int
	value interface{}
}

//	新建一个Graph实例
func New(num int) *Graph {

	var tmp []gvector
	for i := int(0); i < num; i++ {
		var tmp2 gvector
		tmp = append(tmp, tmp2)
	}
	g := &Graph{g: tmp}

	return g
}

//	返回 len
func (g *Graph) Size() int {
	g.length = len(g.g)
	return g.length
}

//	添加一条边
func (g *Graph) Egde(a int, b int, v interface{}) {
	l := g.length
	if a > l || b > l {
		fmt.Println("[Failed] Egde_Id is too large.")
		return
	}

	g.g[a] = append(g.g[a], node{b, v})

	return
}

// 无视吧！这里只是为了应付某些人在外面使用 fmt.Print
func (g *Graph) String() (res string) {
	for k, v := range g.g {
		if len(v) > 0 {
			res = res + fmt.Sprintf("Node %d:\n", k)
			for _, j := range v {
				res = res + fmt.Sprintf("\t[%v -> %v]\t%v\n", k, j.next, j.value)
			}
		}
	}
	return
}

// 求(a,b)两点之间的最短路径
// 根据提供的图自动选择最优的算法
func (g *Graph) ShortesPath(int a, int b) int64 {
	return g.SPFA(a, b)
}

// 利用SPFA算法求(a,b)两点之间的最短路径
func (g *Graph) SPFA(int a, int b) int64 {
	var q []int
	var l, r int

}

// Floyed算法，貌似暂时没用
func (g *Graph) Floyed() (res string) {
	return
}

//	下面基本数据结构的模板
type queue_int struct {
	v []int
}

func (q *queue_int) Push() {

}

func (q *queue_int) Top() {

}
