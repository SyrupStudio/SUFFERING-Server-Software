//go:build windows

/*
Copyright © 2026 Syrup Studios>
*/

package utils

// Because of the import only windows will compile here :/
import "golang.org/x/sys/windows"

func IsElevated() bool {
	var token windows.Token
	err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token)
	if err != nil {
		return false
	}
	defer token.Close()

	adminSID, err := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid)
	if err != nil {
		return false
	}

	member, err := token.IsMember(adminSID)
	if err != nil {
		return false
	}

	return member
}
