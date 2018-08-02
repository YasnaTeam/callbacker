package storage

type MemoryTable struct {
	max uint
	table map[string]interface{}
}

func NewMemoryTable(max uint) *MemoryTable {
	table := make(map[string]interface{})

	return &MemoryTable{
		max,
		table,
	}
}

func (m *MemoryTable) Get(key string) (interface{}, error) {
	if value, ok := m.table[key]; ok {
		return value, nil
	}

	return "", &RouteError{key}
}

func (m *MemoryTable) Set(key string, value interface{}) error {
	if m.CanSet() {
		m.table[key] = value

		return nil
	}

	return &CanNotSetError{"You has reached to the limit."}
}

func (m *MemoryTable) Has(key string) bool {
	if _, ok := m.table[key]; ok {
		return true
	}

	return false
}

func (m *MemoryTable) Unset(key string) error {
	if m.Has(key) {
		delete(m.table, key)

		return nil
	}

	return &RouteError{key}
}

func (m *MemoryTable) CanSet() bool {
	if m.max <= 0 || len(m.table) < int(m.max) {
		return true
	}

	return false
}

func (m *MemoryTable) GetKey(value interface{}) (string, error) {
	for i, v := range m.table {
		if v == value {
			return i, nil
		}
	}

	return "", &RouteError{"Not key has this value."}
}

func (m *MemoryTable) All() map[string]interface{} {
	return m.table
}

func (m *MemoryTable) Truncate() {
	for i, _ := range m.All() {
		m.Unset(i)
	}
}