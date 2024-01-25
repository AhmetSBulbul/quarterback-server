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
    game_player.isCanceled as isCanceled,
    (homeScore - awayScore) > 0 as homeWin,
    (awayScore - homeScore) > 0 as awayWin,
    game_player.isHomeSide > 0 = ((homeScore - awayScore) > 0) as isWin,
    game_player.isHomeSide > 0 = ((awayScore - homeScore) > 0) as isLoss
    from game
    inner JOIN district
    inner join user u ON u.districtID = district.id
    inner join game_player on game_player.playerID = u.ID and game_player.gameID = game.id
    where game.endedAt is not null