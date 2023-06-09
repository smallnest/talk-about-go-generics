package ch3

// import (
// 	"sync"
// 	"sync/atomic"
// 	"unsafe"
// )

// type Map[K comparable, V any] struct {
// 	mu sync.Mutex

// 	read atomic.Value // readOnly

// 	dirty map[K]*entry[V]

// 	misses int
// }

// type readOnly[K comparable, V any] struct {
// 	m       map[K]*entry[V]
// 	amended bool
// }

// var expunged = unsafe.Pointer(new(interface{}))

// type entry[T any] struct {
// 	p unsafe.Pointer // *T
// }

// func newEntry[T any](i T) *entry[T] {
// 	return &entry[T]{p: unsafe.Pointer(&i)}
// }

// func (m *Map[K, V]) Load(key K) (value V, ok bool) {
// 	read, _ := m.read.Load().(readOnly[K, V])
// 	e, ok := read.m[key]
// 	if !ok && read.amended {
// 		m.mu.Lock()
// 		read, _ = m.read.Load().(readOnly[K, V])
// 		e, ok = read.m[key]
// 		if !ok && read.amended {
// 			e, ok = m.dirty[key]
// 			m.missLocked()
// 		}
// 		m.mu.Unlock()
// 	}
// 	if !ok {
// 		var zero V
// 		return zero, false
// 	}
// 	return e.load()
// }

// func (m *entry[T]) load() (value T, ok bool) {
// 	p := atomic.LoadPointer(&m.p)
// 	if p == nil || p == expunged {
// 		var zero T
// 		return zero, false
// 	}
// 	return *(*T)(p), true
// }

// func (m *Map[K, V]) Store(key K, value V) {
// 	read, _ := m.read.Load().(readOnly[K, V])
// 	if e, ok := read.m[key]; ok && e.tryStore(&value) {
// 		return
// 	}

// 	m.mu.Lock()
// 	read, _ = m.read.Load().(readOnly[K, V])
// 	if e, ok := read.m[key]; ok {
// 		if e.unexpungeLocked() {
// 			m.dirty[key] = e
// 		}
// 		e.storeLocked(&value)
// 	} else if e, ok := m.dirty[key]; ok {
// 		e.storeLocked(&value)
// 	} else {
// 		if !read.amended {
// 			m.dirtyLocked()
// 			m.read.Store(readOnly[K, V]{m: read.m, amended: true})
// 		}
// 		m.dirty[key] = newEntry(value)
// 	}
// 	m.mu.Unlock()
// }

// func (m *entry[T]) tryStore(i *T) bool {
// 	for {
// 		p := atomic.LoadPointer(&m.p)
// 		if p == expunged {
// 			return false
// 		}
// 		if atomic.CompareAndSwapPointer(&m.p, p, unsafe.Pointer(i)) {
// 			return true
// 		}
// 	}
// }

// func (m *entry[T]) unexpungeLocked() (wasExpunged bool) {
// 	return atomic.CompareAndSwapPointer(&m.p, expunged, nil)
// }

// func (m *entry[T]) storeLocked(i *T) {
// 	atomic.StorePointer(&m.p, unsafe.Pointer(i))
// }

// func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
// 	read, _ := m.read.Load().(readOnly[K, V])
// 	if e, ok := read.m[key]; ok {
// 		actual, loaded, ok := e.tryLoadOrStore(value)
// 		if ok {
// 			return actual, loaded
// 		}
// 	}

// 	m.mu.Lock()
// 	read, _ = m.read.Load().(readOnly[K, V])
// 	if e, ok := read.m[key]; ok {
// 		if e.unexpungeLocked() {
// 			m.dirty[key] = e
// 		}
// 		actual, loaded, _ = e.tryLoadOrStore(value)
// 	} else if e, ok := m.dirty[key]; ok {
// 		actual, loaded, _ = e.tryLoadOrStore(value)
// 		m.missLocked()
// 	} else {
// 		if !read.amended {
// 			// We're adding the first new key to the dirty map.
// 			// Make sure it is allocated and mark the read-only map as incomplete.
// 			m.dirtyLocked()
// 			m.read.Store(readOnly[K, V]{m: read.m, amended: true})
// 		}
// 		m.dirty[key] = newEntry(value)
// 		actual, loaded = value, false
// 	}
// 	m.mu.Unlock()

