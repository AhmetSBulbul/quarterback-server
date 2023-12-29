INSERT INTO country 
(
    id,
    name
)
VALUES
(
    1,
    'Turkey'
);

INSERT INTO city 
(
    id,
    name,
    countryID
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
);


INSERT INTO district
(
    id,
    name,
    cityID
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
);