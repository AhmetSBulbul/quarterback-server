CREATE VIEW districtPoints AS
select 
    district.id as districtID,
    district.name as districtName,
    game.id as gameID,
    u.id as playerID,
    u.username as username,
    game.endedAt as date,
    homeScore,
    awayScore,
    game_player.isHomeSide as isHomeSide,
    (homeScore - awayScore) > 0 as homeWin,
    (awayScore - homeScore) > 0 as awayWin,
    game_player.isHomeSide > 0 = ((homeScore - awayScore) > 0) as isWin
    from game
    inner join game_player on game.ID = game_player.gameID
    inner join user u ON u.ID = game_player.playerID
    inner join district on u.districtID = district.id
