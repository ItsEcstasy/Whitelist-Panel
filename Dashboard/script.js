/**
 * yeah ik this is aids but im not the best at js soo.. cut me some slack :sob:
 *                    i basically put everythig in here
 */

const users = [
    { username: 'Ecstasy', planType: 'Lifetime', country: 'USA', roles: ['Owner', 'Developer'] },
    { username: 'S409', planType: 'Premium', country: 'UK', roles: ['User'] },
    { username: 'Juice WRLD', planType: 'Basic', country: 'Chicago', roles: ['User'] },
    { username: 'Bakalaka', planType: 'VIP', country: 'Canada', roles: ['User'] },
];

window.onload = function() {
    ShowUsers();
}

function ShowUsers() {
    const tableBody = document.querySelector('#userTable tbody');
    tableBody.innerHTML = '';
    
    users.forEach(user => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${user.username}</td>
            <td>${user.planType}</td>
            <td>${user.country}</td>
            <td>${user.roles.join(', ')}</td>
            <td>
                <button onclick="manageUser('${user.username}', 'blacklist')">Blacklist</button>
                <button onclick="manageUser('${user.username}', 'edit')">Edit</button>
            </td>
        `;
        tableBody.appendChild(row);
    });
}

function showTabs(tabId) {
    const tabs = document.getElementsByClassName('tab');
    for (let i = 0; i < tabs.length; i++) {
        tabs[i].style.display = 'none';
    }
    document.getElementById(tabId).style.display = 'block';
}


// Implementing soon
function LookupKey() {
    const keyInput = document.getElementById('keyInput');
    const key = keyInput.value;
    
    const keyInfo = ValidateKey(key);

    const keyInfoDiv = document.getElementById('keyInfo');
    keyInfoDiv.innerHTML = `
        <p>Key: ${keyInfo.key}</p>
        <p>Duration: ${keyInfo.duration}</p>
        <p>Status: ${keyInfo.status}</p>
        <p>Used: ${keyInfo.used ? 'Yes' : 'No'}</p>
    `;
}

// Implementing soon still testing
function ValidateKey(key) {
    const validKeys = ['123', 'abc', 'xyz'];

    if (validKeys.includes(key)) {
        let duration, status;
        if (key.includes('Lifetime')) {
            duration = 'Lifetime';
            status = 'Active';
        } else if (key.includes('month')) {
            duration = '1 month';
            status = 'Active';
        } else if (key.includes('week')) {
            duration = '1 week';
            status = 'Active';
        } else if (key.includes('day')) {
            duration = '1 day';
            status = 'Active';
        } else {
            duration = 'N/A';
            status = 'Invalid';
        }

        return {
            key: key,
            duration: duration,
            status: status,
            used: false
        };
    } else {
        return {
            key: key,
            duration: 'N/A',
            status: 'Invalid',
            used: false
        };
    }
}

function GenerateKey(event) {
    event.preventDefault();
    var duration = document.getElementById("duration").value;
    var planType = document.getElementById("planType").value;
    var key = GenerateKey();

    ShowKeys(key, duration, planType);
    UpdateKeys();

    var downloadAllButton = document.getElementById("downloadAllButton");
    if (!downloadAllButton) {
        downloadAllButton = document.createElement('button');
        downloadAllButton.textContent = 'Download All Keys';
        downloadAllButton.id = 'downloadAllButton';
        downloadAllButton.onclick = DownloadKeys;
        document.getElementById("generateForm").appendChild(downloadAllButton);
    }
}

function ShowKeys(key, duration, planType) {
    var tableBody = document.getElementById("generatedKeys").getElementsByTagName("tbody")[0];
    var newRow = tableBody.insertRow(tableBody.rows.length);
    var keyCell = newRow.insertCell(0);
    keyCell.textContent = key;

    var copyButton = document.createElement('button');
    copyButton.textContent = 'Copy';
    copyButton.onclick = function() {
        CopyKey(key);
    };
    keyCell.appendChild(copyButton);

    newRow.insertCell(1).textContent = duration;
    newRow.insertCell(2).textContent = planType;

    var manageCell = newRow.insertCell(3);
    manageCell.innerHTML = `
        <button onclick="OpenKeyModel()">Manage</button> <!-- Changed to OpenKeyModel -->
    `;

    UpdateKeys();
}

function OpenKeyModel() {
    var modal = document.getElementById('ManageKeysModal');
    modal.style.display = 'flex';
}

function CloseKeyModel() {
    var modal = document.getElementById('ManageKeysModal');
    modal.style.display = 'none';
}
function ManageKeys(action) {// needs work done
    var key = document.getElementById('keyInput').value;
    switch(action) {
        case 'remove':
            console.log('Removing key:', key);
            break;
        case 'edit':
            console.log('Editing key:', key);
            break;
        case 'reassign':
            console.log('Reassigning key:', key);
            break;
        default:
            console.log('Invalid action');
    }
    CloseKeyModel();
}

function UpdateKeys() {
    var keysData = [];
    var tableRows = document.querySelectorAll("#generatedKeys tbody tr");
    tableRows.forEach(function(row) {
        var cells = row.querySelectorAll("td");
        var key = cells[0].textContent.trim();
        var duration = cells[1].textContent.trim();
        var planType = cells[2].textContent.trim();
        keysData.push(`${key},${duration},${planType}`);
    });
    document.getElementById('keysData').value = keysData.join("\n");
}

// took me fucking forever to get this shit to export properly it kept exporting the full data including "Copy"
function DownloadKeys() {
    var keysData = document.getElementById('keysData').value;
    var formattedKeys = "Key | Duration | Plan Type\n";
    var keysArray = keysData.split('\n');
    var isFirstGroup = true;
    keysArray.forEach(function(keyRow) {
        if (keyRow.trim() !== '') {
            var keyInfo = keyRow.split(',');// tried to use split but it didn't work out well
            var key = keyInfo[0].replace('Copy', '');
            if (isFirstGroup) {
                formattedKeys += `${key} | ${keyInfo[1]} | ${keyInfo[2]}`;
                isFirstGroup = false;
            } else {
                formattedKeys += `\n${key} | ${keyInfo[1]} | ${keyInfo[2]}`;
            }
        }
    });
    var blob = new Blob([formattedKeys], { type: "text/plain;charset=utf-8" });
    var a = document.createElement("a");
    a.href = window.URL.createObjectURL(blob);
    a.download = "Keys.txt";
    a.click();
}

function CopyKey(text) {
    var textarea = document.createElement("textarea");
    textarea.value = text;
    document.body.appendChild(textarea);
    textarea.select();
    document.execCommand("copy");
    document.body.removeChild(textarea);
}

function GenerateKey() {// math.random() isnt the best way to do this but it works for a concept
    var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var result = '';
    var charactersLength = characters.length;
    for (var i = 0; i < 8; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}