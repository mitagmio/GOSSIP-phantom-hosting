package config

import (
	"os"
	"testing"
	"time"
)

func TestGenerateNodeDetails(t *testing.T) {
	masternode := struct {
		Alias            string
		Genkey           string
		IPv4             string
		Port             int
		TransactionID    string
		TransactionIndex int
		EpochTime        int64
	}{
		"331720b1-6d69-404c-b84e-932642c93e92",
		"75eqvNfaEfkd3YTwQ3hMwyxL2BgNSrqHDgWc6jbUh4Gdtnro2Wo",
		"251.84.138.243",
		9998,
		"f8a3e39da2d13e10736a77940a2a78823e30e3ac40140f0a0b1ec31d07989aef",
		1,
		time.Now().Unix(),
	}

	if masternode.Alias == "" {
		t.Errorf("Alias is empty: got :%s, want: %s", masternode.Alias, "331720b1-6d69-404c-b84e-932642c93e92")
	}
	if masternode.Genkey == "" {
		t.Errorf("Genkey is empty: got :%s, want: %s", masternode.Genkey, "75eqvNfaEfkd3YTwQ3hMwyxL2BgNSrqHDgWc6jbUh4Gdtnro2Wo")
	}
	if masternode.IPv4 == "" {
		t.Errorf("IP is empty: got :%s, want: %s", masternode.IPv4, "251.84.138.243")
	}
	if masternode.Port == 0 {
		t.Errorf("Port has no value: got :%d, want: %d", masternode.Port, 9998)
	}
	if masternode.TransactionID == "" {
		t.Errorf("Transaction ID is empty: got :%s, want: %s", masternode.TransactionID, "f8a3e39da2d13e10736a77940a2a78823e30e3ac40140f0a0b1ec31d07989aef")
	}
	if masternode.TransactionIndex < 0 || masternode.TransactionIndex > 9 {
		t.Errorf("Transaction Index is empty or exceeds the min/max values: got :%d, want: %d", masternode.TransactionIndex, 1)
	}
	if masternode.EpochTime != time.Now().Unix() {
		t.Errorf("Epochtime is offset: got :%d, want: %d", masternode.EpochTime, time.Now().Unix())
	}
}

func TestGenerateConfigurationFile(t *testing.T) {
	_, err := os.Create("../masternode.txt")

	if err != nil {
		t.Errorf("An error occurred")
	}

	if _, err := os.Stat("../masternode.txt"); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("Masternode.txt does not exist")
		}
	}
}
