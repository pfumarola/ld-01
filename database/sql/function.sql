/*
* A function to delete authors without books
*/

CREATE OR REPLACE FUNCTION deleteAuthorsWithoutBooks()
RETURNS void AS $$
BEGIN
  DELETE FROM author
  WHERE author.id NOT IN (
    SELECT DISTINCT author_id
    FROM authored
  );
END;
$$ LANGUAGE plpgsql;

SELECT deleteAuthorsWithoutBooks()