package handlers

import "library/services"

type BookHandler struct {
	Service *services.BookService
}

/*
create book - admin
delete book - admin
update book - admin

list All Books - user
get Book By ID - user
list My Books - user
chechout book - user
check in book - user
find book by genre
find book by author
find book by year
find book by title
*/