// const IPFS = IpfsApi;
// const ipfs = new IPFS({ host: 'ipfs.infura.io', port: 5001, protocol: 'https' });

let coinbase;
let LoggerContract;

// Inject our version of web3.js into the DApp.
window.addEventListener('load', async() => {
    // $('#main-popup').modal('show');
    console.log("Welcome to LogTracker");
    await connectWallet();
});

document.getElementById("walletStatus").addEventListener("click", function(event) {
    connectWallet();
    event.preventDefault();
});

async function connectWallet() {
    // Modern dapp browsers...
    if (window.ethereum) {
        window.web3 = new Web3(ethereum);
        try {
            // Request account access if needed
            await ethereum.enable();
            // Acccounts now exposed
            let accounts = await ethereum.request({ method: 'eth_accounts' });
            console.log(`Accounts:\n${accounts.join('\n')}`);
            startDApp();
        } catch (error) {
            // User denied account access...
            console.log("User denied account access");
        }
    }
    // Legacy dapp browsers...
    else if (window.web3) {
        window.web3 = new Web3(web3.currentProvider);
        // Acccounts always exposed
    }
    // Non-dapp browsers...
    else {
        console.log('Non-Ethereum browser detected. Install MetaMask to continue!');
    }
}

// CheckNetwork
function checkNetwork() {
    return new Promise(resolve => {
        web3.eth.net.getNetworkType((error, netId) => {
            switch (netId) {
                case "main":
                    console.log('The Mainnet');
                    break
                case "ropsten":
                    console.log('Ropsten Test Network');
                    break
                case "rinkeby":
                    console.log('Rinkeby Test Network');
                    break
                case "goerli":
                    console.log('Goerli Test Network');
                    break
                case "kovan":
                    console.log('Kovan Test Network');
                    break
                default:
                    console.log('This is an Unknown Network');
            }
            if (!error) {
                console.log("Network: " + netId);
                resolve(netId);
            } else {
                resolve(error);
            }
        });
    });
}

function initDApp() {
    LoggerContract = new web3.eth.Contract(LoggerContractABI, LoggerContractAddress, {
        from: coinbase,
        gasPrice: '200000000000' // default gas price in wei, 20 gwei in this case
    });
}

function getCoinbase() {
    return new Promise(resolve => {
        web3.eth.getCoinbase((error, result) => {
            if (!error) {
                console.log("Coinbase: " + result);
                resolve(result);
            } else {
                resolve(error);
            }
        });
    });
}

// ETH Balance
function getETHBalance() {
    return new Promise(resolve => {
        web3.eth.getBalance(coinbase, (error, result) => {
            if (!error) {
                console.log(web3.utils.fromWei(result, "ether"));
                resolve(web3.utils.fromWei(result, "ether"));
            } else {
                resolve(error);
            }
        });
    });
}

async function fetchAccountDetails() {
    // Fetch the Account Details
    coinbase = web3.utils.toChecksumAddress(await getCoinbase());
    document.getElementById('walletStatus').innerHTML = "Connected";
    document.getElementById('walletAddress').innerHTML = "Connected Wallet: " + coinbase;
    let walletETHBalance = await getETHBalance();
    document.getElementById('walletDetails').innerHTML = "Connected to  " + (await checkNetwork()).toUpperCase() + " Network with Balance: " + walletETHBalance + " ETH";
}

async function startDApp() {
    console.log('Starting LogTracker DApp');
    let connectedNetwork = await checkNetwork();
    if (connectedNetwork != "rinkeby") {
        alert("Please connect to Rinkeby Testnet to Continue!");
    } else {
        console.log("Connected to Rinkeby Testnet!");
    }
    await fetchAccountDetails();
    initDApp();
    console.log("DApp Initialized");
}