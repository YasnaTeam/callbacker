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

func (m *MemoryTable) Get(route string) (interface{}, error) {
	if value, ok := m.table[route]; ok {
		return value, nil
	}

	return "", &RouteError{route}
}

func (m *MemoryTable) Set(route string, value interface{}) error {
	if m.CanSet() {
		m.table[route] = value

		return nil
	}

	return &CanNotSetError{"You has reached to the limit."}
}

func (m *MemoryTable) Has(route string) bool {
	if _, ok := m.table[route]; ok {
		return true
	}

	return false
}

func (m *MemoryTable) Unset(route string) error {
	if m.Has(route) {
		delete(m.table, route)

		return nil
	}

	return &RouteError{route}
}

func (m *MemoryTable) CanSet() bool {
	if m.max > 0 && len(m.table) < int(m.max) {
		return true
	}

	return false
}