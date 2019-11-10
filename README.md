# lol
API Client for League of Legends

[![GoDoc](https://godoc.org/github.com/jonwho/lol?status.svg)](http://godoc.org/github.com/jonwho/lol)
[![Go Report Card](https://goreportcard.com/badge/github.com/jonwho/lol)](https://goreportcard.com/report/github.com/jonwho/lol)
![](https://github.com/jonwho/lol/workflows/tests/badge.svg)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-84%25-brightgreen.svg?longCache=true&style=flat)</a>

# SUPPORTED ENDPOINTS
## CHAMPION-MASTERY-V4
- [x] /lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}
- [x] /lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}/by-champion/{championId}
- [x] /lol/champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}
## CHAMPION-V3
- [x] /lol/platform/v3/champion-rotations
## LEAGUE-EXP-V3
- [x] /lol/league-exp/v4/entries/{queue}/{tier}/{division}
## LEAGUE-V4
- [x] /lol/league/v4/challengerleagues/by-queue/{queue}
- [x] /lol/league/v4/entries/by-summoner/{encryptedSummonerId}
- [x] /lol/league/v4/entries/{queue}/{tier}/{division}
- [x] /lol/league/v4/grandmasterleagues/by-queue/{queue}
- [x] /lol/league/v4/leagues/{leagueId}
- [x] /lol/league/v4/masterleagues/by-queue/{queue}
## LOL-STATUS-V3
- [x] /lol/status/v3/shard-data
## MATCH-V4
- [x] /lol/match/v4/matches/{matchId}
- [x] /lol/match/v4/matchlists/by-account/{encryptedAccountId}
- [ ] /lol/match/v4/timelines/by-match/{matchId}
- [ ] /lol/match/v4/matches/by-tournament-code/{tournamentCode}/ids
- [ ] /lol/match/v4/matches/{matchId}/by-tournament-code/{tournamentCode}
## SPECTATOR-V4
- [ ] /lol/spectator/v4/active-games/by-summoner/{encryptedSummonerId}
- [ ] /lol/spectator/v4/featured-games
## SUMMONER-v4
- [ ] /lol/summoner/v4/summoners/by-account/{encryptedAccountId}
- [x] /lol/summoner/v4/summoners/by-name/{summonerName}
- [x] /lol/summoner/v4/summoners/by-puuid/{encryptedPUUID}
- [ ] /lol/summoner/v4/summoners/{encryptedSummonerId}
## TFT-LEAGUE-V1
- [ ] /tft/league/v1/challenger
- [ ] /tft/league/v1/entries/by-summoner/{encryptedSummonerId}
- [ ] /tft/league/v1/entries/{tier}/{division}
- [ ] /tft/league/v1/grandmaster
- [ ] /tft/league/v1/leagues/{leagueId}
- [ ] /tft/league/v1/master
## TFT-MATCH-V1
- [ ] /tft/match/v1/matches/by-puuid/{encryptedPUUID}/ids
- [ ] /tft/match/v1/matches/{matchId}
## TFT-SUMMONER-v1
- [ ] /tft/summoner/v1/summoners/by-account/{encryptedAccountId}
- [ ] /tft/summoner/v1/summoners/by-name/{summonerName}
- [ ] /tft/summoner/v1/summoners/by-puuid/{encryptedPUUID}
- [ ] /tft/summoner/v1/summoners/{encryptedSummonerId}
## THIRD-PARTY-CODE-V4
- [ ] /lol/platform/v4/third-party-code/by-summoner/{encryptedSummonerId}
## TOURNAMENT-STUB-V4
- [ ] /lol/tournament-stub/v4/codes
- [ ] /lol/tournament-stub/v4/lobby-events/by-code/{tournamentCode}
- [ ] /lol/tournament-stub/v4/providers
- [ ] /lol/tournament-stub/v4/tournaments
## TOURNAMENT-V4
- [ ] /lol/tournament/v4/codes
- [ ] /lol/tournament/v4/lobby-events/by-code/{tournamentCode}
- [ ] /lol/tournament/v4/providers
- [ ] /lol/tournament/v4/tournaments
