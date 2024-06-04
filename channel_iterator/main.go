// This pattern can be applied in the common case where we have to populate a channel with the items of a container type, which contains index-addressable field items. For this, we can define a method Iter() on the content type which returns receive-only channel items (Iter() is a channel factory), as follows:

func (c *container) Iter() <-chan items {
	ch := make(chan item)

	go func() {
		for i := 0; i < c.Len(); i++ {    // for trees/graphs algos, replace for-loop with DFS
			ch <- c.items[i]
		}
	}()
	return ch
}