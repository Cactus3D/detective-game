package locations

import (
	"reflect"
	"strconv"
	"testing"
)

func getGameMap() (gm *GameMap) {
	gm = &GameMap{
		layers: map[int]*Vertex{},
	}

	for i := 1; i < 11; i++ {
		gm.layers[i] = &Vertex{
			name: strconv.Itoa(i),
		}
	}

	gm.layers[1].edges = map[int]*Edge{
		1: {
			destination: gm.layers[2],
			duration:    25,
		},
		2: {
			destination: gm.layers[4],
			duration:    10,
		},
		3: {
			destination: gm.layers[5],
			duration:    5,
		},
	}

	gm.layers[2].edges = map[int]*Edge{
		1: {
			destination: gm.layers[1],
			duration:    25,
		},
		2: {
			destination: gm.layers[3],
			duration:    5,
		},
	}

	gm.layers[3].edges = map[int]*Edge{
		1: {
			destination: gm.layers[1],
			duration:    29,
		},
		2: {
			destination: gm.layers[8],
			duration:    5,
		},
		3: {
			destination: gm.layers[10],
			duration:    10,
		},
	}

	gm.layers[4].edges = map[int]*Edge{
		1: {
			destination: gm.layers[2],
			duration:    10,
		},
		2: {
			destination: gm.layers[1],
			duration:    10,
		},
		3: {
			destination: gm.layers[6],
			duration:    30,
		},
	}

	gm.layers[6].edges = map[int]*Edge{
		1: {
			destination: gm.layers[7],
			duration:    5,
		},
		2: {
			destination: gm.layers[9],
			duration:    5,
		},
	}

	gm.layers[10].edges = map[int]*Edge{
		1: {
			destination: gm.layers[8],
			duration:    5,
		},
	}
	return gm
}

func TestNextWaypoint(t *testing.T) {

	gm := getGameMap()

	want := gm.layers[4]
	got, err := gm.NextWaypoint(gm.layers[1], gm.layers[4])

	if err != nil {
		t.Errorf("Got error %s", err.Error())
	}

	if want != got {
		t.Errorf("Next waypoint = %s; want %s", got.GetInfo(), want.GetInfo())
	}
}

func TestGetDestinationToTarget(t *testing.T) {

	targetA := &Vertex{}
	targetB := &Vertex{}
	current := &Vertex{}
	edge := &Edge{
		destination: targetA,
		duration:    10,
	}
	current.edges = map[int]*Edge{
		1: edge,
	}

	t.Run("Edge exists", func(t *testing.T) {
		want := 10
		got := current.GetDurationToTarget(targetA)
		if want != got {
			t.Errorf("Distance to existing edge = %d; want %d", got, want)
		}
	})
	t.Run("Edge not exists", func(t *testing.T) {
		want := -1
		got := current.GetDurationToTarget(targetB)
		if want != got {
			t.Errorf("Distance to existing edge = %d; want %d", got, want)
		}
	})
	t.Run("Self", func(t *testing.T) {
		want := -1
		got := current.GetDurationToTarget(current)
		if want != got {
			t.Errorf("Distance to existing edge = %d; want %d", got, want)
		}
	})

}

func TestDijkstraShort(t *testing.T) {

	gm := getGameMap()

	t.Run("1->2", func(t *testing.T) {
		want := []*Vertex{
			gm.layers[4],
			gm.layers[2],
		}
		got := gm.dijkstraShort(gm.layers[1], gm.layers[2])
		if !reflect.DeepEqual(want, got) {
			gPath := ""
			for _, g := range got {
				gPath = gPath + "[" + g.name + "]"
			}
			t.Errorf("Got path %s, want path [4]", gPath)
		}
	})
	t.Run("1->8", func(t *testing.T) {
		want := []*Vertex{
			gm.layers[4],
			gm.layers[2],
			gm.layers[3],
			gm.layers[8],
		}
		got := gm.dijkstraShort(gm.layers[1], gm.layers[8])
		if !reflect.DeepEqual(want, got) {
			gPath := ""
			for _, g := range got {
				gPath = gPath + "[" + g.name + "]"
			}
			t.Errorf("Got path %s, want path [4][2][3]", gPath)
		}
	})
	t.Run("5->1", func(t *testing.T) {
		got := gm.dijkstraShort(gm.layers[5], gm.layers[1])
		if got != nil {
			gPath := ""
			for _, g := range got {
				gPath = gPath + "[" + g.name + "]"
			}
			t.Errorf("Got path %s, excepted no path", gPath)
		}
	})

}
