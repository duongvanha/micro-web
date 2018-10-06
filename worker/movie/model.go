package movie

type Chap struct {
	SeverName string
	Urls      []string
}

type Movie struct {
	NameVi       string
	NameEn       string
	Image        string
	Status       string
	IMDb         string
	Director     string
	Country      string
	Year         string
	Date         string
	Time         string
	Cam          string
	Quality      string
	Sub          string
	Categories   string
	Manufacturer string
	View         string
	Url          string
	Movies       []Chap
}
