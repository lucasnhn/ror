package rortypes

type CVEId string

type ResourceVirtualMachineVulnerabilityInfo struct {
	Id                  string  `json:"id"`
	HostName            string  `json:"hostName"`
	HostSeverity        string  `json:"hostSeverity"`
	Severity            string  `json:"severity"`
	SeverityScore       float32 `json:"severityScore"`
	LastCalculationTime int64   `json:"lastCalculationTime"`
	LastReportTime      int64   `json:"lastReportTime"`
	CVEs                []CVEId `json:"cves"`
}
