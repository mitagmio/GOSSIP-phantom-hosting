package config

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/icrowley/fake"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type MasternodeString struct {
	Alias            uuid.UUID `json:"alias"`
	Genkey           string    `json:"genkey"`
	IPv4             string    `json:"ipv6"`
	Port             int       `json:"port"`
	TransactionID    string    `json:"txid"`
	TransactionIndex int       `json:"tx_index"`
	EpochTime        int64     `json:"epoch_time"`
}

func GenerateNodeDetails(m MasternodeString) (mnString string, err error) {
	var alias = uuid.Must(uuid.NewV4())

	m.EpochTime = time.Now().Unix()
	m.IPv4 = fake.IPv4()
	m.Alias = alias
	m.Port = 36001
	txbytes := len(m.TransactionID)
	genkeybytes := len(m.Genkey)

	if m.Genkey == "" || !strings.HasPrefix((m.Genkey), "4") || genkeybytes != 51 {
		return "", errors.New("Masternode Genkey incorrect")
	}

   if m.TransactionID == "" || txbytes != 64 {
		return "", errors.New("Collateral Transaction ID incorrect")
	}
	if m.TransactionIndex < 0 || m.TransactionIndex > 1 {
		return "", errors.New("TX Index incorrect")
	}
	
	if m.Port != 36001 {
		return "", errors.New("Port incorrect")
	}

	mnString = fmt.Sprintf("%v %s:%d %s %s %d %d", m.Alias, m.IPv4, m.Port, m.Genkey, m.TransactionID, m.TransactionIndex, m.EpochTime)

	return mnString, nil
}

func GenerateConfigurationFile(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func ViewConfiguration(path string) (config string, err error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	dataString := string(data)

	return dataString, nil
}

func AddMasternodeToConfigFile(path string, strMasternode string) (err error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.Write([]byte(strMasternode + "\n")); err != nil {
		return err
	}

	return nil
}
