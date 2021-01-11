document.getElementById("loadLogs").addEventListener("click", function(event) {
    loadLogs();
    event.preventDefault();
});

async function loadLogs() {
    // Check for the events fired
    $("#display-logs").show();
    logData = await getLogData();
    let logTableHTML = "";
    for (let index in logData) {
        console.log(logData[index].blockNumber);
        logTableHTML = logTableHTML + "<tr><td>" + logData[index].blockNumber + "</td><td>" + logData[index].returnValues.sender + "</td><td>" + logData[index].returnValues.data + "</td></tr>";
    }
    $("#logTable").html(logTableHTML);
}

function getLogData() {
    return new Promise(resolve => {
        LoggerContract.getPastEvents('Log', {
            filter: {
                sender: web3.utils.toChecksumAddress("0x2dA0a615981C2c9c70E34b8f50Db5f5a905E7928")
            },
            fromBlock: 0,
            toBlock: 'latest'
        }, (error, events) => {
            if (!error) {
                var obj = JSON.parse(JSON.stringify(events));
                var array = Object.keys(obj)

                console.log("returned values", obj[array[0]].returnValues);
                resolve(events);
            } else {
                console.error(error);
                resolve(error);
            }
        });
    });
}