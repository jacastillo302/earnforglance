package domain

// RobotsTxtDefaults represents default values related to robots.txt
type RobotsTxtDefaults struct{}

// RobotsCustomFileName gets the name of the custom robots file
func (r RobotsTxtDefaults) RobotsCustomFileName() string {
	return "robots.custom.txt"
}

// RobotsAdditionsFileName gets the name of the robots additions file
func (r RobotsTxtDefaults) RobotsAdditionsFileName() string {
	return "robots.additions.txt"
}
