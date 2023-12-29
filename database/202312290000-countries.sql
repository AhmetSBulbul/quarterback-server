-- INSERT INTO user 
-- (
--     id,
--     email,
--     password,
--     username
-- )
-- VALUES
-- (
--     1,
--     'admin@example.org',
--     /* 152535 */
--     '$2a$12$LKJcgPyUn4HUuNCGQ6kIBe1WBIIN6uCtHJDDC8W2MXsDwBKNHGKW.',
--     'admin'
-- )

INSERT INTO country 
(
    id,
    name,
)
VALUES
(
    1,
    'Turkey'
)

INSERT INTO city 
(
    id,
    name,
    country_id
)
VALUES
(
    1,
    'Izmir',
    1
),
(
    2,
    'Antalya',
    1
)


INSERT INTO district
(
    id,
    name,
    city_id
)
VALUES
(
    1,
    'Bornova',
    1
),
(
    2,
    'Karsiyaka',
    1
),
(
    3,
    'Konyaalti',
    2
),
(
    4,
    'Muratpasa',
    2
)