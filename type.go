package fmpcloud

func JsonFieldFloat64ToInt64(a []map[string]interface{}, k string) {
	for _, m := range a {
		if oldValue, exists := m[k]; exists {
			if oldValueFloat, ok := oldValue.(float64); ok {
				newValueInt := int64(oldValueFloat)
				m[k] = newValueInt
			}
		}
	}
}
