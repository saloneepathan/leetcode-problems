
class Node:
    def __init__(self, key=0, value=0):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache = {}

        # Dummy nodes
        self.head = Node()
        self.tail = Node()

        self.head.next = self.tail
        self.tail.prev = self.head

    # Remove a node from the linked list
    def remove(self, node):
        prev_node = node.prev
        next_node = node.next

        prev_node.next = next_node
        next_node.prev = prev_node

    # Insert a node just before the tail
    def add_to_end(self, node):
        last = self.tail.prev

        last.next = node
        node.prev = last

        node.next = self.tail
        self.tail.prev = node

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1

        node = self.cache[key]

        # Mark this node as recently used
        self.remove(node)
        self.add_to_end(node)

        return node.value

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            node = self.cache[key]
            node.value = value

            # Updating a key also makes it recently used
            self.remove(node)
            self.add_to_end(node)
        else:
            node = Node(key, value)
            self.cache[key] = node
            self.add_to_end(node)

        # Evict the least recently used entry
        if len(self.cache) > self.capacity:
            oldest = self.head.next
            self.remove(oldest)
            del self.cache[oldest.key]


# Test the implementation
lru = LRUCache(2)

lru.put(1, 1)
lru.put(2, 2)

print(lru.get(1))  # 1

lru.put(3, 3)      # Evicts key 2

print(lru.get(2))  # -1
lru.put(4, 4)      # Evicts key 1

print(lru.get(1))  # -1
print(lru.get(3))  # 3
print(lru.get(4))  # 4
