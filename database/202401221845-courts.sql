INSERT INTO court
(
    ID,
    name,
    districtID,
    coordinate,
    address
)
VALUES
(
    1,
    'Bornova',
    1,
    ST_GeomFromText('POINT(38.456375 27.213445)'),
    'Bornova, Izmir'
)