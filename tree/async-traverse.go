package main

import (
	"fmt"
)

type Tree struct {
	Level    int
	Value    string
	Children []*Tree
}

var tree = Tree{
	Level: 0,
	Value: "root",
	Children: []*Tree{
		{
			Level: 1,
			Value: "A",
			Children: []*Tree{
				{
					Level: 2,
					Value: "A",
					Children: []*Tree{
						{
							Level:    3,
							Value:    "A",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "B",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "C",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "D",
							Children: nil,
						},
					},
				},
				{
					Level: 2,
					Value: "B",
					Children: []*Tree{
						{
							Level:    3,
							Value:    "E",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "F",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "G",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "H",
							Children: nil,
						},
					},
				},
				{
					Level: 2,
					Value: "C",
					Children: []*Tree{
						{
							Level:    3,
							Value:    "I",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "J",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "K",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "L",
							Children: nil,
						},
					},
				},
			},
		},
		{
			Level: 1,
			Value: "B",
			Children: []*Tree{
				{
					Level: 2,
					Value: "D",
					Children: []*Tree{
						{
							Level:    3,
							Value:    "M",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "N",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "O",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "P",
							Children: nil,
						},
					},
				},
				{
					Level: 2,
					Value: "E",
					Children: []*Tree{
						{
							Level:    3,
							Value:    "Q",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "R",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "S",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "T",
							Children: nil,
						},
					},
				},
				{
					Level: 2,
					Value: "F",
					Children: []*Tree{
						{
							Level:    3,
							Value:    "U",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "V",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "W",
							Children: nil,
						},
						{
							Level:    3,
							Value:    "X",
							Children: nil,
						},
					},
				},
			},
		},
	},
}

type Visited struct {
	VisitedNodes map[*Tree]struct{}
	VisitedOrder []*Tree
}

func NewVisited() Visited {
	return Visited{
		VisitedNodes: map[*Tree]struct{}{},
		VisitedOrder: []*Tree{},
	}
}

func (v *Visited) Copy() Visited {
	return Visited{
		VisitedNodes: func() map[*Tree]struct{} {
			c := map[*Tree]struct{}{}
			for n := range v.VisitedNodes {
				c[n] = struct{}{}
			}
			return c
		}(),
		VisitedOrder: func() []*Tree {
			c := []*Tree{}
			for _, n := range v.VisitedOrder {
				c = append(c, n)
			}
			return c
		}(),
	}
}

func (t *Tree) TraverseAsync(
	done chan struct{}, v Visited, completedPaths chan Visited,
) {
	if _, visited := v.VisitedNodes[t]; visited {
		done <- struct{}{}
		return
	}

	v.VisitedNodes[t] = struct{}{}
	v.VisitedOrder = append(v.VisitedOrder, t)

	if t.Children == nil {
		completedPaths <- v
		done <- struct{}{}
		return
	}

	var childrenDone = make(chan struct{}, len(t.Children))
	for _, child := range t.Children {
		go child.TraverseAsync(childrenDone, v.Copy(), completedPaths)
	}

	count := 0
	for {
		select {
		case <-childrenDone:
			count++
		default:
			if count == len(t.Children) {
				done <- struct{}{}
				return
			}
		}
	}
}

func printComplete(
	completedPaths chan Visited,
	finish chan struct{},
	printDone chan struct{},
) {
	paths := 0
	for {
		select {
		case p := <-completedPaths:
			paths++
			fmt.Printf("completed path %d:\n", paths)
			for _, n := range p.VisitedOrder {
				fmt.Println(n.Level, n.Value)
			}
			fmt.Printf("\n")

		case <-finish:
			printDone <- struct{}{}
		}
	}
}

func main() {
	var completedPaths = make(chan Visited)
	var traverseDone = make(chan struct{})
	var finish = make(chan struct{})
	var printDone = make(chan struct{})

	go printComplete(completedPaths, finish, printDone)
	go tree.TraverseAsync(traverseDone, NewVisited(), completedPaths)

	select {
	case <-traverseDone:
		finish <- struct{}{}

	}
	<-printDone

	fmt.Println("done")
}
