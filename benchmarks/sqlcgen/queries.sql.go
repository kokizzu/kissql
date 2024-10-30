// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: queries.sql

package sqlcgen

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id, name, age FROM users
OFFSET $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, offset int32) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, offset)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Age)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users(name, age)
VALUES ($1, $2) RETURNING id
`

type InsertUserParams struct {
	Name string
	Age  int32
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (int32, error) {
	row := q.queryRow(ctx, q.insertUserStmt, insertUser, arg.Name, arg.Age)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const list10Users = `-- name: List10Users :many
SELECT id, name, age FROM users
OFFSET $1 LIMIT 10
`

func (q *Queries) List10Users(ctx context.Context, offset int32) ([]User, error) {
	rows, err := q.query(ctx, q.list10UsersStmt, list10Users, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Age); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
