/*
* Top 10 loaned books of the month
*/
CREATE OR REPLACE VIEW top10books AS
  SELECT count(*) as loans, l.book_id
  FROM loans l
  WHERE EXTRACT(MONTH FROM l.loan_date) = EXTRACT(MONTH FROM CURRENT_DATE)
  GROUP BY l.book_id
  ORDER BY count(*) desc
  LIMIT 10;

SELECT t.loans, b.*
FROM top10books t
	LEFT JOIN books b on t.book_id = b.isbn


/*
* Number of times each user was late in returning a book
*/

CREATE OR REPLACE VIEW late_returns AS
  SELECT L.user_id, count(*) as late_returns
  FROM loans l 
  WHERE L.actual_return > L.due_date 
  GROUP BY L.user_id;


SELECT u.user_id , u.first_name , u.last_name , u.email , COALESCE(l.late_returns,0)
FROM users u
	LEFT JOIN late_returns l on l.user_id = u.user_id;


