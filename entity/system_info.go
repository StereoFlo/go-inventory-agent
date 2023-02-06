package entity

type SystemInfo struct {
	Hostname   string       `bson:"hostname" json:"net_name"`
	Platform   string       `bson:"platform" json:"name"`
	CPU        string       `bson:"cpu" json:"cpu"`
	RAM        uint64       `bson:"ram" json:"ram"`
	Partitions []*Partition `json:"partitions"`
	IP         string       `json:"ip"`
	User       *User        `json:"user"`
}
