-- name: CreateBook :one
INSERT INTO books (id, title, author, publisher, price)
    VALUES (nextval('BOOK_ID_SEQ'), $1, $2, $3, $4)
    RETURNING id, title, author, publisher, price
;

-- name: GetBookByID :one
SELECT id, title, author, publisher, price
    FROM books
    WHERE id = $1
;

-- name: ListBooks :many
SELECT id, title, author, publisher, price
    FROM books
;

-- name: DeleteBookByID :exec
DELETE
    FROM books
    WHERE id = $1
; 
