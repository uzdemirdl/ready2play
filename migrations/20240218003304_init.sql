-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS devices
(
    uuid uuid PRIMARY KEY,
    type varchar NOT NULL,
    name varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS maps
(
    uuid uuid PRIMARY KEY,
    width smallint NOT NULL,
    height smallint NOT NULL,
    name varchar NOT NULL,
    price numeric(10, 2) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS map_points
(
    uuid uuid PRIMARY KEY,
    device_uuid uuid NOT NULL,
    map_uuid uuid NOT NULL,
    x smallint NOT NULL,
    y smallint NOT NULL,
    name varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT FK_map_points_devices FOREIGN KEY(device_uuid)
        REFERENCES devices(uuid),
    CONSTRAINT FK_map_points_maps FOREIGN KEY(map_uuid)
        REFERENCES maps(uuid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS map_points;
DROP TABLE IF EXISTS maps;
DROP TABLE IF EXISTS devices;
-- +goose StatementEnd
