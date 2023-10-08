package helper

func WaterStatus(water int) string {
	var status string

	switch {
	case water <= 5:
		status = "aman"
	case water > 5 && water <= 8:
		status = "siaga"
	default:
		status = "bahaya"
	}

	return status
}

func WindStatus(wind int) string {
	var status string

	switch {
	case wind <= 6:
		status = "aman"
	case wind > 7 && wind <= 15:
		status = "siaga"
	default:
		status = "bahaya"
	}

	return status
}
