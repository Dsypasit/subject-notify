package services

func checkSubjectTime(time string) bool {
	timeValidate := [...]string{"Morning", "Afternoon"}
	for _, t := range timeValidate {
		if t == time {
			return true
		}
	}
	return false
}

func checkSubjectDay(day string) bool {
	dayValidate := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri"}
	for _, d := range dayValidate {
		if d == day {
			return true
		}
	}
	return false
}
