package main

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []Data `json:"data"`
}

// Data represents the prayer timing data
type Data struct {
	Timings Timings `json:"timings"`
	Date    Date    `json:"date"`
	Meta    Meta    `json:"meta"`
}

// Timings represents the prayer times for a day
type Timings struct {
	Fajr       string `json:"Fajr"`
	Sunrise    string `json:"Sunrise"`
	Dhuhr      string `json:"Dhuhr"`
	Asr        string `json:"Asr"`
	Sunset     string `json:"Sunset"`
	Maghrib    string `json:"Maghrib"`
	Isha       string `json:"Isha"`
	Imsak      string `json:"Imsak"`
	Midnight   string `json:"Midnight"`
	Firstthird string `json:"Firstthird"`
	Lastthird  string `json:"Lastthird"`
}

// Date represents date information in multiple calendars
type Date struct {
	Readable  string    `json:"readable"`
	Timestamp string    `json:"timestamp"`
	Hijri     Hijri     `json:"hijri"`
	Gregorian Gregorian `json:"gregorian"`
}

// Hijri represents the Hijri (Islamic) calendar date
type Hijri struct {
	Date             string        `json:"date"`
	Format           string        `json:"format"`
	Day              string        `json:"day"`
	Weekday          HijriWeekday  `json:"weekday"`
	Month            HijriMonth    `json:"month"`
	Year             string        `json:"year"`
	Designation      Designation   `json:"designation"`
	Holidays         []interface{} `json:"holidays"`
	AdjustedHolidays []string      `json:"adjustedHolidays"`
	Method           string        `json:"method"`
}

// HijriWeekday represents the Hijri weekday in different languages
type HijriWeekday struct {
	En string `json:"en"`
	Ar string `json:"ar"`
}

// HijriMonth represents the Hijri month information
type HijriMonth struct {
	Number int    `json:"number"`
	En     string `json:"en"`
	Ar     string `json:"ar"`
	Days   int    `json:"days"`
}

// Gregorian represents the Gregorian calendar date
type Gregorian struct {
	Date          string           `json:"date"`
	Format        string           `json:"format"`
	Day           string           `json:"day"`
	Weekday       GregorianWeekday `json:"weekday"`
	Month         GregorianMonth   `json:"month"`
	Year          string           `json:"year"`
	Designation   Designation      `json:"designation"`
	LunarSighting bool             `json:"lunarSighting"`
}

// GregorianWeekday represents the Gregorian weekday
type GregorianWeekday struct {
	En string `json:"en"`
}

// GregorianMonth represents the Gregorian month
type GregorianMonth struct {
	Number int    `json:"number"`
	En     string `json:"en"`
}

// Designation represents the designation of a calendar
type Designation struct {
	Abbreviated string `json:"abbreviated"`
	Expanded    string `json:"expanded"`
}

// Meta represents metadata about the prayer times calculation
type Meta struct {
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	Timezone                 string  `json:"timezone"`
	Method                   Method  `json:"method"`
	LatitudeAdjustmentMethod string  `json:"latitudeAdjustmentMethod"`
	MidnightMode             string  `json:"midnightMode"`
	School                   string  `json:"school"`
	Offset                   Offset  `json:"offset"`
}

// Method represents the calculation method
type Method struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Params   MethodParams `json:"params"`
	Location Location     `json:"location"`
}

// MethodParams represents parameters for the calculation method
type MethodParams struct {
	Fajr int `json:"Fajr"`
	Isha int `json:"Isha"`
}

// Location represents the geographical location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Offset represents timing adjustments
type Offset struct {
	Imsak    int `json:"Imsak"`
	Fajr     int `json:"Fajr"`
	Sunrise  int `json:"Sunrise"`
	Dhuhr    int `json:"Dhuhr"`
	Asr      int `json:"Asr"`
	Sunset   int `json:"Sunset"`
	Maghrib  int `json:"Maghrib"`
	Isha     int `json:"Isha"`
	Midnight int `json:"Midnight"`
}
