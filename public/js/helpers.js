function CreateTableFromJSON(tableid, jsonurl) {

	var table = document.getElementById(tableid);
		var tbody = document.createElement("tbody");

		$.getJSON(jsonurl,  function( data) {
		for (var row = 0; row < data.length; row++) {

			var tr = document.createElement("tr");

			Object.keys(data[row]).forEach(function(col) {
				var td = document.createElement("td");
				td.innerHTML = data[row][col];
				tr.appendChild(td);

			})
			tbody.appendChild(tr);
		}
	});
	table.appendChild(tbody);
}

// Create table, optionally with head
//
// <table id="mytable">
// 	<thead>
// 		<tr>
// 			<th>Name</th>
// 			<th>Item Name</th>
// 			<th>Item Price</th>
// 			<th>Item Price</th>
// 		</tr>
// 	</thead>
// </table>
//
// Pass url and tableid
//
// url = "http://localhost:9090/test.json"
// CreateTableFromJSON("mytable", url);
