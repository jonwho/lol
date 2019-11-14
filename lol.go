package lol

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dghubble/sling"
)

const (
	baseURL            = "api.riotgames.com/lol"
	defaultRegion      = "na1"
	maxIdleConnections = 10
	requestTimeout     = 5
)

var (
	// DefaultHTTPClient default http client to use
	DefaultHTTPClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
)

type ChampionInfo struct {
	FreeChampionIDs              []int `json:"freeChampionIds"`
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

type ChampionMasteryDTO struct {
	ChampionLevel                int    `json:"championLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	SummonerID                   string `json:"summonerId"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionID                   int    `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
}

// Client API struct to League of Legends
type Client struct {
	Token, Region string
	sling         *sling.Sling
	httpClient    *http.Client
}

type CurrentGameInfo struct {
	GameID            int64                    `json:"gameId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	PlatformID        string                   `json:"platformId"`
	GameMode          string                   `json:"gameMode"`
	MapID             int64                    `json:"mapId"`
	GameType          string                   `json:"gameType"`
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	GameLength        int64                    `json:"gameLength"`
	GameQueueConfigID int64                    `json:"gameQueueConfigId"`
}

type BannedChampion struct {
	PickTurn   int   `json:"pickTurn"`
	ChampionID int64 `json:"championId"`
	TeamID     int64 `json:"teamId"`
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

type CurrentGameParticipant struct {
	ProfileIconID            int64                     `json:"profileIconId"`
	ChampionID               int64                     `json:"championId"`
	SummonerName             string                    `json:"summonerName"`
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
	Bot                      bool                      `json:"bot"`
	Perks                    Perks                     `json:"perks"`
	Spell2ID                 int64                     `json:"spell2Id"`
	Spell1ID                 int64                     `json:"spell1Id"`
	TeamID                   int64                     `json:"teamId"`
	SummonerID               string                    `json:"summonerId"`
}

type Participant struct {
	ProfileIconID int64  `json:"profileIconId"`
	ChampionID    int64  `json:"championId"`
	SummonerName  string `json:"summonerName"`
	Bot           bool   `json:"bot"`
	Spell2ID      int64  `json:"spell2Id"`
	Spell1ID      int64  `json:"spell1Id"`
	TeamID        int64  `json:"teamId"`
	SummonerID    string `json:"summonerId"`
}

type GameCustomizationObject struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

type Perks struct {
	PerkStyle    int64   `json:"perkStyle"`
	PerkIDs      []int64 `json:"perkIds"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

type EntriesParams struct {
	Page string `url:"page,omitempty"`
}

type FeaturedGames struct {
	ClientRefreshInterval int64              `json:"clientRefreshInterval"`
	GameList              []FeaturedGameInfo `json:"gameList"`
}

type FeaturedGameInfo struct {
	GameID            int64                    `json:"gameId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	PlatformID        string                   `json:"platformId"`
	GameMode          string                   `json:"gameMode"`
	MapID             int64                    `json:"mapId"`
	GameType          string                   `json:"gameType"`
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	GameLength        int64                    `json:"gameLength"`
	GameQueueConfigID int64                    `json:"gameQueueConfigId"`
}

type Incident struct {
	Active    bool      `json:"active"`
	CreatedAt string    `json:"created_at"`
	ID        int       `json:"id"`
	Updates   []Message `json:"updates"`
}

type LeagueEntryDTO struct {
	QueueType    string        `json:"queueType"`
	SummonerName string        `json:"summonerName"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	Wins         int           `json:"wins"`
	Veteran      bool          `json:"veteran"`
	Losses       int           `json:"losses"`
	Rank         string        `json:"rank"`
	Tier         string        `json:"tier"`
	Inactive     bool          `json:"inactive"`
	FreshBlood   bool          `json:"freshBlood"`
	LeagueID     string        `json:"leagueId"`
	SummonerID   string        `json:"summonerId"`
	LeaguePoints int           `json:"leaguePoints"`
}

type LeagueExpEntriesParams struct {
	Page string `url:"page,omitempty"`
}

type LeagueItemDTO struct {
	SummonerName string        `json:"summonerName"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	Wins         int           `json:"wins"`
	Veteran      bool          `json:"veteran"`
	Losses       int           `json:"losses"`
	FreshBlood   bool          `json:"freshBlood"`
	Inactive     bool          `json:"inactive"`
	Rank         string        `json:"rank"`
	SummonerID   string        `json:"summonerId"`
	LeaguePoints int           `json:"leaguePoints"`
}

type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Tier     string          `json:"tier"`
	Entries  []LeagueItemDTO `json:"entries"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
}

type PlayerDTO struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  string `json:"currentAccountId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        string `json:"summonerId"`
	AccountID         string `json:"accountId"`
}

type ParticipantIdentityDTO struct {
	Player        PlayerDTO `json:"player"`
	ParticipantID int       `json:"participantId"`
}

type TeamBansDTO struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
}

