package utils

// DerefString return a string from *string
// Or empty string if pointer is nil
func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
