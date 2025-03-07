package objects

type RevenuStructure string

const RevenuStructureFlat = "flat"

const (
	RevenuePeriodQuarter = "quarter"
	RevenuePeriodAnnual  = "annual"
)

type RequestRevenueProductSegmentation struct {
	Symbol    string
	Structure string
	Period    string
}

type RevenueProductSegmentation map[string]interface{}

type RequestRevenueGeoSegmentation struct {
	Symbol    string
	Structure string
}

type RevenueGeoSegmentation map[string]interface{}
