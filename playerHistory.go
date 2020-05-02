package main

// PlayerHistory holds the players history data from the player_history table.
// Stats values that need to be saved over time, are added to this table and
// object.
type PlayerHistory struct {
	//TODO implement
	steamID    int
	time       int
	TotalKills int
}
