package locations

import (
	"container/heap"
	"math"
)

type GameMap struct {
	layers map[int]*Vertex
}

type Vertex struct {
	name   string
	parent *Vertex
	edges  map[int]*Edge
}

type Edge struct {
	destination *Vertex
	duration    int
}

func (this *Vertex) GetInfo() string {
	return this.name
}

func (this *GameMap) GetItem(key int) *Vertex {
	return this.layers[key]
}

func (this *GameMap) AddVertex(key int, name string, parent *Vertex) {
	if _, ok := this.layers[key]; !ok {
		this.layers[key] = &Vertex{name: name, parent: parent, edges: map[int]*Edge{}}
	} else {
		this.layers[key].name = name
		this.layers[key].parent = parent
	}
}

func (this *GameMap) AddEdge(srcKey, destKey int, distance int) {
	// check if src & dest exist
	if _, ok := this.layers[srcKey]; !ok {
		return
	}
	if _, ok := this.layers[destKey]; !ok {
		return
	}

	// add edge src --> dest
	this.layers[srcKey].edges[destKey] = &Edge{duration: distance, destination: this.layers[destKey]}
}

func (this *GameMap) Neighbors(srcKey int) []string {
	result := []string{}

	for _, edge := range this.layers[srcKey].edges {
		result = append(result, edge.destination.name)
	}

	return result
}

func (this *GameMap) NextWaypoint(start, target *Vertex) (*Vertex, error) {

	if start == target {
		return nil, nil
	}

	p := this.dijkstraShort(start, target)
	if p != nil {
		return p[0], nil
	}

	return nil, nil
}

func (this *GameMap) dijkstraShort(start, target *Vertex) (p []*Vertex) {

	d := make(map[*Vertex]int, len(this.layers))
	p = []*Vertex{}

	for _, v := range this.layers {
		d[v] = math.MaxInt
	}

	d[start] = 0

	q := priorityQueue{&queueItem{
		vertex:   start,
		distance: 0,
		index:    0,
	}}

	heap.Init(&q)

	for len(q) > 0 {
		v := heap.Pop(&q).(*queueItem)

		if v.vertex == target {
			p = append(p, v.vertex)
			return p
		}

		if v.distance > d[v.vertex] {
			continue
		}

		futher := false

		for _, e := range v.vertex.edges {
			distance := v.distance + e.duration

			if distance < d[e.destination] {
				d[e.destination] = distance
				heap.Push(&q, &queueItem{
					vertex:   e.destination,
					distance: distance,
				})
				futher = true
			}
		}

		if futher && v.vertex != start {
			p = append(p, v.vertex)
		}
	}

	return nil
}

func (this *Vertex) GetDurationToTarget(target *Vertex) int {

	if this == target {
		return -1
	}

	for _, e := range this.edges {
		if e.destination == target {
			return e.duration
		}
	}

	return -1
}

type GameMapOptions func(this *GameMap)

func NewGameMap(path string) (*GameMap, error) {
	this := &GameMap{}
	yamlMap, err := LoadYamlMap(path)

	if err != nil {
		return nil, err
	}

	for _, l := range yamlMap.Map.Layers {

		var p *Vertex

		if l.Parent != 0 {
			if _, ok := this.layers[l.Parent]; !ok {
				this.AddVertex(l.Parent, "", nil)
			}
			p = this.layers[l.Parent]
		} else {
			p = nil
		}

		this.AddVertex(l.ID, l.Name, p)

		for _, e := range l.Edges {
			if _, ok := this.layers[e.Dst]; !ok {
				this.AddVertex(e.Dst, "", nil)
			}
			this.AddEdge(l.ID, e.Dst, e.Val)
		}
	}

	return this, nil
}
