package query

const (
	CreateBook = `
    INSERT INTO 
     book 
    (
      title,
      author,
      description
    )
    VALUES ($1, $2, $3)
    RETURNING *;
   `

	DeleteBook = `
    DELETE FROM 
      book
    WHERE id = $1;
  `

	GetBookByID = `
    SELECT * FROM
      book
    WHERE id = $1;
  `

	GetAllBooks = `
    SELECT * FROM
      book;
  `
	UpdateBook = `
    UPDATE book 
    SET 
      title = COALESCE(NULLIF($2, ''), title), 
      author = COALESCE(NULLIF($3,''), author), 
      description = COALESCE(NULLIF($4,''), description) 
    WHERE id = $1;
    
  `
)
