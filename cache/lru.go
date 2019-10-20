package cache

// CachedItem represents the item cached.
type CachedItem struct {
	value Value
	key   string
}

func (cachedItem CachedItem) String() string {
	return cachedItem.value.String()
}

// LRUCache that keeps the cached items in access frequency
// order for the eviction.
type LRUCache struct {
	capacity  int
	itemOrder *DoublyLinkedList
	table     map[string]*Node
}

// NewLRUCache create and returns *LRUCache.
func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		capacity:  size,
		itemOrder: &DoublyLinkedList{},
		table:     make(map[string]*Node)}
}

// Length returns the current size of the cache.
func (cache *LRUCache) Length() int {
	return len(cache.table)
}

// Get retries the cached item from the cache.
func (cache *LRUCache) Get(key string) Value {
	node := cache.table[key]
	if node == nil {
		return nil
	}

	cache.itemOrder.MoveToFront(node)
	return node.Value()
}

// Put an item to the cache.
func (cache *LRUCache) Put(key string, value Value) {
	node := cache.table[key]
	if node != nil {
		cache.itemOrder.MoveToFront(node)
		// CachedItem is also a value
		node.SetValue(CachedItem{key: key, value: value})
		return
	}

	newNode := &Node{value: CachedItem{key: key, value: value}}
	cache.table[key] = newNode
	cache.itemOrder.AddToFront(newNode)

	if len(cache.table) > cache.capacity {
		evictionNode := cache.itemOrder.RemoveFromTail()
		cachedItem, ok := evictionNode.Value().(CachedItem)

		if !ok {
			panic("Wrong type item in the table")
		}

		delete(cache.table, cachedItem.key)
	}
}
