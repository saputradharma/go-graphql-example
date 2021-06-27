-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Links(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    title VARCHAR (255) ,
    address VARCHAR (255) ,
    user_id INT ,
    FOREIGN KEY (user_id) REFERENCES users(id) ,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd
