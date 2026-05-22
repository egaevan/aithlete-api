package service

type MockPasswordHasher struct{}

func (m *MockPasswordHasher) Hash(password string) (string, error) {
	return "hashed:" + password, nil
}

func (m *MockPasswordHasher) Verify(password, hash string) bool {
	return hash == "hashed:"+password
}