type TeamStatsDTO struct {
	FirstDragon          bool          `json:"firstDragon"`
	Bans                 []TeamBansDTO `json:"bans"`
	FirstInhibitor       bool          `json:"firstInhibitor"`
	Win                  string        `json:"win"`
	FirstRiftHerald      bool          `json:"firstRiftHerald"`
	FirstBaron           bool          `json:"firstBaron"`
	BaronKills           int           `json:"baronKills"`
	RiftHeraldKills      int           `json:"riftHeraldKills"`
	FirstBlood           bool          `json:"firstBlood"`
	TeamID               int           `json:"teamId"`
	FirstTower           bool          `json:"firstTower"`
	VilemawKills         int           `json:"vilemawKills"`
	InhibitorKills       int           `json:"inhibitorKills"`
	TowerKills           int           `json:"towerKills"`
	DominionVictoryScore int           `json:"dominionVictoryScore"`
	DragonKills          int           `json:"dragonKills"`
}

type ParticipantTimelineDTO struct {
	Lane                        string             `json:"lane"`
	ParticipantID               int                `json:"participantId"`
	CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
	XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
}

type ParticipantStatsDTO struct {
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	VisionScore                     int  `json:"visionScore"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
	Perk1Var1                       int  `json:"perk1Var1"`
	Perk1Var3                       int  `json:"perk1Var3"`
	Perk1Var2                       int  `json:"perk1Var2"`
	TripleKills                     int  `json:"tripleKills"`
	Perk5                           int  `json:"perk5"`
	Perk4                           int  `json:"perk4"`
	PlayerScore9                    int  `json:"playerScore9"`
	PlayerScore8                    int  `json:"playerScore8"`
	Kills                           int  `json:"kills"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore6                    int  `json:"playerScore6"`
	Perk5Var1                       int  `json:"perk5Var1"`
	Perk5Var3                       int  `json:"perk5Var3"`
	Perk5Var2                       int  `json:"perk5Var2"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	StatPerk1                       int  `json:"statPerk1"`
	StatPerk0                       int  `json:"statPerk0"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	Perk2Var2                       int  `json:"perk2Var2"`
	Perk2Var3                       int  `json:"perk2Var3"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	Perk2Var1                       int  `json:"perk2Var1"`
	Perk4Var1                       int  `json:"perk4Var1"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	Perk4Var3                       int  `json:"perk4Var3"`
	WardsKilled                     int  `json:"wardsKilled"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	QuadraKills                     int  `json:"quadraKills"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item6                           int  `json:"item6"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Perk1                           int  `json:"perk1"`
	Perk0                           int  `json:"perk0"`
	Perk3                           int  `json:"perk3"`
	Perk2                           int  `json:"perk2"`
	Perk3Var3                       int  `json:"perk3Var3"`
	Perk3Var2                       int  `json:"perk3Var2"`
	Perk3Var1                       int  `json:"perk3Var1"`
	DamageSelfMitigated             int  `json:"damageSelfMitigated"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	Perk0Var2                       int  `json:"perk0Var2"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	Assists                         int  `json:"assists"`
	Perk4Var2                       int  `json:"perk4Var2"`
	GoldSpent                       int  `json:"goldSpent"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	ParticipantID                   int  `json:"participantId"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	Win                             bool `json:"win"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	Deaths                          int  `json:"deaths"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
	PerkSubStyle                    int  `json:"perkSubStyle"`
	TurretKills                     int  `json:"turretKills"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	GoldEarned                      int  `json:"goldEarned"`
	KillingSprees                   int  `json:"killingSprees"`
	UnrealKills                     int  `json:"unrealKills"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	ChampLevel                      int  `json:"champLevel"`
	DoubleKills                     int  `json:"doubleKills"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	Perk0Var1                       int  `json:"perk0Var1"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	Perk0Var3                       int  `json:"perk0Var3"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	PentaKills                      int  `json:"pentaKills"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
	StatPerk2                       int  `json:"statPerk2"`
}

type ParticipantDTO struct {
	Spell1ID      int                    `json:"spell1Id"`
	ParticipantID int                    `json:"participantId"`
	Timeline      ParticipantTimelineDTO `json:"timeline,omitempty"`
	Spell2ID      int                    `json:"spell2Id"`
	TeamID        int                    `json:"teamId"`
	Stats         ParticipantStatsDTO    `json:"stats"`
	ChampionID    int                    `json:"championId"`
	Masteries     []MasteryDTO           `json:"masteries"`
}

type MasteryDTO struct {
	MasteryID int
	Rank      int
}

type MatchDTO struct {
	SeasonID              int                      `json:"seasonId"`
	QueueID               int                      `json:"queueId"`
	GameID                int64                    `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"participantIdentities"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
	GameMode              string                   `json:"gameMode"`
	MapID                 int                      `json:"mapId"`
	GameType              string                   `json:"gameType"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	Participants          []ParticipantDTO         `json:"participants"`
	GameDuration          int                      `json:"gameDuration"`
	GameCreation          int64                    `json:"gameCreation"`
}

type MatchTimelineDTO struct {
	Frames        []MatchFrameDTO `json:"frames"`
	FrameInterval int
}

type MatchFrameDTO struct {
	Timestamp         int                           `json:"timestamp"`
	ParticipantFrames map[string]MatchParicipantDTO `json:"participantFrames"`
	Events            []MatchEventDTO               `json:"events"`
}

type MatchParicipantDTO struct {
	TotalGold           int              `json:"totalGold"`
	TeamScore           int              `json:"teamScore"`
	ParticipantID       int              `json:"participantId"`
	Level               int              `json:"level"`
	CurrentGold         int              `json:"currentGold"`
	MinionsKilled       int              `json:"minionsKilled"`
	DominionScore       int              `json:"dominionScore"`
	Position            MatchPositionDTO `json:"position"`
	XP                  int              `json:"xp"`
	JungleMinionsKilled int              `json:"jungleMinionsKilled"`
}

type MatchPositionDTO struct {
	Y int `json:"y"`
	X int `json:"x"`
}

type MatchEventDTO struct {
	EventType               string           `json:"eventType"`
	TowerType               string           `json:"towerType"`
	TeamID                  int              `json:"teamId"`
	AscendedType            string           `json:"ascendedType"`
	KillerID                int              `json:"killerId"`
	LevelUpType             string           `json:"levelUpType"`
	PointCaptured           string           `json:"pointCaptured"`
	AssistingParticipantIDs []int            `json:"assistingParticipantIDs"`
	WardType                string           `json:"wardType"`
	MonsterType             string           `json:"monsterType"`
	Type                    string           `json:"type"`
	SkillSlot               int              `json:"skillSlot"`
	VictimID                int              `json:"victimId"`
	Timestamp               int64            `json:"timestamp"`
	AfterID                 int              `json:"afterId"`
	MonsterSubType          string           `json:"monsterSubType"`
	LaneType                string           `json:"laneType"`
	ItemID                  int              `json:"itemId"`
	ParticipantID           int              `json:"participantId"`
	BuildingType            string           `json:"buildingType"`
	CreatorID               int              `json:"creatorId"`
	Position                MatchPositionDTO `json:"position"`
	BeforeID                int              `json:"beforeId"`
}

type MatchlistDTO struct {
	Matches    []MatchReferenceDTO `json:"matches"`
	TotalGames int                 `json:"totalGames"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
}

