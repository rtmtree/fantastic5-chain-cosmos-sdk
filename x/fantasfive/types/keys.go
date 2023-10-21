package types

const (
	// ModuleName defines the module name
	ModuleName = "fantasfive"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fantasfive"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MwInfoKey = "MwInfo-value-"
)

const (
	TeamInfoKey = "TeamInfo-value-"
)

const (
	TeamCreatedEventType         = "new-team-created"
	TeamCreatedEventCreator      = "creator"
	TeamCreatedEventTeamId       = "team-id"
	TeamCreatedEventMwId         = "mw-id"
	TeamCreatedEventPlayer0      = "player-0"
	TeamCreatedEventPlayer1      = "player-1"
	TeamCreatedEventPlayer2      = "player-2"
	TeamCreatedEventPlayer3      = "player-3"
	TeamCreatedEventPlayer4      = "player-4"
	TeamCreatedEventCaptainIndex = "captain-index"
)

const (
	MatchWeekAnnouncedEventType         = "match-week-announced"
	MatchWeekAnnouncedEventMwId         = "mw-id"
	MatchWeekAnnouncedEventAnnoucer     = "announcer"
	MatchWeekAnnouncedEventTeamCount    = "team-count"
	MatchWeekAnnouncedEventWinnerTeamId = "winner-team-id"
	MatchWeekAnnouncedEventPlayerPref   = "players-performance"
)

const (
	MatchWeekCreatedEventType    = "match-week-created"
	MatchWeekCreatedEventCreator = "creator"
	MatchWeekCreatedEventMwId    = "mw-id"
)
