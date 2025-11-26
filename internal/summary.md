# Internal Package Summary

This document provides a detailed description of each method within the `internal` package.

## internal/container

### bag.go

-   `type Bag[T comparable] struct`: A generic bag data structure that stores elements of type T.
-   `func NewBag[T comparable]() *Bag[T]`: Creates a new empty Bag.
-   `func (b *Bag[T]) Add(data T)`: Adds an element to the bag.
-   `func (b *Bag[T]) Remove(data T)`: Removes an element from the bag.
-   `func (b *Bag[T]) Contains(data T) bool`: Checks if the bag contains a specific element.
-   `func (b *Bag[T]) Count(data T) int`: Returns the number of times an element appears in the bag.
-   `func (b *Bag[T]) Values() map[T]int`: Returns a map of all elements in the bag and their counts.

### priority_queue.go

-   `type PriorityQueue[T any] struct`: A generic priority queue data structure that stores elements of type T.
-   `func NewPriorityQueue[T any]() PriorityQueue[T]`: Creates a new empty PriorityQueue.
-   `func (pq PriorityQueue[T]) IsEmpty() bool`: Checks if the priority queue is empty.
-   `func (pq PriorityQueue[T]) Len() int`: Returns the number of elements in the priority queue.
-   `func (pq PriorityQueue[T]) Push(value T, priority int)`: Adds an element to the priority queue with a given priority.
-   `func (pq PriorityQueue[T]) Pop() T`: Removes and returns the element with the highest priority.
-   `func (pq PriorityQueue[T]) Peek() T`: Returns the element with the highest priority without removing it.
-   `type item[T any] struct`: Represents an item in the priority queue with a value and priority.
-   `type priorityQueue[T any] []*item[T]`: A slice of items used to implement the priority queue.
-   `func (pq priorityQueue[T]) Len() int`: Returns the length of the priority queue.
-   `func (pq priorityQueue[T]) Less(i, j int) bool`: Compares the priority of two items in the priority queue.
-   `func (pq priorityQueue[T]) Swap(i, j int)`: Swaps two items in the priority queue.
-   `func (pq *priorityQueue[T]) Push(x any)`: Pushes an item onto the priority queue.
-   `func (pq *priorityQueue[T]) Pop() any`: Pops an item from the priority queue.

### queue.go

-   `type Queue[T any] struct`: A generic queue data structure that stores elements of type T.
-   `func NewQueue[T any]() *Queue[T]`: Creates a new empty Queue.
-   `func (q *Queue[T]) IsEmpty() bool`: Checks if the queue is empty.
-   `func (q *Queue[T]) Len() int`: Returns the number of elements in the queue.
-   `func (q *Queue[T]) Push(data T)`: Adds an element to the back of the queue.
-   `func (q *Queue[T]) Pop() T`: Removes and returns the element from the front of the queue.
-   `func (q *Queue[T]) resize()`: Resizes the underlying array of the queue.
-   `func (q *Queue[T]) shrink()`: Shrinks the underlying array of the queue.

### set.go

-   `type Set[T comparable] struct`: A generic set data structure that stores unique elements of type T.
-   `func NewSet[T comparable]() *Set[T]`: Creates a new empty Set.
-   `func (s *Set[T]) Add(data T)`: Adds an element to the set.
-   `func (s *Set[T]) AddAll(datas []T)`: Adds all elements from a slice to the set.
-   `func (s *Set[T]) Remove(data T)`: Removes an element from the set.
-   `func (s *Set[T]) Contains(data T) bool`: Checks if the set contains a specific element.
-   `func (s *Set[T]) Len() int`: Returns the number of elements in the set.
-   `func (s *Set[T]) Values() []T`: Returns a slice of all elements in the set.
-   `func (s *Set[T]) Copy() *Set[T]`: Returns a new set with the same elements as the original set.
-   `func (s *Set[T]) Intersection(otherSet *Set[T]) *Set[T]`: Returns a new set with elements that are present in both sets.

## internal/conv

### conv.go

-   `func MustSscanf(s string, format string, a ...interface{})`: Parses a string according to a format specifier, logs a fatal error and terminates the program on error.
-   `func MustAtoi(s string) int`: Converts a string to an integer, logs a fatal error and terminates the program on error.
-   `func MustAtoi64(s string) int64`: Converts a string to an int64, logs a fatal error and terminates the program on error.
-   `func SplitNewline(s string) []string`: Splits a string by newlines.
-   `func ToIntSlice(s []string) []int`: Converts a slice of strings to a slice of integers.
-   `func ToIntSliceComma(s string) []int`: Converts a comma-separated string to a slice of integers.
-   `func ToInt64SliceComma(s string) []int64`: Converts a comma-separated string to a slice of int64s.

## internal/download

### download_input.go

-   `func ReadInput(year, day int) (string, error)`: Reads input for a given year and day.
-   `func DownloadInput(year, day int) (string, error)`: Downloads input for a given year and day.

## internal/geomutil

### geometry.go

