-- name: ListProducts :many
SELECT
    *
FROM
    products;


-- name: FindProductByID :one
SELECT * FROM PRODUCTS WHERE id = $1;