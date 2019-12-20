
function loadFormattedEmployeeTable() {
    var tableRef = document.getElementById('formatted_employee_table').getElementsByTagName('tbody')[0];
    var values = httpGet("http://localhost:8080/listAllFormattedEmployees");
    var rows = JSON.parse(values);

    for (var i = 0; i < rows.length; i++) {
        var newRow = tableRef.insertRow();

        for (j = 0; j < rows[i].length; j++) {
            var newCell = newRow.insertCell(j);

            var newText  = document.createTextNode(rows[i][j]);
            newCell.appendChild(newText);
        }
    }
    
}

function loadEmployeeTable() {
    var tableRef = document.getElementById('employee_table').getElementsByTagName('tbody')[0];
    var values = httpGet("http://localhost:8080/employees");
    var rows = JSON.parse(values);

    for (var i = 0; i < rows.length; i++) {
        var newRow = tableRef.insertRow();

        for (j = 0; j < rows[i].length; j++) {
            var newCell = newRow.insertCell(j);

            var newText  = document.createTextNode(rows[i][j]);
            newCell.appendChild(newText);
        }
    }
}

function httpGet(url) {
    var xmlHttp = new XMLHttpRequest();

    xmlHttp.open( "GET", url, false );
    xmlHttp.send( null );
    
    return xmlHttp.responseText;
}

function createEmployee() {
    name = document.getElementById('new_employee_name').value;
    id = document.getElementById('new_employee_id').value;
    manager_id = document.getElementById('new_employee_manager_id').value;
    

    if (name=="" || id=="" || manager_id == "") {
        alert("Please fill in the employee values !");
    
        return;
    }

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/createEmployee", false);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        "name": name,
        "id": parseInt(id, 10),
        "manager_id": parseInt(manager_id, 10)
    }));

    alert(xhr.response);
    refreshTables();
}

function refreshTables() {
    var Table1 = document.getElementById("employee_table");
    Table1.innerHTML = "<th>Name</th><th>ID</th><th>ManagerID</th><tbody></tbody>";

    var Table2 = document.getElementById("formatted_employee_table");
    Table2.innerHTML = "<tbody></tbody>";

    loadFormattedEmployeeTable();
    loadEmployeeTable();
}