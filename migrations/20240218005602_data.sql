-- +goose Up
-- +goose StatementBegin
INSERT INTO maps (uuid,width,height,name,price,created_at,updated_at)
VALUES (
        '122bad95-61c5-4d6f-b6e9-4b92b7ca5d17',
        10,
        10,
        'Room 1',
        100.00,
        now(),
        now()
);
INSERT INTO devices (uuid,type,name,created_at,updated_at)
VALUES (
        'ffd8b514-5d77-4d2e-9c71-6944cb79e4d8',
        'pc',
        'pc 1',
        now(),
        now()
);
INSERT INTO map_points (uuid,device_uuid,map_uuid,x,y,name,created_at,updated_at)
VALUES (
           'e903907b-3cff-4dab-bfc2-359ba5bed7d2',
           'ffd8b514-5d77-4d2e-9c71-6944cb79e4d8',
           '122bad95-61c5-4d6f-b6e9-4b92b7ca5d17',
           2,
           2,
            'Place 1',
           now(),
           now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE FROM maps;
TRUNCATE FROM devices;
TRUNCATE FROM map_points;
-- +goose StatementEnd
