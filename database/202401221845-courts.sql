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
    'Bornova Spor Salonu',
    1,
    ST_GeomFromText('POINT(38.456375 27.213445)'),
    'Bornova, Izmir'
),
(
    2,
    'Dr. Cavit Ozyegin Spor Salonu',
    1,
    ST_GeomFromText('POINT(38.46752422573013 27.22935507045534)'),
    'Dr. Cavit Özyeğin Spor Salonu, Erzene, 116/19. Sk. 2/B, 35040 Bornova/İzmir'
),
(
    3,
    'Izmir Tofas Basketbol Okullari',
    1,
    ST_GeomFromText('POINT(38.466815923856274 27.238966095533364)'),
    'İzmir Tofaş Basketbol Okulları Ve Kursları Bornova Buca Şubeleri, Evka 3, 119/1. Sk. NO:3, 35050 Bornova/İzmir'
),
(
    4,
    'Borbasketbol Sahasi',
    1,
    ST_GeomFromText('POINT(38.46837231564899 27.21101493272493)'),
    'Kızılay, Aşık Veysel Rekreasyon Alanı Açık Hava Tiyatrosu, 35030 Bornova/İzmir'
),
(
    5,
    'Ogrenci Koyu Basketbol Sahasi',
    1,
    ST_GeomFromText('POINT(38.45425480188003 27.215131517819206)'),
    'Kazımdirik, Ege Ünv. Ögrenci Köyü İç Yol, 35100 Bornova/İzmir'
),
(
    6,
    'Esin Basketbol Sahasi',
    2,
    ST_GeomFromText('POINT(38.494033260691396 27.092755416078628)'),
    'Mustafa Kemal, 6753/10. Sk., 35570 Karşıyaka/İzmir'
),
(
    7,
    'Mustafa Kemal Ataturk Karsiyaka Spor Salonu',
    2,
    ST_GeomFromText('POINT(38.47619587691626 27.075939213676918)'),
    'Mavişehir, 2040/3. Sk., 35540 Karşıyaka/İzmir'
),
(
    8,
    'Konyaalti Sahili 3. Basketbol Sahasi',
    3,
    ST_GeomFromText('POINT(36.863858605570506 30.640913069420886)'),
    'Altınkum, 07070 Konyaaltı/Antalya'
),
(
    9,
    'Lara Basketbol Sahasi',
    4,
    ST_GeomFromText('POINT(36.86070061782464 30.7480464580622)'),
    'Fener, Bülent Ecevit Blv. 11 292, 07160 Muratpaşa/Antalya'
)