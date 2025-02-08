# Data Structures and Algorithms

## Data Structures

- Arrays allow fast reads
- Linked lists allow fast inserts and deletes

## Algorithms

### Binary Search
  - Only works on sorted arrays
  - O(log2 n) time complexity

## Problems

### The traveling salesperson problem
  - Given a list of cities and the distances between each pair of cities, what is the shortest route that visits each city exactly once and returns to the origin city?
  - O(n!) time complexity

  | Cities (n) | Operations (n!) |
  |------------|----------------|
  | 6          | 720           |
  | 7          | 5,040         |
  | 8          | 40,320        |
  | 15         | 1,307,674,368,000 |
  | 30         | 2.652528598121911e+32 |

  The number of operations grows factorially, making it computationally infeasible for large numbers of cities. With 100+ cities, the computation time would exceed the lifespan of our sun!

