package gridutil

import "aoc/internal/container"

// DFSCallbacks contains callback functions for customizing DFS behavior
type DFSCallbacks[T comparable, S any] struct {
	// OnVisit is called when visiting a node. It receives the current coordinate,
	// its value, and the current state. It returns the updated state.
	OnVisit func(coord Coordinate, value T, state S) S

	// GetNextNodes determines which nodes to visit next from the current position.
	// It receives the current coordinate, its value, the grid, and returns a slice
	// of next coordinates to visit.
	GetNextNodes func(current Coordinate, value T, grid *Grid2D[T]) []Coordinate
}

// DFS performs a depth-first search traversal of the grid starting from the given coordinate.
// It uses the provided callbacks to customize the traversal behavior and maintains a state of type S.
func DFS[T comparable, S any](
	grid *Grid2D[T],
	start Coordinate,
	initialState S,
	callbacks DFSCallbacks[T, S],
) S {
	visited := container.NewSet[Coordinate]()
	return dfsHelper(grid, start, initialState, visited, callbacks)
}

func dfsHelper[T comparable, S any](
	grid *Grid2D[T],
	current Coordinate,
	state S,
	visited *container.Set[Coordinate],
	callbacks DFSCallbacks[T, S],
) S {
	visited.Add(current)

	value, exists := grid.GetC(current)
	if !exists {
		return state
	}

	state = callbacks.OnVisit(current, value, state)
	nextNodes := callbacks.GetNextNodes(current, value, grid)
	for _, next := range nextNodes {
		if !visited.Contains(next) {
			state = dfsHelper[T, S](grid, next, state, visited, callbacks)
		}
	}
	visited.Remove(current)

	return state
}

// Common DFS patterns

// CountPaths counts all possible paths from start to nodes matching the target condition
func CountPaths[T comparable](
	grid *Grid2D[T],
	start Coordinate,
	isTarget func(value T) bool,
	isValidNext func(current, next T) bool,
) int {
	callbacks := DFSCallbacks[T, int]{
		OnVisit: func(coord Coordinate, value T, count int) int {
			if isTarget(value) {
				return count + 1
			}
			return count
		},
		GetNextNodes: func(current Coordinate, value T, grid *Grid2D[T]) []Coordinate {
			var nextNodes []Coordinate
			for _, dir := range Get4Directions() {
				next := Coordinate{Row: current.Row + dir.Row, Col: current.Col + dir.Col}
				if nextVal, exists := grid.GetC(next); exists {
					if isValidNext(value, nextVal) {
						nextNodes = append(nextNodes, next)
					}
				}
			}
			return nextNodes
		},
	}

	return DFS(grid, start, 0, callbacks)
}

// CollectNodes collects all unique nodes that match a given condition during DFS traversal
func CollectNodes[T comparable](
	grid *Grid2D[T],
	start Coordinate,
	shouldCollect func(value T) bool,
	isValidNext func(current, next T) bool,
) []Coordinate {
	type collectionState struct {
		collected *container.Set[Coordinate]
	}

	callbacks := DFSCallbacks[T, *collectionState]{
		OnVisit: func(coord Coordinate, value T, state *collectionState) *collectionState {
			if shouldCollect(value) {
				state.collected.Add(coord)
			}
			return state
		},
		GetNextNodes: func(current Coordinate, value T, grid *Grid2D[T]) []Coordinate {
			var nextNodes []Coordinate
			for _, dir := range Get4Directions() {
				next := Coordinate{Row: current.Row + dir.Row, Col: current.Col + dir.Col}
				if nextVal, exists := grid.GetC(next); exists {
					if isValidNext(value, nextVal) {
						nextNodes = append(nextNodes, next)
					}
				}
			}
			return nextNodes
		},
	}

	state := DFS(grid, start, &collectionState{collected: container.NewSet[Coordinate]()}, callbacks)
	return state.collected.Values()
}
