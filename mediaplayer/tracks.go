package mediaplayer

var tracks = []track{
	{name: "Tanz es raus", urlSlug: "na8oj1", fileName: "01 Tanz es raus.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Andertal", urlSlug: "btovoz", fileName: "02 Andertal.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Finsteraarhorn", urlSlug: "qthzsz", fileName: "03 Finsteraarhorn.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Gletscherfee", urlSlug: "pcoaaw", fileName: "04 Gletscherfee.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Feuerkuss", urlSlug: "c5arvn", fileName: "05 Feuerkuss.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Gratzug", urlSlug: "i3awuz", fileName: "06 Gratzug.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Ke Troum", urlSlug: "k9x4sa", fileName: "07 Ke Troum.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Drachenreiter", urlSlug: "dvraya", fileName: "08 Drachenreiter.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Imagonomanie", urlSlug: "7iorfc", fileName: "09 Imagonomanie.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Hier & Jetzt", urlSlug: "iigvdv", fileName: "10 Hier Jetzt.wav", backgroundFile: "song_trs_bg.jpg"},
	{name: "Gratzug DE", urlSlug: "4gibew", fileName: "11 Gratzug DE.wav", backgroundFile: "song_trs_bg.jpg"},
}

type track struct {
	name           string
	urlSlug        string
	fileName       string
	backgroundFile string
}
