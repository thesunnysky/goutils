package table

type tableEntry map[interface{}]interface{}

type Table struct {
	table map[interface{}]tableEntry
}

func NewTable() *Table {
	return &Table{
		table: make(map[interface{}]tableEntry),
	}
}

func (t *Table) Load(row, col interface{}) (value interface{}, ok bool) {
	te, ok := t.table[row]
	if !ok {
		return nil, false
	}

	val, ok := te[col]
	return val, ok
}

func (t *Table) LoadOrStore(row, col, val interface{}) (actual interface{}, loaded bool) {
	te, ok := t.table[row]
	if !ok {
		te = tableEntry{}
		t.table[row] = te
	}

	actual, loaded = te[col]
	if loaded {
		return actual, loaded
	}

	te[col] = val
	return val, false
}

func (t *Table) Store(row, col, val interface{}) {
	te, ok := t.table[row]
	if !ok {
		te = tableEntry{}
		t.table[row] = te
	}

	te[col] = val
}

func (t *Table) Remove(row, col interface{}) {
	te, ok := t.table[row]
	if !ok {
		return
	}

	delete(te, col)

	if len(te) == 0 {
		delete(t.table, row)
	}
}

func (t *Table) RemoveRow(row interface{}) {
	_, ok := t.table[row]
	if !ok {
		return
	}

	delete(t.table, row)
}
