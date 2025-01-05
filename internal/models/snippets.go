package models

import (
	"database/sql"
	"errors"
	"time"
)

// define snippet matching database models
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	// write the sql statement we want to create
	// split into two line for readability
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// use the exec method on the embedded connection pool
	// and execute the statement
	// first parameeter is SQL statement, the next one is placeholder
	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	// to get the result of our newly inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// id returned in int64 so we convert it to int
	return int(id), nil
}

// return specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	// write the SQL statement we want to exec

	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	//queryRow execute our SQL statement
	// return pointer to a sql.Row object
	row := m.DB.QueryRow(stmt, id)

	//pointer to a zeroed snipped struct
	s := &Snippet{}

	//copy each value in sql.Row to corresponding Snippet struct
	// argments of rowScan are pointers to the place u want to copy the data
	// number of arguments = number of columns returned
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	// write the SQL statement
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	//Query() to execute SQL statement
	//returns sql.Rows contain the result of our query
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	//ensure the sql.rows resultset is always properly closed
	//before the latest() method returns
	//should come after check for an error on query()
	defer rows.Close()

	//empty slice to hold snippet structs
	snippets := []*Snippet{}

	//iterate through rows in resultset
	//if iteration all the rows completed then resultset automatically closes itself
	for rows.Next() {
		//pointer to a new zeroed snippet struct
		s := &Snippet{}

		//copy the values from each field in the row
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		//append it to the slice of snippets
		snippets = append(snippets, s)
	}

	// retrieve error during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//return snippets slice
	return snippets, nil
}
