-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    username VARCHAR (127) NOT NULL UNIQUE,
    password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
