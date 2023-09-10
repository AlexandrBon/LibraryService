DROP TABLE IF EXISTS book CASCADE;
CREATE TABLE book(
    id              INT AUTO_INCREMENT PRIMARY KEY,
    title           VARCHAR(99),
    page_count      INT,
    publishing_year INT CHECK ( publishing_year <= 2023 )
);

DROP TABLE IF EXISTS author CASCADE;
CREATE TABLE author(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(99),
    birth_date DATE,
    book_count INT
);

DROP TABLE IF EXISTS author_x_book CASCADE;
CREATE TABLE author_x_book(
    author_id INT NOT NULL,
    book_id   INT NOT NULL,

    FOREIGN KEY (author_id) REFERENCES author(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES book(id) ON DELETE CASCADE,

    PRIMARY KEY (author_id, book_id)
);
