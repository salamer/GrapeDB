package GrapeDB

import "errors"

var (
	NameError             = errors.New("The name of edge must sematic and its length lagger than one letter")
	EdgeHasExistedError   = errors.New("The edge has already exited")
	VertexHasExistedError = errors.New("The vertex has already exited")
)
