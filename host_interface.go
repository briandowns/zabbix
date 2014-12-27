package zabbix

type (
	InterfaceType int
)

const (
	Agent InterfaceType = 1 << iota
	SNMP
	IPMI
	JMX
)

// https://www.zabbix.com/documentation/2.0/manual/appendix/api/hostinterface/definitions
type HostInterface struct {
	DNS   string        `json:"dns"`
	IP    string        `json:"ip"`
	Main  int           `json:"main"`
	Port  string        `json:"port"`
	Type  InterfaceType `json:"type"`
	UseIP int           `json:"useip"`
}

type HostInterfaces []HostInterface
