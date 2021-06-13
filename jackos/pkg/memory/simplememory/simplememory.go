package simplememory

// simple memory
// 簡易版メモリ管理

// Manager .
type Manager struct {
	Free int64
}

// New .
func New() *Manager {
	return &Manager{Free: 0}
}

// Alloc sizeで指定されたメモリブロックの割当を行う
func (m *Manager) Alloc(size int64) int64 {
	pointer := m.Free
	m.Free += size
	return pointer
}

// DeAlloc 与えられたオブジェクトについて、そのメモリ領域の破棄を行う
func (m *Manager) DeAlloc(obj interface{}) {
	// Do Nothing
}
