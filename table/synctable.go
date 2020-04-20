package table

import "sync"

type SyncTable struct {
	sync.Mutex
	table *Table
}

func NewSyncTable() *SyncTable {
	return &SyncTable{
		Mutex: sync.Mutex{},
		table: NewTable(),
	}
}

func (t *SyncTable) Load(row, col interface{}) (value interface{}, ok bool) {
	t.Lock()
	defer t.Unlock()

	return t.table.Load(row, col)
}

func (t *SyncTable) LoadOrStore(row, col, val interface{}) (actual interface{}, loaded bool) {
	t.Lock()
	defer t.Unlock()

	return t.table.LoadOrStore(row, col, val)
}

func (t *SyncTable) Store(row, col, val interface{}) {
	t.Lock()
	defer t.Unlock()

	t.table.Store(row, col, val)
}

func (t *SyncTable) Remove(row, col interface{}) {
	t.Lock()
	defer t.Unlock()

	t.table.Remove(row, col)
}

func (t *SyncTable) RemoveRow(row interface{}) {
	t.Lock()
	defer t.Unlock()

	t.table.RemoveRow(row)
}
