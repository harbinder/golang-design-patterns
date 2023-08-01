package misc

import (
	"testing"
)

/*
Declare an interface for database
*/
type Database interface {
	Get(input interface{}) (interface{}, error)
}

/*
Implementation of Database interface
For actual DB
*/
type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (m *Mysql) Get(input interface{}) (output interface{}, err error) {
	output = "Mysql: Get"
	return
}

/*
Implementation of Database interface
For Testing
*/
type MockDB struct {
	GetFunc func(ip interface{}) (op interface{}, err error)
}

func (m *MockDB) Get(input interface{}) (output interface{}, err error) {
	return m.GetFunc(input)
}

type Repository struct {
	db Database
}

// Create instance of repository with db instance
func NewRepository(db Database) *Repository {
	return &Repository{db: db}
}
func (r *Repository) Get() (output interface{}, err error) {
	output, err = r.db.Get(struct{}{})
	return
}

func TestDatabase(t *testing.T) {
	mockDb := &MockDB{
		GetFunc: func(ip interface{}) (op interface{}, err error) {
			op = "MockDB"
			return
		},
	}

	testCases := []struct {
		name            string
		input           interface{}
		output_expected interface{}
	}{
		{name: "MockDB-TC-1", input: mockDb, output_expected: "MockDB-TC-1"},
		{name: "MockDB-TC-2", input: mockDb, output_expected: "MockDB"},
	}

	for _, tc := range testCases {
		tc := tc                            // avoid closure variable issue
		t.Run(tc.name, func(t *testing.T) { // executes test cases in go routines BUT not in Parallel
			t.Parallel() // this method will run tests in parallel
			r := NewRepository(tc.input.(*MockDB))
			output_actual, err := r.Get()

			if err != nil {
				t.Errorf("Error in testing: ")
			}
			if output_actual != tc.output_expected {
				t.Errorf("Failed: Actual output: %v, Expected output: %v", output_actual, tc.output_expected)
			}
		})
	}

}
