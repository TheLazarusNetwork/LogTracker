package log

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"strings"

	"github.com/TheLazarusNetwork/LogTracker/logger"
	"github.com/TheLazarusNetwork/LogTracker/utility"
	"github.com/TheLazarusNetwork/LogTracker/wallet"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/log")
	{
		g.GET("/all", readAllLogs)
		g.GET("/version", readVersion)
	}
}

func readAllLogs(c *gin.Context) {
	resp := utility.MessageList(http.StatusOK, []string{"log1", "log2", "log3", "log4"})

	mnemonic := viper.Get("MNEMONIC").(string)
	privateKey, publicKey, path, err := wallet.HDWallet(mnemonic)
	utility.CheckError("Error in computing Hierarchical Deterministic Wallet:", err)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes) // hexutil.Encode(privateKeyBytes)[2:] for without 0x
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	publicKeyHex := hexutil.Encode(publicKeyBytes[1:]) // As Ethereum does not DER encode its public keys, public keys in Ethereum are only 64 bytes long
	walletAddress := crypto.PubkeyToAddress(*publicKey).Hex()

	// Display mnemonic and keys
	log.WithFields(utility.StandardFields).Infof("Mnemonic: %s", mnemonic)
	log.WithFields(utility.StandardFields).Infof("ETH Private Key: %s", privateKeyHex)
	log.WithFields(utility.StandardFields).Infof("ETH Public Key: %s", publicKeyHex)
	log.WithFields(utility.StandardFields).Infof("ETH Wallet Address: %s", walletAddress)
	log.WithFields(utility.StandardFields).Infof("Path: %s", *path)

	// ECIES Encryption and Decryption
	ecdsaPrivateKey, err := crypto.HexToECDSA(hexutil.Encode(privateKeyBytes)[2:])
	eciesPrivateKey := ecies.ImportECDSA(ecdsaPrivateKey)
	eciesPublicKey := eciesPrivateKey.PublicKey
	_ = eciesPublicKey

	infuraEndPoint := viper.Get("INFURA_ENDPOINT").(string)
	client, err := ethclient.Dial(infuraEndPoint)
	utility.CheckError("Error in connecting to Infura EndPoint:", err)

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKey))
	utility.CheckError("Error in fetching nonce:", err)
	log.WithFields(utility.StandardFields).Infof("Nonce: %d", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	utility.CheckError("Error in fetching Gas Price:", err)
	log.WithFields(utility.StandardFields).Infof("Gas Price: %d", gasPrice)

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(walletAddress), nil)
	utility.CheckError("Error in fetching account balance:", err)
	ethbalance := new(big.Float)
	ethbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(ethbalance, big.NewFloat(math.Pow10(18)))
	log.WithFields(utility.StandardFields).Infof("ETH Balance: %f", ethValue)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	loggerAddress := common.HexToAddress(viper.Get("LOGGER_CONTRACT_ADDRESS").(string))
	instance, err := logger.NewLogger(loggerAddress, client)
	utility.CheckError("Unable to load instance of the deployed contract:", err)
	_ = instance

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(7842388),
		ToBlock:   nil,
		Addresses: []common.Address{
			loggerAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	utility.CheckError("Failed to query logs:", err)

	loggerABI, err := abi.JSON(strings.NewReader(string(logger.LoggerABI)))
	utility.CheckError("Failed to parse logs:", err)
	_ = loggerABI

	for _, vLog := range logs {
		fmt.Println(vLog.BlockNumber)  // 7876312
		fmt.Println(vLog.TxHash.Hex()) // 0xb56a6e82fb1bddb96d3825ec21a0105df9625afdf5cf97d25de3efef940722be

		// ```indexed``` event types become a topic rather than part of the data property of the log. Total 4 topics with 3 indexed event types.
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}
		// The first topic is always the signature of the event
		// topic0 -> 0x0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e0 corresponding to Log(address,string) https://emn178.github.io/online-tools/keccak_256.html
		// topic1 -> 0x0000000000000000000000002da0a615981c2c9c70e34b8f50db5f5a905e7928 corresponding to indexed event types as its a topic rather than part of the data property of the log
		fmt.Println(common.HexToAddress(topics[1]))

		parsedLogs, err := loggerABI.Unpack("Log", vLog.Data)
		utility.CheckError("Failed to parse logs:", err)
		fmt.Println(parsedLogs[0])
	}

	// Decryption
	// decryptedLogData, err := eciesPrivateKey.Decrypt(encryptedLogData, nil, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Infof("TX Hash: %s --> Decrypted DataLog: %s", tx.Hash().Hex(), string(decryptedLogData))
	c.JSON(http.StatusOK, resp)
}

func readRecentLog(c *gin.Context) {
	resp := utility.Message(http.StatusOK, utility.Version)
	c.JSON(http.StatusOK, resp)
}

func readVersion(c *gin.Context) {
	resp := utility.Message(http.StatusOK, utility.Version)
	c.JSON(http.StatusOK, resp)
}