-   `func ManhattanDistance(p1, p2 gridutil.Coordinate) int`: Calculates the Manhattan distance between two coordinates.

## internal/graphutil

### graph.go

-   `type Node struct`: Represents a node in a graph.
-   `type Edge struct`: Represents an edge in a graph.
-   `type Graph struct`: Represents a graph data structure.
-   `func NewGraph() *Graph`: Creates a new empty Graph.
-   `func (g *Graph) AddNode(id string)`: Adds a node to the graph.
-   `func (g *Graph) AddEdge(from, to string, weight int)`: Adds an edge to the graph.
-   `func (g *Graph) GetNeighbors(nodeID string) []*Node`: Returns the neighbors of a given node.
-   `func (g *Graph) FindCycle(startID string) []string`: Finds a cycle in the graph starting from a given node.
-   `func (g *Graph) FindAllPaths(start, end string, condition func(node *Node) bool) [][]string`: Finds all paths between two nodes that satisfy a given condition.

## internal/gridutil

### coordinate3d.go

-   `type Coordinate3D struct`: Represents a 3D coordinate.
-   `func (c Coordinate3D) Add(other Coordinate3D) Coordinate3D`: Adds two 3D coordinates.
-   `func (c Coordinate3D) Sub(other Coordinate3D) Coordinate3D`: Subtracts two 3D coordinates.
-   `func (c Coordinate3D) ManhattanDistance(other Coordinate3D) int`: Calculates the Manhattan distance between two 3D coordinates.

### dfs.go

-   `type DFSCallbacks[T comparable, S any] struct`: Represents callbacks for DFS.
-   `func DFS[T comparable, S any](...)`: Performs a Depth-First Search.
-   `func dfsHelper[T comparable, S any](...)`: Helper function for DFS.
-   `func CountPaths[T comparable](...)`: Counts the number of paths.
-   `func CollectNodes[T comparable](...)`: Collects nodes during DFS.
-   `type collectionState struct`: Represents the state of the collection during DFS.

### grid2d.go

-   `type Coordinate struct`: Represents a 2D coordinate.
-   `type Direction struct`: Represents a direction.
-   `type Grid2D[T comparable] struct`: Represents a 2D grid data structure.
-   `func Get8Directions() []Direction`: Returns all 8 directions.
-   `func Get4Directions() []Direction`: Returns the 4 cardinal directions.
-   `func TurnLeft(direction Direction) Direction`: Turns a direction 90 degrees left.
-   `func TurnRight(direction Direction) Direction`: Turns a direction 90 degrees right.
-   `func NewGrid2D[T comparable](wrap bool) Grid2D[T]`: Creates a new empty Grid2D.
-   `func NewNumberGrid2D(lines []string) Grid2D[int]`: Creates a new Grid2D from a slice of strings, parsing integers.
-   `func NewCharGrid2D(lines []string) Grid2D[rune]`: Creates a new Grid2D from a slice of strings, parsing runes.
-   `func (g *Grid2D[T]) SetMinRow(row int)`: Sets the minimum row of the grid.
-   `func (g *Grid2D[T]) SetMinRowCol(row, col int)`: Sets the minimum row and column of the grid.
-   `func (g *Grid2D[T]) SetMaxRowCol(row, col int)`: Sets the maximum row and column of the grid.
-   `func (g *Grid2D[T]) SetMaxRowColC(coord Coordinate)`: Sets the maximum row and column of the grid using a coordinate.
-   `func (g Grid2D[T]) Width() int`: Returns the width of the grid.
-   `func (g Grid2D[T]) Height() int`: Returns the height of the grid.
-   `func (g *Grid2D[T]) Count() int`: Returns the number of elements in the grid.
-   `func (g Grid2D[T]) GetMinMaxCol() (int, int)`: Returns the minimum and maximum column of the grid.
-   `func (g Grid2D[T]) GetMinMaxRow() (int, int)`: Returns the minimum and maximum row of the grid.
-   `func (g Grid2D[T]) Get(row, col int) (T, bool)`: Returns the value at a given row and column.
-   `func (g Grid2D[T]) GetC(coord Coordinate) (T, bool)`: Returns the value at a given coordinate.
-   `func (g *Grid2D[T]) Set(row, col int, value T)`: Sets the value at a given row and column.
-   `func (g *Grid2D[T]) SetC(coord Coordinate, value T)`: Sets the value at a given coordinate.
-   `func (g *Grid2D[T]) updateMinMax(row, col int)`: Updates the minimum and maximum row and column of the grid.
-   `func (g *Grid2D[T]) Peek(row, col int, direction Direction) (T, bool)`: Returns the value at a given row and column in a given direction.
-   `func (g *Grid2D[T]) PeekC(coord Coordinate, direction Direction) (T, bool)`: Returns the value at a given coordinate in a given direction.
-   `func (g Grid2D[T]) Copy() Grid2D[T]`: Returns a copy of the grid.
-   `func (g Grid2D[T]) GetNeighbours8(row, col int) []T`: Returns the 8 neighbors of a given cell.
-   `func (g Grid2D[T]) GetNeighbours8C(coord Coordinate) []T`: Returns the 8 neighbors of a given coordinate.
-   `func (g Grid2D[T]) GetNeighbours4(row, col int) []T`: Returns the 4 neighbors of a given cell.
-   `func (g Grid2D[T]) GetNeighbours4C(coord Coordinate) []T`: Returns the 4 neighbors of a given coordinate.
-   `func (g Grid2D[T]) GetNeighbours(row, col int, directions []Direction) []T`: Returns the neighbors of a given cell in the given directions.
-   `func (g Grid2D[T]) GetNeighboursC(coord Coordinate, directions []Direction) []T`: Returns the neighbors of a given coordinate in the given directions.
-   `func (g *Grid2D[T]) RotateRow(row, amount int)`: Rotates a row by a given amount.
-   `func (g *Grid2D[T]) RotateColumn(col, amount int)`: Rotates a column by a given amount.
-   `func (g *Grid2D[T]) String() string`: Returns a string representation of the grid.

