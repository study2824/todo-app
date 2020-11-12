CREATE TABLE todo
(
    id         SERIAL       NOT NULL PRIMARY KEY,
    title      varchar(255) NOT NULL,
    text       text         NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL
)