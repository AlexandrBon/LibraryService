INSERT INTO book(id, title, page_count, publishing_year) VALUES
    (1, 't_1', 40, 1988),
    (2, 't_2', 410, 1894),
    (3, 't_3', 67, 2008),
    (4, 't_4', 530, 2013),
    (5, 't_5', 148, 1795);

INSERT INTO author(id, name, birth_date, book_count) VALUES
    (1, 'a_1', '1967-12-04', 97),
    (2, 'a_2', '1959-02-27', 33),
    (3, 'a_3', '1868-03-23', 54),
    (4, 'a_4', '1983-08-14', 134),
    (5, 'a_5', '1805-05-30', 28);

INSERT INTO author_x_book(author_id, book_id) VALUES
    (1, 1),
    (2, 1),
    (2, 3),
    (3, 2),
    (4, 3),
    (4, 4);
