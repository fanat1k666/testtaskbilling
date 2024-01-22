-- +goose Up
-- +goose StatementBegin
CREATE TABLE tariffs
(
    id   serial primary key,
    name varchar(255) NOT NULL,
    price int NOT NULL,
    speed varchar(255),
    CONSTRAINT tariffs_unique UNIQUE (name)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE "user"
(
    id      serial primary key,
    user_id int NOT NULL,
    CONSTRAINT user_unique UNIQUE (user_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE user_tariffs
(
    id             serial primary key,
    user_id        int       NOT NULL,
    tariffs_id     int       NOT NULL,
    address varchar(255),
    constraint user_fk foreign key (user_id) references "user" on update cascade on delete cascade,
    constraint tariffs_fk foreign key (tariffs_id) references tariffs on update cascade on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tariffs;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE "user";
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE user_tariffs;
-- +goose StatementEnd
