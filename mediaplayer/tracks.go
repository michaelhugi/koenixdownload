package mediaplayer

var tracks = []track{
	{name: "Tanz es raus", urlSlug: "na8oj1", fileName: "01 Tanz es raus.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Andertal", urlSlug: "btovoz", fileName: "02 Andertal.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Finsteraarhorn", urlSlug: "qthzsz", fileName: "03 Finsteraarhorn.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Gletscherfee", urlSlug: "pcoaaw", fileName: "04 Gletscherfee.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Feuerkuss", urlSlug: "c5arvn", fileName: "05 Feuerkuss.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Gratzug", urlSlug: "i3awuz", fileName: "06 Gratzug.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Ke Troum", urlSlug: "k9x4sa", fileName: "07 Ke Troum.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Drachenreiter", urlSlug: "dvraya", fileName: "08 Drachenreiter.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Imagonomanie", urlSlug: "7iorfc", fileName: "09 Imagonomanie.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Hier & Jetzt", urlSlug: "iigvdv", fileName: "10 Hier Jetzt.wav", backgroundFile: "tanzesraus.jpg"},
	{name: "Gratzug DE", urlSlug: "4gibew", fileName: "11 Gratzug DE.wav", backgroundFile: "tanzesraus.jpg"},
}

type track struct {
	name           string
	urlSlug        string
	fileName       string
	backgroundFile string // e.g., "/embed/forest.jpg"
}
