package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/sys/windows/registry"
)

func print(args interface{}) {
	fmt.Println(args)
}

func check(e error, errMsg string) bool {
	if e != nil {
		print(errMsg)
		os.Exit(1)
	}
	return true
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func getNumberOfSubKeysAndValues(k registry.Key) (uint32, uint32) {
	keyInfo, err := k.Stat()
	check(err, "Unable to fetch Stat info from registry object...")
	return keyInfo.SubKeyCount, keyInfo.ValueCount
}

func openKey(hive registry.Key, subkey string, access uint32) registry.Key {
	key, err := registry.OpenKey(hive, subkey, access)
	check(err, "Unable to open registry key...")
	return key
}

func getComputerInfo() {
	key := openKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		registry.ALL_ACCESS,
	)
	defer key.Close()

	blu := color.New(color.FgBlue)
	boldBlue := blu.Add(color.Bold)
	boldBlue.Println("◎ ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ Computer Build Info ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ ◎")
	print("")

	productName, _, err := key.GetStringValue("ProductName")
	check(err, "ProductName value not found in registry...")
	print("Product Name : " + productName)
	currentVersion, _, err := key.GetStringValue("CurrentVersion")
	check(err, "CurrentVersion value not found in registry...")
	print("Current Version : " + currentVersion)
	currentBuildNumber, _, err := key.GetStringValue("CurrentBuildNumber")
	check(err, "CurrentBuildNumber Value not found in registry...")
	print("Build Number : " + currentBuildNumber)
	registeredOwner, _, err := key.GetStringValue("RegisteredOwner")
	check(err, "RegisteredOwner value not found in registry...")
	print("Registered Owner : " + registeredOwner)
	print("")
}

func getInstalledApps() {
	key := openKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
		registry.ALL_ACCESS,
	)
	defer key.Close()

	numOfSubKeys, numOfValues := getNumberOfSubKeysAndValues(key)
	subkeys, err := key.ReadSubKeyNames(int(numOfSubKeys))
	check(err, "Unable to read subkeys...")

	rd := color.New(color.FgRed)
	boldRed := rd.Add(color.Bold)
	boldRed.Println("◎ ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ Installed Applications ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ ◎")
	print("")
	for _, skey := range subkeys {
		k := openKey(
			registry.LOCAL_MACHINE,
			`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`+"\\"+skey,
			registry.ALL_ACCESS,
		)
		values, err := k.ReadValueNames(int(numOfValues))
		check(err, "Unable to read values from registry key...")
		if exist := find(values, "DisplayName"); exist {
			val, _, err := k.GetStringValue("DisplayName")
			check(err, "Unable to retrieve data from value DisplayName...")
			print("\u2022 " + val)
		} else {
			print("\u2022 " + skey)
		}
	}
}

func getEnVars() {
	key := openKey(
		registry.LOCAL_MACHINE,
		`SYSTEM\CurrentControlSet\Control\Session Manager\Environment`,
		registry.ALL_ACCESS,
	)
	defer key.Close()

	_, numOfValues := getNumberOfSubKeysAndValues(key)
	environmentVariables, err := key.ReadValueNames(int(numOfValues))
	check(err, "Unable to read values from registry key...")

	grn := color.New(color.FgGreen)
	boldGreen := grn.Add(color.Bold)
	boldGreen.Println("\n◎ ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ Environment Variables ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ ◎")
	print("")

	for _, envar := range environmentVariables {
		envarValue, _, err := key.GetStringValue(envar)
		check(err, "Unable to retrieve data from value in registry key...")
		print(envar + " ☰☰ " + envarValue)
	}
	print("")
}

func getStartUpApps() {

}

func getJumpLists() {
	currentUser, err := user.Current()
	check(err, "Unable to fetch username...")
	username := strings.Split(currentUser.Username, `\`)
	jumpListPath := fmt.Sprintf(
		`C:\Users\%s\AppData\Roaming\Microsoft\Windows\Recent\AutomaticDestinations`,
		username[1],
	)
	jumpListFiles, err := ioutil.ReadDir(jumpListPath)
	check(err, "Unable to read files in jump list directory...")

	yel := color.New(color.FgYellow)
	boldYellow := yel.Add(color.Bold)
	boldYellow.Println("◎ ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ Jump List Files ☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶☶ ◎")
	print("")
	for _, file := range jumpListFiles {
		print(file.Name())
	}
	print("")
}

func getLNKFiles() {

}

func getShellBags() {

}

func getPrefetchFiles() {

}

func main() {
	getComputerInfo()
	getInstalledApps()
	getEnVars()
	getStartUpApps()
	getJumpLists()
	getLNKFiles()
}
