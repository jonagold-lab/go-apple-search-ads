package searchads

type ReportingResponse struct {
	Row []ReportingRow
}

type ReportingRow struct {
	Metadata    Metadata
	Total       Total
	Granularity []GranularityObj
}

type GranularityObj struct {
	Date string
	Statistics
}

type Metadata struct {
}

type Total struct {
	LocalSpend Amount
	Statistics
}

type Statistics struct {
	Taps                    int
	Conversions             int
	AvgCPA                  Amount
	AvgCPT                  Amount
	TTR                     float64
	ConversionRate          float64
	Impressions             int
	ConversionsLATOn        int
	ConversionsLATOff       int
	ConversionsNewDownloads int
	ConversionsRedownloads  int
}

type ReportFilter struct {
	StartTime                  string      `json:"startTime,omitempty"`
	EndTime                    string      `json:"endTime,omitempty"`
	TimeZone                   TimeZone    `json:"timeZone,omitempty"`
	Granularity                Granularity `json:"granularity,omitempty"`
	Selector                   Selector    `json:"selector,omitempty"`
	GroupBy                    string      `json:"groupBy,omitempty"`
	ReturnRowTotals            bool        `json:"returnRowTotals,omitempty"`
	ReturnGrandTotals          bool        `json:"returnGrandTotals,omitempty"`
	ReturnRecordsWithNoMetrics bool        `json:"returnRecordsWithNoMetrics,omitempty"`
}

type Selector struct {
	Conditions []Condition      `json:"conditions,omitempty"`
	Fields     []string         `json:"fields,omitempty"`
	OrderBy    OrderBy          `json:"orderBy,omitempty"`
	Pagination FilterPagination `json:"pagination,omitempty"`
}

type FilterPagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type OrderBy struct {
	Field     string    `json:"field,omitempty"`
	SortOrder SortOrder `json:"sortOrder,omitempty"`
}

type Condition struct {
	Field    string   `json:"field"`
	Operator Operator `json:"operator"`
	Values   []string `json:"values"`
}
