// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package linkedlist

type Item[T any] struct {
	prev *Item[T]
	next *Item[T]

	Value T
}

func NewItem[T any](value T) *Item[T] {
	return &Item[T]{Value: value}
}