// 	return actual, loaded
// }
// func (m *entry[T]) tryLoadOrStore(i T) (actual T, loaded, ok bool) {
// 	p := atomic.LoadPointer(&m.p)
// 	if p == expunged {
// 		var zero T
// 		return zero, false, false
// 	}
// 	if p != nil {
// 		return *(*T)(p), true, true
// 	}

// 	ic := i
// 	for {
// 		if atomic.CompareAndSwapPointer(&m.p, nil, unsafe.Pointer(&ic)) {
// 			return i, false, true
// 		}
// 		p = atomic.LoadPointer(&m.p)
// 		if p == expunged {
// 			var zero T
// 			return zero, false, false
// 		}
// 		if p != nil {
// 			return *(*T)(p), true, true
// 		}
// 	}
// }

// func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
// 	read, _ := m.read.Load().(readOnly[K, V])
// 	e, ok := read.m[key]
// 	if !ok && read.amended {
// 		m.mu.Lock()
// 		read, _ = m.read.Load().(readOnly[K, V])
// 		e, ok = read.m[key]
// 		if !ok && read.amended {
// 			e, ok = m.dirty[key]
// 			delete(m.dirty, key)
// 			m.missLocked()
// 		}
// 		m.mu.Unlock()
// 	}
// 	if ok {
// 		return e.delete()
// 	}

// 	var zero V
// 	return zero, false
// }

// func (m *Map[K, V]) Delete(key K) {
// 	m.LoadAndDelete(key)
// }

// func (m *entry[T]) delete() (value T, ok bool) {
// 	for {
// 		p := atomic.LoadPointer(&m.p)
// 		if p == nil || p == expunged {
// 			var zero T
// 			return zero, false
// 		}
// 		if atomic.CompareAndSwapPointer(&m.p, p, nil) {
// 			return *(*T)(p), true
// 		}
// 	}
// }

// func (m *Map[K, V]) Range(f func(key K, value V) bool) {
// 	read, _ := m.read.Load().(readOnly[K, V])
// 	if read.amended {
// 		m.mu.Lock()
// 		read, _ = m.read.Load().(readOnly[K, V])
// 		if read.amended {
// 			read = readOnly[K, V]{m: m.dirty}
// 			m.read.Store(read)
// 			m.dirty = nil
// 			m.misses = 0
// 		}
// 		m.mu.Unlock()
// 	}

// 	for k, e := range read.m {
// 		v, ok := e.load()
// 		if !ok {
// 			continue
// 		}
// 		if !f(k, v) {
// 			break
// 		}
// 	}
// }

// func (m *Map[K, V]) missLocked() {
// 	m.misses++
// 	if m.misses < len(m.dirty) {
// 		return
// 	}
// 	m.read.Store(readOnly[K, V]{m: m.dirty})
// 	m.dirty = nil
// 	m.misses = 0
// }

// func (m *Map[K, V]) dirtyLocked() {
// 	if m.dirty != nil {
// 		return
// 	}

// 	read, _ := m.read.Load().(readOnly[K, V])
// 	m.dirty = make(map[K]*entry[V], len(read.m))
// 	for k, e := range read.m {
// 		if !e.tryExpungeLocked() {
// 			m.dirty[k] = e
// 		}
// 	}
// }

// func (m *entry[T]) tryExpungeLocked() (isExpunged bool) {
// 	p := atomic.LoadPointer(&m.p)
// 	for p == nil {
// 		if atomic.CompareAndSwapPointer(&m.p, nil, expunged) {
// 			return true
// 		}
// 		p = atomic.LoadPointer(&m.p)
// 	}
// 	return p == expunged
// }

// func (m *Map[K, V]) Len() int {
// 	var length int
// 	m.Range(func(K, V) bool {
// 		length++
// 		return true
// 	})
// 	return length
// }
