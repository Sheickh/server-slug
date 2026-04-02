package database

type Record struct{
	ID int
	URL string
	Slug string
}

var records []Record

func SaveRecords(newRecord []Record) error{
	records = make([]Record, len(newRecord))
	copy(records, newRecord)
	return nil
}

func LoadRecords() ([]Record, error){
	result := make([]Record, len(records))
	copy(result, records)
	return result, nil
}