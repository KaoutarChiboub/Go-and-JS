const userICCID = document.getElementById("user-iccid")
const userIMSI = document.getElementById("user-imsi")
const userLAI = document.getElementById("user-lai")
const userK = document.getElementById("user-k")
const userOEN = document.getElementById("user-oen")
const userLoginButton = document.getElementById("user-login-button")
const userFetch = document.getElementById("fetch-users")
const userTable = document.getElementById("users-table")


userLoginButton.addEventListener("click", () => {
    console.log("Login in process")
    verifyUser();
})

userFetch.addEventListener("click", () => {
    console.log("Listing all logged-in users...")
    listUsers();
})

// userTable.add

function verifyUser() {
    let data = {
        "ICCID": userICCID.value,
        "IMSI": userIMSI.value,
        "LAI": userLAI.value,
        "K": userK.value,
        "OEN": userOEN.value

    };
    fetch("http://localhost:8008/sim-login", {
        method: "POST",
        body: JSON.stringify(data), 
    }).then((response) => {
        // console.log(response.json())
    }).catch((error) => {
        console.log(error)
    });
}

function listUsers() {
    fetch("http://localhost:8008/simlist", {
        method:"GET",
    }).then((response)=> response.json())
    .then((data) => {
        console.log("Logged-in users: " + JSON.stringify(data));
        showTitle();
        renderTable(data);
    }).catch((error) => {
        console.log(error);
    });
}

function renderTable(data) {
    // Clear previous table data
     userTable.innerHTML = "";

    // Create table header
    const tableHeader = document.createElement("thead");
    const headerRow = document.createElement("tr");
    const headers = ["ID", "ICCID", "IMSI", "LAI", "K", "OEN"];
    headers.forEach((headerText) => {
      const headerCell = document.createElement("th");
      headerCell.textContent = headerText;
      headerRow.appendChild(headerCell);
    });
    tableHeader.appendChild(headerRow);
    userTable.appendChild(tableHeader);
  
    // Create table body
    const tableBody = document.createElement("tbody");
    data.forEach((user) => {
      const row = document.createElement("tr");
      Object.values(user).forEach((value) => {
        const cell = document.createElement("td");
        cell.textContent = value;
        row.appendChild(cell);
      });
      tableBody.appendChild(row);
    });
    userTable.appendChild(tableBody);
}

function showTitle(){
    // Create table title
    const caption = document.getElementById("table-caption");
    caption.style.display = "block";
}