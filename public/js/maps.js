var ctx = document.getElementById('mapsChart');
var myChart = new Chart(ctx, {
    type: 'pie',
    data: {
        labels: [
            "Baggage",
            "Bank",
            "CBBL",
            "Dust",
            "Dust 2",
            "Inferno",
            "Italy",
            "Lake",
            "Nuke",
            "Safehouse",
            "St. Marc",
            "Sugarcane",
            "Train",
            "Vertigo"
        ],
        datasets: [{
            label: 'Maps',
            data: [
                {{ index .UserStatsForGame.Stats "total_wins_map_ar_baggage"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_bank"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_cbble"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_dust"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_dust2"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_inferno"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_cs_italy"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_lake"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_nuke"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_safehouse"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_stmarc"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_sugarcane"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_train"}},
                {{ index .UserStatsForGame.Stats "total_wins_map_de_vertigo"}},
            ],
            backgroundColor: [
                'rgba(255, 99, 132, 0.9)',
                'rgba(255, 99, 132, 0.85)',
                'rgba(255, 99, 132, 0.8)',
                'rgba(255, 99, 132, 0.75)',
                'rgba(255, 99, 132, 0.7)',
                'rgba(255, 99, 132, 0.65)',
                'rgba(255, 99, 132, 0.6)',
                'rgba(255, 99, 132, 0.55)',
                'rgba(255, 99, 132, 0.5)',
                'rgba(255, 99, 132, 0.45)',
                'rgba(255, 99, 132, 0.4)',
                'rgba(255, 99, 132, 0.35)',
                'rgba(255, 99, 132, 0.3)',
                'rgba(255, 99, 132, 0.25)',      
            ],
            borderColor: [
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',
                'rgba(255, 99, 132, 1)',       
            ],
            borderWidth: 1
        }]
    },
    options: {
        legend: {
            labels: {
                // This more specific font property overrides the global property
                fontColor: 'white'
                
            },
            position: 'bottom',
            align: 'start'
        }
    }
});