type MatchlistsParams struct {
	Champion   []int `url:"champion"`
	Queue      []int `url:"queue"`
	Season     []int `url:"season"`
	EndTime    int   `url:"endTime"`
	BeginTime  int   `url:"beginTime"`
	EndIndex   int   `url:"endIndex"`
	BeginIndex int   `url:"beginIndex"`
}

type MatchReferenceDTO struct {
	Lane       string `json:"lane"`
	GameID     int    `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Timestamp  int    `json:"timestamp"`
}

type Message struct {
	Severity     string        `json:"severity"`
	Author       string        `json:"author"`
	CreatedAt    string        `json:"created_at"`
	Translations []Translation `json:"translations"`
	UpdatedAt    string        `json:"updated_at"`
	Content      string        `json:"content"`
	ID           string        `json:"id"`
}

type MiniSeriesDTO struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type Service struct {
	Status    string     `json:"status"`
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
}

type ShardStatus struct {
	Name      string    `json:"name"`
	RegionTag string    `json:"region_tag"`
	Hostname  string    `json:"hostname"`
	Services  []Service `json:"services"`
	Slug      string    `json:"slug"`
	Locales   []string  `json:"locales"`
}

type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

type Translation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

// ClientOption is a func that operates on *Client
type ClientOption func(*Client) error

// NewClient returns interface to League of Legends API
func NewClient(token string, options ...ClientOption) (*Client, error) {
	cli := &Client{}
	WithToken(token)(cli)
	WithRegion(defaultRegion)(cli)
	cli.sling = sling.New().Base("https://" + cli.Region + "." + baseURL)
	cli.sling.Set("User-Agent", "jonwho/lol")

	for _, option := range options {
		if err := option(cli); err != nil {
			return nil, err
		}
	}

	cli.sling.Set("X-Riot-Token", cli.Token)

	return cli, nil
}

// WithToken set the client token
func WithToken(token string) ClientOption {
	return func(c *Client) error {
		c.Token = token
		return nil
	}
}

// WithRegion set the client region
func WithRegion(region string) ClientOption {
	return func(c *Client) error {
		c.Region = region
		return nil
	}
}

// WithHTTPClient set the client http.Client and the sling http.Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		c.sling.Client(httpClient)
		return nil
	}
}

// AllChampionMastery GET /lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerID}
func (c *Client) AllChampionMastery(encryptedSummonerID string) (*[]ChampionMasteryDTO, *http.Response, error) {
	dtos := new([]ChampionMasteryDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/champion-mastery/v4/champion-masteries/by-summoner/"+encryptedSummonerID).Receive(dtos, reqErr)

	if err != nil {
		return nil, resp, err
	}

	return dtos, resp, reqErr
}

// ChampionMastery GET /lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerID}/by-champion/{championID}
func (c *Client) ChampionMastery(encryptedSummonerID, championID string) (*ChampionMasteryDTO, *http.Response, error) {
	dto := new(ChampionMasteryDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/champion-mastery/v4/champion-masteries/by-summoner/"+encryptedSummonerID+"/by-champion/"+championID).Receive(dto, reqErr)

	if err != nil {
		return nil, resp, err
	}

	return dto, resp, reqErr
}

// MasteryScore GET /lol/champion-mastery/v4/scores/by-summoner/{encryptedSummonerID}
func (c *Client) MasteryScore(encryptedSummonerID string) (int, *http.Response, error) {
	req, err := c.sling.Get("lol/champion-mastery/v4/scores/by-summoner/" + encryptedSummonerID).Request()
	if err != nil {
		return 0, nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, resp, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, resp, err
	}
	score, err := strconv.Atoi(string(bodyBytes))
	if err != nil {
		return 0, resp, err
	}
	return score, resp, err
}

// ChampionRotations GET /lol/platform/v3/champion-rotations
func (c *Client) ChampionRotations() (*ChampionInfo, *http.Response, error) {
	ci := new(ChampionInfo)
	var reqErr error
	resp, err := c.sling.Get("lol/platform/v3/champion-rotations").Receive(ci, reqErr)

	if err != nil {
		return nil, resp, err
	}

	return ci, resp, reqErr
}

// LeagueExpEntries GET /lol/league-exp/v4/entries/{queue}/{tier}/{division}
func (c *Client) LeagueExpEntries(queue, tier, division string, params *LeagueExpEntriesParams) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	endpoint := fmt.Sprintf("lol/league-exp/v4/entries/%s/%s/%s", queue, tier, division)
	resp, err := c.sling.Get(endpoint).QueryStruct(params).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// ChallengerLeagues GET /lol/league/v4/challengerleagues/by-queue/{queue}
func (c *Client) ChallengerLeagues(queue string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/challengerleagues/by-queue/"+queue).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// EntriesBySummoner GET /lol/league/v4/entries/by-summoner/{encryptedSummonerId}
func (c *Client) EntriesBySummoner(encryptedSummonerID string) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/entries/by-summoner/"+encryptedSummonerID).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// Entries GET /lol/league/v4/entries/{queue}/{tier}/{division}
func (c *Client) Entries(queue, tier, division string, params *EntriesParams) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	endpoint := fmt.Sprintf("lol/league/v4/entries/%s/%s/%s", queue, tier, division)
	resp, err := c.sling.Get(endpoint).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// GrandmasterLeagues GET /lol/league/v4/grandmasterleagues/by-queue/{queue}
func (c *Client) GrandmasterLeagues(queue string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/grandmasterleagues/by-queue/"+queue).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Leagues GET /lol/league/v4/leagues/{leagueId}
func (c *Client) Leagues(leagueID string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/leagues/"+leagueID).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// MasterLeagues GET /lol/league/v4/masterleagues/by-queue/{queue}
func (c *Client) MasterLeagues(queue string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/masterleagues/by-queue/"+queue).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Status GET /lol/status/v3/shard-data
func (c *Client) Status() (*ShardStatus, *http.Response, error) {
	shardStatus := new(ShardStatus)
	var reqErr error
	resp, err := c.sling.Get("lol/status/v3/shard-data").Receive(shardStatus, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return shardStatus, resp, reqErr
}

// Matches GET /lol/match/v4/matches/{matchID}
func (c *Client) Matches(matchID string) (*MatchDTO, *http.Response, error) {
	dto := new(MatchDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/match/v4/matches/"+matchID).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Matchlists GET /lol/match/v4/matchlists/by-account/{encryptedAccountID}
func (c *Client) Matchlists(encryptedAccountID string, params *MatchlistsParams) (*MatchlistDTO, *http.Response, error) {
	dto := new(MatchlistDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/match/v4/matchlists/by-account/"+encryptedAccountID).QueryStruct(params).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Timelines GET /lol/match/v4/timelines/by-match/{matchID}
func (c *Client) Timelines(matchID string) (*MatchTimelineDTO, *http.Response, error) {
	dto := new(MatchTimelineDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/match/v4/timelines/by-match/"+matchID).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// ActiveGames GET /lol/spectator/v4/active-games/by-summoner/{encryptedSummonerId}
func (c *Client) ActiveGames(encryptedSummonerID string) (*CurrentGameInfo, *http.Response, error) {
	info := new(CurrentGameInfo)
	var reqErr error
	resp, err := c.sling.Get("lol/spectator/v4/active-games/by-summoner/"+encryptedSummonerID).Receive(info, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return info, resp, reqErr
}

// FeaturedGames GET /lol/spectator/v4/featured-games
func (c *Client) FeaturedGames() (*FeaturedGames, *http.Response, error) {
	info := new(FeaturedGames)
	var reqErr error
	resp, err := c.sling.Get("lol/spectator/v4/featured-games").Receive(info, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return info, resp, reqErr
}

// SummonerByAccount GET /lol/summoner/v4/summoners/by-account/{encryptedAccountID}
func (c *Client) SummonerByAccount(encryptedAccountID string) (*SummonerDTO, *http.Response, error) {
	sd := new(SummonerDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/summoner/v4/summoners/by-account/"+encryptedAccountID).Receive(sd, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return sd, resp, reqErr
}

// SummonerByName GET /lol/summoner/v4/summoners/by-name/{summonerName}
func (c *Client) SummonerByName(summonerName string) (*SummonerDTO, *http.Response, error) {
	sd := new(SummonerDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/summoner/v4/summoners/by-name/"+summonerName).Receive(sd, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return sd, resp, reqErr
}

// SummonerByPUUID GET /lol/summoner/v4/summoners/by-puuid/{encryptedPUUID}
func (c *Client) SummonerByPUUID(encryptedPUUID string) (*SummonerDTO, *http.Response, error) {
	sd := new(SummonerDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/summoner/v4/summoners/by-puuid/"+encryptedPUUID).Receive(sd, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return sd, resp, reqErr
}

// SummonerByID GET /lol/summoner/v4/summoners/{encryptedID}
func (c *Client) SummonerByID(encryptedID string) (*SummonerDTO, *http.Response, error) {
	sd := new(SummonerDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/summoner/v4/summoners/"+encryptedID).Receive(sd, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return sd, resp, reqErr
}
