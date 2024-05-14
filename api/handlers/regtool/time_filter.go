package v1

func UpdateTimeFilter(up string) string {
	if up != "0001-01-01 00:00:00" {
		return up
	}
	return ""
}