### pathfinding.go

-   `type Node[T comparable] struct`: Represents a node in a pathfinding graph.
-   `type PathResult[T comparable] struct`: Represents the result of a pathfinding algorithm.
-   `type IsGoal[T comparable] func(currentPos Coordinate, current T) bool`: Represents a function that checks if a given position is the goal.
-   `type IsObstacle[T comparable] func(currentPos Coordinate, current T) bool`: Represents a function that checks if a given position is an obstacle.
-   `type GetCost[T comparable] func(from, to Coordinate, fromValue, toValue T) float64`: Represents a function that calculates the cost of moving between two positions.
-   `func (g *Grid2D[T]) ShortestPathWithBFS(start Coordinate, isGoal IsGoal[T], isObstacle IsObstacle[T], directions []Direction) (PathResult[T], bool)`: Finds the shortest path using Breadth-First Search.
-   `func (g *Grid2D[T]) ShortestPathWithDijkstra(...)`: Finds the shortest path using Dijkstra's algorithm.

## internal/mathx

### mathx.go

-   `func Combinations[T any](input []T) [][]T`: Generates all combinations of elements in a slice.
-   `func Permutations[T any](input []T) [][]T`: Generates all permutations of elements in a slice.
-   `func CartesianProductSelf[T any](n int, values []T) [][]T`: Generates the Cartesian product of a slice with itself n times.
-   `func Abs[E constraints.Float | constraints.Integer](input E) E`: Returns the absolute value of a number.
-   `func Round(x float64) int`: Rounds a float64 to the nearest integer.
-   `func Lcm(n []int) int`: Calculates the least common multiple of a slice of integers.
-   `func Gcd(a, b int) int`: Calculates the greatest common divisor of two integers.
-   `func MannhattanDistance(x1, y1, x2, y2 int) int`: Calculates the Manhattan distance between two points.
-   `func Factors(n int) []int`: Returns all factors of a given integer.

## internal/rangeutil

### range.go

-   `type Range struct`: Represents a range of integers.
-   `func NewRange(start, end int) Range`: Creates a new Range.
-   `func (r Range) Length() int`: Returns the length of the range.
-   `func (r Range) Contains(value int) bool`: Checks if the range contains a given value.
-   `func (r Range) Overlaps(other Range) bool`: Checks if the range overlaps with another range.
-   `func (r Range) Split(at int) (before, after *Range)`: Splits the range at a given value.
-   `func (r Range) Intersection(other Range) *Range`: Returns the intersection of two ranges.
-   `func (r Range) Union(other Range) *Range`: Returns the union of two ranges.
-   `func (r Range) Adjacent(other Range) bool`: Checks if the range is adjacent to another range.
-   `func (r Range) Map(transform func(int) int) Range`: Maps the range using a given transform function.
-   `func Merge(ranges []Range) []Range`: Merges a slice of ranges.

## internal/slicesx

### slicesx.go

-   `type Number interface`: A constraint that permits any numeric type that supports addition (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64).
-   `func Sum[T Number](slice []T) T`: Returns the sum of all elements in the slice. Works with any numeric type. Returns zero for an empty slice.
-   `func SumIter[T Number](seq iter.Seq[T]) T`: Returns the sum of all elements from an iterator. Works with any numeric type. Returns zero for an empty iterator.
-   `func Unique[T comparable](slice []T) []T`: Returns a new slice with duplicate values removed. The order of elements is not guaranteed to be preserved.

## internal/stringutil

### stringutil.go

-   `func Reverse(s string) string`: Reverses a string.
-   `func FindAllOccurrences(str, substr string) []int`: Finds all occurrences of a substring in a string.
-   `func CountVowels(s string) int`: Counts the number of vowels in a string.
-   `func HasRepeatedChar(s string) bool`: Checks if a string has repeated characters.
-   `func HasRepeatedPair(s string) bool`: Checks if a string has repeated pairs of characters.
-   `func HasSandwichedChar(s string) bool`: Checks if a string has sandwiched characters.
