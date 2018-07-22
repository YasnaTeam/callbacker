package storage

type MemoryTable struct {
	max uint
	table map[string]string
}

func (m *MemoryTable) Get(route string) (string, error) {
	if value, ok := m.table[route]; ok {
		return value, nil
	}

	return "", &RouteError{route}
}

func (m *MemoryTable) Set(route, value string) error {
	if m.CanSet() {
		m.table[route] = value

		return nil
	}

	return &CanNotSetError{"You has reached to the limit."}
}

func (m *MemoryTable) Unset(route string) error {
	if _, ok := m.table[route]; ok {
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