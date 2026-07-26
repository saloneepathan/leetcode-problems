# MovieRentingSystem Design

## Core Data Structures

| Name       | Type                   | Purpose                                                        |
|------------|------------------------|----------------------------------------------------------------|
| `price`    | map[[2]int]int         | Maps (shop, movie) → price                                     |
| `state`    | map[[2]int]int         | Version counter for each (shop, movie) to track updates        |
| `isRented` | map[[2]int]bool        | Tracks if (shop, movie) is currently rented                    |
| `avail`    | map[int]*AvailHeap     | For each movie, a min-heap of available shops by price, shop   |
| `rentedH`  | *RentedHeap            | Min-heap of currently rented movies by price, shop, movie      |

---

## Data Structure Interaction

- **Constructor:**  
  Initializes all maps and heaps. Each entry is pushed into the corresponding movie’s `AvailHeap`.

- **Search(movie):**  
  Pops up to 5 shops from `AvailHeap` for the movie. Checks `state` and `isRented` to ensure the shop has the latest version and is not rented. Valid shops are returned; all popped nodes are pushed back.

- **Rent(shop, movie):**  
  Marks the movie as rented in `isRented`. Increments the `state` version for (shop, movie). Pushes a new node into `RentedHeap`.

- **Drop(shop, movie):**  
  Marks the movie as not rented in `isRented`. Increments the `state` version for (shop, movie). Pushes a new node into the movie’s `AvailHeap`.

- **Report():**  
  Pops up to 5 rented nodes from `RentedHeap`. Checks `state` and `isRented` for validity. Valid rented movies are returned; all popped nodes are pushed back.

---

## Summary Table

| Operation | Reads/Writes | Heaps Used      | Maps Used               |
|-----------|--------------|-----------------|-------------------------|
| Search    | Reads        | AvailHeap       | price, state, isRented  |
| Rent      | Writes       | RentedHeap      | state, isRented, price  |
| Drop      | Writes       | AvailHeap       | state, isRented, price  |
| Report    | Reads        | RentedHeap      | state, isRented         |

---

**In essence:**  
- Heaps efficiently provide the top results for search/report.  
- Maps track the state and rental status for correctness and versioning.  
- Versioning (`state`) ensures outdated heap nodes are ignored.  
- All operations maintain heap integrity by pushing back nodes after temporary pops.