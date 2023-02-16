package address

import "strings"

func TypeAddress(address string) string {
	typeValids := []string{"Rua", "AV", "Avenida", "Alameda", "Rodovia"}
	fistWork := strings.Split(address, " ")[0]

	addressValids := false

	for _, types := range typeValids {
		if types == fistWork {
			addressValids = true
		}
	}

	if addressValids {
		return fistWork
	}

	return "type invalid"
}
