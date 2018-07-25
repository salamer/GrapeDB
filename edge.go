package GrapeDB

type Edge struct {
	hash_code uint32
	name      string
	from      *Vertex
	to        *Vertex
}

func NewEdge(from *Vertex, to *Vertex, name string) (*Edge, error) {
	if len(name) > 0 {
		e := &Edge{
			hash_code: Hash(name),
			name:      name,
			from:      from,
			to:        to,
		}
		_, from_exit_key := from.out_edges[e.hash_code]
		_, to_exit_key := to.in_edges[e.hash_code]
		if !from_exit_key && !to_exit_key {
			from.out_edges[e.hash_code] = e
			to.in_edges[e.hash_code] = e
			return e, nil
		} else {
			return nil, EdgeHasExistedError
		}
	}
	return nil, NameError
}

func (e *Edge) IsInVertex(v *Vertex) bool {
	if e.from == v {
		return true
	}
	return false
}

func (e *Edge) IsOutVertex(v *Vertex) bool {
	if e.to == v {
		return true
	}
	return false
}
