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

func (t *Tree) TraverseLevel() {
	var queue = make([]*Tree, 0)

	queue = append(queue, t)

	for len(queue) != 0 {
		node := queue[0]

		fmt.Printf("Level: %d, Value: %s\n", node.Level, node.Value)

		queue = append(queue, node.Children...)

		queue = queue[1:]
	}

}

func (t *Tree) TraverseDepth() {
	fmt.Printf("Level: %d, Value: %s\n", t.Level, t.Value)

	for _, child := range t.Children {
		child.TraverseDepth()
	}
}

func main() {
	fmt.Println("Breadth first:")
	tree.TraverseLevel()

	fmt.Println("Depth first:")
	tree.TraverseDepth()
}
