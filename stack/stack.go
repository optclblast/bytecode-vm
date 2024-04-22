package stack

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type stack struct {
	mu        sync.Mutex
	stack_ptr atomic.Int32
	pool      []unsafe.Pointer
}

func (s *stack) Push(ptr unsafe.Pointer) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pool[s.stack_ptr.Load()] = ptr
	_ = s.stack_ptr.Add(1)
}

func (s *stack) Pop() unsafe.Pointer {
	s.mu.Lock()
	defer s.mu.Unlock()

	ptr := s.pool[s.stack_ptr.Load()-1]
	s.stack_ptr.Add(-1)
	return ptr
}

// allocstack allocates a vm's stack
func AllocStack(size uint32) Stack {
	return &stack{
		pool: make([]unsafe.Pointer, size),
	}
}

type Stack interface {
	Push(unsafe.Pointer)
	Pop() unsafe.Pointer
}
