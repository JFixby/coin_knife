package str

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"github.com/picfight/coin_knife/lang"
	"github.com/picfight/coin_knife/projectops"
	"os"
	"path/filepath"
	"strings"
)

func TransferFiles(set *settings.Settings, from string, to string) {

	inputFiles := projectops.ListInputProjectFiles(from)

	for _, f := range inputFiles {
		postfix := strings.TrimPrefix(f, from)
		postfix = set.FileNameProcessor(postfix, set)
		//pin.D("postfix", postfix)
		newpath := filepath.Join(to, postfix)
		//pin.D("newpath", newpath)
		if fileops.IsFolder(f) {
			err := os.MkdirAll(newpath, 0700)
			lang.CheckErr(err)
			pin.D("make", newpath)
			continue
		}
		if fileops.IsFile(f) {
			ProcessFile(f, newpath, set)
			continue
		}
	}

}

func ProcessFile(from string, to string, set *settings.Settings) {
	if fileops.IsFolder(from) {
		lang.ReportErr("This is not a file: %v", from)
	}
	if set.IsProcessableFile(from) {
		data := fileops.ReadFileToString(from)
		if !set.DoNotProcessFiles {
			data = set.TextFileProcessor(data, set)
		}
		fileops.WriteStringToFile(to, data)
	} else {
		fileops.Copy(from, to)
	}
}

func Replace(data, from, to string) string {
	return strings.Replace(data, from, to, -1)
}

func PicfightCoinFileNameGenerator(data string, set *settings.Settings) string {
	data = Replace(data, "dcr", "pfc")
	data = Replace(data, "decred", "picfight")
	data = Replace(data, "decrediton", "pfcredit")
	data = Replace(data, "DecredLoading", "PicFightLoading")
	return data
}

