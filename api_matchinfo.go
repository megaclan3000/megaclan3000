package main

import (
	// "github.com/gorilla/mux"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// func handlerAPI(w http.ResponseWriter, r *http.Request) {
// 	byt := apiHandler(mux.Vars(r))
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(byt)
// }

func handlerAPIMatchinfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("API request to:", r.RequestURI)

	var byt []byte

	vars := mux.Vars(r)
	switch vars["endpoint"] {

	case "scoreboard":
		byt = []byte(`
{
    "clan": [
        {
            "name": "Player1",
            "clantag": "megaclan3000",
            "avatar_url": "/public/TODO",
            "rank": "7",

            "steamid64": "TODO",
            "kills": "TODO",
            "deaths": "TODO",
            "assists": "TODO",
            "kddiff": "TODO",
            "kd": "TODO",
            "adr": "TODO",
            "hsprecent": "TODO",
            "firstkills": "TODO",
            "firstdeaths": "TODO",
            "tradekills": "TODO",
            "tradedeaths": "TODO",
            "tradefirstkills": "TODO",
            "tradefirstdeaths": "TODO",
            "roundswonv5": "TODO",
            "roundswonv4": "TODO",
            "roundswonv3": "TODO",
            "rounds5k": "TODO",
            "rounds4k": "TODO",
            "rounds3k": "TODO"
        },
        {
            "name": "Player2",
            "clantag": "megaclan3000",
            "avatar_url": "/public/TODO",
            "rank": "7",

            "steamid64": "TODO",
            "kills": "TODO",
            "deaths": "TODO",
            "assists": "TODO",
            "kddiff": "TODO",
            "kd": "TODO",
            "adr": "TODO",
            "hsprecent": "TODO",
            "firstkills": "TODO",
            "firstdeaths": "TODO",
            "tradekills": "TODO",
            "tradedeaths": "TODO",
            "tradefirstkills": "TODO",
            "tradefirstdeaths": "TODO",
            "roundswonv5": "TODO",
            "roundswonv4": "TODO",
            "roundswonv3": "TODO",
            "rounds5k": "TODO",
            "rounds4k": "TODO",
            "rounds3k": "TODO"
        }
    ],
    "enemy": [
        {
            "name": "Player3",
            "clantag": "enemyclan",
            "avatar_url": "/public/TODO",
            "rank": "7",

            "steamid64": "TODO",
            "kills": "TODO",
            "deaths": "TODO",
            "assists": "TODO",
            "kddiff": "TODO",
            "kd": "TODO",
            "adr": "TODO",
            "hsprecent": "TODO",
            "firstkills": "TODO",
            "firstdeaths": "TODO",
            "tradekills": "TODO",
            "tradedeaths": "TODO",
            "tradefirstkills": "TODO",
            "tradefirstdeaths": "TODO",
            "roundswonv5": "TODO",
            "roundswonv4": "TODO",
            "roundswonv3": "TODO",
            "rounds5k": "TODO",
            "rounds4k": "TODO",
            "rounds3k": "TODO"
        },
        {
            "name": "Player4",
            "clantag": "enemyclan",
            "avatar_url": "/public/TODO",
            "rank": "7",

            "steamid64": "TODO",
            "kills": "TODO",
            "deaths": "TODO",
            "assists": "TODO",
            "kddiff": "TODO",
            "kd": "TODO",
            "adr": "TODO",
            "hsprecent": "TODO",
            "firstkills": "TODO",
            "firstdeaths": "TODO",
            "tradekills": "TODO",
            "tradedeaths": "TODO",
            "tradefirstkills": "TODO",
            "tradefirstdeaths": "TODO",
            "roundswonv5": "TODO",
            "roundswonv4": "TODO",
            "roundswonv3": "TODO",
            "rounds5k": "TODO",
            "rounds4k": "TODO",
            "rounds3k": "TODO"
        }
    ]
}`)
		// TODO
	case "rounds":
		byt = []byte(`
"rounds": [
	{
		"score_clan": 2,
		"score_enemy": 1,
		"win_reason": 1,
		"total_damage_taken": 100,
		"total_damage_given": 200,
		"winner_team": "CT",
		"kills_clan": [
			{
				"killer": "Player1",
				"victim": "Player2",
				"weapon": "USP"
			},
			{
				"killer": "Player1",
				"victim": "Player3",
				"weapon": "AK-47"
			}
		],
		"kills_enemy": [
			{
				"killer": "Player4",
				"victim": "Player1",
				"weapon": "Glock-18"
			},
			{
				"killer": "Player4",
				"victim": "Player5",
				"weapon": "M4A4"
			}
		],
		"duration": "1:20"
	},
	{
		"score_clan": 3,
		"score_enemy": 1,
		"win_reason": 2,
		"total_damage_taken": 100,
		"total_damage_given": 200,
		"winner_team": "T",
		"kills_clan": [
			{
				"killer": "Player1",
				"victim": "Player2",
				"weapon": "USP"
			},
			{
				"killer": "Player1",
				"victim": "Player3",
				"weapon": "AK-47"
			}
		],
		"kills_enemy": [],
		"duration": "1:30"
	}
]

`)
	// TODO
	case "weapons":
		byt = []byte(`
[
    {
        "name": "AWP",
        "kills": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        },
        "headshots": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        },
        "accuracy": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        },
        "damage": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        }
    },
    {
        "name": "AK-47",
        "kills": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        },
        "headshots": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        },
        "accuracy": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        },
        "damage": {
            "clan": [
                {
                    "name": "Player1",
                    "amount": 9
                },
                {
                    "name": "Player2",
                    "amount": 4
                }
            ],
            "enemy": [
                {
                    "name": "Player5",
                    "amount": 2
                },
                {
                    "name": "Player6",
                    "amount": 1
                }
            ]
        }
    }
]`)
	// TODO
	case "duels":
		byt = []byte(`
[
	"PlayerClan1": [
		{
			"PlayerEnemy1": "40",
			"PlayerEnemy2": "41",
			"PlayerEnemy3": "42",
			"PlayerEnemy4": "43",
			"PlayerEnemy5": "44"
		},
	"PlayerClan2": [
		{
			"PlayerEnemy1": "40",
			"PlayerEnemy2": "41",
			"PlayerEnemy3": "42",
			"PlayerEnemy4": "43",
			"PlayerEnemy5": "44"
		},
	"PlayerClan3": [
		{
			"PlayerEnemy1": "40",
			"PlayerEnemy2": "41",
			"PlayerEnemy3": "42",
			"PlayerEnemy4": "43",
			"PlayerEnemy5": "44"
		},
	"PlayerClan4": [
		{
			"PlayerEnemy1": "40",
			"PlayerEnemy2": "41",
			"PlayerEnemy3": "42",
			"PlayerEnemy4": "43",
			"PlayerEnemy5": "44"
		},
	"PlayerClan5": [
		{
			"PlayerEnemy1": "40",
			"PlayerEnemy2": "41",
			"PlayerEnemy3": "42",
			"PlayerEnemy4": "43",
			"PlayerEnemy5": "44"
		}
	]
`)
		//TODO
	case "heatmaps":
		//TODO
	case "megacoins":
		//TODO
	default:
		byt = []byte(`
[
	["PlayerA1", "PlayerB1", 5],
	["PlayerA1", "PlayerB2", 5],
	["PlayerA1", "PlayerB3", 5],
	["PlayerA1", "PlayerB4", 5],
	["PlayerA1", "PlayerB5", 5],
	["PlayerA2", "PlayerB1", 5],
	["PlayerA2", "PlayerB2", 5],
	["PlayerA2", "PlayerB3", 5],
	["PlayerA2", "PlayerB4", 5],
	["PlayerA2", "PlayerB5", 5],
	["PlayerA3", "PlayerB1", 5],
	["PlayerA3", "PlayerB2", 5],
	["PlayerA3", "PlayerB3", 5],
	["PlayerA3", "PlayerB4", 5],
	["PlayerA3", "PlayerB5", 5],
	["PlayerA4", "PlayerB1", 5],
	["PlayerA4", "PlayerB2", 5],
	["PlayerA4", "PlayerB3", 5],
	["PlayerA4", "PlayerB4", 5],
	["PlayerA4", "PlayerB5", 5],
	["PlayerA5", "PlayerB1", 5],
	["PlayerA5", "PlayerB2", 5],
	["PlayerA5", "PlayerB3", 5],
	["PlayerA5", "PlayerB4", 5],
	["PlayerA5", "PlayerB5", 5]
]`)

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
