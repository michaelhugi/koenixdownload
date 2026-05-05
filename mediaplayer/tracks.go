package mediaplayer

var tracks = []track{
	{name: "Tanz es raus", urlSlug: "na8oj1", fileName: "01 Tanz es raus.wav", backgroundFile: "AND_01_Tanz-es-raus.jpg"},
	{name: "Andertal", urlSlug: "btovoz", fileName: "02 Andertal.wav", backgroundFile: "AND_02_Andertal.jpg"},
	{name: "Finsteraarhorn", urlSlug: "qthzsz", fileName: "03 Finsteraarhorn.wav", backgroundFile: "AND_03_Finsteraarhorn.jpg"},
	{name: "Gletscherfee", urlSlug: "pcoaaw", fileName: "04 Gletscherfee.wav", backgroundFile: "AND_04_Gletscherfee.jpg"},
	{name: "Feuerkuss", urlSlug: "c5arvn", fileName: "05 Feuerkuss.wav", backgroundFile: "AND_05_Feuerkuss.jpg"},
	{name: "Gratzug", urlSlug: "i3awuz", fileName: "06 Gratzug.wav", backgroundFile: "AND_06_Gratzug-CH.jpg"},
	{name: "Ke Troum", urlSlug: "k9x4sa", fileName: "07 Ke Troum.wav", backgroundFile: "AND_07_Ke-Troum.jpg"},
	{name: "Drachenreiter", urlSlug: "dvraya", fileName: "08 Drachenreiter.wav", backgroundFile: "AND_08_Drachenreiter.jpg"},
	{name: "Imagonomanie", urlSlug: "7iorfc", fileName: "09 Imagonomanie.wav", backgroundFile: "AND_09_Imagonomanie.jpg"},
	{name: "Hier & Jetzt", urlSlug: "iigvdv", fileName: "10 Hier Jetzt.wav", backgroundFile: "AND_10_Hier-und-Jetzt.jpg"},
	{name: "Gratzug DE", urlSlug: "4gibew", fileName: "11 Gratzug DE.wav", backgroundFile: "AND_11_Gratzug-DE.jpg"},
}

type track struct {
	name           string
	urlSlug        string
	fileName       string
	backgroundFile string
}
