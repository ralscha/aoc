package gridutil

import (
	"aoc/internal/container"
	"math"
)

// Node represents a node in the pathfinding algorithms
type Node[T comparable] struct {
	Coord  Coordinate
	Cost   float64
	Parent *Node[T] // Used to reconstruct path
	Value  T        // The actual value in the grid
}

// PathResult holds the result of a pathfinding operation
type PathResult[T comparable] struct {
	Path []Coordinate // The path from start to goal
	Cost float64     // Total cost of the path
}

// IsGoal is a function type that determines if a position is the goal
type IsGoal[T comparable] func(currentPos Coordinate, current T) bool

// IsObstacle is a function type that determines if a position is an obstacle.
// It returns true if the position is an obstacle and false if it's free to move.
type IsObstacle[T comparable] func(currentPos Coordinate, current T) bool

// GetCost is a function type that determines the cost of moving from one position to another
type GetCost[T comparable] func(from, to Coordinate, fromValue, toValue T) float64

// ShortestPathWithBFS performs a Breadth-First Search starting from the given coordinate.
// It finds the shortest path to the goal position by exploring nodes level by level,
// avoiding obstacles defined by the IsObstacle function.
func (g *Grid2D[T]) ShortestPathWithBFS(start Coordinate, isGoal IsGoal[T], isObstacle IsObstacle[T], directions []Direction) (PathResult[T], bool) {
	visited := container.NewSet[Coordinate]()
	queue := container.NewQueue[*Node[T]]()

	// Create start node
	startValue, exists := g.GetC(start)
	if !exists {
		return PathResult[T]{}, false
	}
	startNode := &Node[T]{
		Coord:  start,
		Cost:   0,
		Parent: nil,
		Value:  startValue,
	}

	// Add start node to queue and visited set
	queue.Push(startNode)
	visited.Add(start)

	// BFS
	for !queue.IsEmpty() {
		current := queue.Pop()
		if isGoal(current.Coord, current.Value) {
			var path []Coordinate
			node := current
			for node != nil {
				path = append([]Coordinate{node.Coord}, path...)
				node = node.Parent
			}
			return PathResult[T]{
				Path: path,
				Cost: current.Cost,
			}, true
		}

		for _, dir := range directions {
			neighbor := Coordinate{Row: current.Coord.Row + dir.Row, Col: current.Coord.Col + dir.Col}

			// Skip if out of bounds
			if neighbor.Row < 0 || neighbor.Row >= g.Height() ||
				neighbor.Col < 0 || neighbor.Col >= g.Width() {
				continue
			}

			// Skip if already visited
			if visited.Contains(neighbor) {
				continue
			}

			// Get neighbor value and check if it's an obstacle
			neighborValue, exists := g.GetC(neighbor)
			if !exists || isObstacle(neighbor, neighborValue) {
				continue
			}

			// Create neighbor node
			neighborNode := &Node[T]{
				Coord:  neighbor,
				Cost:   current.Cost + 1,
				Parent: current,
				Value:  neighborValue,
			}

			// Add to queue and mark as visited
			queue.Push(neighborNode)
			visited.Add(neighbor)
		}
	}

	// No path found
	return PathResult[T]{}, false
}

// ShortestPathWithDijkstra performs Dijkstra's algorithm starting from the given coordinate.
// It finds the shortest path to the goal position considering edge weights defined by the getCost function.
func (g *Grid2D[T]) ShortestPathWithDijkstra(
	start Coordinate,
	isGoal IsGoal[T],
	isObstacle IsObstacle[T],
	getCost GetCost[T],
	directions []Direction,
) (PathResult[T], bool) {
	// Initialize distance map with infinity for all coordinates
	distances := make(map[Coordinate]float64)
	for row := 0; row < g.Height(); row++ {
		for col := 0; col < g.Width(); col++ {
			distances[Coordinate{Row: row, Col: col}] = math.Inf(1)
		}
	}
	distances[start] = 0

	// Create priority queue and start node
	pq := container.NewPriorityQueue[*Node[T]]()
	startValue, exists := g.GetC(start)
	if !exists {
		return PathResult[T]{}, false
	}
	startNode := &Node[T]{
		Coord:  start,
		Cost:   0,
		Parent: nil,
		Value:  startValue,
	}
	pq.Push(startNode, 0)

	// Track visited nodes
	visited := container.NewSet[Coordinate]()

	// Dijkstra's algorithm
	for !pq.IsEmpty() {
		current := pq.Pop()

		// Skip if we've already found a better path to this node
		if visited.Contains(current.Coord) {
			continue
		}
		visited.Add(current.Coord)

		// Check if we've reached the goal
		if isGoal(current.Coord, current.Value) {
			var path []Coordinate
			node := current
			for node != nil {
				path = append([]Coordinate{node.Coord}, path...)
				node = node.Parent
			}
			return PathResult[T]{
				Path: path,
				Cost: current.Cost,
			}, true
		}

		// Explore neighbors
		for _, dir := range directions {
			neighbor := Coordinate{Row: current.Coord.Row + dir.Row, Col: current.Coord.Col + dir.Col}

			// Skip if out of bounds
			if neighbor.Row < 0 || neighbor.Row >= g.Height() ||
				neighbor.Col < 0 || neighbor.Col >= g.Width() {
				continue
			}

			// Get neighbor value and check if it's an obstacle
			neighborValue, exists := g.GetC(neighbor)
			if !exists || isObstacle(neighbor, neighborValue) {
				continue
			}

			// Calculate cost to reach neighbor through current node
			edgeCost := getCost(current.Coord, neighbor, current.Value, neighborValue)
			totalCost := current.Cost + edgeCost

			// If we found a better path to the neighbor
			if totalCost < distances[neighbor] {
				distances[neighbor] = totalCost
				neighborNode := &Node[T]{
					Coord:  neighbor,
					Cost:   totalCost,
					Parent: current,
					Value:  neighborValue,
				}
				// Priority is based on total cost (convert to int as required by PriorityQueue)
				pq.Push(neighborNode, int(totalCost*1000)) // Multiply by 1000 to preserve some decimal precision
			}
		}
	}

	// No path found
	return PathResult[T]{}, false
}
