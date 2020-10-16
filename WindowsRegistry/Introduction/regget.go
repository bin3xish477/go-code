package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/sys/windows/registry"
)

func print(args interface{}) {
	fmt.Println(args)
}

func check(e error) bool {
	if e != nil {
		print("An error occured ...")
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
	check(err)
	return keyInfo.SubKeyCount, keyInfo.ValueCount
}

func openKey(hive registry.Key, subkey string, access uint32) registry.Key {
	key, err := registry.OpenKey(hive, subkey, access)
	check(err)
	return key
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
	check(err)

	color.Red("◎ ≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡ Installed Applications ≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡ ◎")
	for _, skey := range subkeys {
		k := openKey(
			registry.LOCAL_MACHINE,
			`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`+"\\"+skey,
			registry.ALL_ACCESS,
		)
		values, err := k.ReadValueNames(int(numOfValues))
		check(err)
		if exist := find(values, "DisplayName"); exist {
			val, _, err := k.GetStringValue("DisplayName")
			check(err)
			print("\u2022 " + val)
		} else {
			print("\u2022 " + skey)
		}
	}
}

func getEnVars() {
	/*
		For system environment variables:
		HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment

		For User environment variables:
		HKEY_CURRENT_USER\Environment
	*/

}

func getStartUpApps() {

}

func main() {
	getInstalledApps()
	getEnVars()
	getStartUpApps()
}
