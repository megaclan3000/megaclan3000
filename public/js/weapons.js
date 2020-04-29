var ctx = document.getElementById('weaponsChart');
var myChart = new Chart(ctx, {
    type: 'pie',
    data: {
        labels: [
            // --------  Close Range  --------  
            'Knife', 
            // 'Zeus x27',
            
            // --------  Pistoles  --------  
            // 'CZ75-Auto',
            'Desert Eagle',
            // 'Dual Berettas',
            'Five Seven',
            'Glock', 
            'P250',
            'P2000',
            // 'R8 Revolver',
            'Tec-9', 
            // 'USP-S',

            // --------  SMG  --------  
            'MAC-10',
            // 'MP5-SD',
            'MP7',
            'MP9',
            // 'M4A4',
            'P90',
            'PP-Bizon', 
            // 'SSG 08',
            'UMP-45',

            // --------  Rifles  --------
            'AK 47',
            'AUG', 
            'AWP',
            'Famas',
            'Galil AR',
            'G3SG1',
            'M4A1-S',
            'SG 556',
            'SCAR-20', 

            // --------  Heavy  --------
            'MAG-7',
            'M249',
            'Negev',
            'Nova',
            'Sawed off',
            'XM1014',

            // --------  Grenades  --------
            'HE Grenade',
            'Molotov'
        ],
        datasets: [{
            label: '# of kills',
            data: [
                // --------  Close Range  --------  
                {{ index .UserStatsForGame.Stats "total_kills_knife"}},
                // 'Zeus x27',

                // --------  Pistoles  --------  
                // 'CZ75-Auto',
                {{ index .UserStatsForGame.Stats "total_kills_deagle"}},
                // 'Dual Berettas',
                {{ index .UserStatsForGame.Stats "total_kills_fiveseven"}},
                {{ index .UserStatsForGame.Stats "total_kills_glock"}},
                {{ index .UserStatsForGame.Stats "total_kills_p250"}},
                {{ index .UserStatsForGame.Stats "total_kills_hkp2000"}},
                // 'R8 Revolver',
                {{ index .UserStatsForGame.Stats "total_kills_tec9"}},
                // 'USP-S',

                // --------  SMG  --------  
                {{ index .UserStatsForGame.Stats "total_kills_mac10"}},
                 // 'MP5-SD',
                {{ index .UserStatsForGame.Stats "total_kills_mp7"}},
                {{ index .UserStatsForGame.Stats "total_kills_mp9"}},
                // 'M4A4',
                {{ index .UserStatsForGame.Stats "total_kills_p90"}},
                {{ index .UserStatsForGame.Stats "total_kills_bizon"}},
                // 'SSG 08',
                {{ index .UserStatsForGame.Stats "total_kills_ump45"}},

                // --------  Rifles  --------
                {{ index .UserStatsForGame.Stats "total_kills_ak47"}},
                {{ index .UserStatsForGame.Stats "total_kills_aug"}},
                {{ index .UserStatsForGame.Stats "total_kills_awp"}},
                {{ index .UserStatsForGame.Stats "total_kills_famas"}},
                {{ index .UserStatsForGame.Stats "total_kills_galilar"}},
                {{ index .UserStatsForGame.Stats "total_kills_g3sg1"}},
                {{ index .UserStatsForGame.Stats "total_kills_m4a1"}},
                {{ index .UserStatsForGame.Stats "total_kills_sg556"}},
                {{ index .UserStatsForGame.Stats "total_kills_scar20"}}, 
                
                // --------  Heavy  --------
                {{ index .UserStatsForGame.Stats "total_kills_mag7"}},
                {{ index .UserStatsForGame.Stats "total_kills_m249"}},
                {{ index .UserStatsForGame.Stats "total_kills_negev"}},
                {{ index .UserStatsForGame.Stats "total_kills_nova"}},
                {{ index .UserStatsForGame.Stats "total_kills_sawedoff"}},
                {{ index .UserStatsForGame.Stats "total_kills_xm1014"}},  
                
                // --------  Grenades  --------
                {{ index .UserStatsForGame.Stats "total_kills_hegrenade"}},
                {{ index .UserStatsForGame.Stats "total_kills_molotov"}}
            ],
            backgroundColor: [
                // --------  Close Range  --------
                'rgba(255, 99, 132, 0.9)',
                // 'rgba(255, 99, 132, 0.8)',

                // --------  Pistoles  -------- 
                // 'rgba(54, 162, 235, 0.9)', 
                'rgba(54, 162, 235, 0.8)', 
                // 'rgba(54, 162, 235, 0.7)', 
                'rgba(54, 162, 235, 0.6)', 
                'rgba(54, 162, 235, 0.5)', 
                'rgba(54, 162, 235, 0.4)', 
                'rgba(54, 162, 235, 0.3)', 
                // 'rgba(54, 162, 235, 0.2)', 
                'rgba(54, 162, 235, 0.1)', 
                // 'rgba(54, 162, 235, 0.8)', 

                // --------  SMG  --------  
                'rgba(255, 206, 86, 0.9)',
                // 'rgba(255, 206, 86, 0.8)',
                'rgba(255, 206, 86, 0.7)',
                'rgba(255, 206, 86, 0.6)',
                // 'rgba(255, 206, 86, 0.5)',
                'rgba(255, 206, 86, 0.4)',
                'rgba(255, 206, 86, 0.3)',
                // 'rgba(255, 206, 86, 0.2)',
                'rgba(255, 206, 86, 0.1)',

                // --------  Rifles  --------
                'rgba(75, 192, 192, 0.9)',
                'rgba(75, 192, 192, 0.8)',
                'rgba(75, 192, 192, 0.7)',
                'rgba(75, 192, 192, 0.6)',
                'rgba(75, 192, 192, 0.5)',
                'rgba(75, 192, 192, 0.4)',
                'rgba(75, 192, 192, 0.3)',
                'rgba(75, 192, 192, 0.2)',
                'rgba(75, 192, 192, 0.1)',

                // --------  Heavy  --------
                'rgba(153, 102, 255, 0.9)',
                'rgba(153, 102, 255, 0.8)',
                'rgba(153, 102, 255, 0.7)',
                'rgba(153, 102, 255, 0.6)',
                'rgba(153, 102, 255, 0.5)',
                'rgba(153, 102, 255, 0.4)',

                // --------  Grenades  --------
                'rgba(255, 159, 64, 0.9)', 
                'rgba(255, 159, 64, 0.8)' 
            ],
            borderColor: [
                // --------  Close Range  --------
                'rgba(255, 99, 132, 1)',
                // 'rgba(255, 99, 132, 1)',

                // --------  Pistoles  -------- 
                // 'rgba(54, 162, 235, 1)', 
                'rgba(54, 162, 235, 1)', 
                // 'rgba(54, 162, 235, 1)', 
                'rgba(54, 162, 235, 1)', 
                'rgba(54, 162, 235, 1)', 
                'rgba(54, 162, 235, 1)', 
                'rgba(54, 162, 235, 1)', 
                // 'rgba(54, 162, 235, 1)', 
                'rgba(54, 162, 235, 1)', 
                // 'rgba(54, 162, 235, 1)', 

                // --------  SMG  --------  
                'rgba(255, 206, 86, 1)',
                // 'rgba(255, 206, 86, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(255, 206, 86, 1)',
                // 'rgba(255, 206, 86, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(255, 206, 86, 1)',
                // 'rgba(255, 206, 86, 1)',
                'rgba(255, 206, 86, 1)',

                // --------  Rifles  --------
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(75, 192, 192, 1)',

                // --------  Heavy  --------
                'rgba(153, 102, 255, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(153, 102, 255, 1)',

                // --------  Grenades  --------
                'rgba(255, 159, 64, 1)', 
                'rgba(255, 159, 64, 1)' 
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