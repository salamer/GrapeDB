package GrapeDB

type Vertex struct {
	hash_code uint32
	name      string
	in_edges  map[uint32]*Edge
	out_edges map[uint32]*Edge
}

func NewVertex(name string) (*Vertex, error) {
	if len(name) > 0 {
		return &Vertex{
			name:      name,
			hash_code: Hash(name),
			in_edges:  make(map[uint32]*Edge),
			out_edges: make(map[uint32]*Edge),
		}, nil
	}
	return nil, NameError
}

func (v *Vertex) NewInEdges(e *Edge) error {
	if _, ok := v.in_edges[e.hash_code]; !ok {
		v.in_edges[e.hash_code] = e
		return nil
	} else {
		return EdgeHasExistedError
	}
}

func (v *Vertex) NewOutEdges(e *Edge) error {
	if _, ok := v.out_edges[e.hash_code]; !ok {
		v.out_edges[e.hash_code] = e
		return nil
	} else {
		return EdgeHasExistedError
	}
}

func (v *Vertex) NewTargetVertex(target *Vertex, edge_name string) error {
	edge, err := NewEdge(v, target, edge_name)
	if err != nil {
		return err
	}
	target.NewInEdges(edge)
	v.NewOutEdges(edge)
	return nil
}

func (v *Vertex) NewRelation(target_name string, edge_name string) error {
	target, err := NewVertex(target_name)
	if err != nil {
		return err
	}
	edge, err := NewEdge(v, target, edge_name)
	if err != nil {
		return err
	}
	target.NewInEdges(edge)
	v.NewOutEdges(edge)
	return nil
}

// relation means the edge meaning, but it is more sematic
func (v *Vertex) HasRelation(relation_name string) bool {
	_, from_exit_key := v.out_edges[Hash(relation_name)]
	_, to_exit_key := v.in_edges[Hash(relation_name)]
	if from_exit_key || to_exit_key {
		return true
	}
	return false
}