func PicfightCoinFileGenerator(data string, set *settings.Settings) string {
	data = Replace(data, "DCR", "PFC")
	data = Replace(data, `decred/decrediton`, `picfight/pfcredit`)
	data = Replace(data, `decred\decrediton`, `picfight\pfcredit`)
	data = Replace(data, `decred\dcrd`, `picfight\pfcd`)
	data = Replace(data, `decred\dcrwallet`, `picfight\pfcwallet`)
	data = Replace(data, `decred\\dcrd`, `picfight\\pfcd`)
	data = Replace(data, `decred\\dcrwallet`, `picfight\\pfcwallet`)
	data = Replace(data, "decred/dcrd", "picfight/pfcd")
	data = Replace(data, "decrediton", "pfcredit")
	data = Replace(data, "Decrediton", "Pfcredit")
	data = Replace(data, "a decred", "a picfight")
	data = Replace(data, "decred/dcrutil", "picfight/pfcutil")
	data = Replace(data, "dcrutil", "pfcutil")
	data = Replace(data, "dcrd", "pfcd")
	data = Replace(data, "buy-decred.svg", "buy-picfight.svg")
	data = Replace(data, "/decred-", "/picfight-")
	data = Replace(data, "dcrec", "pfcec")
	data = Replace(data, "dcrj", "pfcj")
	data = Replace(data, "dcrpg", "pfcpg")
	data = Replace(data, "dcrsqlite", "pfcsqlite")
	data = Replace(data, "dcricon", "pfcicon")
	data = Replace(data, "dcrpassword", "pfcpassword")
	data = Replace(data, `json:\"dcr_`, `json:\"pfc_`)
	data = Replace(data, ".dcr-total", ".pfc-total")
	data = Replace(data, ".dcr ", ".pfc ")
	data = Replace(data, "decred - symbol", "picfight - symbol")
	data = Replace(data, "-decred:", "-picfight:")
	data = Replace(data, "-decred ", "-picfight ")
	data = Replace(data, "-decred-", "-picfight-")
	data = Replace(data, `name="decred"`, `name="picfight"`)
	data = Replace(data, "deCRED", "PicFight coin")
	data = Replace(data, "right dcr mono", "right pfc mono")
	data = Replace(data, "Dcrd", "Pfcd")
	data = Replace(data, "Dcrctl", "Pfcctl")
	data = Replace(data, "DcrWallet", "PfcWallet")
	data = Replace(data, "Dcrwallet", "Pfcwallet")
	data = Replace(data, "decred network", "picfight network")
	data = Replace(data, "decred supports", "picfight supports")
	data = Replace(data, "the decred", "the picfight")
	data = Replace(data, "for decred", "for picfight")
	data = Replace(data, "decred RPC", "picfight RPC")
	data = Replace(data, "and decred", "and picfight")
	data = Replace(data, "decred wallet", "picfight wallet")
	data = Replace(data, "when decred", "when picfight")
	data = Replace(data, "decred/decred-release", "picfight/picfight-release")
	data = Replace(data, "decred-release", "picfight-release")
	data = Replace(data, "#decred-golang", "#picfight-golang")
	data = Replace(data, "decred/$DOCKER_IMAGE_TAG", "picfight/$DOCKER_IMAGE_TAG")
	data = Replace(data, "/decred/bin", "/picfight/bin")
	data = Replace(data, "github.com/decred/$REPO", "github.com/picfight/$REPO")
	data = Replace(data, "github.com/decred/decrediton", "github.com/pfcredit")
	data = Replace(data, "with decred", "with picfight")
	data = Replace(data, "#decred", "#picfight")
	data = Replace(data, "=decred", "=picfight")
	data = Replace(data, "decred.slack.com", "picfight.slack.com")
	data = Replace(data, "unmined decred", "unmined picfight")
	data = Replace(data, "associated decred", "associated picfight")
	data = Replace(data, "of decred", "of picfight")
	data = Replace(data, "to decred", "to picfight")
	data = Replace(data, "supported decred", "supported picfight")
	data = Replace(data, "handling decred", "handling picfight")
	data = Replace(data, "wrong decred", "wrong picfight")
	data = Replace(data, "decred.org", "picfight.org")
	data = Replace(data, "decred/dcrctl", "picfight/pfcctl")
	data = Replace(data, "dcrctl", "pfcctl")
	data = Replace(data, "decred/dcrwallet", "picfight/pfcwallet")
	data = Replace(data, "decred/dcrwallet", "picfight/pfcwallet")
	data = Replace(data, "dcrwallet", "pfcwallet")
	data = Replace(data, "decredaddress", "picfightaddress")
	data = Replace(data, "decred transaction", "picfight coin transaction")
	data = Replace(data, "spend decred", "spend picfight coin")
	data = Replace(data, "dcrnet", "pfcnet")
	data = Replace(data, "decreds", "picfight coins")
	data = Replace(data, "decred-specific", "picfight-specific")
	data = Replace(data, "in decred", "in picfight coins")
	data = Replace(data, "dcrwire", "pfcwire")
	data = Replace(data, "dcrPK", "pfcPK")

	data = Replace(data, DCR_P2P_MAINNET_PORT(set)+":", PFC_P2P_MAINNET_PORT(set)+":")
	data = Replace(data, DCR_P2P_TESTNET_PORT(set)+":", PFC_P2P_TESTNET_PORT(set)+":")
	data = Replace(data, DCR_RPC_MAINNET_PORT(set)+":", PFC_RPC_MAINNET_PORT(set)+":")
	data = Replace(data, DCR_RPC_TESTNET_PORT(set)+":", PFC_RPC_TESTNET_PORT(set)+":")

	data = Replace(data, DCR_P2P_MAINNET_PORT(set)+"|", PFC_P2P_MAINNET_PORT(set)+"|")
	data = Replace(data, DCR_P2P_TESTNET_PORT(set)+"|", PFC_P2P_TESTNET_PORT(set)+"|")
	data = Replace(data, DCR_RPC_MAINNET_PORT(set)+"|", PFC_RPC_MAINNET_PORT(set)+"|")
	data = Replace(data, DCR_RPC_TESTNET_PORT(set)+"|", PFC_RPC_TESTNET_PORT(set)+"|")

	data = Replace(data, ":"+DCR_P2P_MAINNET_PORT(set), ":"+PFC_P2P_MAINNET_PORT(set)+"")
	data = Replace(data, ":"+DCR_P2P_TESTNET_PORT(set)+"", ":"+PFC_P2P_TESTNET_PORT(set)+"")
	data = Replace(data, ":"+DCR_RPC_MAINNET_PORT(set)+"", ":"+PFC_RPC_MAINNET_PORT(set)+"")
	data = Replace(data, ":"+DCR_RPC_TESTNET_PORT(set)+"", ":"+PFC_RPC_TESTNET_PORT(set)+"")

	data = Replace(data, ": "+DCR_P2P_MAINNET_PORT(set), ": "+PFC_P2P_MAINNET_PORT(set)+"")
	data = Replace(data, ": "+DCR_P2P_TESTNET_PORT(set)+"", ": "+PFC_P2P_TESTNET_PORT(set)+"")
	data = Replace(data, ": "+DCR_RPC_MAINNET_PORT(set)+"", ": "+PFC_RPC_MAINNET_PORT(set)+"")
	data = Replace(data, ": "+DCR_RPC_TESTNET_PORT(set)+"", ": "+PFC_RPC_TESTNET_PORT(set)+"")

	data = Replace(data, " "+DCR_P2P_MAINNET_PORT(set)+" ", " "+PFC_P2P_MAINNET_PORT(set)+" ")
	data = Replace(data, " "+DCR_P2P_TESTNET_PORT(set)+" ", " "+PFC_P2P_TESTNET_PORT(set)+" ")
	data = Replace(data, " "+DCR_RPC_MAINNET_PORT(set)+" ", " "+PFC_RPC_MAINNET_PORT(set)+" ")
	data = Replace(data, " "+DCR_RPC_TESTNET_PORT(set)+" ", " "+PFC_RPC_TESTNET_PORT(set)+" ")

	data = Replace(data, "\""+DCR_P2P_MAINNET_PORT(set)+"\"", "\""+PFC_P2P_MAINNET_PORT(set)+"\"")
	data = Replace(data, "\""+DCR_P2P_TESTNET_PORT(set)+"\"", "\""+PFC_P2P_TESTNET_PORT(set)+"\"")
	data = Replace(data, "\""+DCR_RPC_MAINNET_PORT(set)+"\"", "\""+PFC_RPC_MAINNET_PORT(set)+"\"")
	data = Replace(data, "\""+DCR_RPC_TESTNET_PORT(set)+"\"", "\""+PFC_RPC_TESTNET_PORT(set)+"\"")

	data = Replace(data, "dcrchain", "pfcchain")
	data = Replace(data, "dcr work", "pfc work")
	data = Replace(data, "dcr/kb", "pfc/kb")
	data = Replace(data, "decred/dcrrpcclient", "picfight/pfcrpcclient")
	data = Replace(data, "dcrrpcclient", "pfcrpcclient")
	data = Replace(data, "The Decred developers", "#DEVS#")
	data = Replace(data, "The Decred Developers", "#DEVS#")
	data = Replace(data, "The Decred Authors", "#DEVS#")
	data = Replace(data, "DecredLoading", "PicFightLoading")
	data = Replace(data, "Decred", "PicFight")
	data = Replace(data, "#DEVS#", "The Decred developers")
	return data
}

func DCR_RPC_TESTNET_PORT(set *settings.Settings) string {
	return set.DecredNetwork.Testnet.NodeRPCPort
}

func DCR_RPC_MAINNET_PORT(set *settings.Settings) string {
	return set.DecredNetwork.Mainnet.NodeRPCPort
}

func DCR_P2P_TESTNET_PORT(set *settings.Settings) string {
	return set.DecredNetwork.Testnet.NodeP2PPort
}

func DCR_P2P_MAINNET_PORT(set *settings.Settings) string {
	return set.DecredNetwork.Mainnet.NodeP2PPort
}

func PFC_RPC_TESTNET_PORT(set *settings.Settings) string {
	return set.OutputNetwork.Testnet.NodeRPCPort
}

func PFC_RPC_MAINNET_PORT(set *settings.Settings) string {
	return set.OutputNetwork.Mainnet.NodeRPCPort
}

func PFC_P2P_TESTNET_PORT(set *settings.Settings) string {
	return set.OutputNetwork.Testnet.NodeP2PPort
}

func PFC_P2P_MAINNET_PORT(set *settings.Settings) string {
	return set.OutputNetwork.Mainnet.NodeP2PPort
}
