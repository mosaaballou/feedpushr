package whatlanggo

//Lang represents a language following ISO 639-3 standard.
type Lang int

//Aka ...
const (
	Aka Lang = iota
	Amh
	Arb
	Azj
	Bel
	Ben
	Bho
	Bul
	Ceb
	Ces
	Cmn
	Dan
	Deu
	Ell
	Eng
	Epo
	Est
	Fin
	Fra
	Guj
	Hat
	Hau
	Heb
	Hin
	Hrv
	Hun
	Ibo
	Ilo
	Ind
	Ita
	Jav
	Jpn
	Kan
	Kat
	Khm
	Kin
	Kor
	Kur
	Lav
	Lit
	Mai
	Mal
	Mar
	Mkd
	Mlg
	Mya
	Nep
	Nld
	Nno
	Nob
	Nya
	Ori
	Orm
	Pan
	Pes
	Pol
	Por
	Ron
	Run
	Rus
	Sin
	Skr
	Slv
	Sna
	Som
	Spa
	Srp
	Swe
	Tam
	Tel
	Tgl
	Tha
	Tir
	Tuk
	Tur
	Uig
	Ukr
	Urd
	Uzb
	Vie
	Ydd
	Yor
	Zul
)

//CodeToLang gets enum by ISO 639-3 code as a string.
func CodeToLang(code string) Lang {
	switch code {
	case "aka":
		return Aka
	case "amh":
		return Amh
	case "arb":
		return Arb
	case "azj":
		return Azj
	case "bel":
		return Bel
	case "ben":
		return Ben
	case "bho":
		return Bho
	case "bul":
		return Bul
	case "ceb":
		return Ceb
	case "ces":
		return Ces
	case "cmn":
		return Cmn
	case "dan":
		return Dan
	case "deu":
		return Deu
	case "ell":
		return Ell
	case "eng":
		return Eng
	case "epo":
		return Epo
	case "est":
		return Est
	case "fin":
		return Fin
	case "fra":
		return Fra
	case "guj":
		return Guj
	case "hat":
		return Hat
	case "hau":
		return Hau
	case "heb":
		return Heb
	case "hin":
		return Hin
	case "hrv":
		return Hrv
	case "hun":
		return Hun
	case "ibo":
		return Ibo
	case "ilo":
		return Ilo
	case "ind":
		return Ind
	case "ita":
		return Ita
	case "jav":
		return Jav
	case "jpn":
		return Jpn
	case "kan":
		return Kan
	case "kat":
		return Kat
	case "khm":
		return Khm
	case "kin":
		return Kin
	case "kor":
		return Kor
	case "kur":
		return Kur
	case "lav":
		return Lav
	case "lit":
		return Lit
	case "mai":
		return Mai
	case "mal":
		return Mal
	case "mar":
		return Mar
	case "mkd":
		return Mkd
	case "mlg":
		return Mlg
	case "mya":
		return Mya
	case "nep":
		return Nep
	case "nld":
		return Nld
	case "nno":
		return Nno
	case "nob":
		return Nob
	case "nya":
		return Nya
	case "ori":
		return Ori
	case "orm":
		return Orm
	case "pan":
		return Pan
	case "pes":
		return Pes
	case "pol":
		return Pol
	case "por":
		return Por
	case "ron":
		return Ron
	case "run":
		return Run
	case "rus":
		return Rus
	case "sin":
		return Sin
	case "skr":
		return Skr
	case "slv":
		return Slv
	case "sna":
		return Sna
	case "som":
		return Som
	case "spa":
		return Spa
	case "srp":
		return Srp
	case "swe":
		return Swe
	case "tam":
		return Tam
	case "tel":
		return Tel
	case "tgl":
		return Tgl
	case "tha":
		return Tha
	case "tir":
		return Tir
	case "tuk":
		return Tuk
	case "tur":
		return Tur
	case "uig":
		return Uig
	case "ukr":
		return Ukr
	case "urd":
		return Urd
	case "uzb":
		return Uzb
	case "vie":
		return Vie
	case "ydd":
		return Ydd
	case "yor":
		return Yor
	case "zul":
		return Zul
	default:
		return -1
	}
}

//LangToStringShort converts enum into ISO 639-1 code as a string.
//
// Return empty string when there is no ISO 639-1 code.
func LangToStringShort(lang Lang) string {
	switch lang {
	case Aka:
		return "ak"
	case Amh:
		return "am"
	case Arb:
		return "ar"
	case Azj:
		return "az" // Azerbaijani iso 639-3 is aze, iso 639-1 az
	case Bel:
		return "be"
	case Ben:
		return "bn"
	case Bho:
		return "bh"
	case Bul:
		return "bg"
	case Ceb:
		return "" // No iso 639-1 code
	case Ces:
		return "cs"
	case Cmn:
		return "zh" // No iso 639-1, but http://www.loc.gov/standards/iso639-2/faq.html#24
	case Dan:
		return "da"
	case Deu:
		return "de"
	case Ell:
		return "el"
	case Eng:
		return "en"
	case Epo:
		return "eo"
	case Est:
		return "et"
	case Fin:
		return "fi"
	case Fra:
		return "fr"
	case Guj:
		return "gu"
	case Hat:
		return "ht"
	case Hau:
		return "ha"
	case Heb:
		return "he"
	case Hin:
		return "hi"
	case Hrv:
		return "hr"
	case Hun:
		return "hu"
	case Ibo:
		return "ig"
	case Ilo:
		return "" // No iso639-1
	case Ind:
		return "id"
	case Ita:
		return "it"
	case Jav:
		return "jv"
	case Jpn:
		return "ja"
	case Kan:
		return "kn"
	case Kat:
		return "ka"
	case Khm:
		return "km"
	case Kin:
		return "rw"
	case Kor:
		return "ko"
	case Kur:
		return "ku"
	case Lav:
		return "lv"
	case Lit:
		return "lt"
	case Mai:
		return "" // No iso639-1
	case Mal:
		return "ml"
	case Mar:
		return "mr"
	case Mkd:
		return "mk"
	case Mlg:
		return "mg"
	case Mya:
		return "my"
	case Nep:
		return "ne"
	case Nld:
		return "nl"
	case Nno:
		return "nn"
	case Nob:
		return "nb"
	case Nya:
		return "ny"
	case Ori:
		return "or"
	case Orm:
		return "om"
	case Pan:
		return "pa"
	case Pes:
		return "" // No iso639-1
	case Pol:
		return "pl"
	case Por:
		return "pt"
	case Ron:
		return "ro"
	case Run:
		return "rn"
	case Rus:
		return "ru"
	case Sin:
		return "si"
	case Skr:
		return "" // No iso639-1
	case Slv:
		return "sl"
	case Sna:
		return "sn"
	case Som:
		return "so"
	case Spa:
		return "es"
	case Srp:
		return "sr"
	case Swe:
		return "sv"
	case Tam:
		return "ta"
	case Tel:
		return "te"
	case Tgl:
		return "tl"
	case Tha:
		return "th"
	case Tir:
		return "ti"
	case Tuk:
		return "tk"
	case Tur:
		return "tr"
	case Uig:
		return "ug"
	case Ukr:
		return "uk"
	case Urd:
		return "ur"
	case Uzb:
		return "uz"
	case Vie:
		return "vi"
	case Ydd:
		return "" // No iso639-1
	case Yor:
		return "yo"
	case Zul:
		return "zu"
	default:
		return ""
	}
}

//LangToString converts enum into ISO 639-3 code as a string.
func LangToString(lang Lang) string {
	switch lang {
	case Aka:
		return "aka"
	case Amh:
		return "amh"
	case Arb:
		return "arb"
	case Azj:
		return "azj"
	case Bel:
		return "bel"
	case Ben:
		return "ben"
	case Bho:
		return "bho"
	case Bul:
		return "bul"
	case Ceb:
		return "ceb"
	case Ces:
		return "ces"
	case Cmn:
		return "cmn"
	case Dan:
		return "dan"
	case Deu:
		return "deu"
	case Ell:
		return "ell"
	case Eng:
		return "eng"
	case Epo:
		return "epo"
	case Est:
		return "est"
	case Fin:
		return "fin"
	case Fra:
		return "fra"
	case Guj:
		return "guj"
	case Hat:
		return "hat"
	case Hau:
		return "hau"
	case Heb:
		return "heb"
	case Hin:
		return "hin"
	case Hrv:
		return "hrv"
	case Hun:
		return "hun"
	case Ibo:
		return "ibo"
	case Ilo:
		return "ilo"
	case Ind:
		return "ind"
	case Ita:
		return "ita"
	case Jav:
		return "jav"
	case Jpn:
		return "jpn"
	case Kan:
		return "kan"
	case Kat:
		return "kat"
	case Khm:
		return "khm"
	case Kin:
		return "kin"
	case Kor:
		return "kor"
	case Kur:
		return "kur"
	case Lav:
		return "lav"
	case Lit:
		return "lit"
	case Mai:
		return "mai"
	case Mal:
		return "mal"
	case Mar:
		return "mar"
	case Mkd:
		return "mkd"
	case Mlg:
		return "mlg"
	case Mya:
		return "mya"
	case Nep:
		return "nep"
	case Nld:
		return "nld"
	case Nno:
		return "nno"
	case Nob:
		return "nob"
	case Nya:
		return "nya"
	case Ori:
		return "ori"
	case Orm:
		return "orm"
	case Pan:
		return "pan"
	case Pes:
		return "pes"
	case Pol:
		return "pol"
	case Por:
		return "por"
	case Ron:
		return "ron"
	case Run:
		return "run"
	case Rus:
		return "rus"
	case Sin:
		return "sin"
	case Skr:
		return "skr"
	case Slv:
		return "slv"
	case Sna:
		return "sna"
	case Som:
		return "som"
	case Spa:
		return "spa"
	case Srp:
		return "srp"
	case Swe:
		return "swe"
	case Tam:
		return "tam"
	case Tel:
		return "tel"
	case Tgl:
		return "tgl"
	case Tha:
		return "tha"
	case Tir:
		return "tir"
	case Tuk:
		return "tuk"
	case Tur:
		return "tur"
	case Uig:
		return "uig"
	case Ukr:
		return "ukr"
	case Urd:
		return "urd"
	case Uzb:
		return "uzb"
	case Vie:
		return "vie"
	case Ydd:
		return "ydd"
	case Yor:
		return "yor"
	case Zul:
		return "zul"
	default:
		return ""
	}
}

//Langs represents a map of Lang to language name.
var Langs = map[Lang]string{
	Aka: "Akan",
	Amh: "Amharic",
	Arb: "Arabic",
	Azj: "Azerbaijani",
	Bel: "Belarusian",
	Ben: "Bengali",
	Bho: "Bhojpuri",
	Bul: "Bulgarian",
	Ceb: "Cebuano",
	Ces: "Czech",
	Cmn: "Mandarin",
	Dan: "Danish",
	Deu: "German",
	Ell: "Greek",
	Eng: "English",
	Epo: "Esperanto",
	Est: "Estonian",
	Fin: "Finnish",
	Fra: "French",
	Guj: "Gujarati",
	Hat: "Haitian Creole",
	Hau: "Hausa",
	Heb: "Hebrew",
	Hin: "Hindi",
	Hrv: "Croatian",
	Hun: "Hungarian",
	Ibo: "Igbo",
	Ilo: "Ilocano",
	Ind: "Indonesian",
	Ita: "Italian",
	Jav: "Javanese",
	Jpn: "Japanese",
	Kan: "Kannada",
	Kat: "Georgian",
	Khm: "Khmer",
	Kin: "Kinyarwanda",
	Kor: "Korean",
	Kur: "Kurdish",
	Lav: "Latvian",
	Lit: "Lithuanian",
	Mai: "Maithili",
	Mal: "Malayalam",
	Mar: "Marathi",
	Mkd: "Macedonian",
	Mlg: "Malagasy",
	Mya: "Burmese",
	Nep: "Nepali",
	Nld: "Dutch",
	Nno: "Nynorsk",
	Nob: "Bokmal",
	Nya: "Chewa",
	Ori: "Oriya",
	Orm: "Oromo",
	Pan: "Punjabi",
	Pes: "Persian",
	Pol: "Polish",
	Por: "Portuguese",
	Ron: "Romanian",
	Run: "Rundi",
	Rus: "Russian",
	Sin: "Sinhalese",
	Skr: "Saraiki",
	Slv: "Slovene",
	Sna: "Shona",
	Som: "Somali",
	Spa: "Spanish",
	Srp: "Serbian",
	Swe: "Swedish",
	Tam: "Tamil",
	Tel: "Telugu",
	Tgl: "Tagalog",
	Tha: "Thai",
	Tir: "Tigrinya",
	Tuk: "Turkmen",
	Tur: "Turkish",
	Uig: "Uyghur",
	Ukr: "Ukranian",
	Urd: "Urdu",
	Uzb: "Uzbek",
	Vie: "Vietnamese",
	Ydd: "Yiddish",
	Yor: "Yoruba",
	Zul: "Zulu",
}

//langProfileList ...
type langProfileList map[Lang][]string

//LatinLangs ...
var latinLangs = langProfileList{
	Spa: []string{" de", "os ", "de ", " la", "la ", " y ", " a ", "es ", "ón ", "ión", "rec", "ere", "der", " co", "e l", "el ", "en ", "ien", "cho", "ent", "ech", "ció", "aci", "o a", "a p", " el", "a l", "al ", "as ", "e d", " en", "na ", "ona", "s d", "da ", "nte", " to", "ad ", "ene", "con", " pr", " su", "tod", " se", "ho ", "los", " pe", "per", "ers", " lo", "o d", " ti", "cia", "n d", "cio", " es", "ida", "res", "a t", "tie", "ion", "rso", "te ", "do ", " in", "son", " re", " li", "to ", "dad", "tad", "e s", "est", "pro", "que", "men", " po", "a e", "oda", "nci", " qu", " un", "ue ", "ne ", "n e", "s y", "lib", "su ", " na", "s e", "nac", "ia ", "e e", "tra", " pa", "or ", "ado", "a d", "nes", "ra ", "se ", "ual", "a c", "er ", "por", "com", "nal", "rta", "a s", "ber", " o ", "one", "s p", "dos", "rá ", "sta", "les", "des", "ibe", "ser", "era", "ar ", "ert", "ter", " di", "ale", "l d", "nto", "hos", "del", "ica", "a a", "s n", "n c", "oci", "imi", "io ", "o e", "re ", "y l", "e c", "ant", "cci", " as", "las", "par", "ame", " cu", "ici", "ara", "enc", "s t", "ndi", " so", "o s", "mie", "tos", "una", "bre", "dic", "cla", "s l", "e a", "l p", "pre", "ntr", "o t", "ial", "y a", "nid", "n p", "a y", "man", "omo", "so ", "n l", " al", "ali", "s a", "no ", " ig", "s s", "e p", "nta", "uma", "ten", "gua", "ade", "y e", "soc", "mo ", " fu", "igu", "o p", "n t", "hum", "d d", "ran", "ria", "y d", "ada", "tiv", "l e", "cas", " ca", "vid", "l t", "s c", "ido", "das", "dis", "s i", " hu", "s o", "nad", "fun", " ma", "rac", "nda", "eli", "sar", "und", " ac", "uni", "mbr", "a u", "die", "e i", "qui", "a i", " ha", "lar", " tr", "odo", "ca ", "tic", "o y", "cti", "lid", "ori", "ndo", "ari", " me", "ta ", "ind", "esa", "cua", "un ", "ier", "tal", "esp", "seg", "ele", "ons", "ito", "ont", "iva", "s h", "d y", "nos", "ist", "rse", " le", "cie", "ide", "edi", "ecc", "ios", "l m", "r e", "med", "tor", "sti", "n a", "rim", "uie", "ple", "tri", "ibr", "sus", "lo ", "ect", "pen", "y c", "an ", "e h", "n s", "ern", "tar", "l y", "egu", "gur", "ura", "int", "ond", "mat", "l r", "r a", "isf", "ote"},
	Eng: []string{" th", "the", " an", "he ", "nd ", "and", "ion", " of", "of ", "tio", " to", "to ", "on ", " in", "al ", "ati", "igh", "ght", "rig", " ri", "or ", "ent", "as ", "ed ", "is ", "ll ", "in ", " be", "e r", "ne ", "one", "ver", "all", "s t", "eve", "t t", " fr", "s a", " ha", " re", "ty ", "ery", " or", "d t", " pr", "ht ", " co", " ev", "e h", "e a", "ng ", "ts ", "his", "ing", "be ", "yon", " sh", "ce ", "ree", "fre", "ryo", "n t", "her", "men", "nat", "sha", "pro", "nal", "y a", "has", "es ", "for", " hi", "hal", "f t", "n a", "n o", "nt ", " pe", "s o", " fo", "d i", "nce", "er ", "ons", "res", "e s", "ect", "ity", "ly ", "l b", "ry ", "e e", "ers", "e i", "an ", "e o", " de", "cti", "dom", "edo", "eed", "hts", "ter", "ona", "re ", " no", " wh", " a ", " un", "d f", " as", "ny ", "l a", "e p", "ere", " en", " na", " wi", "nit", "nte", "d a", "any", "ted", " di", "ns ", "sta", "th ", "per", "ith", "e t", "st ", "e c", "y t", "om ", "soc", " ar", "ch ", "t o", "d o", "nti", "s e", "equ", "ve ", "oci", "man", " fu", "ote", "oth", "ess", " al", " ac", "wit", "ial", " ma", "uni", " se", "rea", " so", " on", "lit", "int", "r t", "y o", "enc", "thi", "ual", "t a", " eq", "tat", "qua", "ive", " st", "ali", "e w", "l o", "are", "f h", "con", "te ", "led", " is", "und", "cia", "e f", "le ", " la", "y i", "uma", "by ", " by", "hum", "f a", "ic ", " hu", "ave", "ge ", "r a", " wo", "o a", "ms ", "com", " me", "eas", "s d", "tec", " li", "n e", "en ", "rat", "tit", "ple", "whe", "ate", "o t", "s r", "t f", "rot", " ch", "cie", "dis", "age", "ary", "o o", "anc", "eli", "no ", " fa", " su", "son", "inc", "at ", "nda", "hou", "wor", "t i", "nde", "rom", "oms", " ot", "g t", "eme", "tle", "iti", "gni", "s w", "itl", "duc", "d w", "whi", "act", "hic", "aw ", "law", " he", "ich", "min", "imi", "ort", "o s", "se ", "e b", "ntr", "tra", "edu", "oun", "tan", "e d", "nst", "l p", "d n", "ld ", "nta", "s i", "ble", "n p", " pu", "n s", " at", "ily", "rth", "tho", "ful", "ssi", "der", "o e", "cat", "uca", "unt", "ien", " ed", "o p", "h a", "era", "ind", "pen", "sec", "n w", "omm", "r s"},
	Por: []string{"os ", "de ", " de", " a ", " e ", "o d", "to ", "ão ", " di", "ent", "da ", "ito", "em ", " co", "eit", "as ", "dir", "es ", "ire", "rei", " se", "ção", "ade", "a p", "dad", "e d", "s d", "men", "nte", "do ", "s e", " pr", " pe", "dos", " to", " da", "a a", "o e", " o ", "o a", "ess", "con", "tod", "que", " qu", "te ", "e a", " do", "al ", "res", "ida", "m d", " in", " ou", "er ", "sso", " na", " re", " po", "a s", " li", "uma", "cia", "ar ", "pro", "e e", "a d", " te", "açã", "a t", " es", " su", "ou ", "ue ", "s p", "tos", "a e", "des", "ra ", "com", "no ", "ame", "ia ", "e p", "tem", "nto", " pa", "is ", "est", "tra", "ões", "na ", "s o", "oda", "das", "ser", "soa", "s n", "pes", "o p", "s a", "o s", "e o", " em", " as", " à ", "o o", "ais", "ber", "ado", "oa ", "o t", "e s", "man", "sua", "ua ", " no", " os", "a c", "ter", "çõe", "erd", "lib", "rda", "s s", "nci", "ibe", "e n", "ica", "odo", "so ", "nal", "ntr", "s t", "hum", "ura", " ao", "ona", "ual", " so", "or ", "ma ", "sta", "o c", "a n", "pre", "ara", "era", "ons", "e t", "r a", "par", "o à", " hu", "ind", "por", "cio", "ria", "m a", "s c", " um", "a l", "gua", "ran", " en", "ndi", "o i", "e c", "raç", "ion", "nid", "aci", "ano", "soc", "e r", "oci", " ac", "und", "sen", "nos", "nsi", "rec", "ime", "ali", "int", "um ", "per", "nac", " al", "m o", "r p", " fu", "ndo", "ont", "açõ", " ig", "igu", "fun", "nta", " ma", "uni", "cçã", "ere", " ex", "a i", " me", "ese", "rio", "l d", "a o", "s h", "pel", "ada", "pri", "ide", "am ", "m p", "pod", "s f", "ém ", "a f", "io ", "ode", "ca ", "ita", "lid", "tiv", "e f", "vid", "r e", "esp", "nda", "omo", "e l", "naç", "o r", "ant", "a q", "tad", "lic", "iva", " fa", "ver", "s l", "ial", "cla", "ngu", "ing", " ca", "mo ", "der", " vi", "eli", "ist", "ta ", "se ", "ati", "ios", "ido", "r o", "eci", "dis", " un", "e i", "r d", "ecç", "o q", "s i", "qua", "ênc", "a m", "seu", "sti", "nin", "uer", "rar", "cas", "aos", "ens", "gué", "ias", "sid", "uém", "tur", "dam", "sse", "ao ", "ela", "l e", "for", "tec", "ote", " pl", "ena", " tr", "m c", "tro", " ni", "ico", "rot"},
	Ind: []string{"an ", "ang", " da", "ng ", " pe", "ak ", " ke", " me", "ata", " se", "dan", "kan", " di", " be", "hak", "ber", "per", "ran", "nga", "yan", "eng", " ya", " ha", "asa", "gan", "men", "ara", "nya", "n p", "n d", "n k", "a d", "tan", " at", "at ", "ora", "ala", "san", " ba", "ap ", "erh", "n b", "rha", "ya ", " ma", "g b", "a s", "pen", "eba", "as ", "aan", "uk ", "ntu", " or", "eti", "tas", "aka", "tia", "ban", "set", " un", "n s", "ter", "n y", " te", "k m", "tuk", "bas", "iap", "lam", "beb", "am ", " de", "k a", "keb", "n m", "i d", "unt", "ama", "dal", "ah ", "ika", "dak", "ebe", "p o", "sa ", "pun", "mem", "n h", "end", "den", "ra ", "ela", "ri ", "nda", " sa", "di ", "ma ", "a m", "n t", "k d", "n a", "ngg", "tau", "man", "gar", "eri", "asi", " ti", "un ", "al ", "ada", "um ", "a p", "lak", "ari", "au ", " ne", "neg", "a b", "ngs", "ta ", "ole", "leh", "ert", "ers", "ida", "k h", "ana", "gsa", "dar", "uka", "tid", "bat", "sia", "era", "eh ", "dap", "ila", "dil", "h d", "atu", "sam", "ia ", "i m", " in", "lan", "aha", "uan", "tu ", "ai ", "t d", "a a", "g d", "har", "sem", "na ", "apa", "ser", "ena", "kat", "uat", "erb", "erl", "mas", "rta", "ega", "ung", "nan", "emp", "n u", "kum", "l d", "g s", " hu", "ka ", "ent", "pat", "mba", "aga", "nta", "adi", " su", "eni", "uku", "n i", "huk", "ind", "ar ", "rga", "i s", "aku", "ndi", "sua", "ni ", "rus", "han", "si ", "car", "nny", " la", "in ", "u d", "ik ", "ua ", "lah", "rik", "usi", "emb", "ann", "mer", "ian", "gga", "lai", "min", "a u", "lua", "ema", "emu", "arg", "dun", "dip", "a t", "mat", "aya", "rbu", "aru", "erk", "rka", "ini", "eka", "a k", "rak", "kes", "yat", "iba", "nas", "rma", "ern", "ese", "s p", "nus", " pu", "anu", "ina", " ta", "mel", "mua", "kel", "k s", "us ", "ndu", "nak", "da ", "sya", "das", "pem", "lin", "ut ", "yar", "ami", "upu", "seo", "aik", "eor", "iny", "aup", "tak", "ipe", "ing", "tin", " an", "dik", "uar", "ili", "g t", "rse", "sar", "ant", "g p", "a n", "aks", "ain", " ja", "t p", " um", "g m", "dir", "ksa", "umu", "kep", "mum", "i k", "eca", "rat", "m p", "h p", "aba", "ses", "m m"},
	Fra: []string{" de", "es ", "de ", "ion", "nt ", "et ", "tio", " et", "ent", " la", "la ", "e d", "on ", "ne ", "oit", "e l", "le ", " le", "s d", "e p", "t d", "ati", "roi", " dr", "dro", "it ", " à ", " co", "té ", "ns ", "te ", "e s", "men", "re ", " to", "con", " l’", "tou", "que", " qu", "les", " so", "des", "son", " pe", "ons", " un", "s l", "s e", " pr", "ue ", " pa", "e c", "t l", "ts ", "onn", " au", "e a", "eme", "e e", " li", "ont", "ant", "out", "ute", "t à", "res", "ers", " sa", "ce ", " a ", "tre", "per", "a d", "cti", "er ", "lib", "ité", " en", "ux ", " re", "en ", "rso", "à l", " ou", " in", "lle", "un ", "nat", "ou ", "nne", "n d", "une", " d’", " se", "par", "nte", "us ", "ur ", "s s", "ans", "dan", "a p", "r l", "pro", "its", "és ", "t p", "ire", "e t", "s p", "sa ", " dé", "ond", "é d", "a l", "nce", "ert", "aux", "omm", "nal", "me ", " na", " fo", "iqu", " ce", "rté", "ect", "ale", "ber", "t a", "s a", " da", "mme", "ibe", "san", "e r", " po", "com", "al ", "s c", "qui", "our", "t e", " ne", "e n", "ous", "r d", "ali", "ter", " di", "fon", "e o", "au ", " ch", "air", "ui ", "ell", " es", "lit", "s n", "iss", "éra", "tes", "soc", "aut", "oci", "êtr", "ien", "int", "du ", "est", "été", "tra", "pou", " pl", "rat", "ar ", "ran", "rai", "s o", "ona", "ain", "cla", "éga", "anc", "rs ", "eur", "pri", "n c", "e m", "s t", "à u", " do", "ure", "bre", "ut ", " êt", "age", " ét", "nsi", "sur", "ein", "sen", "ser", "ndi", "ens", "ess", "ntr", "ir ", " ma", "cia", "n p", "st ", "a c", " du", "l e", " su", "bli", "ge ", "rés", " ré", "e q", "ass", "nda", "peu", "ée ", "l’a", " te", "a s", "tat", "il ", "tés", "ais", "u d", "ine", "ind", "é e", "qu’", " ac", "s i", "n t", "t c", "n a", "l’h", "t q", "soi", "t s", "cun", "rit", " ég", "oir", "’en", "nta", "hom", " on", "n e", " mo", "ie ", "ign", "rel", "nna", "t i", "l n", " tr", "ill", "ple", "s é", "l’e", "rec", "a r", "ote", "sse", "uni", "idé", "ive", "s u", "t ê", "ins", "act", " fa", "n s", " vi", "gal", " as", "lig", "ssa", "pré", "leu", "e f", "lic", "dis", "ver", " nu", "ten", "ssi", "rot", "tec", "s m", "abl"},
	Deu: []string{"en ", "er ", "der", " un", "nd ", "und", "ein", "ung", "cht", " de", "ich", "sch", "ng ", " ge", "ie ", "che", "ech", " di", "die", "rec", "gen", "ine", "eit", " re", "ch ", " da", "n d", "ver", "hen", " zu", "t d", " au", "ht ", " ha", "lic", "it ", "ten", "rei", " be", "in ", " ve", " in", " ei", "nde", "auf", "den", "ede", "zu ", "n s", "uf ", "fre", "ne ", "ter", "es ", " je", "jed", "n u", " an", "sei", "and", " fr", "run", "at ", " se", "e u", "das", "hei", "s r", "hte", "hat", "nsc", "nge", "r h", "as ", "ens", " al", "ere", "lle", "t a", " we", "n g", "rde", "nte", "ese", "men", " od", "ode", "ner", "g d", "all", "t u", "ers", "te ", "nen", " so", "d d", "n a", "ben", "lei", " gr", " vo", "wer", "e a", "ege", "ion", " st", "ige", "le ", "cha", " me", "haf", "aft", "n j", "ren", " er", "erk", "ent", "bei", " si", "eih", "ihe", "kei", "erd", "tig", "n i", "on ", "lun", "r d", "len", "gem", "ies", "gru", "tli", "unt", "chu", "ern", "ges", "end", "e s", "ft ", "st ", "ist", "tio", "ati", " gl", "sta", "gun", "mit", "sen", "n n", " na", "n z", "ite", " wi", "r g", "eic", "e e", "ei ", "lie", "r s", "n w", "gle", "mei", "de ", "uch", "em ", "chl", "nat", "rch", "t w", "des", "n e", "hre", "ale", "spr", "d f", "ach", "sse", "r e", " sc", "urc", "r m", "nie", "e f", "fen", "e g", "e d", " ni", "dur", "dar", "int", " du", "geh", "ied", "t s", " mi", "alt", "her", "hab", "f g", "sic", "ste", "taa", "aat", "he ", "ang", "ruc", "hli", "tz ", "eme", "abe", "h a", "n v", "nun", "geg", "arf", "rf ", "ehe", "pru", " is", "erf", "e m", "ans", "ndl", "e b", "tun", "n o", "d g", "n r", "r v", "wie", "ber", "r a", "arb", "bes", "t i", "h d", "r w", "r b", " ih", "d s", "igk", "gke", "nsp", "dig", "ema", "ell", "eru", "n f", "ins", "rbe", "ffe", "esc", "igu", "ger", "str", "ken", "e v", "gew", "han", "ind", "rt ", " ar", "ieß", "n h", "rn ", "man", "r i", "hut", "utz", "d a", "ls ", "ebe", "von", "lte", "r o", "rli", "etz", "tra", "aus", "det", "hul", "e i", "one", "nne", "isc", "son", "sel", "et ", "ohn", "t g", "sam", " fa", "rst", "rkl", "ser", "iem", "g v", "t z", "err"},
	Jav: []string{"ng ", "an ", "ang", " ka", "ing", "kan", " sa", "ak ", "lan", " la", "hak", " ha", " pa", " ma", "ngg", "ara", "sa ", "abe", "ne ", " in", "n k", "ant", " ng", "tan", "nin", " an", "nga", "ata", "en ", "ran", " ba", "man", "ban", "ane", "hi ", "n u", "ong", "ra ", "nth", "ake", "ke ", "thi", " da", "won", "uwo", "ung", "ngs", " uw", "asa", "gsa", "ben", "sab", "ana", "aka", "beb", "a k", "g p", "nan", "nda", "adi", "at ", "awa", "san", "ni ", "dan", "g k", "pan", "eba", " be", "e k", "g s", "ani", "bas", " pr", "dha", "aya", "gan", "ya ", "wa ", "di ", "mar", "n s", " wa", "ta ", "a s", "g u", " na", "e h", "arb", "a n", "a b", "a l", "n n", " ut", "yan", "n p", "asi", "g d", "han", "ah ", "g n", " tu", " um", "as ", "wen", "dak", "rbe", "dar", " di", "ggo", "sar", "mat", "k h", "a a", "iya", " un", "und", "eni", "kab", "be ", "art", "ka ", "uma", "ora", "n b", "ala", "n m", "ngk", "rta", "i h", " or", "gar", "yat", "kar", "al ", "a m", "n i", "na ", "g b", "ega", "pra", "ina", "kak", "g a", "a p", "tum", "nya", "kal", "ger", "gge", " ta", "kat", "i k", "ena", "oni", "kas", " pe", "dad", "aga", "g m", "duw", "k k", "uta", "uwe", " si", " ne", "adh", "pa ", "n a", "go ", "and", "i l", " ke", "nun", "nal", "ngu", "uju", "apa", "a d", "t m", "i p", "min", "iba", "er ", " li", "anu", "sak", "per", "ama", "gay", "war", "pad", "ggu", "ha ", "ind", "taw", "ras", "n l", "ali", "eng", "awi", "a u", " bi", "we ", "bad", "ndu", "uwa", "awe", "bak", "ase", "eh ", " me", "neg", "pri", " ku", "ron", "ih ", "g t", "bis", "iji", "i t", "e p", " pi", "aba", "isa", "mba", "ini", "a w", "g l", "ika", "n t", "ebu", "ndh", "ar ", "sin", "lak", "ur ", "mra", "men", "ku ", " we", "e s", "a i", "liy", " ik", "ayo", "rib", "ngl", "ami", "arg", "nas", "yom", "wae", "ut ", "kon", "ae ", "rap", "aku", " te", "dil", "tin", "rga", "jud", "umu", " as", "rak", "bed", "k b", "il ", "kap", "h k", "jin", "k a", " nd", "e d", "i s", " lu", "i w", "eka", "mum", "um ", "uha", "ate", " mi", "k p", "gon", "eda", " ti", "but", "n d", "r k", "ona", "uto", "tow", "wat", "gka", "si ", "umr", "k l", "oma"},
	Vie: []string{"ng ", "̣c ", "́c ", " qu", " th", "à ", "nh ", " ng", "̣i ", " nh", "và", " va", "̀n ", "uyê", " ph", " ca", "quy", "ền", "yề", "̀i ", " ch", "̀nh", " tr", " cu", "ngư", "i n", "gươ", "ườ", "́t ", "ời", " gi", "ác", " co", "̣t ", "ó ", "c t", "ự ", "n t", "cá", "ông", " kh", "ượ", "ợc", " tư", " đư", "iệ", "đươ", "ìn", "́i ", " ha", "có", "i đ", "gia", " đê", "pha", " mo", "ọi", "mọ", "như", "n n", "củ", " ba", "̣n ", "̉a ", "ủa", "n c", "̀u ", "̃ng", "ân ", "ều", "ất", " bi", "tự", "hôn", " vi", "g t", " la", "n đ", "đề", "nhâ", " ti", "t c", " đô", "ên ", "bả", "hiê", "u c", " tô", "do ", "hân", " do", "ch ", "́ q", "̀ t", " na", "́n ", "ay ", " hi", "àn", "̣ d", "ới", "há", " đi", "hay", "g n", " mô", "ốc", "uố", "n v", "ội", "hữ", "thư", "́p ", "quô", " ho", "̣p ", "nà", "ào", "̀ng", "̉n ", "ị ", "́ch", "ôn ", "̀o ", "khô", "c h", "i c", "c đ", " hô", "i v", "tro", " đa", "́ng", "mộ", "i t", "ột", "g v", "ia ", "̣ng", "ản", "ướ", "ữn", "̉ng", "h t", "hư ", "ện", "n b", "ộc", "ả ", "là", "c c", "g c", " đo", "̉ c", "n h", "hà", "hộ", " bâ", "ã ", "̀y ", " vơ", "̣ t", "̉i ", "iế", " cô", "t t", "g đ", "ức", "iên", " vê", "viê", "vớ", "h v", "ớc", "ực", "ật", "tha", "̉m ", "ron", "ong", "áp", "g b", "hươ", " sư", "a c", "sự", "̉o ", "ảo", "h c", "ể ", "o v", "uậ", "a m", "ế ", "iá", "̀ c", "cho", "qua", "hạ", "ục", " mi", "̀ n", "phâ", "c q", "côn", "o c", "á ", "i h", "ại", " hơ", "̃ h", " cư", "n l", "bị", " lu", "bấ", "cả", "ín", "h đ", " xa", "độ", "g h", "c n", "c p", "thu", "ải", "ệ ", " hư", "́ c", "o n", " nư", "ốn", "́o ", "áo", "xã", "oà", "y t", "hả", "tộ", "̣ c", " tâ", "thô", " du", "m v", "mì", "ho ", "hứ", "ệc", "́ t", "hợ", "án", "n p", "cũ", "ũn", "iể", "ối", "tiê", "ề ", "hấ", "ợp", "hoa", "y đ", "chi", "o h", "ở ", "ày", "̉ t", "đó", "c l", "về", "̀ đ", "i b", "kha", "c b", " đâ", "luâ", "ai ", "̉ n", "đố", "ết", "hự", "tri", "p q", "nươ", "dụ", "hí", "g q", "yên", "họ", "́nh", " ta", " bă", "c g", "n g", "thê", "o t", "c v", "am ", "c m", "an "},
	Ita: []string{" di", "to ", " de", "ion", " in", "la ", "e d", "di ", "ne ", " e ", "zio", "re ", "le ", "ni ", "ell", "one", "lla", "rit", "a d", "o d", "del", "itt", "iri", "dir", " co", "ti ", "ess", "ent", " al", "azi", "tto", "te ", "i d", "i i", "ere", "tà ", " pr", "ndi", "e l", "ale", "o a", "ind", "e e", "e i", "gni", "nte", "con", "i e", "li ", "a s", " un", "men", "ogn", " ne", "uo ", " og", "idu", "e a", "ivi", "duo", "vid", " es", "tti", " ha", "div", " li", "a p", "no ", "all", "pro", "za ", "ato", "per", "sse", "ser", " so", "i s", " la", " su", "e p", " pe", "ibe", "na ", "a l", " il", "ber", "e n", "il ", "ali", "lib", "ha ", "che", "in ", "o s", "e s", " qu", "o e", "ia ", "e c", " ri", "nza", "ta ", "nto", "he ", "oni", "o i", " o ", "sta", "o c", "nel", " a ", "o p", "naz", "e o", "so ", " po", "o h", "gli", "i u", "ond", "i c", "ers", "ame", "i p", "lle", "un ", "era", "ri ", "ver", "ro ", "el ", "una", "a c", " ch", "ert", "ua ", "i a", "ssi", "rtà", "a e", "ei ", "dis", "ant", " l ", "tat", "a a", "ona", "ual", " le", "ità", "are", "ter", " ad", "nit", " da", "pri", "dei", "à e", "cia", " st", " si", "nal", "est", "tut", "ist", "com", "uni", " ed", "ono", " na", "sua", "al ", "si ", "anz", " pa", " re", "raz", "gua", "ita", "res", "der", "soc", "man", "o o", "ad ", "i o", "ese", "que", "enz", "ed ", " se", "io ", "ett", "on ", " tu", "dic", "à d", "sia", "i r", "rso", "oci", "rio", "ari", "qua", "ial", "pre", "ich", "rat", "ien", "tra", "ani", "uma", "se ", "ll ", "eri", "a n", "o n", " um", "do ", "ara", "a t", "zza", "er ", "tri", "att", "ico", "pos", "sci", "i l", "son", "nda", "par", "e u", "fon", " fo", "nti", "uzi", "str", "utt", "ati", "sen", "int", "nes", "iar", " i ", "hia", "n c", "sti", "chi", "ann", "ra ", " eg", "egu", "isp", "bil", "ont", "a r", " no", "rop", " me", "opr", "ost", " ma", "ues", "ica", "sso", "tal", "cie", "sun", "lit", "ore", "ina", "ite", "tan", " ra", "non", "gio", "d a", "e r", "dev", "i m", "l i", "ezz", "izi", " cu", "nno", "rà ", "a i", "tta", "ria", "lia", "cos", "ssu", "dal", "l p", " as", "ass", "opo", "ve ", "eve"},
	Tur: []string{" ve", " ha", "ve ", "ler", "lar", "ir ", "in ", "hak", " he", "her", "bir", "er ", "an ", "arı", "eri", "ya ", " bi", "ak ", "r h", "eti", "ın ", "iye", "yet", " ka", "ası", "ını", " ol", "tle", "eya", "kkı", "ara", "akk", "etl", "sın", "esi", "na ", "de ", "ek ", " ta", "nda", "ini", " bu", "ile", "rın", "rin", "vey", "ne ", "kla", "e h", "ine", "ır ", "ere", "ama", "dır", "n h", " sa", "ına", "sin", "e k", "le ", " ge", "mas", "ınd", "nın", "ı v", " va", "lan", "lma", "erk", "rke", "nma", "tin", "rle", " te", "nin", "akl", "a v", "da ", " de", "let", "ill", "e m", "ard", "en ", "riy", "aya", "nı ", " hü", " şa", "e b", "k v", "kın", "k h", " me", "mil", "san", " il", "si ", "rdı", "e d", "dan", "hür", "var", "ana", "e a", "kes", "et ", "mes", "şah", "dir", " mi", "ret", "rri", " se", "ola", "ürr", "irl", "bu ", "mak", " ma", "mek", "n e", "kı ", "n v", "n i", "lik", "lle", " ed", " hi", "n b", "a h", " ba", "nsa", " iş", "eli", "kar", " iç", "ı h", "ala", "li ", "ulu", "rak", "evl", "e i", "ni ", "re ", "r ş", "eme", "etm", "e t", "ik ", "e s", "a b", "iş ", "n k", "hai", "nde", "aiz", " eş", "izd", "un ", "olm", "hiç", "zdi", "ar ", "unm", "ma ", " gö", "ilm", "lme", "im ", "n t", "tir", "dil", "mal", "e g", "i v", " ko", "lun", "e e", "mel", "ket", "ık ", "n s", "ele", "la ", "el ", "r v", "ede", "şit", "ili", "eşi", "yla", "a i", " an", "anı", " et", "rı ", "ahs", " ya", "sı ", "edi", "siy", "t v", "i b", "se ", "içi", "çin", "bul", "ame", " da", "miş", "may", "tim", "a k", "tme", "r b", "ins", "yan", "nla", "mle", " di", "eye", "ger", "ye ", "uğu", "erd", "din", "ser", " mü", "mem", "vle", " ke", "nam", "ind", "len", "eke", "es ", " ki", "n m", "it ", " in", " ku", "rşı", "a s", "arş", " ay", "eml", "lek", "oru", "rme", "kor", "rde", "i m", " so", "tür", "al ", "lam", "eni", "nun", " uy", "ken", "hsı", "i i", "a d", "ri ", "dev", "ün ", "a m", "r a", "mey", "cak", "ıyl", "maz", "e v", "ece", "ade", "iç ", "şma", "mse", "te ", "tün", "ims", "kim", "e y", "şı ", "end", "k g", "ndi", "alı", " ce", "lem", "öğr", "ütü", "k i", "r t", " öğ", "büt", "anl", " bü"},
	Pol: []string{" pr", "nie", " i ", "ie ", "pra", " po", "ani", "raw", "ia ", "nia", "wie", "go ", " do", "ch ", "ego", "iek", "owi", " ni", "ści", "ci ", "a p", "do ", "awo", " cz", "ośc", "ych", " ma", "ek ", "rze", " na", "prz", " w ", "wo ", "ej ", " za", "noś", "czł", "zło", "eni", "wa ", " je", "łow", "i p", "wol", "oln", " lu", "rod", " ka", " wo", "lno", "wsz", "y c", "ma ", "ny ", "każ", "ażd", "o d", "stw", "owa", "dy ", "żdy", " wy", "rzy", "sta", "ecz", " sw", "dzi", "i w", "e p", "czn", "twa", "na ", "zys", "ów ", "szy", "ub ", "lub", "a w", "est", "kie", "k m", "wan", " sp", "ają", " ws", "e w", "pow", "pos", "nyc", "rac", "spo", "ać ", "a i", "cze", "sze", "neg", "yst", "jak", " ja", "o p", "pod", "acj", "ne ", "ńst", "aro", "mi ", " z ", "i i", "nar", " ko", "obo", "awa", " ro", "i n", "jąc", "zec", "zne", "zan", "dow", " ró", "iej", "zy ", "zen", "nic", "ony", "aw ", "i z", "czy", "no ", "nej", "o s", "rów", "odn", "cy ", "ówn", "odz", "o w", "o z", "jeg", "edn", "o o", "aki", "mie", "ien", "kol", " in", "zie", "bez", "ami", "eńs", "owo", "dno", " ob", " or", " st", "a s", "ni ", "orz", "o u", "ym ", "stę", "tęp", "łec", "jed", "i k", " os", "w c", "lwi", "ez ", "olw", "ołe", "poł", "cji", "y w", "o n", "wia", " be", "któ", "a j", "zna", "zyn", "owe", "wob", "ka ", "wyc", "owy", "ji ", " od", "aln", "inn", "jes", "icz", "h p", "i s", "się", "a o", "ją ", "ost", "kra", "st ", "sza", "swo", "war", "cza", "roz", "y s", "raz", "nik", "ara", "ora", "lud", "i o", "a z", "zes", " kr", "ran", "ows", "ech", "w p", "dów", "ą p", "pop", "a n", "tki", "stk", "gan", "zon", "raj", "e o", "iec", "i l", " si", "że ", "eka", " kt", " de", "em ", "tór", "ię ", "wni", "lni", "ejs", "ini", "odo", "dni", "ełn", "kow", "peł", "a d", "ron", "dek", "pie", "udz", "bod", "nan", "h i", "dst", "ieg", "taw", "z p", "z w", "zeń", "god", "iu ", "ano", "lar", " to", "y z", "a k", "ale", "kla", "trz", "zaw", "ich", "e i", "ier", "iko", "dzy", "chn", "w z", "by ", "ków", "adz", "ekl", "ywa", "ju ", "och", "kor", "sob", "ocz", "oso", "u p", "du ", "tyc", "tan", "ędz", " mi", "e s", " ta", "ki "},
	Orm: []string{"aa ", "an ", "uu ", " ka", "ni ", "aan", "umm", "ii ", "mma", "maa", " wa", "ti ", "nam", " fi", "ta ", "tti", " na", "saa", "fi ", " mi", "rga", "i k", "a n", " qa", "dha", "iyy", "oot", "in ", "mir", "irg", "raa", "qab", "a i", "a k", "kan", "akk", "isa", "chu", "amu", "a f", "huu", "aba", "kka", " ta", "kam", "a a", " is", "amn", "ami", "att", "ach", "mni", "yaa", " bi", "yuu", "yyu", "ee ", "wal", "miy", "waa", "ga ", "ata", "aat", "tii", "oo ", "a e", "moo", " ni", " ee", "ba ", " ak", "ota", "a h", "i q", " ga", " dh", "daa", "haa", "a m", "ama", "yoo", "a b", "i a", "ka ", "kaa", " hi", "sum", "aas", "arg", "man", " hu", " uu", "u n", " yo", " ar", " ke", " ha", "ees", " ba", "uf ", "i i", "taa", "uuf", "iin", "ada", "a w", "i f", "ani", "rra", "na ", "isu", " ad", "i w", "a u", "nya", "irr", "da ", "hun", "hin", "ess", " ho", " ma", "i m", "und", "i b", "bar", "ana", "een", "mu ", "is ", "bu ", "f m", " ir", " sa", "u a", "add", "aad", " la", "i d", "n h", "eeg", "i h", "sa ", "hoj", "abu", " ya", "kee", "al ", "udh", "ook", "goo", "ala", "ira", "nda", "itt", "gac", "as ", "n k", "mum", "see", "rgo", "uum", "ra ", "n t", "n i", "ara", "muu", "ums", "mat", "nii", "sii", "ssa", "a d", "a q", " da", "haw", "a g", "yya", "asu", "eef", "u h", "tum", "biy", " mo", "a t", "ati", "eny", "gam", "abs", "awa", "roo", "uma", "n b", "n m", "u y", "a s", "sat", "baa", "gar", "n a", "mmo", "nis", " qo", "nna", " ku", "eer", " to", "kko", "bil", "ili", "lis", "bir", "otu", "tee", "ya ", "msa", "aaf", "suu", "n d", "jii", "n w", "okk", "rka", "gaa", "ald", "un ", "rum", " ye", "ame", " fu", "mee", "yer", "ero", "amm", "era", "kun", "i y", "oti", "tok", "ant", "ali", "nni", " am", "lda", "lii", "n u", "lee", "ura", "lab", "aal", "tan", "laa", "i g", "ila", "ddu", "aru", "u m", "oji", "gum", "han", "ega", " se", "ffa", "dar", "faa", "ark", "n y", "hii", "qix", "gal", "ndi", " qi", "asa", "art", "ef ", "uud", " bu", "jir", " ji", "arb", "n g", "chi", "tam", "u b", "dda", "bat", "di ", "kar", "lam", "a l", " go", "bsi", "sad", "oka", "a j", "egu", "u t", "bee", "u f", "uun"},
	Ron: []string{" de", "și ", " și", "re ", " în", "are", "te ", "de ", "ea ", "ul ", "rep", "le ", "ept", "dre", "e d", " dr", "ie ", "în ", "e a", "ate", "ptu", " sa", "tul", " pr", "or ", "e p", " pe", "la ", "e s", "ori", " la", " co", "lor", " or", "ii ", "rea", "ce ", "au ", "tat", "ați", " a ", " ca", "ent", " fi", "ale", "ă a", "a s", " ar", "ers", "per", "ice", " li", "uri", "a d", "al ", " re", "e c", "ric", "nă ", "i s", "e o", "ei ", "tur", " să", "lib", "con", "men", "ibe", "ber", "rso", "să ", "tăț", "sau", " ac", "ilo", "pri", "ăți", "i a", "i l", "car", "l l", "ter", " in", "ție", "că ", "soa", "oan", "ții", "lă ", "tea", "ri ", "a p", " al", "ril", "e ș", "ană", "in ", "nal", "pre", "i î", "uni", "ui ", "se ", "e f", "ere", "i d", "e î", "ita", " un", "ert", "ile", "tă ", "a o", " se", "i ș", "pen", "ia ", "ele", "fie", "i c", "a l", "ace", "nte", "ntr", "eni", " că", "ală", " ni", "ire", "ă d", "pro", "est", "a c", " cu", " nu", "n c", "lui", "eri", "ona", " as", "sal", "ând", "naț", "ecu", "i p", "rin", "inț", " su", "ră ", "e n", " om", "ici", "nu ", "i n", "oat", "ări", "l d", " to", "tor", " di", " na", "iun", " po", "oci", "tre", "ni ", "ste", "soc", "ega", "i o", "gal", " so", " tr", "ă p", "a a", "n m", "sta", "va ", "ă î", "fi ", "res", "rec", "ulu", "nic", "din", "sa ", "cla", "nd ", " mo", " ce", " au", "ara", "lit", "int", "i e", "ces", "uie", "at ", "rar", "rel", "iei", "ons", "e e", "leg", "nit", "ă f", " îm", "a î", "act", "e l", "ru ", "u d", "nta", "a f", "ial", "ra ", "ă c", " eg", "ță ", " fa", "i f", "rtă", "tru", "tar", "ți ", "ă ș", "ion", "ntu", "dep", "ame", "i i", "reb", "ect", "ali", "l c", "eme", "nde", "n a", "ite", "ebu", "bui", "ât ", "ili", "toa", "dec", " o ", "pli", "văț", "nt ", "e r", "u c", "ța ", "t î", "l ș", "cu ", "rta", "cia", "ane", "țio", "ca ", "ită", "poa", "cți", "împ", "bil", "r ș", " st", "omu", "ăță", "țiu", "rie", "uma", "mân", " ma", "ani", "nța", "cur", "era", "u a", "tra", "oar", " ex", "t s", "iil", "ta ", "rit", "rot", "mod", "tri", "riv", "od ", "lic", "rii", "eze", "man", "înv", "ne ", "nvă", "a ș", "cti"},
	Hau: []string{"da ", " da", "in ", "a k", "ya ", "an ", "a d", "a a", " ya", " ko", " wa", " a ", "sa ", "na ", " ha", "a s", "ta ", "kin", "wan", "wa ", " ta", " ba", "a y", "a h", "n d", "n a", "iya", "ko ", "a t", "ma ", "ar ", " na", "yan", "ba ", " sa", "asa", " za", " ma", "a w", "hak", "ata", " ka", "ama", "akk", "i d", "a m", " mu", "su ", "owa", "a z", "iki", "a b", "nci", " ƙa", " ci", " sh", "ai ", "kow", "anc", "nsa", "a ƙ", "a c", " su", "shi", "ka ", " ku", " ga", "ci ", "ne ", "ani", "e d", "uma", "‘ya", "cik", "kum", "uwa", "ana", " du", " ‘y", "ɗan", "ali", "i k", " yi", "ada", "ƙas", "aka", "kki", "utu", "n y", "a n", "hi ", " ra", "mut", " do", " ad", "tar", " ɗa", "nda", " ab", "man", "a g", "nan", "ars", "and", "cin", "ane", "i a", "yi ", "n k", "min", "sam", "ke ", "a i", "ins", "yin", "ki ", "nin", "aɗa", "ann", "ni ", "tum", "za ", "e m", "ami", "dam", "kan", "yar", "en ", "um ", "n h", "oka", "duk", "mi ", " ja", "ewa", "abi", "kam", "i y", "dai", "mat", "nna", "waɗ", "n s", "ash", "ga ", "kok", "oki", "re ", "am ", "ida", "sar", "awa", "mas", "abu", "uni", "n j", "una", "ra ", "i b", " ƙu", "dun", "a ‘", "cew", "a r", "aba", "ƙun", "ce ", "e s", "a ɗ", "san", "she", "ara", "li ", "kko", "ari", "n w", "m n", "buw", "aik", "u d", "kar", " ai", "niy", " ne", "hal", "rin", "bub", "zam", "omi", " la", "rsa", "ubu", "han", "are", "aya", "a l", "i m", "zai", "ban", "o n", "add", "n m", "i s", " fa", "bin", "r d", "ake", "n ‘", "uns", "sas", "tsa", "dom", " ce", "ans", " hu", "me ", "kiy", "ƙar", " am", "ɗin", " an", "ika", "jam", "i w", "wat", "n t", "yya", "ame", "n ƙ", "abb", "bay", "har", "din", "hen", "dok", "yak", "n b", "nce", "ray", "gan", "fa ", "on ", " ki", "aid", " ts", "rsu", " al", "aye", " id", "n r", "u k", "ili", "nsu", "bba", "aur", "kka", "ayu", "ant", "aci", "dan", "ukk", "ayi", "tun", "aga", "fan", "unc", " lo", "o d", "lok", "sha", "un ", "lin", "kac", "aɗi", "fi ", "gam", "i i", "yuw", "sun", "aif", "aja", " ir", "yay", "imi", "war", " iy", "riy", "ace", "nta", "uka", "o a", "bat", "mar", "bi ", "sak", "n i", " ak", "tab", "afi", "sab"},
	Hrv: []string{" pr", " i ", "je ", "rav", "pra", "ma ", " na", "ima", " sv", "na ", "ti ", "a p", "nje", " po", "a s", "anj", "a i", "vo ", "ko ", "da ", "vat", "va ", "no ", " za", "i s", "o i", "ja ", "avo", " u ", " im", "sva", "i p", " bi", "e s", "ju ", "tko", "o n", "li ", "ili", "van", "ava", " sl", "ih ", "ne ", "ost", " dr", "ije", " ne", "jed", "slo", " ra", "u s", "lob", "obo", " os", "bod", " da", " ko", "ova", "nja", "koj", "i d", "atk", "iti", " il", "stv", "pri", "om ", "im ", " je", " ob", " su", " ka", "i i", "i n", "e i", "vje", "i u", "se ", "dru", "bit", "voj", "ati", "i o", "ćen", "a o", "o p", "a b", "a n", "ući", " se", "enj", "sti", "a u", "edn", "dje", "lo ", "ćav", " mo", "raz", "u p", " od", "ran", "ni ", "rod", "a k", "su ", "aro", "drć", "svo", "ako", "u i", "rća", "a j", "mij", "ji ", "nih", "eni", "e n", "e o", " nj", "pre", "pos", "ćiv", "oje", "eno", "e p", "nar", "oda", "nim", "ovo", "aju", "ra ", "ći ", "og ", "nov", "iva", "a d", "nos", "bra", "bil", "i b", "avn", "a z", "jen", "e d", "ve ", "ora", "tva", "jel", "sta", "mor", "u o", "cij", "pro", "ovi", "za ", "jer", "ka ", "sno", "ilo", "jem", "red", "em ", "lju", "osn", "oji", " iz", "aci", " do", "lje", "i m", " ni", "odn", "nom", "jeg", " dj", "vno", "vim", "elj", "u z", "o d", "rad", "o o", "m i", "du ", "uje", " sa", "nit", "e b", " st", "oj ", "tit", "a ć", "dno", "e u", "o s", "u d", "eću", "ani", "dna", "nak", "nst", "stu", " sm", "e k", "u u", "an ", "gov", "nju", "juć", "aln", "m s", "tu ", "a r", "ćov", "jan", "u n", "o k", "ist", "ću ", "te ", "tvo", "ans", "šti", "nu ", "ara", "nap", "m p", "nić", "olj", "bud", " bu", "edi", "ovj", "i v", "pod", "sam", "obr", "tel", " mi", "ina", "zaš", "e m", "ašt", " vj", "ona", "nji", "jek", " ta", "duć", "ija", " ćo", "tup", "h p", "oja", "smi", "ada", " op", "oso", "una", "sob", "odu", "dni", "rug", "udu", "ao ", "di ", "avi", "tno", "jim", "itu", "itk", "će ", "odr", "ave", "meć", "nog", "din", "svi", " ći", "kak", "kla", "rim", "akv", "elo", "štv", "ite", "vol", "jet", "opć", "pot", "tan", "ak ", "nic", "nac", "uće", " sk", " me", "ven"},
	Nld: []string{"en ", "de ", "an ", " de", "van", " va", " en", " he", "ing", "cht", "der", "ng ", "n d", "n v", "et ", "een", " ge", "ech", "n e", "ver", "rec", "nde", " ee", " re", " be", "ede", "er ", "e v", "gen", "den", "het", "ten", " te", " in", " op", "n i", " ve", "lij", " zi", "ere", "eli", "zij", "ijk", "te ", "oor", "ht ", "ens", "n o", "and", "t o", "ijn", "ied", "ke ", " on", "eid", "op ", " vo", "jn ", "id ", "ond", "in ", "sch", " vr", "aar", "n z", "aan", " ie", "rde", "rij", "men", "ren", "ord", "hei", "hte", " we", "eft", "n g", "ft ", "n w", "or ", "n h", "eef", "vri", "wor", " me", "hee", "al ", "t r", "of ", "le ", " of", "ati", "g v", "e b", "eni", " aa", "lle", " wo", "n a", "e o", "nd ", "r h", "voo", " al", "ege", "n t", "erk", " da", " na", "t h", "sta", "jke", "at ", "nat", "nge", "e e", "end", " st", "om ", "e g", "tie", "n b", "ste", "die", "e r", "erw", "wel", "e s", "r d", " om", "ij ", "dig", "t e", "ige", "ter", "ie ", "gel", "re ", "jhe", "t d", " za", "e m", "ers", "ijh", "nig", "zal", "nie", "d v", "ns ", "d e", "e w", "e n", "est", "ele", "bes", " do", "g e", "che", "vol", "ge ", "eze", "e d", "ig ", "gin", "dat", "hap", "cha", "eke", " di", "ona", "e a", "lke", "nst", "ard", " gr", "tel", "min", " to", "waa", "len", "elk", "lin", "eme", "jk ", "n s", "del", "str", "han", "eve", "gro", "ich", "ven", "doo", " wa", "t v", "it ", "ove", "rin", "aat", "n n", "wet", "uit", "ijd", "ze ", " zo", "ion", " ov", "dez", "gem", "met", "tio", "bbe", "ach", " ni", "hed", "st ", "all", "ies", "per", "heb", "ebb", "e i", "toe", "es ", "taa", "n m", "nte", "ien", "el ", "nin", "ale", "ben", "daa", "sti", " ma", "mee", "kin", "pen", "e h", "wer", "ont", "iet", "tig", "g o", "s e", " er", "igd", "ete", "ang", "lan", "nsc", "ema", "man", "t g", "is ", "beg", "her", "esc", "bij", "d o", "ron", "tin", "nal", "eer", "p v", "edi", "erm", "ite", "t w", "t a", " hu", "rwi", "wij", "ijs", "r e", "weg", "js ", "rmi", "naa", "t b", "app", "rwe", " bi", "t z", "ker", "ame", "eri", "ken", " an", "ar ", " la", "tre", "ger", "rdi", "tan", "eit", "gde", "g i", "d z", "oep"},
	Srp: []string{" pr", " i ", "rav", "pra", " na", "na ", "ma ", " po", "je ", " sv", "da ", "a p", "ima", "ja ", "a i", "vo ", "nje", "va ", "ko ", "anj", "ti ", "i p", " u ", "ako", "a s", " da", "avo", "i s", "ju ", "ost", " za", "sva", "o i", "vak", " im", "e s", "o n", "ava", " sl", "nja", " ko", "no ", "ne ", "li ", "om ", " ne", "ili", " dr", "u s", "slo", "koj", "a n", "obo", "ih ", "lob", "bod", "im ", "sti", "stv", "a o", " bi", " il", " ra", "pri", "a u", "og ", " je", "jed", "e p", "enj", "ni ", "van", "u p", "nos", "a d", "iti", "a k", "edn", "i u", "pro", "o d", "ova", " su", "ran", "cij", "i i", "sta", "se ", " os", "e i", "dru", " ob", "i o", "rod", "aju", "ove", " de", "i n", " ka", "aci", "e o", " ni", " od", "ovo", "i d", "ve ", " se", "eni", "voj", "ija", "su ", "u i", "žav", "avn", "uje", " st", "red", "m i", "dna", "a b", "odi", "ara", "drž", "ji ", "nov", "lju", "e b", "rža", "tva", "što", "u o", "oja", " ov", "a j", "odn", "u u", "jan", "poš", "jen", " nj", "nim", "ka ", "ošt", "du ", "raz", "a z", " iz", "sno", "o p", "vu ", "u n", "u d", "šti", "osn", "e d", "pre", "u z", "de ", "ave", "nih", "bit", "aro", "oji", "bez", "tu ", "gov", "lje", "ičn", " sa", "lja", "svo", "lo ", "za ", "vno", "e n", "eđu", " tr", "nar", " me", "vim", "čno", "oda", "ani", "đen", "nac", "nak", "an ", "to ", "tre", "ašt", " kr", "stu", "nog", "o k", "m s", "tit", "aln", "nom", "oj ", "pos", "e u", "reb", " vr", "olj", "dno", "iko", "ku ", "me ", "nik", " do", "ika", "e k", "jeg", "nst", "tav", "em ", "i m", "sme", "o s", "dni", "bra", "nju", "šen", "ovi", "tan", "te ", "avi", "vol", " li", "zaš", "ilo", "rug", "var", "kao", "ao ", "riv", "tup", "st ", "živ", "ans", "eno", "čov", "štv", "kla", "vre", "bud", "ena", " ve", "ver", "odu", "međ", "oju", "ušt", " bu", "kom", "kri", "pod", "ruš", "m n", "i b", "ba ", "a t", "ugi", "edi", " mo", "la ", "u v", "kak", " sm", "ego", "akv", "o j", "rad", "dst", "jav", "del", "tvo", " op", "nu ", "por", "vlj", "avl", "m p", "od ", "jem", "oje", " čo", "a r", "sam", "i v", "ere", "pot", "o o", "šte", "rem", "vek", "svi", " on", "rot", "e r"},
	Kur: []string{" he", " û ", "ên ", " bi", " ma", "in ", "na ", " di", "maf", "an ", "ku ", " de", " ku", " ji", "xwe", "her", " xw", "iya", "ya ", "kes", "kir", "rin", "iri", " ne", "ji ", "bi ", "yên", "afê", "e b", "de ", "tin", "e h", "iyê", "ke ", "es ", "ye ", " we", "er ", "di ", "we ", "ê d", "i b", " be", "erk", "ina", " na", " an", "î û", "yê ", "eye", "î y", "kî ", "rke", "nê ", "diy", "ete", "eke", "ber", "hem", "hey", " li", " ci", "wek", "li ", "n d", "fê ", " bê", " te", "ne ", "yî ", " se", "net", "rî ", "tew", "yek", "sti", "af ", " ki", "re ", "yan", "n b", "kar", "hev", "e k", "aza", "n û", "wî ", " ew", "i h", "n k", "û b", "î b", " mi", " az", "dan", " wî", "ekî", "î a", "a m", "zad", "e d", "mir", "bin", "est", "ara", "iro", "nav", "ser", "a w", "adi", "rov", "n h", "anê", "tê ", "ewe", "be ", "ewl", "ev ", "mû ", " ya", "tî ", "ta ", "emû", " yê", "ast", "wle", " tê", "n m", " bo", "wey", "s m", "bo ", " tu", "n j", "ras", " da", " me", "din", "î d", "ê h", "n n", "n w", "ing", "st ", " ke", " ge", "în ", "ar ", " pê", "iye", "îna", "bat", "r k", "ema", "cih", "ê b", "wed", "û m", "dî ", "û a", "vak", "ê t", "ekh", "par", " ye", "vî ", "civ", "n e", "ana", "î h", "ê k", "khe", "geh", "nge", "ûna", "fên", "ane", "av ", "î m", "bik", "eyê", "eyî", "e û", " re", "man", "erb", "a x", "vê ", "ê m", "iva", "e n", "hî ", "bûn", "kê ", " pa", "erî", "jî ", "end", " ta", "ela", "nên", "n x", "a k", "ika", "f û", "f h", "î n", "ari", "mî ", "a s", "e j", "eza", "tên", "nek", " ni", "ra ", "ehî", "tiy", "n a", "bes", "rbe", "û h", "rwe", "zan", " a ", "erw", "ov ", "inê", "ama", "ek ", "nîn", "bê ", "ovî", "ike", "a n", " ra", "riy", "i d", "anî", "û d", "e e", "etê", "ê x", "yet", "aye", "ê j", "tem", "e t", "erd", "i n", "eta", "ibe", "a g", "u d", "xeb", "atê", "i m", "tu ", " wi", "dew", "mal", "let", "nda", "ewa", " ên", "awa", "e m", "a d", "mam", "han", "u h", "a b", "pêş", "ere", " ba", "lat", "ist", " za", "bib", "uke", "tuk", "are", "asî", "rti", "arî", "i a", "hîn", " hî", "edi", "nûn", "anû", "qan", " qa", " hi", " şe", "ine", "n l", "mên", "ûn ", "e a"},
	Yor: []string{"ti ", " ní", "ó̩ ", " è̩", "ní ", " lá", "̩n ", "o̩n", "é̩ ", "wo̩", "àn ", " e̩", "kan", "an ", "tí ", " tí", "tó̩", " kò", "ò̩ ", "̩tó", " àw", " àt", "è̩ ", "è̩t", "e̩n", "bí ", "àti", "lát", "áti", " gb", "lè̩", "s̩e", " ló", " ó ", "àwo", "gbo", "̩nì", "n l", " a ", " tó", "í è", "ra ", " s̩", "n t", "ò̩k", "sí ", "tó ", "̩ka", "kò̩", "ìyà", "o̩ ", " sí", "ílè", "orí", "ni ", "yàn", "dè ", "̩‐è", "ì k", "̩ à", "èdè", " or", "ún ", "ríl", "è̩‐", "í à", "jé̩", "‐èd", "àbí", "̩ò̩", "ò̩ò", "tàb", "nì ", "í ó", "n à", " tà", "̩ l", "jo̩", " ti", "̩e ", "̩ t", " wo", "nìy", "í ì", "ó n", " jé", " sì", "ló ", "kò ", "n è", "wó̩", " bá", "n n", "sì ", " fú", "̩ s", "í a", "rè̩", "fún", " pé", " òm", "̩ni", "gbà", " kí", " èn", "ènì", "in ", "òmì", "ìí ", "ba ", "nir", "pé ", "ira", "mìn", "ìni", "n o", "ràn", "ìgb", " ìg", "bá ", "e̩ ", " rè", "̩ n", "kí ", "n e", "un ", "gba", "̩ p", "í ò", "nú ", " o̩", "nín", "gbé", "yé ", " ka", "ínú", "a k", "fi ", " fi", "mo̩", "bé̩", "o̩d", "dò̩", "̩dò", "ó s", "i l", "̩ o", "̩ ì", "wà ", "í i", "i ì", "hun", "bò ", "i ò", "dá ", "bo̩", "o̩m", "̩mo", "̩wó", "bo ", "áà ", "̩ k", "ó j", "ló̩", "àgb", "ohu", " oh", " bí", " ò̩", "bà ", "ara", "yìí", "ogb", "írà", "n s", "ú ì", " ìb", "pò̩", "í k", " lè", "bog", "i t", "à t", "óò ", "yóò", "kó̩", "gé̩", "à l", "ó̩n", "rú ", "lè ", " yó", "̩ ò", "̩ e", "a w", "̩ y", "ò̩r", "̩ f", " wà", "ò l", "í t", "ó b", "i n", "ó̩w", "̩gb", "yí ", "í w", "ìké", "̩ a", "láà", "wùj", "àbò", "i è", "ùjo", "fin", "é̩n", "n k", "í e", "i j", "ú à", " ìk", "òfi", " òf", " ar", "i s", "mìí", "ìír", " mì", " ir", "rin", "náà", " ná", "jú ", "̩ b", " yì", "ó t", "̩é̩", " i ", "̩ m", "fé̩", "kàn", "rí ", "ú è", "à n", "wù ", "s̩é", "é à", " mú", " èt", "áyé", "í g", "̩kó", "̩dá", "è̩d", "àwù", "è̩k", " ìd", "irú", "í o", "i o", "i à", "láì", "í n", "ípa", " kú", "níp", " ìm", "a l", "ké̩", "bé ", "i g", "de ", "ábé", "ìn ", "báy", "̩è̩", "ígb", "wò̩", "níg", "mú ", "láb", " àà", "n f", "è̩s", "̩ w", "ùn ", "i a", "ayé", "èyí", " èy", "mó̩", "á è", " ni", "n b", " wó", "je̩", " ìj", "gbá", "ò̩n", "ó̩g"},
	Uzb: []string{"lar", "ish", "an ", "ga ", "ar ", " va", " bi", "da ", "va ", "ir ", " hu", "iga", "sh ", "uqu", "shi", "bir", "quq", "huq", "gan", " bo", " ha", "ini", "ng ", "a e", "r b", " ta", "lis", "ni ", "ing", "lik", "ida", "oʻl", "ili", "ari", "nin", "on ", "ins", " in", "adi", "nso", "son", "iy ", " oʻ", "lan", " ma", "dir", "hi ", "kin", "har", "i b", "ash", " yo", "boʻ", " mu", "dan", "uqi", "ila", "ega", "qla", "r i", "qig", "oʻz", " eg", "kla", "a b", "qil", "erk", "ki ", " er", "oli", "nli", "at ", " ol", "gad", "lga", "rki", "oki", "i h", "a o", " qa", "yok", "lig", "osh", "igi", "ib ", "las", "n b", "atl", "n m", " ba", "ara", " qi", "ri ", " sh", "iya", "ala", "lat", "in ", "ham", "bil", "a t", "a y", "bos", "r h", "siy", "n o", "yat", "inl", "ik ", "a q", "cha", "a h", " et", "eti", "nis", "a s", "til", "ani", "h h", "i v", "mas", "tla", "osi", "asi", " qo", "ʻli", "ati", "i m", "rni", "im ", "uql", "arn", "ris", "qar", "a i", "gi ", " da", "n h", "ha ", "sha", "i t", "mla", "rch", " xa", "i o", "li ", "hun", "bar", "lin", "ʻz ", "arc", "rla", " bu", "a m", "a a", " as", "mum", " be", " tu", "aro", "r v", "ikl", "lib", "taʼ", "h v", "tga", "tib", "un ", "lla", "mda", " ke", "shg", " to", "n q", "sid", "n e", "mat", "amd", "shu", "hga", " te", "tas", "ali", "umk", "oya", "hla", "ola", "aml", "iro", "ill", "tis", "iri", "rga", "mki", "irl", " ya", "xal", "dam", " de", "gin", "eng", "rda", "tar", "ush", "rak", "ayo", " eʼ", " so", "ten", "alq", " sa", "ur ", " is", "imo", "r t", " ki", "mil", " mi", "era", "zar", "hqa", "aza", "k b", " si", "nda", "hda", "kat", "ak ", "oʻr", "n v", "a k", "or ", "rat", "ada", "ʻlg", "miy", "tni", "i q", "shq", "oda", "shl", "bu ", "dav", "nid", "y t", "ch ", "asl", "sos", "ilg", "aso", "n t", "atn", "sin", "am ", "ti ", "as ", "ana", "rin", "siz", "yot", "lim", "uni", "nga", "lak", "n i", "a u", "qon", "i a", "h k", "vla", "avl", "ami", "dek", " ja", "ema", "a d", "na ", " em", "ekl", "gʻi", "si ", "i e", "ino", " ka", "uch", "bor", "ker", " ch", "lma", "liy", "a v", "ʼti", "lli", "aka", "muh", "rig", "ech", "i y", "uri", "ror"},
	Ibo: []string{"a n", "e n", "ke ", " na", "na ", " ọ ", " bụ", " n ", "nwe", "ere", "ọ b", "re ", "nye", " nk", "ya ", "la ", " nw", " ik", " ma", "ye ", "e ọ", "ike", "a o", "nke", "a m", "ụ n", " ya", "a ọ", "ma ", "bụl", "ụla", " on", " a ", "e i", "kik", "iki", "ka ", "ony", "ta ", "bụ ", "kwa", " nd", "a i", "i n", "di ", "a a", "wa ", "wer", "do ", " mm", "dụ ", "e a", "ha ", " ga", "any", " ob", "ndi", " ok", "he ", "e m", "e o", "a e", "ọ n", "ite", "rụ ", "hi ", "mma", "ga‐", "wu ", "ara", " dị", "aka", "che", "oke", "we ", "o n", " ih", "n o", "adụ", "mad", "obo", "bod", "a g", "odo", " ka", " ez", "te ", "hị ", "be ", "ụta", "dị ", " an", "zi ", " oh", "a‐e", "akw", "gba", "i m", "me ", " ak", "u n", "nya", "ihe", "ala", "ohe", "ghi", "ri ", " ọz", "her", "ra ", "weg", " nt", " iw", " mb", "ba ", "pụt", " si", "ro ", "oro", "iwu", "chi", "a‐a", "rị ", "ụ i", "ụ ọ", " eb", "iri", "ebe", "ụrụ", "zọ ", " in", "a y", "ezi", "e ị", "kpa", "le ", "ile", "ịrị", "n e", "kpe", "mba", " ha", "bi ", "sit", "e e", "inw", "nil", "asị", " en", "mak", "a u", " ni", "apụ", "chị", "i i", "ghị", "i ọ", "i o", "si ", " e ", "ide", "o i", "e y", "ụ m", "a s", "u o", "kwu", "ozu", "yer", "ru ", "enw", "ụ o", "ọzọ", "gid", "hụ ", "n a", "ahụ", "nkw", "sor", "egh", "edo", "a ụ", "tar", "n i", "toz", "ị o", "pa ", "i a", " me", "ime", "uru", "kwe", " mk", "tu ", "ama", "eny", "uso", "de ", " im", "ọ d", "osi", "hed", "a d", " kw", "mkp", "wet", " ọr", " ọn", "obi", "ọrụ", " ịk", " to", "gas", " ch", "ịch", "nha", "ọnọ", "nọd", " nc", " al", "n ụ", "ị m", " us", "nọ ", "u ọ", "nch", " o ", "eta", "n u", " ot", "otu", "sir", "sịr", " nh", "a k", "ali", "o m", " ag", " gb", "e s", "ọta", "nwa", "ị n", "lit", "ega", "ji ", "ọdụ", "e k", "ban", "e g", "ị k", "esi", "agb", "eme", "hu ", "ikp", "zu ", "pe ", "nta", "na‐", "chọ", "u a", "a b", "uch", "n ọ", "onw", "ram", "kwụ", "ekọ", "i e", " nọ", " ug", "ọch", "u m", "gwu", "a h", "zụz", "ugw", "meg", "ị e", "nat", "e h", "dịg", "o y", "kpu", "pụr", "cha", "zụ ", "hịc", "ich", " ng", "ach", " og", "wap", "wan", "ịgh", "uwa", " di", " nn", "i ị"},
	Ceb: []string{"sa ", " sa", "ng ", "ang", " ka", "an ", " pa", "ga ", " ma", "nga", "pag", " ng", "a p", "on ", "kat", "a k", "ug ", "od ", " ug", "g m", " an", "ana", "n s", "ay ", "ung", "ata", "ngo", "a m", "atu", "ala", "san", "ag ", "tun", "g s", "g k", "god", "d s", "a s", "ong", "mga", " mg", "g p", "n u", "yon", "a a", "pan", "ing", "usa", "tan", "tag", "una", "aga", "mat", "ali", "g u", "han", "nan", " us", "man", "y k", "ina", "non", "kin", " na", "syo", "lan", "a b", "asa", "nay", "n n", "a i", "awa", " ta", "taw", "gaw", "nsa", "a n", "nas", " o ", "ban", "agp", "isa", "dun", "was", "iya", " gi", "asy", "adu", "ini", "bis", " ad", "ili", "o s", " bi", "g a", "nah", "nag", "a t", " ki", "lin", "lay", "ahi", "sam", "al ", "wal", " di", "nal", "asu", " ba", "ano", "agt", " wa", "ama", "yan", "a u", " iy", "kan", "him", "n k", "gan", "ags", "n a", "kag", " un", "ya ", "kas", "gpa", "g t", " su", "aha", "wha", "agk", "awh", "gka", "a g", "kal", "l n", "gla", "gsa", "sud", "gal", "imo", "ud ", "d u", "ran", "uka", "ig ", "aka", "aba", "ika", "g d", "ara", "ipo", "ngl", "g n", "uns", "n o", "kau", "i s", "y s", "og ", "uta", "d n", "li ", " si", "gik", "g i", "mta", "ot ", "iin", " la", " og", "o a", "ayo", "ok ", "awo", "aki", "kab", "aho", "n m", "hat", "o p", "gpi", "a w", "apa", "lip", "ip ", " hu", " ga", "a h", "uba", "na ", " ti", "bal", "gon", "la ", "ati", "wo ", "ad ", "hin", "sal", "gba", "buh", " bu", " ub", "uha", "agb", "hon", "ma ", "nin", "uga", "t n", "ihi", " pi", "may", " pu", "mak", "ni ", " ni", "d a", "pin", "abu", "agh", "ahu", "uma", "as ", "dil", "say", " in", "at ", "ins", "lak", "hun", "ila", "mo ", "s s", "sak", "amt", "o u", "pod", "ngp", "tin", "a d", "but", "ura", "lam", "aod", "t s", "bah", "ami", "aug", "mal", "sos", "os ", "k s", " il", "tra", " at", "gta", "bat", "aan", "ulo", "iha", "ha ", "n p", " al", "g b", "lih", "kar", "lao", "agi", "amb", "mah", "ho ", "sya", "ona", "aya", "ngb", "in ", "inu", "a l", " hi", "mag", "iko", "it ", "agl", "mbo", "oon", "tar", "o n", "til", "ghi", "rab", "y p", " re", "yal", "aw ", "nab", "osy", "dan"},
	Tgl: []string{"ng ", "ang", " pa", "an ", "sa ", " sa", "at ", " ka", " ng", " ma", "ala", "g p", "apa", " na", "ata", "pag", "pan", " an", " at", "ay ", "ara", "ga ", "a p", "tan", "g m", "mga", " mg", "n n", "pat", " ba", "n a", "aya", "na ", "ama", "g k", "awa", "kar", "a k", "lan", "rap", "gka", "nga", "n s", "g n", "aha", "g b", "a a", " ta", "agk", "gan", "tao", "asa", "aka", "yan", "ao ", "a m", "may", "man", "kal", "ing", "a s", "nan", "aga", " la", "ban", "ali", "g a", "ana", "y m", "kat", "san", "kan", "g i", "ong", "pam", "mag", "a n", "o a", "baw", "isa", "wat", " y ", "lay", "g s", "y k", "in ", "ila", "t t", " ay", "aan", "o y", "kas", "ina", "t n", "ag ", "t p", "wal", "una", "yon", " o ", " it", "nag", "lal", "tay", "pin", "ili", "ans", "ito", "nsa", "lah", "kak", "any", "a i", "nta", "nya", "to ", "hay", "gal", "mam", "aba", "ran", "ant", "agt", "on ", "t s", "agp", " wa", " ga", "gaw", "han", "kap", "o m", "lip", "ya ", "as ", "g t", "hat", "y n", "ngk", "ung", "no ", "g l", "gpa", "wa ", "lag", "gta", "t m", "kai", "yaa", "sal", "ari", "lin", "a l", "pap", "ahi", " is", " di", "ita", " pi", "pun", "agi", "ipi", "mak", "a b", "y s", "bat", "yag", "ags", "o n", "aki", "tat", "pah", "la ", "gay", "hin", " si", "di ", "i n", "sas", "iti", "a t", "t k", "mal", "ais", "s n", "t a", "al ", "ipu", "ika", "lit", "gin", " ip", "ano", "gsa", "alo", "nin", "uma", "hal", "ira", "ap ", "ani", "od ", "i a", "gga", "y p", "par", "tas", "ig ", "sap", "ihi", "nah", "ini", " bu", "ngi", "syo", "o s", "nap", "o p", "a g", " ha", "uka", "a h", "aru", "a o", "mah", "iba", "asy", "li ", "usa", "g e", "uha", "ipa", "mba", "lam", "kin", "kil", "duk", "n o", "iga", " da", "dai", "aig", "igd", "gdi", "pil", "dig", "pak", " tu", "d n", "sam", "nas", "nak", "ba ", "ad ", "lim", "sin", "buh", "ri ", "lab", "it ", "tag", "g g", "lun", "ain", "and", "nda", "pas", "kab", "aho", "lig", "nar", "ula", " ed", "edu", " ib", "git", "ma ", "mas", "agb", "ami", "agg", "gi ", "sar", "i m", "siy", "g w", "api", "pul", "iya", "amb", "nil", "agl", "sta", "uli", "ino", "abu", "aun", "ayu", " al", "iyo"},
	Hun: []string{" sz", " a ", "en ", " va", "és ", " és", "min", "ek ", " mi", " jo", "jog", "ind", "an ", "nek", "sze", "ság", " az", "gy ", "sza", "nde", "ala", "az ", "den", "a v", "val", "ele", " el", "oga", "mél", "egy", " eg", "n a", "ga ", "zab", " me", "zem", "emé", "aba", "int", "van", "bad", "tel", "tet", " te", "ak ", "tás", "ény", "t a", " ne", "gye", "ély", "tt ", "n s", "ben", "ség", "zet", "lam", "meg", "nak", "ni ", " se", "ete", "sen", "agy", "let", "lyn", "s a", "yne", "ra ", "z e", "et ", " al", "mel", "kin", "k j", "eté", "ok ", "tek", " ki", "vag", "re ", "n m", "oz ", "hoz", "ez ", "s s", "ett", "gok", "ogy", " kö", "mbe", "es ", "em ", "nem", "ely", " le", "ell", "emb", "hog", "k a", "atá", "köz", "nt ", " ho", "yen", "hez", "el ", "z a", "len", "dsá", "ásá", "tés", "ads", "k m", " ál", " em", "a s", "nte", "a m", "szt", "a t", "áll", "ás ", "y a", "ogo", "sem", "a h", "enk", "nye", "ese", "nki", "ágo", "t s", "lap", "ame", "ber", "ló ", "k é", "nyi", "ban", "mén", "s e", "i m", "t m", " vé", "lla", "ly ", "ébe", "lat", "ág ", "ami", "on ", "mze", "n v", "emz", "fel", "a n", "lő ", "a a", "eki", "eri", "yes", " cs", "lle", "tat", "elő", "nd ", "i é", "ég ", "ésé", "lis", "yil", "vet", "át ", "kül", "ért", " ke", "éte", "rés", "l a", "het", "szo", "art", "alá", " ny", "tar", "koz", " am", "a j", "ész", "enl", "elé", "ól ", "s k", "tár", "s é", "éle", "s t", "lem", "sít", "ges", "ott", " fe", "n k", "tko", "zás", "t é", "kel", "ja ", " ha", "aló", "zés", "nlő", "ése", "ot ", "ri ", "lek", "más", "tő ", "vel", "i j", "se ", "ehe", "tes", "eve", "ssá", "tot", "t k", "olg", "eze", "i v", "áza", "leh", "n e", "ül ", "tte", "os ", "ti ", "atk", "zto", "e a", "tos", "ány", "ána", "zte", "fej", "del", "árs", "k k", "kor", "ége", "szá", "t n", " bi", "zat", "véd", "nev", "elm", "éde", "zer", "téb", "biz", "rra", "ife", "izt", "ere", "at ", "ll ", "k e", "ny ", "sel", " né", "ába", "lt ", "ai ", "sül", "ház", "kif", "t e", " ar", "leg", "d a", "is ", "i e", "arr", "t t", "áso", "it ", "ető", "al ", " má", "t v", " bá", "bár", "a é", "esü", "lye", "m l", " es", "nyo"},
	Azj: []string{" və", "və ", "ər ", "lar", " hə", "in ", "ir ", " ol", " hü", " bi", "hüq", "üqu", "quq", "na ", "lər", "də ", "hər", " şə", "bir", "an ", "lik", " tə", "r b", "mal", "lma", "ası", "ini", "r h", "əxs", "şəx", "ən ", "arı", "qla", "a m", "dir", "aq ", "uqu", "ali", " ma", "una", "ilə", "ın ", "yət", " ya", "ara", "ikd", "əri", "ar ", "əsi", "əti", "r ş", "rin", "yyə", "n h", " az", "dən", "nin", "ərə", "tin", "iyy", "mək", "zad", " mü", "sin", " mə", "ni ", "nda", "ət ", "ndə", "aza", "rın", "ün ", "ını", "ə a", "i v", "nın", "olu", "qun", " qa", " et", "ilm", "lıq", "ə y", "ək ", "lmə", "lə ", "kdi", "ind", "ına", "olm", "lun", "mas", "xs ", "sın", "ə b", " in", "n m", "q v", "nə ", "əmi", "n t", "ya ", "da ", " bə", "tmə", "dlı", "adl", "bər", " on", "əya", "ə h", "sı ", "nun", "maq", "dan", "inə", "etm", "un ", "ə v", "rlə", "n b", "si ", "raq", " va", "ə m", "n a", "ınd", "rı ", "anı", " öz", "əra", "nma", "n i", "ama", "a b", "irl", "ala", "li ", "ins", "bil", "ik ", " al", " di", "ığı", "ə d", "lət", "il ", "ələ", "ə i", "ıq ", "nı ", "nla", "dil", "müd", "n v", "ə e", "unm", "alı", " sə", "xsi", "ə o", "uq ", "uql", "nsa", "ətl", " də", "ili", "üda", "asi", " he", "ola", "san", "əni", "məs", " da", "lan", " bu", "tər", "həm", "dır", "kil", "iş ", "u v", " ki", "min", "eyn", "mi ", "yin", " ha", "sos", "heç", "bu ", "eç ", " ed", "kim", "lığ", "alq", "xal", " as", "sia", "osi", "r v", "q h", "rə ", "yan", "i s", " əs", "daf", "afi", " iş", "ı h", "fiə", " ta", "ə q", "ıql", "a q", "yar", "sas", "lı ", "ill", "mil", "əsa", "liy", "tlə", "siy", "a h", "məz", "tün", "ə t", " is", "ist", "iyi", " so", "n ə", "al ", "ifa", "ina", "lıd", "ı o", "ıdı", "əmə", "ır ", "ədə", "ial", " mi", "əyi", "miy", "çün", "n e", "iya", "edi", " cə", " bü", "büt", "ütü", "xil", "üçü", "mən", "adə", "t v", "a v", "axi", "dax", "r a", "onu", " üç", "seç", " nə", " se", "man", "ril", "sil", "əz ", "iə ", "öz ", "ılı", "aya", "qan", "i t", "şər", "təm", "ulm", "rəf", "məh", " xa", "ğın", " dö", " ni", "sti", "ild", "amə", "qu ", "nam", "n o", "n d", "var", "ad ", "zam", "tam", "təh"},
	Ces: []string{" pr", " a ", "ní ", " ne", "prá", "ráv", "ost", " sv", " po", "na ", "ch ", "ho ", " na", "nos", "o n", " ro", "ání", "ti ", "vo ", "neb", "ávo", "má ", "bo ", "ebo", " má", "kaž", " ka", "ou ", "ažd", " za", " je", "dý ", "svo", "ždý", " př", "a s", " st", "sti", "á p", " v ", "obo", "vob", " sp", "bod", " zá", "ých", "pro", "rod", "ván", "ení", "né ", "ý m", "ého", " by", " ná", "spo", "ně ", "o p", "mi ", "í a", "ter", "roz", "ová", "to ", " ja", " li", "áro", "nár", "by ", "jak", "a p", "a z", "ny ", " vš", "kte", "i a", "lid", "ím ", "o v", "í p", "u p", "mu ", "at ", " vy", "odn", " so", " ma", "a v", " kt", "í n", "zák", "li ", "oli", "ví ", "kla", "tní", "pod", "stá", "en ", "do ", "t s", "mí ", "je ", "em ", "áva", " do", "byl", " se", "být", "í s", "rov", " k ", "čin", " ve", "ýt ", "í b", "it ", "dní", "vše", "pol", "o s", " bý", "tví", "nýc", "stn", "nou", "ejn", "sou", "ran", "ci ", "vol", "se ", "nes", "a n", "pří", "eho", "ným", "tát", "va ", "ním", "mez", "ají", "i s", "stv", "ké ", "ích", "ečn", "žen", "e s", "vé ", "ova", "své", "ým ", "kol", "du ", "u s", "jeh", "kon", "ave", "ech", "eré", "nu ", " ze", "i v", "o d", "í v", "hra", "ids", "m p", "ému", "ole", "y s", " i ", "maj", "o z", " to", "aby", "sta", " ab", "m a", "pra", " ta", "chn", " ni", "že ", "ovn", "ako", "néh", "len", "dsk", "rac", "lad", "chr", " že", "vat", " os", "sob", "aké", "i p", "smí", "esm", "st ", "i n", "m n", "a m", "lně", "lní", "při", "bez", "dy ", "áln", "ens", "zem", "t v", "čen", "leč", "kdo", "ými", " ji", "oci", "i k", " s ", "í m", "jí ", " či", "áv ", "ste", "och", " oc", "vou", "ákl", " vz", "rav", "odu", "nez", "inn", "ský", "nit", "ivo", "a j", "u k", "iál", " me", "ezi", "ské", "ven", "stu", "u a", "tej", "oln", "slu", "zen", "í z", "y b", "oko", "zac", "níc", "jin", "ky ", "a o", "řís", "obe", "u v", "tak", "věd", "oje", " vý", "ikd", "h n", " od", "čno", "oso", "ciá", "h p", " de", "a t", "ům ", "soc", "jíc", "odů", "něn", "adn", "tup", "dů ", "děl", "jno", "kéh", "por", "ože", "hov", "aci", "nem", "é v", "rok", "i j", "u o", "od ", "ího", "vin", "odi"},
	Mlg: []string{"ny ", "na ", "ana", " ny", "y f", "a n", "sy ", "aha", "ra ", "a a", " fa", "n n", "y n", "a m", "an ", " fi", "tra", "any", " ma", "han", "nan", "ara", "y a", " am", "ka ", "in ", "y m", "ami", "olo", " ts", "lon", "min", " mi", " sy", " na", "a t", " ol", "fan", " ha", "a i", "man", "iza", " iz", "ina", "ona", "y h", "aka", "o a", "ian", "a h", "reh", "etr", "a s", "het", "on ", "a f", "ire", "fah", "tsy", "mba", " ar", " hi", "zan", "ay ", "ndr", "y o", "ira", "y t", " an", "ehe", "o h", "afa", "y i", "ren", "ran", " zo", "ena", "amb", "dia", "ala", "amp", "zo ", "ika", " di", "tan", "y s", "y z", " az", "ia ", "m p", "rin", "jo ", "n j", " jo", " dr", "zy ", "ry ", "a d", "ao ", "and", "dre", "haf", "nen", "mpi", "rah", " ka", "eo ", "n d", " ir", "ho ", "am ", "rai", "fa ", "elo", "ene", "oan", "omb", " ta", " pi", " ho", "ava", "azo", "dra", "itr", "iny", "ant", "tsi", "zon", "asa", "tsa", " to", "ari", "ha ", "a k", "van", "n i", "fia", "ray", " fo", "mbe", "ony", "sa ", "isy", "azy", "o f", "lal", "ly ", "ova", "lom", " vo", "nat", "fir", "sam", "oto", "zay", "mis", "ham", "bel", " ra", "a r", "ban", "kan", "iha", "nin", "a e", "ary", "ito", " he", " re", " no", "ita", "voa", "nam", "fit", "iar", " ko", "tok", "isa", "fot", "no ", "otr", "mah", "aly", "har", "y v", "y r", " sa", "o n", "ain", "kam", "aza", "n o", "oka", "ial", "ila", "ano", "atr", "oa ", " la", "y l", "eri", "y d", "ata", "hev", "sia", "pia", "its", "reo", " ao", "pan", "anj", "aro", "tov", "nja", "o s", "fam", "pir", " as", "ty ", "nto", "oko", "y k", "sir", "air", "tin", "hia", "ais", "mit", "ba ", " it", " eo", "o t", "mpa", "kon", "a z", "a v", "ity", "ton", "rak", "era", "ani", "ive", "mik", "ati", "tot", "vy ", "hit", "hoa", "aho", "ank", "ame", "ver", "vah", "tao", "o m", "ino", "dy ", "dri", "oni", "ori", " mo", "hah", "nao", "koa", "ato", "end", "n t", " za", "eha", "nga", "jak", "bar", "lah", "mia", "lna", "aln", "va ", " mb", "lan", " pa", "aov", "ama", "eve", "za ", "dro", "ria", "to ", "nar", "izy", "ifa", "adi", "via", "aja", " va", "ind", "n k", "idi", "fiv", "rov", "vel"},
	Nya: []string{"ndi", "ali", "a k", "a m", " ku", " nd", "wa ", "na ", "nth", " mu", " al", "yen", "thu", "se ", "ra ", "nse", "hu ", "di ", "a n", "la ", " pa", "mun", " wa", "nga", "unt", " la", "a u", "u a", "e a", "ons", "za ", " ma", " lo", "iye", "ace", "ce ", "a l", "idw", "ang", " ka", "kha", "liy", "ens", "li ", "ala", "ira", "ene", "pa ", "i n", "we ", "e m", "ana", "dwa", "era", "hal", "ulu", "lo ", "ko ", "dzi", " ci", "yo ", "o w", "iko", "ga ", "a p", "chi", " mo", "lu ", "o l", "o m", "oyo", "ufu", " um", "moy", "zik", " an", "ner", "and", "umo", "ena", " uf", "dan", "iri", "ful", "a a", "ka ", "to ", "hit", "nch", " nc", "a c", "ito", "fun", "dwe", " da", "kuk", "wac", " dz", "e l", "a z", "ape", "kap", "u w", "e k", "ere", "ti ", "lir", " za", "pen", "tha", "aye", "kut", "mu ", "ro ", "ofu", "ing", "lid", " zo", "amu", "o c", "i m", "mal", "kwa", "mwa", "o a", "eza", "i p", "o n", "so ", "i d", "lin", "nso", " mw", "iro", "zo ", " a ", "ati", " li", "i l", "a d", "ri ", "edw", "kul", "una", "uti", "lan", "a b", "iki", "i c", "alo", "i k", " ca", "lam", "o k", "dza", "ung", "o z", "mul", "ulo", "uni", "gan", "ant", "nzi", " na", "nkh", "e n", "san", "oli", "wir", "tsa", "u k", "ome", "ca ", "gwi", "unz", "lon", "dip", "ipo", "yan", "gwe", "pon", "akh", "uli", "aku", "mer", "ngw", "cit", " po", " ko", "kir", "mba", "ukh", "tsi", "bun", "iya", "ope", "kup", "bvo", "han", " bu", "pan", "ame", "vom", "ama", " ya", "siy", " am", "rez", "u n", "zid", "men", "osa", "ao ", "pez", "i a", " kw", " on", "u o", "lac", "ezo", "aka", "nda", "hun", "u d", "ank", "diz", "ina", "its", "adz", " kh", "ne ", "nik", "e p", "o o", "ku ", "phu", "eka", " un", "eze", "mol", "ma ", " ad", "pat", "oma", "ets", "wez", "kwe", "kho", "ya ", "izo", "sa ", "o p", "kus", "oci", "khu", "okh", "ans", "awi", "izi", "zi ", "ndu", "iza", "no ", "say", " si", "i u", "aik", "jir", "ats", "ogw", "du ", "mak", "ukw", "nji", "mai", "ja ", "sam", "ika", "aph", "sid", "isa", "amb", "ula", "osi", "haw", "u m", " zi", "oye", "lok", "win", "lal", "ani", " ba", "si ", " yo", "e o", "opa", "ha ", "map", "emb"},
	Kin: []string{"ra ", " ku", " mu", "se ", "a k", "ntu", "nga", "tu ", "umu", "ye ", "li ", " um", "mun", "unt", "a n", "ira", " n ", "ere", "wa ", "we ", " gu", "mu ", "ko ", "a b", "e n", "o k", "e a", "a u", "a a", "u b", "e k", "ose", "uli", "aba", "ro ", " ab", "gom", "e b", "ba ", "ugu", " ag", "omb", "ang", " ib", "eng", "mba", "o a", "gu ", " ub", "ama", " by", " bu", "za ", "ihu", "ga ", "e u", "o b", " ba", "kwi", "hug", "ash", "ren", "yo ", "ndi", "e i", " ka", " ak", " cy", "iye", " bi", "ora", "re ", "gih", "igi", "ban", "ubu", " nt", " kw", "di ", "gan", "a g", "a m", "aka", "nta", "aga", " am", "a i", "ku ", "iro", "i m", "ta ", "ka ", "ago", "byo", "ali", "and", "ibi", "na ", "uba", "ili", " bw", "sha", "cya", "u m", "yan", "o n", " ig", "ese", "no ", "obo", "ana", "ish", "kan", "sho", " we", "era", "ya ", "aci", "wes", "ura", "i a", "uko", "e m", "n a", "o i", "kub", "uru", "hob", "ber", "ran", "bor", " im", "ure", "u w", "wo ", "cir", "gac", "ani", "bur", "u a", "o m", "ush", " no", "e y", " y ", "rwa", "eke", "nge", "ara", "wiy", "uga", "zo ", "ne ", "ho ", "bwa", "yos", "anz", "aha", "ind", "mwe", "teg", "ege", "are", "ze ", "n i", "rag", "ane", "u n", "ge ", "mo ", "u k", "bul", " uk", "bwo", "bye", "iza", "age", "ngo", "u g", "gir", "ger", "zir", "kug", "ite", "bah", " al", " ki", "uha", "go ", "mul", "ugo", "n u", "tan", "guh", "y i", " ry", "gar", "bih", "iki", "atu", "ha ", "mbe", "bat", "o g", "akw", "iby", "imi", "kim", "ate", "abo", "e c", "aho", "o u", "eye", "tur", "kir", " ni", "je ", "bo ", "ata", "u u", " ng", "shy", "a s", "gek", " ru", "iko", " bo", "bos", "i i", " gi", "nir", "i n", "gus", "eza", "nzi", "i b", "kur", " ya", "o r", "ung", "rez", "ugi", "ngi", "nya", " se", "mat", "eko", "o y", " in", "uki", " as", "any", "bis", "ako", "gaz", "imw", "rer", "bak", "ige", "mug", "ing", "byi", "kor", "eme", "nu ", " at", "bit", " ik", "hin", "ire", "kar", "shi", "yem", "yam", " yi", "gen", "tse", "ets", "ihe", "hak", "ubi", "key", "rek", "icy", " na", "bag", "yer", " ic", "eze", "awe", "but", "irw", " ur", "fit", "ruk", "ubw", "rya", "uka", "afi"},
	Zul: []string{"nge", "oku", "lo ", " ng", "a n", "ung", "nga", "le ", "lun", " no", "elo", "wa ", "la ", "e n", "ele", "ntu", "gel", "tu ", "we ", "ngo", " um", "e u", "thi", "uth", "ke ", "hi ", "lek", "ni ", "ezi", " ku", "ma ", "nom", "o n", "pha", "gok", "nke", "onk", "a u", "nel", "ulu", "oma", "o e", "o l", "kwe", "unt", "ang", "lul", "kul", " uk", "a k", "eni", "uku", "hla", " ne", " wo", "mun", " lo", "kel", "ama", "ath", "umu", "ho ", "ela", "lwa", "won", "zwe", "ban", "elw", "ule", "a i", " un", "ana", "une", "lok", "ing", "elu", "wen", "aka", "tho", "aba", " kw", "gan", "ko ", "ala", "enz", "o y", "khe", "akh", "thu", "u u", "na ", "enk", "kho", "a e", "zin", "gen", "i n", "kun", "alu", "mal", "lel", "e k", "nku", "e a", "eko", " na", "kat", "lan", "he ", "hak", " ez", "o a", "kwa", "o o", "ayo", "okw", "kut", "kub", "lwe", " em", "yo ", "nzi", "ane", "obu", " ok", "eth", "het", "ise", "so ", "ile", "nok", " ba", "ben", "eki", "nye", "ike", "i k", "isi", " is", "aph", "esi", "nhl", "mph", " ab", "fan", "e i", "isa", " ye", "nen", "ini", "ga ", "zi ", "fut", " fu", "uba", "ukh", "ka ", "ant", "uhl", "hol", "ba ", "and", "do ", "kuk", "abe", "za ", "nda", " ya", "e w", "kil", "the", " im", "eke", "a a", "olo", "sa ", "olu", "ith", "kuh", "o u", "ye ", "nis", " in", "ekh", "e e", " ak", "i w", "any", "khu", "eng", "eli", "yok", "ne ", "no ", "ume", "ndl", "iph", "amb", "emp", " ko", "i i", " le", "isw", "zo ", "a o", "emi", "uny", "mel", "eka", "mth", "uph", "ndo", "vik", " yo", "hlo", "alo", "kuf", "yen", "enh", "o w", "nay", "lin", "hul", "ezw", "ind", "eze", "ebe", "kan", "kuz", "phe", "kug", "nez", "ake", "nya", "wez", "wam", "seb", "ufa", "bo ", "din", "ahl", "azw", "fun", "yez", "und", "a l", "li ", "bus", "ale", "ula", "kuq", "ola", "izi", "ink", "i e", "da ", "nan", "ase", "phi", "ano", "nem", "hel", "a y", "hut", "kis", "kup", "swa", "han", "ili", "mbi", "kuv", "o k", "kek", "omp", "pho", "kol", "i u", "oko", "izw", "lon", "e l", " el", "uke", "kus", "kom", "ulo", "zis", "hun", "nje", "lak", "u n", "huk", "sek", "ham", " ol", "ani", "o i", "ubu", "mba", " am"},
	Swe: []string{" oc", "och", "ch ", "er ", "ing", "för", "tt ", "ar ", "en ", "ätt", "nde", " fö", "rät", "ill", "et ", "and", " rä", " en", " ti", " de", "til", "het", "ll ", "de ", "om ", "var", "lig", "gen", " fr", "ell", "ska", "nin", "ng ", "ter", " ha", "as ", " in", "ka ", "att", "lle", "der", "sam", " i ", "und", "lla", "ghe", "fri", "all", "ens", "ete", "na ", "ler", " at", "ör ", "den", " el", "av ", " av", " so", "igh", "r h", "nva", "ga ", "r r", "env", "la ", "tig", "nsk", "iga", "har", "t a", "som", "tti", " ut", "ion", "t t", "a s", "nge", "ns ", "a f", "r s", "män", "a o", " sk", " si", "rna", "isk", "an ", " st", "är ", "ra ", " vi", " al", "t f", " sa", "a r", "ati", " är", " me", " be", "n s", " an", "tio", "nna", "lan", "ern", "t e", "med", " va", "ig ", "äns", " åt", "sta", "ta ", "nat", " un", "kli", "ten", " gr", "vis", "äll", " la", "one", "han", "änd", "t s", "stä", "t i", "ner", "ans", "gru", " ge", "ver", " må", " li", "lik", "ihe", "ers", "rih", "r a", " re", "må ", "sni", "n f", "t o", " mä", " na", "r e", "ri ", "ad ", "ent", "kla", "det", " vä", "run", "rkl", "da ", "h r", "upp", "dra", "rin", "igt", "dig", "n e", "erk", "kap", "tta", "ed ", "d f", "ran", "e s", "tan", "uta", "nom", "lar", "gt ", "s f", " på", " om", "kte", "lin", "r u", "vid", "g o", "änn", "erv", "ika", "ari", "a i", "lag", "rvi", "id ", "r o", "s s", "vil", "r m", "örk", "ot ", "ndl", "str", "els", "ro ", "a m", "mot", " mo", "i o", "på ", "r d", "on ", "del", "isn", "sky", "e m", "ras", " hä", "r f", "i s", "a n", "nad", "n o", "gan", "tni", "era", "ärd", "a d", "täl", "ber", "nga", "r i", "enn", "nd ", "n a", " up", "sin", "dd ", "örs", "je ", "itt", "kal", "n m", "amt", "n i", "kil", "lse", "ski", "nas", "end", "s e", " så", "inn", "tat", "per", "t v", "arj", "e f", "l a", "rel", "t b", "int", "tet", "g a", "öra", "l v", "kyd", "ydd", "rje", " fa", "bet", "se ", "t l", "lit", "sa ", "när", "häl", "l s", "ndr", "nis", "yck", "h a", "llm", "lke", "h f", "arb", "lmä", "nda", "bar", "ckl", "v s", "rän", "gar", "tra", "re ", "ege", "r g", "ara", "ess", "d e", "vär", "mt ", "ap "},
	Som: []string{" ka", "ay ", "ka ", "an ", "uu ", "oo ", "da ", "yo ", "aha", " iy", "ada", "aan", "iyo", "a i", " wa", " in", "sha", " ah", " u ", "a a", " qo", "ama", " la", "hay", "ga ", "ma ", "aad", " dh", " xa", "ah ", "qof", "in ", " da", "a d", "aa ", "iya", "a s", "a w", " si", " oo", "isa", "yah", "eey", "xaq", "ku ", " le", "lee", " ku", "u l", "la ", "taa", " ma", "q u", "dha", "y i", "ta ", "aq ", "eya", "sta", "ast", "a k", "of ", "ha ", "u x", "kas", "wux", " wu", "doo", "sa ", "ara", "wax", "uxu", " am", "xuu", "inu", "nuu", "a x", "iis", "ala", "a q", "ro ", "maa", "o a", " qa", "nay", "o i", " sh", " aa", "kal", "loo", " lo", "le ", "a u", " xo", " xu", "o x", "f k", " ba", "ana", "o d", " uu", "iga", "a l", "yad", "dii", "yaa", "si ", "a m", "gu ", "ale", "u d", "ash", "ima", "adk", "do ", "aas", " ca", "o m", "lag", "san", "dka", "xor", "adi", "add", " so", "o k", " is", "lo ", " mi", "aqa", "na ", " fa", "soo", "baa", " he", "kar", "mid", "dad", "rka", "had", "iin", "a o", "aro", "ado", "aar", "u k", "qaa", " ha", "ad ", "nta", "o h", "har", "axa", "quu", " sa", "n k", " ay", "mad", "u s", " ga", "eed", "aga", "dda", "hii", "aal", "haa", "n l", "daa", "xuq", "o q", "o s", "uqu", "uuq", "aya", "i k", "hel", "id ", "n i", " ee", "nka", " ho", "ina", "waa", "dan", "nim", "elo", "agu", "ihi", "naa", "mar", "ark", "saa", "riy", "rri", "qda", "uqd", " bu", "ax ", "a h", "o w", "ya ", "ays", "gga", "ee ", "ank", " no", "n s", "oon", "u h", "n a", "ab ", "haq", "iri", "o l", " gu", "uur", "lka", "laa", "u a", "ida", "int", "lad", "aam", "ood", "ofk", "dhi", "dah", "orr", "eli", " xi", "ysa", "arc", "rci", "to ", "yih", "ool", "kii", "h q", "a f", " ug", "ayn", "asa", " ge", "sho", "n x", "siy", "ido", "a g", "gel", "ami", "hoo", "i a", "jee", "n q", "agg", "al ", " di", " ta", "e u", "o u", " ji", "goo", "a c", "sag", "alk", "aba", "sig", " mu", "caa", "aqo", "u q", "ooc", "oob", "bar", "ii ", "ra ", "a b", "ago", "xir", "aaq", " ci", "dal", "oba", "mo ", "iir", "hor", "fal", "qan", " du", "dar", "ari", "uma", "d k", "ban", "y d", "qar", "ugu", " ya", "xay", "a j"},
	Ilo: []string{"ti ", "iti", "an ", "nga", "ga ", " ng", " pa", " it", "en ", " ka", " ke", " ma", "ana", " a ", " ti", "pan", "ken", "agi", "ang", "a n", "a k", "aya", "gan", "n a", "int", "lin", "ali", "n t", "a m", "dag", "git", "a a", "i p", "teg", "a p", " na", "nte", "man", "awa", "kal", "da ", "ng ", "ega", "ada", "way", "nag", "n i", " da", "na ", "i k", "sa ", "n k", "ysa", "n n", "no ", "a i", "al ", "add", "aba", " me", "i a", "eys", "nna", "dda", "ngg", "mey", " sa", "pag", "ann", "ya ", "gal", " ba", "mai", " tu", "gga", "kad", "i s", "yan", "ung", "nak", "tun", "wen", "aan", "nan", "aka", " ad", "enn", " ag", "asa", " we", "yaw", "i n", "wan", "nno", "ata", " ta", "l m", "i t", "ami", "a t", " si", "ong", "apa", "kas", "li ", "i m", "ina", " an", "aki", "ay ", "n d", "ala", "gpa", "a s", "g k", "ara", "et ", "n p", "at ", "ili", "eng", "mak", "ika", "ama", "dad", "nai", "g i", "ipa", "in ", " aw", "toy", "oy ", "ao ", "yon", "ag ", "on ", "aen", "ta ", "ani", "ily", "bab", "tao", "ket", "lya", "sin", "aik", " ki", "bal", "oma", "agp", "ngi", "a d", "y n", "iwa", "o k", "kin", "naa", "uma", "daa", "o t", "gil", "bae", "i i", "g a", "mil", " am", " um", "aga", "kab", "pad", "ram", "ags", "syo", "ar ", "ida", "yto", "i b", "gim", "sab", "ino", "n w", " wa", " de", "a b", "nia", "dey", "n m", "o n", "min", "nom", "asi", "tan", "aar", "eg ", "agt", "san", "pap", "eyt", "iam", "i e", "saa", "sal", "pam", "bag", "nat", "ak ", "sap", "ed ", "gsa", "lak", "t n", "ari", "i u", " gi", "o p", "nay", "kan", "t k", "sia", "aw ", "g n", "day", "i l", "kit", "uka", "lan", "i d", "aib", "pak", "imo", "y a", "ias", "mon", "ma ", " li", "den", "i g", "to ", "dum", "sta", "apu", "o i", "ubo", "ged", "lub", "agb", "pul", "bia", "i w", "ita", "asy", "mid", "umi", "abi", "akd", "kar", "kap", "kai", " ar", "gin", "kni", " id", "ban", "bas", "ad ", "bon", "agk", "nib", "o m", "ibi", "ing", "ran", "kda", "din", "abs", "iba", "akn", "nnu", "t i", "isu", "o a", "aip", "as ", "inn", "sar", " la", "maa", "nto", "amm", "idi", "g t", "ulo", "lal", "bsa", "waw", "kip", "w k", "ura", "d n", "y i"},
	Uig: []string{"ish", " he", "ini", "ing", "nin", "gha", "ng ", "ili", " we", "we ", "sh ", "in ", " bo", "quq", "oqu", "ni ", "hoq", " ho", "ush", "shi", "lik", "qil", "bol", "shq", "en ", "lis", "qa ", "hqa", "n b", "hem", " qi", "ki ", "dem", "iy ", " ad", "ade", "igh", "e a", "em ", "han", "liq", "et ", "ge ", "uq ", "nda", "din", " te", " bi", "idi", "let", "qan", "nli", "ige", "ash", "tin", "ha ", "kin", "iki", "her", "de ", " er", " ba", "and", "iti", "olu", "an ", " dö", "döl", "aq ", "luq", " ya", "me ", "lus", "öle", "mme", "emm", " qa", "daq", "rki", "lgh", "erq", "erk", "shk", "esh", "rqa", "iq ", "uqi", "ile", "rim", "i w", "er ", "ik ", "yak", "aki", "ara", "a h", " be", "men", " ar", "du ", "shu", "uql", "hri", "hi ", "qlu", "q h", "inl", "lar", "da ", "i b", "ime", " as", "ler", "etl", "nis", " öz", "ehr", "lin", "e q", "ar ", "ila", " mu", "len", " me", "qi ", "asi", "beh", "a b", "ayd", "q a", "bir", "bil", " sh", "che", "rli", "ke ", "bar", "hke", "yet", "éli", "shl", "tni", "u h", "ek ", "may", "e b", " ké", "h h", " ig", "ydu", "isi", "ali", "hli", "k h", " qo", "iri", "emd", "ari", "e h", "ida", "e t", "tle", "rni", " al", "siy", "lid", "olm", "iye", "anl", " tu", "iqi", "lma", "ip ", "mde", "e e", "tur", "a i", "uru", "i k", "raw", "hu ", "mus", "kil", " is", "i a", "ir ", "éti", "r b", "özi", "ris", "asa", "i h", "sas", " je", "he ", " ch", "qig", "bas", "n q", "alg", "ett", "les", " xi", "tid", " él", "tes", "ti ", "awa", "ima", "nun", "a a", " xe", " bu", "hil", "n h", " xa", "adi", "dig", "anu", "uni", "mni", " sa", "arl", "rek", "ére", " hö", "kér", " ji", "min", "i q", "tis", "rqi", " iy", "elq", "xel", "p q", " qe", "y i", "i s", "lig", " ma", "iya", "i y", "siz", "ani", " ki", "qti", " de", "q w", "emn", "met", "jin", "niy", "i i", "tim", "irl", " ti", "rin", "éri", "i d", "ati", "si ", "tew", "i t", "tli", "eli", "e m", "rus", "oli", "ami", "gen", "ide", "ina", "chi", "dil", "nay", "ken", "ern", "n w", " to", "ayi", " ij", "elg", "she", "tti", "arq", "hek", "e i", "n a", "zin", "r a", "ijt", "g b", "atn", "qar", "his", "uch", "lim", "hki", "dik"},
	Hat: []string{"ou ", "an ", " li", "on ", "wa ", "yon", " po", "li ", "pou", "te ", " yo", "oun", " mo", "un ", "mou", "ak ", " na", "en ", "n p", "nan", "tou", "syo", " dw", " to", "yo ", " fè", "dwa", " ak", " ki", "ki ", " pa", " sa", "out", " la", " ko", " ge", "ut ", "n s", "gen", " de", "se ", "asy", "èt ", "i p", "n d", " a ", " so", "n l", "a a", "fè ", "n k", " se", "pa ", "e d", "u l", " re", "ite", "sa ", " ch", "kon", "n n", "e l", "t p", "ni ", "cha", "a p", "nn ", "ans", "pi ", "t m", " ka", " an", "nm ", "fèt", "i s", "son", "man", " me", "n m", "n a", "e p", "swa", "sou", "e k", "hak", "òt ", "n y", "men", "i l", "epi", " pe", "ote", "san", " ep", "i k", " si", "yen", "eyi", "a l", " ap", "i a", "yi ", "pey", "je ", "n t", "e a", "k m", "e s", " ni", "lib", "e n", "i t", "lit", "ran", "lè ", "enn", "al ", "a s", " pr", "a f", "ns ", " lò", "ap ", "lòt", "enm", "k l", "n e", "t l", "kla", "anm", "e y", "a k", " ma", "e t", "ay ", "i m", "ali", " lè", "è a", "ye ", "a y", "ant", " os", " ba", "i g", " tè", "aso", "u t", "a n", " pw", "ras", " pè", "n f", "nas", "ka ", "n g", "osw", " ta", "dek", "i d", "pwo", "e m", " di", " vi", "la ", "i n", "u s", "sos", "bli", " te", "o t", " tr", "lwa", "ète", "a t", "le ", "u y", "i f", "tan", "a c", "lar", "a m", "ete", "ara", "t k", " pi", "ibè", "bèt", "re ", "osy", "de ", "ati", "ke ", "res", "tis", "i y", "tè ", "nen", " fa", "ekl", "ze ", "nal", "ons", "ksy", "ini", "che", " le", "e r", "a d", " en", "aye", "he ", "o p", "alw", " kò", "lal", " no", "esp", "a g", "ava", "kou", "las", "way", "u f", "isy", " za", " ok", "oke", "kal", "ken", "sye", "ta ", "onn", "k k", "nje", "pra", "van", "esi", "pès", "kot", "ret", "sya", "n v", "lek", "jan", "ik ", "a b", "eks", "wot", "è n", "di ", "òl ", "tra", "u k", "i r", "nou", " as", "k a", "u d", "ist", "èso", "ib ", " ne", "iti", "ti ", "is ", "y a", "des", "è l", "a r", "ont", " ke", "nsa", "pat", "rit", "sit", "pòt", "ona", "ab ", "è s", " sw", "ond", "ide", " ja", "rav", "t a", "ri ", "bon", "viv", " sè", "pre", "vay", "k p", "l l", "kòm", "i o", " ra", "era", "fan", "dev"},
	Aka: []string{"sɛ ", "a a", " sɛ", "ne ", "ra ", "a n", " wɔ", " a ", "ara", "an ", "eɛ ", "no ", " ne", " bi", " no", " as", "iar", "bia", "yɛ ", "mu ", "aa ", " an", "ɛ s", "e a", "ma ", " ho", "bi ", "man", "deɛ", " mu", "ho ", "ɛ a", "na ", "a ɛ", " ob", "obi", "e n", "a b", "n a", "so ", "o n", "pa ", "ama", "ɛ o", "o a", "ipa", "nip", "ɛ n", "naa", " na", "a w", "ana", " so", " ad", " nn", "ɛ ɔ", "ɛde", "asɛ", "kwa", " on", "oni", "wan", " am", "a ɔ", "sɛd", "wɔ ", " ah", "ɛyɛ", " ny", "oɔ ", " n ", "mma", "i a", " mm", "nni", " kw", "ie ", "wɔn", "ɛ w", "de ", " ɛy", " ba", "ase", "ɔ n", "o b", "i m", "ɔ a", "uo ", "n n", "a m", "o s", "iri", " yi", "ni ", "e s", "nyi", "di ", "u n", "a o", "aho", " de", "tum", " ɛn", "ɔn ", "nya", "i n", "ɔma", "e m", "adw", " yɛ", "umi", "die", "mi ", "ɛ ɛ", "o k", " ab", "ɛm ", "a s", " ma", "nam", " ɔm", " ɛs", "yin", " at", " bɔ", "o d", "ina", "pɛ ", "sɛm", "ua ", "n s", "bɔ ", "adi", "ya ", "e h", "aso", "mar", "ani", "kuo", "rɛ ", "fa ", "a k", "ɔde", "a h", "ba ", "n b", "re ", "uma", "wum", "om ", "ɔ h", "m n", "yi ", "u a", " sa", "se ", "dwu", "ɔ b", " nt", "m a", "erɛ", " kɔ", "a y", "orɔ", " nk", " bɛ", " ɔd", "ten", "rɔ ", "hyɛ", "saa", "ka ", "ɛ b", "e b", "i s", "ade", "am ", "nka", "kor", "i ɛ", "ene", "ena", " ns", "ban", "ɛns", " ku", "ɛsɛ", "ane", "nsɛ", "fof", "ɛɛ ", " fi", "gye", "ɔtu", " di", "ano", "i k", "o m", " ɔt", " ko", "yɛɛ", "bir", " ak", "im ", "kye", " pɛ", "a d", "yie", "ko ", "nti", "i b", "ete", "ofo", "amm", "ye ", "ri ", "foɔ", "kɔ ", "bom", "abo", "ɔ s", "ɔne", " ɛb", "soɔ", "for", "isɛ", "m k", "asa", "nod", "ɛ m", "fir", "ti ", " da", "e y", "sua", " be", "nii", "seɛ", "wa ", "ber", " aw", "dwe", "n f", " fo", "o ɛ", "i h", "u b", "ɔ m", " mf", "hɔ ", "kab", "wɛ ", "to ", "rib", "hwɛ", "ibi", " dw", "dis", "nso", "ans", "tir", "u ɛ", " ti", " hɔ", "sa ", "e o", " tu", "odi", "ɛ y", "ia ", "ofa", " ɔn", "o w", "ɛbɛ", "aba", " ka", "ii ", "wen", "ɛsi", "m m", "sia", "ada", "yer", "ian", "da ", "set", " gy", "dua", "i d", "som", "mfa", "ɔ w", " af", "i y", "any", "ora", "rim", "wɔd", "dwa", "nsi"},
	Sna: []string{"wa ", "a k", "ana", "ro ", "na ", " ku", " mu", "nhu", "dze", "hu ", "a m", " zv", "mun", "oku", "chi", "a n", "aka", "dzi", "ka ", "zer", "ero", " ch", "che", "se ", "unh", "odz", "rwa", "ra ", "kod", "zvi", " ne", " pa", "kan", " we", " dz", " no", "ika", "va ", "iri", " an", "kut", "nyi", "o y", "yik", "van", "nek", "ese", "eko", "zva", "idz", "e a", " ka", "ane", "ano", "ngu", "eku", "cha", "ung", " yo", "ri ", "ake", "ke ", "ach", "udz", "iro", "a z", "u w", " va", "ira", "wes", "ang", "ech", "nge", "i p", "eng", "yok", "nok", "edz", "o i", "irw", "ani", "ino", "uva", "ich", "nga", "ti ", "zir", "anh", "rir", "ko ", "dza", "o n", "wan", "wo ", "tan", "sun", "ipi", "dzw", "eny", "asi", "hen", "zve", "kur", "vak", "a p", "sha", "unu", "zwa", "ita", "kwa", "e k", "rud", "nun", "uru", "guk", "a c", "a d", " ya", "a y", "bat", "pas", "ezv", "ta ", "e n", "uti", " kw", "o k", "o c", "o m", "ara", " ma", "si ", "ga ", "uko", "ata", "ose", "ema", "dzo", "uch", "hip", "kuv", "no ", "rus", "hec", "omu", "i z", "wak", "o r", "kus", "kwe", "ere", "re ", " rw", " po", "o a", "mwe", "yak", "mo ", "usu", "isi", "za ", "sa ", "e z", "uta", "gar", " in", "hin", "nem", "pac", "kuc", "we ", "ete", " ye", "twa", "pos", "o d", "a i", "hur", "get", "ari", "ong", "pan", "erw", "uka", "rwo", "vo ", " ak", "tem", "zo ", "emu", "emo", "oru", " ha", "uit", "wen", "uye", "kui", " uy", "vin", "hak", "kub", "i m", "a a", "kud", " se", " ko", "yo ", "and", "da ", "nor", "sin", "uba", "a s", "a u", " ic", "zvo", "mut", "mat", "nez", "e m", "a w", "adz", "ura", "eva", "ava", "pi ", "a r", "era", "ute", "oko", "vis", " iy", "ha ", "u a", "han", "cho", "aru", "asa", "fan", "aan", "pir", "ina", "guv", "ush", "ton", " hu", "uny", "enz", "ran", "yor", "ted", "ait", "hek", " ny", "uri", "hok", "nen", "osh", " ac", "ngi", "muk", "ngo", "o z", "azv", "kun", "nid", "uma", "i h", "vem", "a h", "mir", "usa", "o p", "i n", "a v", "i k", "amb", "zan", "nza", "kuz", "zi ", "kak", "ing", "u v", "ngw", "mum", "mba", "nir", "sar", "ewo", "e p", "uwa", "vic", "i i", "gwa", "aga", "ama", "go ", "yew", "pam"},
	Fin: []string{"en ", "ise", "ja ", "ist", " ja", "on ", "ta ", "sta", "an ", "n j", "ais", "sen", "n o", "keu", "ike", "oik", "lis", " va", "ell", "lla", "n t", "uks", " on", "ksi", " oi", "n k", " ka", "aan", "een", "la ", "lli", "kai", "a j", " ta", "sa ", "in ", "mis", " jo", "a o", "ään", "än ", "sel", "n s", "kse", "a t", "a k", "tai", "us ", "tta", "ans", "ssa", "kun", "den", "tä ", "eus", "nen", "kan", "nsa", "apa", "all", "est", " se", "eis", "ill", "ien", "see", "taa", " yh", "jok", "n y", "vap", "a v", "ttä", "oka", "n v", "ai ", "itt", "aa ", "aik", "ett", "tuk", "ti ", "ust", " ku", "isi", "stä", "ses", " tä", " tu", "lai", "n p", "sti", "ast", "n e", "n m", "tää", "sia", "unn", "ä j", "ude", "ä o", "ste", "si ", "tei", "ine", "per", "a s", "ia ", "kä ", "äne", " mi", "maa", " pe", "a p", "ess", "a m", "ain", "ämä", "tam", "yht", " ju", "jul", "yks", "hän", "ä t", " hä", "utt", "ide", "et ", "llä", "val", "sek", "stu", "n a", "lä ", "ami", "hmi", " ke", "ikk", "lle", "iin", "sä ", "euk", "täm", "ihm", "tee", " ih", "lta", "pau", " sa", "isk", "mää", "ois", "un ", "tav", "ten", "dis", "hte", "n h", "iss", "ssä", "a h", "ava", " ma", "a y", " ei", " te", " si", " ol", "ekä", "sty", "alt", "toi", "att", "oll", "tet", " jä", " ra", "vat", " mu", "iel", " to", "mai", "sal", "isu", "a a", "kki", "at ", "suu", "n l", "väl", "ää ", "uli", "tun", "tie", "eru", " yk", "etu", "vaa", "rus", "muk", " he", "ei ", "a e", "kie", "sku", "eid", "iit", " su", "nna", "sil", "oma", "min", " yl", "lin", "aut", "uut", "sko", " ko", "tti", "le ", "sie", "kaa", "a r", " ri", "sii", "nno", "eli", "tur", "saa", "aat", "lei", "oli", "na ", " la", "oon", "urv", "lma", "rva", "ite", "mie", "vas", "ä m", " ed", "tus", "iaa", "itä", "ä v", "uol", "yle", " al", "lit", "suo", "ama", "joi", "unt", "ute", "i o", "tyk", "n r", "ali", "lii", "nee", "paa", "avi", "omi", "oit", "jen", "kää", "voi", "yhd", "ä k", " ki", "eet", "eks", " sy", "ity", "ilö", "ilm", "oim", "ole", "sit", "ita", "uom", "vai", "usk", "ala", "hen", "ope", " pu", "auk", "pet", "oja", "i s", "rii", "uud", "hdi", "äli", "va ", " om"},
	Run: []string{"ra ", "we ", "wa ", " mu", "e a", "se ", " n ", "a k", "ira", "ntu", "tu ", " ku", " um", "ko ", "a i", "mu ", "iri", "mun", "hir", "ye ", "unt", "ing", "ash", "ere", "shi", "a n", "umu", "zwa", " bi", "gu ", "ege", "a a", "za ", "teg", "ama", "e k", "go ", "uba", "aba", "ngo", "ora", "o a", "ish", " ba", " ar", "ung", "a m", " we", "e n", "na ", "sho", "ese", "nga", " ab", "e m", "mwe", "ugu", " kw", "ndi", " gu", "ate", "kwi", "wes", "riz", "ger", "u w", " at", "di ", "gih", "iza", "n u", "ngi", "ban", "yo ", "ka ", "e b", "a b", " am", " ca", "ara", "e i", "obo", "hob", "ri ", "u b", "can", "nke", "ro ", "bor", " in", "bah", "ahi", "ezw", "a u", "gir", "ke ", "igi", "iki", "iwe", "rez", "ihu", "hug", "aku", "ari", "ang", "a g", "ank", "ose", "u n", "o n", "rwa", "kan", " ak", "nta", "and", "ngu", " vy", "aka", "n i", "ran", " nt", " ub", "kun", "ata", "i n", "kur", "ana", "e u", " ko", "gin", "nye", "re ", " ka", "any", "ta ", "uko", "amw", "iye", " zi", "ga ", "ite", " ib", "aha", " ng", "era", "o b", "ako", "o i", " bu", "o k", "o u", "o z", " ig", "o m", "ho ", "mak", "sha", " as", " iv", "ivy", "n a", "i b", "izw", "o y", " uk", "ubu", "aga", "ba ", "kir", "vyi", "aho", " is", "nya", "gan", "uri", " it", " im", "u m", "kub", "rik", "hin", "guk", "ene", "bat", "nge", "jwe", "imi", " y ", "vyo", "imw", "ani", "kug", "u a", "ina", "gek", "ham", "i i", "e c", "ze ", "ush", "e y", "uru", "bur", "amb", "ibi", "agi", "uza", "zi ", "eye", "u g", "gus", "i a", " nk", "no ", "abi", "ha ", "rah", "ber", "eme", "ras", "ura", "kiz", "ne ", "tun", "ron", " zu", "ma ", "gen", "wo ", "zub", "w i", "kor", "zin", "wub", "ind", " gi", "y i", "ugi", "je ", "iro", "mbe", " mw", "bak", " ma", "ryo", "eka", "mat", " ic", "onk", "a z", " bo", "ika", "eko", "ihe", "ukw", "wir", "bwa", " ry", " ha", "bwo", " ag", "umw", "yiw", "tse", " ya", "he ", "eng", " ki", "nka", "bir", "ant", "aro", "gis", "ury", "twa", " yo", "bik", "rek", "ni ", " ah", " bw", "uro", "mw ", "tan", "i y", "nde", "ejw", " no", "zam", "puz", "ku ", "y a", "a c", "bih", "ya ", "mur", "utu", "eny", "uki", "bos"},
	Tuk: []string{"lar", " we", "we ", " bi", "yň ", "ary", "ada", "da ", " he", " ha", "an ", "yny", "kla", "dam", "de ", " ad", "yna", "er ", "na ", " ýa", "ir ", "dyr", "iň ", "bir", "r b", "ydy", "ler", "ara", "am ", "yr ", "ini", "lan", "r a", "kly", "lyd", " öz", "mag", "nyň", "öz ", "her", "gyn", "aga", "en ", "ryn", "akl", "ala", "dan", "hak", "eri", "ne ", "uku", "ar ", "r h", "ga ", "ny ", "huk", " de", "ili", "ygy", "li ", "kuk", "a h", "nda", "asy", "len", " ed", "bil", "atl", "ine", "edi", "niň", "lyg", " hu", " ga", "e h", "nde", "dil", "ryň", "aza", "zat", "a g", "‐da", "a‐d", "eti", "ukl", " gö", "ly ", " bo", "tly", "gin", " az", "lma", "ama", "hem", "dir", "ykl", "‐de", "e d", "ile", "ýan", "a d", "ýet", "ýa‐", "ynd", "lyk", "aýy", "e a", "ünd", "ge ", " go", "egi", "ilm", "sy ", "ni ", "etm", "em‐", "lme", "m‐d", "aly", "any", " be", "tle", "syn", "rin", "y b", "let", "mak", "a w", "a ý", "den", "äge", "ra ", " äh", "mäg", " du", "n e", "bol", "meg", "ele", "ň h", " et", "igi", "ň w", "im ", "iýa", " ýe", " di", "r e", "ek ", " ba", "ak ", "esi", "ril", "a b", "in ", "p b", "deň", "etl", "agy", " bu", " je", "bu ", "e ö", "y d", " hi", "mez", " es", "ard", " sa", "ähl", "e b", "yly", " ka", "esa", "mek", " gu", "n a", "e t", "lik", " do", "e g", "sas", "ill", "nma", "ň a", "ram", "ola", "hal", "y w", "ýar", " ar", "anm", "mel", "iri", "siý", "ndi", "ede", "gal", "end", "mil", "rla", "göz", " ma", "n b", "e ý", "öňü", "ňün", "n h", " tu", "hiç", "yýe", " ge", "my ", "iç ", " öň", "n ý", "tla", "ň ý", "lin", "rda", "al ", "lig", "gar", " mi", "i g", "dal", "rle", "mal", "kan", "gat", "tme", "sin", "and", "ň g", "gor", " ta", "öwl", "ýle", "y g", "e w", "ora", "tiň", "ekl", " yn", "alk", "döw", " dö", "ere", "m h", " me", "dur", " er", "asi", "tut", "at ", "çin", "irl", "umy", "eli", "erk", "nme", "wle", "gur", "a ö", "aýa", " çä", "nun", " ki", "ras", "aml", "up ", "ýaş", "tyn", " aý", "ry ", "ň d", "baş", "ip ", "gi ", "z h", "kin", "z ö", "n w", "ter", "inm", "eýl", "i ý", "kim", "nam", "eň ", "beý", "dol", " se", " te", "r d", "utu", "gyý", "ez ", "umu", "mum"},
	Dan: []string{"er ", "og ", " og", "der", " de", "for", "en ", "et ", "til", " fo", " ti", "ing", "de ", "nde", "ret", " re", "hed", "il ", "lig", " ha", "lle", "den", " en", "ed ", "ver", "els", "und", "ar ", " fr", " me", "se ", "lse", "and", "har", "gen", "ede", "ge ", "ell", "ng ", "at ", " af", "nne", "le ", "nge", "e f", "ghe", "e o", "igh", "es ", "af ", "enn", " at", "ler", " i ", "ske", "hve", "e e", "r h", "ne ", "enh", "t t", "ige", "esk", " el", " be", "ig ", "tig", "fri", "or ", "ska", "nin", "e s", "ion", " er", "nhv", "re ", "men", "r o", "e a", " st", "ati", " sk", " in", "l a", "tio", " på", "ett", "ens", "al ", "tti", "med", "r f", "om ", "end", "r e", "del", "g f", "ke ", " so", "på ", "eli", "g o", " an", "r r", "ns ", " al", "nat", "han", " ve", "r s", "r a", " un", " he", "t f", "lin", " si", "r d", "ter", "ere", "nes", "det", "e r", " ud", "ale", "sam", "ihe", "lan", "tte", "rin", "rih", "ent", "ndl", "e m", "isk", "erk", "ans", "t s", "kal", " na", "som", "hol", "lde", "ind", "e n", "ren", "n s", "ner", "kel", "old", "dig", "te ", "ors", "e i", " hv", "sni", "sky", "ene", "vær", " li", " sa", "s f", "d d", "ers", "ste", "nte", "mme", "ove", "e h", "nal", "ona", "ger", " gr", "age", "g a", "vil", "all", "e d", "fre", "tel", "s o", "g h", "t o", "t d", "r i", "e t", " om", "arb", "d e", "ern", "r u", " væ", "d o", "res", "g t", "klæ", "øre", "n f", " vi", " må", "ven", "sk ", " la", "gte", "kab", "str", "n m", "rel", "e b", "run", "rbe", "bej", "t i", "ejd", "kke", "t e", "g d", "rkl", "ilk", "gru", "ved", "bes", " da", "nd ", " fu", "lær", "æri", "rdi", "ærd", "ld ", "t m", "dli", "fun", "sig", " mo", "sta", "nst", "rt ", "od ", " ar", " op", "vis", "igt", "ære", "tet", "t a", "emm", "g e", "mod", "rho", "ie ", "g u", "ker", "rem", " no", "n h", " fa", "rsk", "orm", "e u", "s s", "em ", "d h", " ge", "ets", "e g", "g s", "per", " et", "lem", " tr", "i s", "da ", "dre", "n a", "des", "dt ", "kyt", "rde", "ytt", "eri", "hen", "erv", "l e", "rvi", "ffe", "off", "isn", "r t", " of", "ken", "l h", "rke", "g i", "tal", "må ", "r k", "lke", "gt ", "t v", "t b"},
	Nob: []string{"er ", " og", "og ", "en ", " de", "for", "til", "ing", "ett", " ti", "et ", " ha", " fo", " re", "ret", "il ", "het", "lle", "ver", "tt ", "ar ", "nne", " en", "om ", "ell", "ng ", "har", " me", "enn", "ter", "de ", "lig", " fr", " so", "r h", "ler", "av ", "le ", "den", "and", " i ", " er", "som", " å ", "hve", "or ", "t t", "ne ", " el", "els", "re ", " av", "se ", "esk", "enh", "nge", "ska", "nde", "e o", "ete", "gen", "ke ", "lse", "ghe", "ten", "men", " st", "r s", "fri", "igh", "ig ", " be", "e e", "nhv", "r r", "tte", "ske", "te ", " på", " ut", " sk", "al ", " in", "sjo", "på ", "der", "e s", "ner", "rin", "jon", "t o", "unn", "e f", "han", "asj", "tig", "ed ", "es ", "g f", "sam", "ent", "tti", "ene", "nes", "med", "ge ", " al", "r o", "ens", "r e", "eli", "isk", "lin", " ve", "nin", "g o", " sa", " an", "t f", "itt", "lik", "end", "kal", "r f", "t s", "rih", "ihe", "nas", "nte", "e r", "ns ", " si", "lan", "g s", "mme", "ige", "l å", "erk", "dig", " gr", "n s", "ren", "r a", "all", " na", "kte", "erd", "ere", "e m", "und", "r u", "res", "tel", "ste", "gru", "inn", "lær", "ers", " un", "det", "t e", "arb", "ale", "del", "ekt", "ven", "t i", "g e", "bei", "eid", "e a", "n m", "e d", " ar", "rbe", "e g", " bl", "ans", "klæ", " li", " he", "g t", "æri", "sky", "run", "rkl", " la", "sta", "sni", "kke", "m e", "rt ", "mot", " mo", "e n", "tat", "at ", "e h", "e b", "ove", "e t", "jen", "t d", "str", " må", "r m", "n e", "ors", "rel", "ker", " et", "n a", "bes", "one", " vi", "nn ", "g r", "e i", "kap", "sk ", "ot ", "ndi", "nnl", "i s", " da", "s o", " no", "id ", "ger", "g h", "vis", "n o", "bar", "s f", "ndl", "t m", "g a", "opp", "t a", "dis", "nal", "r d", "per", "dre", "ona", "ære", "rdi", "da ", "ute", "nse", "bli", "ore", "tet", "rit", " op", "kra", "eri", "hol", "old", " kr", "ytt", "kyt", "ffe", "emm", "g d", "l f", " om", "isn", " gj", "å d", "ser", "r b", " di", " fa", "n t", "r k", "lt ", "set", " sl", "dom", "rvi", "me ", "l e", "gre", "å s", "må ", " tr", "nd ", "m s", "g i", "ikk", "n h", " at", "tes", "vil", "dli", "g b", "d d", " hv", "rav"},
	Nno: []string{" og", "og ", " de", " ha", "er ", "en ", "ar ", "til", " ti", "lle", "ett", "il ", "ret", "om ", "et ", " re", "le ", "har", "enn", " me", " al", "all", " fr", "ne ", "tt ", "re ", " å ", " i ", "nne", "and", "ing", "ska", " sk", "men", " fo", "det", "den", "ver", "for", "ell", "t t", "dom", " so", "de ", "e s", " ve", " ei", "ere", " på", "al ", "an ", "e o", "e h", "fri", "sam", " sa", "l å", "på ", "leg", " el", "ler", "som", "ein", "ei ", "nde", "av ", " st", "dei", "or ", "ten", "esk", "kal", "gje", "n s", "tte", "je ", "ske", "rid", "r r", "i s", "te ", "nes", " gj", "eg ", "ido", "med", "e f", "r s", "st ", "ke ", "jon", " in", "r f", "sjo", "asj", "nas", "ter", "unn", "ed ", "kje", "han", "ona", " er", "t o", "t e", "g f", "ski", "e m", "ast", "ane", "e t", " av", " gr", "lan", "ste", "tan", "å f", " na", "der", " sl", "t s", "seg", "n o", "r k", "nga", "ge ", " an", "g o", "at ", "na ", "ern", "nte", "ng ", " ut", "lik", "e a", "bei", "gru", "e i", "arb", "kil", "g s", "lag", "eid", "r a", "e d", "g d", " si", " få", "ame", "a s", "e r", "rbe", "jen", "n m", "r d", "n e", "nn ", "e n", "erd", " tr", " må", " bl", " mo", "ren", "run", "nin", "bli", "kra", " kr", " at", "ege", "n i", "me ", "nsk", "ins", "år ", "frå", "in ", "lov", "v p", "end", "mot", "ale", "e v", "å a", "få ", "rav", "int", "nal", " ar", "sta", "e k", "t f", "ome", " la", "ot ", "t a", "sla", " ik", "nle", "itt", " li", " kv", "id ", "kkj", "ikk", " lo", "nad", "å v", "tta", " fa", " se", "gen", "ld ", "å s", "kan", "g t", " ka", "r l", "god", "n a", "lin", "jel", "ild", "dig", "ha ", "l d", "kap", "ve ", "ndr", "g i", "g a", "inn", "var", "rna", "r m", "r g", "a o", "dre", "d a", "n t", "ag ", "kår", "mål", "ig ", "va ", "i d", "t m", "e e", "n d", "tyr", " om", "g e", "eve", "då ", "e u", " då", "und", " no", "ir ", "gar", "g g", "l h", "se ", "ga ", "d d", "l f", "ker", "r o", "å d", "eld", "ige", "t d", "t i", "t h", "oko", "nnl", "rel", "nok", "rt ", "lt ", "åse", "jer", "ta ", "ik ", "ial", "eig", "r p", "i e", "olk", "bar", "osi", "kte", "sos", "lir", "opp", " un", "ad ", " be"},
	Lit: []string{"as ", "ir ", " ir", "eis", "tei", " te", "s t", "os ", "uri", "ti ", "us ", "is ", "iek", " pa", "ai ", " vi", "vie", "tur", " ki", "ri ", "žmo", " tu", " žm", "ien", "ės ", "ių ", "ali", "ais", "mog", "vis", " ka", "lai", " la", "ini", "i t", "s i", "s ž", "sę ", " į ", "isę", "ena", " ne", " pr", " bū", " jo", "pri", "kie", " ta", "kvi", "nas", " su", "ekv", "mas", "gus", "būt", "tin", "isv", "s s", "ogu", "isi", "mą ", "mo ", "ant", " ar", "s k", "ama", "kai", "ūti", "s a", "s v", "aci", " ti", "s n", " sa", "s p", "oki", "cij", "inė", "ar ", "val", "ms ", "tai", "jo ", "i b", " na", "gal", "sav", "kur", "aus", "men", "rin", " ap", "imą", "ma ", "sta", "ę į", "ina", "i p", "imo", "nim", "i k", " nu", "ima", "oti", "mis", " ku", "jos", "lyg", "dar", "išk", "je ", " at", "tas", "kad", "r t", "tų ", "ad ", "tik", "i i", "nės", "arb", "i v", "ijo", "eik", "aut", "s b", " įs", " re", "iam", "sin", "suo", " be", "isu", " va", "li ", "sty", "asi", "tie", "ara", "lin", "isė", "i s", "ą i", "jų ", " ly", " ga", "vo ", "si ", "r p", "tuo", "aik", "rie", " mo", "din", "pas", "mok", "ip ", "i n", "rei", "ybė", "mos", "aip", "r l", "ntu", "įst", "į t", "gyv", " iš", "nti", "tyb", "ų i", "pag", "kia", "kit", "es ", "uot", " sk", "jim", "tis", " or", "aud", "yve", "ven", "mų ", "als", "ų t", "nac", "avo", "dam", "ą k", "i a", "s j", "oje", "agr", "kla", "gau", "neg", "nių", "o k", "ega", "iki", "aug", "ek ", "tat", "ieš", "tar", "ia ", " ši", "ios", "ška", "sva", " to", "tau", "int", "sau", "uti", " as", "io ", "oga", "san", "mon", "omi", "kin", "ito", "s g", "ome", "r j", " ve", "aty", "kim", "nt ", "iai", "lst", " da", "ją ", "min", "r k", "o t", "nuo", "tu ", "ver", "kal", "am ", "usi", "o n", "o a", "ymo", "tym", "vę ", "ati", " ji", "o p", "tim", "ų n", "paž", "ter", "s š", " vy", "alt", "ksl", "ing", "ų s", "oma", "šal", "ran", "e t", " ni", " ša", "ava", "avi", "nie", "uom", "irt", "elg", "jam", "ipa", "kių", "tok", "eka", "tos", "oja", "kio", "eny", "nam", "s d", "ndi", "amo", "yti", "gri", "svę", " gy", "lie", "ėmi", "ats", "ygi", "soc", "sie", "oci", "pat", "cia"},
	Slv: []string{" pr", "in ", " in", "rav", "pra", "do ", "anj", "ti ", "avi", "je ", "nje", "no ", "vic", " do", "ih ", " po", "li ", "o d", " za", " vs", "ost", "a p", "ega", "o i", "ne ", " dr", " na", " v ", "ga ", " sv", "ja ", "van", "svo", "ako", "pri", "co ", "ico", "i s", "e s", "o p", " ka", "ali", "stv", "sti", "vsa", " ne", " im", "sak", "ima", "jo ", "dru", "nos", "kdo", "i d", "akd", "i p", "nja", "o s", "nih", " al", "o v", "ma ", "i i", " de", "e n", "pre", "vo ", "i v", "ni ", "red", "obo", "vob", "avn", "neg", " bi", "ova", " iz", "ove", "iti", "lov", "ki ", "jan", "a v", "na ", " so", "em ", " nj", "a i", "se ", " te", "tva", "oli", "bod", "ruž", "e i", " ra", " sk", "ati", "e p", "aro", "i k", " ob", "a d", " čl", "eva", "rža", "drž", " sp", "ko ", "i n", " se", " ki", "ena", "sto", "e v", "žen", "nak", "kak", "i z", "var", "ter", "žav", " mo", "di ", "gov", "imi", "va ", "kol", "n s", " z ", "mi ", "ovo", "rod", "voj", " en", "nar", "ve ", " je", "pos", "a s", "ego", "vlj", "jeg", " st", "h p", "er ", "kat", "člo", "ate", "a z", "enj", "n p", "del", "i o", "lja", "pol", "čin", "a n", "ed ", "sme", "jen", "eni", " ta", "odn", " ve", " ni", "e b", "en ", " me", "jem", "kon", "nan", "elj", "sam", "da ", "lje", "zak", "ovi", "šči", "raz", "ans", "ju ", "bit", "ic ", " sm", "ji ", "nsk", "v s", " s ", "n v", "tvo", "ene", "a k", "me ", "vat", "ora", "krš", "nim", "sta", "živ", "ebn", "ev ", "ri ", "eko", "o k", "n n", "so ", "za ", "ičn", "ski", "e d", " va", "o z", "aci", "cij", "eja", "elo", "dej", "si ", "nju", "vol", "kih", "i m", "nst", "kup", "kov", "uži", "la ", "mor", "vih", " da", "h i", "lju", "otr", "med", "o a", "sku", "rug", "odo", "ijo", "dst", "spo", "tak", "zna", "edn", "vne", "ara", "ršn", "itv", "odi", "u s", "čen", "boš", "nik", "avl", "akr", "e o", "vek", "dno", "oln", "o o", "ošč", "e m", "ta ", "vič", "bi ", "pno", "čno", "mel", "eme", "olj", "ode", "rst", "rem", "ov ", "ars", " bo", "n d", "ere", "dov", "ajo", "kla", "ice", "vez", "vni", " ko", "ose", "tev", "bno", "užb", "ava", "ver", "e z", "ljn", "mu ", "a b", "vi ", "dol", "ker", "r s"},
	Epo: []string{"aj ", " la", "la ", "kaj", " ka", "oj ", " de", "on ", "de ", "raj", " ra", "iu ", "ajt", "as ", "o k", " ĉi", "e l", "j k", " li", " pr", "eco", "aŭ ", "ĉiu", "jn ", "ia ", "jto", "est", " es", " al", "an ", " ki", "pro", "io ", " ko", "en ", "n k", "kon", " ti", "co ", "j p", "o d", " po", "ibe", " aŭ", "ro ", "tas", "lib", "ber", "aci", "toj", " en", "a p", " ne", "cio", "ere", "ta ", " in", "to ", "do ", "o e", "j l", "n a", "j d", " se", "a k", "j r", "ala", "j e", "taj", " re", "rec", "iuj", "kiu", " pe", "o a", "ita", "ajn", "ado", "n d", "sta", "nac", "a a", "nta", "lia", "ekt", "eni", "iaj", "ter", "uj ", "per", "ton", "int", " si", "cia", " ha", "stu", "a l", "je ", " je", "al ", "o ĉ", "n p", "jta", "tu ", " ri", "vas", "sen", "hav", "hom", " di", " ho", "nte", "a e", "ali", "ent", " so", "nec", "tra", "a s", "ava", "por", "a r", " na", "igi", "tiu", "sia", "o p", "n l", "ega", "or ", " aj", "soc", "j ĉ", "s l", "oci", "no ", " pl", "j n", "kto", "evi", "s r", "j s", "ojn", "laj", "u a", "re ", " eg", "j a", "gal", "ers", "ke ", "pre", "igo", "er ", "lan", "n j", "pri", " ku", "era", "ian", "rim", " fa", "e s", " ju", "e a", "ika", "ata", "ntr", "el ", "is ", "u h", "li ", "ioj", "don", "ont", "tat", "ons", " el", " su", "go ", "un ", " ke", "ebl", "bla", "n s", "oma", "ĉi ", "raŭ", "kla", "u r", "ne ", "ili", "iĝo", "o t", "s e", "tek", "men", "nen", "j i", "nda", "con", "a d", "ena", "cev", "moj", "ice", "ric", "ple", "son", "art", "a h", "o r", "res", " un", "u s", "coj", "e p", "ĝi ", "for", "ato", "ren", "ara", "ame", "tan", " pu", "ote", "rot", " ma", "vi ", "j f", "len", "dis", "ive", "ant", "n r", " vi", "ami", "iĝi", "sti", "ĝo ", "r l", "n ĉ", "u l", " ag", "erv", "u e", "unu", "gno", " ce", " me", "niu", "iel", "duk", "ern", " ŝt", "laŭ", "o n", "lab", "olo", "abo", "tio", "bor", "ŝta", "imi", " ed", "lo ", "kun", "edu", "kom", "dev", "enc", "ndo", "lig", "e e", "a f", "tig", "i e", " kr", " pa", "na ", "n i", "kad", "and", "e d", "mal", "ono", "dek", "pol", "oro", "eri", "edo", "e k", "rso", "ti ", "rac", "ion", "loj", "j h", "pli", "j m"},
	Lav: []string{"as ", "ība", " un", "un ", "tie", "ies", "bas", "ai ", " ti", "esī", "sīb", "ien", " vi", "bu ", "vie", "ir ", " ir", "ību", "iem", " va", " pa", "em ", " ne", "s u", "am ", "m i", "šan", "u u", "r t", "pie", " ci", " sa", "ās ", " uz", "vai", " ka", " pi", "brī", " iz", "rīv", " br", "uz ", "cij", "dzī", "ena", " ar", "ar ", "isk", "s p", "es ", " at", "āci", " ap", "ot ", "nam", "viņ", "inā", "ikv", "kvi", " no", "s v", " ie", "vis", " ik", "i i", "pār", "u a", "ju ", "nu ", " pr", "edr", "vīb", "īvī", "iju", "drī", "u p", "dar", " st", "lvē", "cil", "ilv", "s t", " la", "iņa", "ana", "s i", "n i", "īdz", "s s", "kā ", "tīb", "i a", "ija", "bai", "ībā", "ied", "s n", "arb", "val", "līd", "s b", "aiz", "tu ", "iec", "cie", "ām ", "gu ", "vēk", "īgu", "īgi", "ka ", "jas", "umu", "mu ", "t p", " jā", "u v", "zīb", "ska", "lst", "als", "kum", "gi ", "s l", " tā", "jot", "stā", "st ", "n v", "vēr", "a p", "arī", "aut", "n p", "ama", "kas", "u k", " da", " ta", "nīg", "izs", "ojo", "anu", "ņa ", "u n", "sta", "s a", "ba ", " ai", " so", "s d", "a u", "ā a", "stī", "cīb", "m u", "i u", "son", "not", "mat", "sav", "iev", "ā v", "jum", " kā", "u t", "ned", "ajā", "s k", "u i", "i v", "līt", "ēro", " pe", " dz", "i n", "per", "u d", "īks", "kat", "nāt", "līb", "nāc", "rdz", "nīb", "pil", "rīk", "kst", "a s", "cit", "pam", " pā", "ekl", "tau", "u s", "bie", "jā ", " re", "i p", "kur", "a a", "t v", " li", "evi", "tis", "evē", "bā ", "ma ", "rīb", "a v", "os ", "ras", "abi", "nev", "iku", "skā", " ve", "lik", " lī", "nas", "t k", "ant", "uma", "roš", "kād", "zsa", "sar", "ciā", "mie", "ais", "eci", "oci", "oša", " je", "jeb", "būt", "atr", "n b", "ieš", "rso", "ers", "soc", "enā", "a t", "t s", "īša", " be", "bez", "āda", "ebk", " ku", "glī", "isp", "tot", "spā", "roj", "lie", "pre", "ret", "aul", "na ", "tra", "iet", "du ", "zgl", "āt ", "ard", "kt ", "ier", "izg", "ikt", "paš", "iāl", "nod", "ts ", "eja", "ā u", "sab", "eno", "ēt ", "ta ", "tik", "tīt", "ecī", " de", "īga", "tar", "arp", "r j", "īst", "tās", "ja ", "enī", "atv", "vu ", "ārē", "rēj", "rie", "oši", "dro"},
	Est: []string{"sel", "ja ", " ja", "le ", "se ", "ust", "ste", "use", "ise", "õig", "mis", " va", "gus", "ele", "te ", "igu", "us ", "st ", "dus", " õi", " võ", " on", "on ", "e j", " in", "ini", "nim", "ma ", "el ", "a v", "iga", "ist", "ime", "al ", "või", "da ", " te", "lik", " ig", "adu", "mes", "ami", "end", "e k", "e v", "l o", " ka", "est", " ra", " se", "õi ", "iku", " ko", "vab", "aba", "tus", "ud ", "a k", "ese", " ku", "l i", "gal", "tsi", "lt ", "es ", "ema", "ida", "ks ", "a i", "n õ", "lis", "atu", "rah", "tam", "ast", "sta", "e t", "s s", " mi", "ta ", "ole", "stu", "bad", "ga ", "val", "ine", " ta", "ne ", " pe", "nda", "ell", "a t", "ali", "ava", "ada", "a p", "ik ", "kus", "e s", "ioo", "tes", "ahe", "ing", "lus", " ol", "a a", "is ", "vah", "a s", "ei ", " ei", "kon", "vas", "tud", "ahv", "t k", "as ", "a r", "s t", "e e", "i v", "eks", "oon", "t v", "oni", "kõi", "s k", "sio", "sus", "e a", "gi ", "mat", "min", " pi", "s v", "oma", "kul", "dad", " ni", "e p", " om", "igi", "tel", "a j", "e o", "ndu", "dse", "lle", "ees", "tse", "uta", "vus", "aal", "aja", "i t", "dam", "ats", "ni ", "ete", "pid", "pea", "e õ", "its", "lma", "lev", "nis", "dis", "ühi", "sli", "i s", "nen", "iel", "des", "de ", "t i", "et ", "nin", "eva", "teg", "usl", "elt", "ili", "i m", "ng ", " ee", "tem", "ses", "ilm", "sek", "ab ", " põ", "ait", " ne", "õrd", "sed", "võr", "ul ", " üh", " ki", "abi", " kõ", "ega", "rds", " vä", "ots", " et", " ri", "põh", "ed ", "töö", "si ", "ad ", "i k", " tä", "ata", " ab", " su", "eli", " sa", "s o", "s j", "sil", "nni", "ari", "asu", "nna", " al", "nud", "uma", "sik", "hvu", "onn", "eab", "emi", "rid", "ara", "set", "e m", " ke", "a e", "täi", "d k", "s p", "i e", "imi", "eis", "e r", "na ", " ül", "a ü", "koh", "a o", "aks", "s e", "e n", " so", "õik", "saa", "and", "isi", "nde", "tum", "hel", "lii", "kin", "äär", "sea", "isk", "een", "ead", "dum", " kä", "rii", "rat", "lem", "umi", "kor", "sa ", "idu", "mus", "rit", "har", " si", "vad", "ita", "ale", "kai", "teo", " mõ", "ade", "üks", "mas", "lse", "als", "iaa", "sia", "sot", "jal", "iig", "ite"},
}

var cyrillicLangs = langProfileList{
	Rus: []string{" пр", " и ", "рав", "ств", " на", "пра", "го ", "ени", "ове", "во ", " ка", "ани", "ть ", " в ", " по", " об", "ия ", "сво", " св", "лов", "на ", " че", "ело", "о н", " со", "ост", "чел", "ие ", "ого", "ет ", "ния", "ест", "аво", "ый ", "ажд", " им", "ние", "век", " не", "льн", "ли ", "ова", "име", "ать", "при", "т п", "и п", "каж", "или", "обо", " ра", "ых ", "жды", " до", "дый", "воб", "ек ", "бод", "ва ", "й ч", "его", "ся ", "и с", "ии ", "аци", "еет", "но ", "мее", "и и", "лен", "ой ", "тва", "ных", "то ", " ил", "к и", "енн", " бы", "ию ", " за", "ми ", "тво", "и н", "о п", "ван", "о с", "сто", "аль", " вс", "ом ", "о в", "ьно", "их ", "ног", "и в", "нов", "ако", "про", "ий ", "сти", "и о", "пол", "олж", "дол", "ое ", "бра", "я в", " ос", "ным", "жен", "раз", "ти ", "нос", "я и", " во", "тор", "все", " ег", "ей ", "тел", "не ", "и р", "ред", "ель", "тве", "оди", " ко", "общ", "о и", " де", "има", "а и", "чес", "ним", "сно", "как", " ли", "щес", "вле", "ься", "нны", "аст", "тьс", "нно", "осу", "е д", " от", "пре", "шен", "а с", "бще", "осн", "одн", "быт", "сов", "ыть", "лжн", "ран", "нию", "иче", "ак ", "ым ", "ват", "что", "сту", "чен", "е в", " ст", "рес", "оль", " ни", "ном", "род", "ля ", "нар", "вен", "ду ", "оже", "ны ", "е и", " то", "вер", "а о", "зов", "м и", "нац", "ден", "рин", "туп", "ежд", "стр", " чт", "я п", "она", "дос", "х и", "й и", "тоя", "есп", "лич", "бес", "обр", "ото", "о б", "ьны", "ь в", "нии", "е м", "ую ", " мо", "ем ", " ме", "аро", " ре", "ава", "кот", "ав ", " вы", "ам ", "жно", "ста", "ая ", "под", "и к", "ное", " к ", " та", " го", "гос", "суд", "еоб", "я н", "ен ", "и д", "мож", "еск", "ели", "авн", "ве ", "ече", "уще", "печ", "дно", "о д", "ход", "ка ", " дл", "для", "ово", "ате", "льс", "ю и", "в к", "нен", "ции", "ной", "уда", "вов", " бе", "оро", "нст", "ами", "циа", "кон", "сем", "е о", "вно", " эт", "азо", "х п", "ни ", "жде", "м п", "ког", "от ", "дст", "вны", "сть", "ые ", "о о", "пос", "сре", "тра", "ейс", "так", "и б", "дов", "му ", "я к", "нал", "дру", " др", "кой", "тер", "ь п", "арс", "изн", "соц", "еди", "олн"},
	Ukr: []string{"на ", " пр", " і ", "пра", "рав", " на", "ня ", "ння", " за", "ого", " по", "ти ", "го ", "люд", " лю", "во ", " ко", " ма", "льн", "юди", "их ", "о н", " не", "аво", "анн", "дин", " св", "сво", "ожн", "кож", "енн", "пов", "жна", " до", "ати", "ина", "ає ", "а л", " бу", "аці", "не ", "ува", "обо", " ос", " як", "має", " ви", "них", "аль", "або", "є п", " та", "ні ", "ть ", "ови", "бо ", " ві", " аб", "ере", "і п", "а м", "вин", "без", "при", "іль", "ног", "о п", "ми ", "та ", "ом ", "ою ", "бод", "ста", "воб", " бе", "до ", "ва ", "ті ", " об", "о в", "ост", " в ", " що", "ий ", "ся ", "і с", " сп", "инн", "від", "ств", "и п", "ван", "нов", "нан", "кон", " у ", "ват", "она", "ії ", "но ", "дно", "ій ", "езп", "пер", " де", "ути", "ьно", "ист", "під", "сті", "бут", " мо", "и і", "ідн", "ако", "нні", "ід ", "тис", "що ", "род", "і в", "а з", "ава", " пе", "му ", "і н", "а п", "соб", "ої ", "а в", "спр", "ів ", "ний", "яко", "ду ", "вно", "і д", "ну ", "аро", "и с", " ін", "ля ", "рів", "у в", " рі", "и д", "нар", "нен", "ова", "ому", "лен", "нац", "ним", "ися", "чи ", "ав ", "і р", "ном", " ро", "нос", "ві ", "вни", "овн", " її", "ові", "мож", "віл", "у п", " пі", " су", "її ", "одн", " вс", "ово", "ють", "іст", "сть", "і з", " ст", "буд", " ра", "чен", "про", "роз", "івн", "оду", "а о", "ьни", "ни ", "о с", "сно", "зна", "рац", "им ", "о д", "ими", "я і", "ції", "х п", "дер", "чин", " со", "а с", "ерж", "и з", "и в", "е п", "ди ", "заб", "осо", "у с", "е б", "сі ", "тер", "ніх", "я н", "і б", "кла", "спі", "в і", " ні", "о з", "ржа", "сту", "їх ", "а н", "нна", "так", "я п", "зпе", " од", "абе", "для", "ту ", "і м", "печ", " дл", "же ", "ки ", "віт", "ніс", "гал", "ага", "е м", "ами", "зах", "рим", "ї о", "тан", "ког", "рес", "удь", " ре", "то ", "ков", "тор", "ара", "сві", "тва", "а б", "оже", "соц", "оці", "ціа", "осн", "роб", "дь‐", "ь‐я", "‐як", "і і", "заг", "ахи", "хис", "піл", "цій", "х в", "лив", "осв", "іал", "руч", "ь п", "інш", "в я", "ги ", "аги", " ді", "ком", "ини", "а і", "оди", "нал", "тво", "кої", "всі", "я в", "ною", "об ", "о у", "о о", "і о"},
	Srp: []string{" пр", " и ", "рав", "пра", " на", "на ", " по", "ма ", " св", "да ", "има", "а п", "а и", "во ", "ко ", "ва ", "ти ", "и п", " у ", "ако", " да", "а с", "аво", "и с", "ост", " за", "о и", "сва", " им", "вак", "ава", "је ", "е с", " сл", " ко", "о н", "ња ", "но ", "не ", " не", "ом ", "ли ", " др", "или", "у с", "сло", "обо", "кој", "их ", "лоб", "бод", "им ", "а н", "ју ", " ил", "ств", " би", "сти", "а о", "при", "а у", " ра", "јед", "ог ", " је", "е п", "ње ", "ни ", "у п", "а д", "едн", "ити", "а к", "нос", "и у", "о д", "про", " су", "ање", "ова", "е и", "вањ", "и и", "циј", " ос", "се ", "дру", "ста", "ају", "ања", "и о", " об", "род", "ове", " ка", " де", "е о", "аци", "ја ", "ово", " ни", " од", "и д", " се", "ве ", "ује", "ени", "ија", "авн", "жав", " ст", "у и", "м и", "дна", "су ", "ред", "и н", "оја", "е б", "ара", "што", "нов", "ржа", "вој", "држ", "тва", "оди", "у о", "а б", "одн", "пош", "ошт", "ним", "а ј", "ка ", "ран", "у у", " ов", "аро", "е д", "сно", "ења", "у з", "раз", " из", "осн", "а з", "о п", "аве", "пре", "де ", "бит", "них", "шти", "ву ", "у д", "ду ", "ту ", " тр", "нар", " са", "гов", "за ", "без", "оји", "у н", "вно", "ичн", "еђу", "ло ", "ан ", "чно", "ји ", "нак", "ода", " ме", "вим", "то ", "сво", "ани", "нац", " ње", "ник", "њег", "тит", "ој ", "ме ", "ном", "м с", "е у", "о к", "ку ", " до", "ика", "ико", "е к", "пос", "ашт", "тре", "алн", "ног", " вр", "реб", "нст", " кр", "сту", "дно", "ем ", "вар", "е н", "рив", "туп", "жив", "те ", "чов", "ст ", "ови", "дни", "ао ", "сме", "бра", "ави", " ли", "као", "вољ", "ило", "о с", "штв", "и м", "заш", "њу ", "руг", "тав", "анс", "ено", "пор", "кри", "и б", "оду", "а р", "ла ", " чо", "а т", "руш", "ушт", " бу", "буд", "ављ", "уги", "м п", "ком", "оје", "вер", " ве", "под", "и в", "међ", "его", "вре", "акв", "еди", "тво", " см", "од ", "дел", "ена", "рад", "ба ", " мо", "ну ", "о ј", "дст", "кла", " оп", "как", "сам", "ере", "рим", "вич", "ива", "о о", " он", "вни", "тер", "збе", "х п", "ниц", "еба", "е р", "у в", "ист", "век", "рем", "сви", "бил", "ште", "езб", "јућ", "њен", "гла"},
	Azj: []string{" вә", "вә ", "әр ", "лар", " һә", "ин ", "ир ", " ол", " һү", " би", "һүг", "үгу", "гуг", "на ", "ләр", "дә ", "һәр", " шә", "бир", "ан ", " тә", "лик", "р б", "мал", "лма", "асы", "ини", "р һ", "шәх", "ән ", "әхс", "ары", "гла", "дир", "а м", "али", "угу", "аг ", " ма", "ын ", "илә", "уна", "јәт", " ја", "икд", "ара", "ар ", "әри", "әси", "рин", "әти", "р ш", "нин", "дән", "јјә", "н һ", " аз", "ни ", "әрә", " мә", "зад", "мәк", "ијј", " мү", "син", "тин", "үн ", "олу", "и в", "ндә", "гун", "рын", "аза", "нда", "ә а", "әт ", "ыны", "нын", "лыг", "илм", " га", " ет", "ә ј", "кди", "әк ", "лә ", "лмә", "олм", "ына", "инд", "лун", " ин", "мас", "хс ", "сын", "ә б", "г в", "н м", "адл", "ја ", "тмә", "н т", "әми", "нә ", "длы", "да ", " бә", "нун", "бәр", "сы ", " он", "әја", "ә һ", "маг", "дан", "ун ", "етм", "инә", "н а", "рлә", "си ", " ва", "ә в", "раг", "н б", "ә м", "ама", "ры ", "н и", "әра", "нма", "ынд", "инс", " өз", "аны", "ала", " ал", "ик ", "ә д", "ләт", "ирл", "ил ", " ди", "бил", "ығы", "ли ", "а б", "әлә", "дил", "ә е", "унм", "алы", "мүд", " сә", "ны ", "ә и", "н в", "ыг ", "нла", "үда", "аси", "или", " дә", "нса", "сан", "угл", "уг ", "әтл", "ә о", "хси", " һе", "ола", "кил", "ејн", "тәр", "јин", " бу", "ми ", "мәс", "дыр", "һәм", " да", "мин", "иш ", " һа", " ки", "у в", "лан", "әни", " ас", "хал", "бу ", "лығ", "р в", " ед", "јан", "рә ", "һеч", "алг", " та", "еч ", "и с", "ы һ", "сиа", "оси", "сос", "фиә", "г һ", "афи", "ким", "даф", " әс", "ә г", " иш", "н ә", "ији", "ыгл", "әмә", "ы о", "әдә", "әса", " со", "а г", "лыд", "илл", "мил", "а һ", "ыды", "сас", "лы ", "ист", " ис", "ифа", "мәз", "ыр ", "јар", "тлә", "лиј", "түн", "ина", "ә т", "сиј", "ал ", "рил", " бү", "иә ", "бүт", " үч", "үтү", "өз ", "ону", " ми", "ија", " нә", "адә", "ман", "үчү", "чүн", "сеч", "ылы", "т в", " се", "иал", "дах", "сил", "еди", "н е", "әји", "ахи", "хил", " ҹә", "миј", "мән", "р а", "әз ", "а в", "илд", "и һ", "тәһ", "әһс", "ы в", "һси", "вар", "шәр", "абә", "гу ", "раб", "аја", "з һ", "амә", "там", "ғын", "ад ", "уғу", "н д", "мәһ", "тәм", " ни", "и т", " ха"},
	Bel: []string{" і ", " пр", "пра", "ава", " на", "на ", " па", "рав", "ны ", "ць ", "або", " аб", "ва ", "ацы", "аве", "ае ", " ча", "ння", "анн", "льн", " ма", " св", "сва", "ала", "не ", "чал", "лав", "ня ", "ай ", "ых ", " як", "га ", "век", "е п", " ад", "а н", " не", "пры", "ага", " ко", "а п", " за", "кож", "ожн", "ы ч", "бод", "дна", "жны", "ваб", "цца", "ца ", " ў ", "а а", "ек ", "мае", "і п", "нне", "ных", "асц", "а с", "пав", "бо ", "ам ", "ста", " са", " вы", "ван", "ьна", " да", "ара", "дзе", "одн", "го ", "наг", "він", "аць", "оўн", "цыя", "мі ", "то ", " ра", "і а", "тва", " ас", "ств", "лен", "аві", "ад ", "і с", "енн", "і н", "аль", "най", "аво", "рац", "аро", "ці ", "сці", "пад", "ама", " бы", " яг", "яго", "к м", "іх ", "рым", "ым ", "энн", "што", "і і", "род", " та", "нан", " дз", "ні ", "я а", "гэт", "нас", "ана", " гэ", "інн", "а б", "ыць", "да ", "ыі ", "оў ", "чын", " шт", "а ў", "цыі", "які", "дзя", "а і", "агу", "я п", "ным", "нац", " у ", " ўс", "ыя ", "ьны", "оль", "нар", "ўна", "х п", "і д", "ў і", " гр", "амі", "ымі", "ах ", " ус", "адз", " ні", "эта", "ля ", "воў", "ыма", "рад", "ы п", "зна", "чэн", "нен", "аба", " ка", "ўле", "іна", "быц", "ход", " ін", "о п", " ст", "ера", "уль", "аў ", "асн", "сам", "рам", "ры ", " су", "нал", "ду ", "ь с", "чы ", "кла", "аны", "жна", "і р", "пер", "і з", "ь у", "маю", "ако", "ыцц", "яко", "для", "ую ", "гра", "ука", "е і", "нае", "адс", "і ў", "кац", "ўны", "а з", " дл", "яўл", "а р", "аюч", "ючы", "оду", " пе", " ро", "ы і", "вы ", "і м", "аса", "е м", "аду", "х н", "ода", "адн", "нні", "кі ", " шл", "але", "раз", "ада", "х і", "авя", "нав", "алі", "раб", "ы ў", "нна", "мад", "роў", "кан", "зе ", "дст", "жыц", "ані", "нст", "зяр", "ржа", "зак", "дзі", "люб", "аюц", "бар", "ім ", "ены", "бес", "тан", "м п", "дук", "е а", "гул", "я ў", " дэ", "ве ", "жав", "ацц", "ахо", "заб", "а в", "авы", "ган", "о н", "ваг", "я і", "чна", "я я", "сац", "так", "од ", "ярж", "соб", "м н", "се ", "чац", "ніч", "ыял", "яль", "цця", "ь п", "о с", "вол", "дэк", " бе", "ну ", "ога", " рэ", "рас", "буд", "а т", "асо", "сно", "ейн"},
	Bul: []string{" на", "на ", " пр", "то ", " и ", "рав", "да ", "пра", " да", "а с", "ств", "ва ", "та ", "а п", "ите", "но ", "во ", "ени", "а н", "е н", " за", "о и", "ото", "ван", "не ", " вс", "те ", "ки ", " не", "о н", "ове", " по", "а и", "ава", "чов", "ни ", "ане", "ия ", " чо", "аво", "ие ", " св", "е п", "а д", " об", "век", "ест", "сво", " им", "има", "ост", "и д", "и ч", "ани", "или", "все", "ли ", "тво", "и с", "ние", "вот", "а в", "ват", "ма ", " ра", "и п", "и н", " в ", "ек ", "сек", "еки", "а о", " ил", "е и", "при", " се", "ова", "ето", "ата", "воб", "обо", "бод", "аци", "ат ", "пре", "оди", "к и", " бъ", " съ", "раз", " ос", "ред", " ка", "а б", "о д", "се ", " ко", "бъд", "лно", "ния", "о п", " от", "ъде", "о в", "за ", "ята", " е ", " тр", "и и", "о с", "тел", "и в", "нит", "е с", "ран", " де", "от ", "общ", "де ", "ка ", "бра", "ен ", "ява", "ция", "про", "алн", "и о", "ият", "ст ", "нов", " до", "его", "как", "ато", " из", "нег", "а т", "ден", "а к", "щес", "а р", "тря", "а ч", "ряб", "о о", "вен", "ябв", "бва", "дър", "гов", "нац", "ено", "тве", "ърж", "е д", "нос", "ржа", "а з", "вит", "зи ", "акв", "лен", " та", "ежд", "и з", "род", "е о", "обр", "нот", " ни", " с ", "т с", "нар", "о т", "она", "ез ", "йст", "кат", "иче", " бе", "жав", "е т", "е в", "тва", "зак", "аро", "кой", "осн", " ли", "ува", "авн", "ейс", "сно", "рес", "пол", "нен", "вни", "без", "ри ", "стр", " ст", "сто", "под", "чки", "вид", "ган", "си ", "ди ", "и к", "нст", " те", "а е", "вси", "еоб", " дъ", "сич", "ичк", "едв", "жен", "ник", "ода", "т н", "о р", "ака", "ели", "одн", "елн", "лич", " че", "чес", "бще", " ре", "и м", " ср", "сре", "и р", "са ", "лни", " си", "дви", "ичн", "жда", " къ", "оет", "ира", "я н", "дей", " ме", "еди", "дру", "ход", "еме", "кри", "че ", "дос", "ста", "гра", " то", "ой ", "тъп", "въз", "ико", "и у", "нет", " со", "ави", "той", "елс", "меж", "чит", "ита", "що ", "ъм ", "азо", "зов", "нич", "нал", "дно", " мо", "ине", "а у", "тно", "таз", "кон", "лит", "ан ", "клю", "люч", "пос", "тви", "а м", "й н", "т и", "изв", "рез", "ази", "ра ", "оят", "нео", "чре"},
	Tuk: []string{" би", "лар", " ве", "ве ", "да ", "ада", "ары", " хе", "ир ", " ад", "бир", "дам", "кла", "ер ", "р б", "ың ", " ха", "ара", "га ", "ен ", "лан", "ыны", "или", "дыр", "ам ", "ала", " бо", "хер", "р а", "ыр ", "лы ", "лер", "ан ", "бил", "иң ", "ыды", "р х", "акл", "нда", " өз", "клы", "ны ", "хук", "ери", " ху", "уку", "ага", "не ", "лыд", "ине", "ына", "лен", "на ", "хак", "де ", "‐да", "ин ", "рын", "атл", " эд", "маг", "өз ", " де", "асы", "лыг", "кук", "е а", "ынд", "алы", "лма", "бол", "дан", "ини", "а х", " я‐", "е х", "ге ", "иле", "я‐д", "ар ", "ама", "ли ", "ыгы", "ети", " ба", " га", "гын", "ере", "укл", "лиг", "ның", "зат", "лык", "тлы", "нде", "ни ", "лик", "ден", "мак", "сын", "дил", "ры ", "аны", "кин", "әге", "п б", "а г", "хем", "иги", "эрк", "аза", "а д", "мек", " эр", "мал", "ыкл", "мәг", "сас", " эс", "екл", " ма", "рин", "эса", "ола", "ы б", "айы", "н э", "эди", " гө", " хи", "сы ", " аз", "баш", "ы д", "йда", "шга", "ашг", "а в", " до", "ыет", "ы в", "дак", "ниң", "рки", "гал", "чин", "гда", "ак ", " җе", "а б", " эт", "этм", "кы ", "лет", "йән", " та", "гин", "ян ", "тме", "хич", "ич ", "мез", " гу", "хал", "ылы", "үнд", "илм", "дай", "ягд", " яг", "и в", "им ", "акы", "ы г", "ән ", "а а", "рың", "ги ", "тле", "н м", " го", "ип ", "ал ", "еси", " се", "лме", " ка", "м х", "дең", "ң х", "е д", "дир", "илл", "рил", " ал", "кан", "е г", "лин", "ра ", "дол", " бе", " ми", "мил", "ң д", "н х", "ели", "н а", "е м", " ге", "ы х", " дө", "ик ", " со", "ң а", "чил", "дөв", "е б", " са", "гар", "е в", "ең ", "н б", "рма", " ме", "кли", "үчи", " дә", " үч", "ция", "н в", " дү", "и б", "айд", "кле", "сер", "а я", "соц", "гор", "оци", "дал", "мы ", "олм", "циа", "уң ", " он", "уп ", "кда", "дәл", "ири", " ди", "еле", "лип", "алк", "лим", "гур", "үни", "нме", " әх", "н г", " иш", "ы ө", "ң э", "нун", "еги", "тин", "ы а", "рле", "аци", "ыз ", "з х", "сыз", "аха", "м э", "олы", "рам", " ту", " ни", "ып ", "ерт", "алм", "ора", "и х", "хли", "әхл", "к э", "өвл", "вле", "тмә", "ет ", "нли", "ахс", "гөз", "гы ", "етл", "ы ү", "нуң", "ону", "сиз", "емм", "ек "},
	Mkd: []string{" на", "на ", " пр", " и ", "во ", " се", "то ", "ите", "те ", "рав", "та ", "а с", "пра", "ува", "да ", " да", " не", "ва ", "а п", "а н", "и с", "ата", "о н", "еко", "а и", " по", "но ", "ој ", "кој", " со", " за", " во", "ств", "ја ", "ње ", "ање", "аво", "ни ", " им", "от ", "е п", "е н", "ма ", "ат ", "вањ", "ост", "а д", "о с", "е и", "се ", "ова", "ија", "и п", " сл", "а о", "има", "сек", "сло", "ото", "ли ", "о д", "ава", "обо", "о и", " ил", "или", " би", "бод", "и н", "лоб", " од", "бид", "ред", "ен ", "при", "вот", "иде", "а в", "ста", " об", "и и", "и д", "пре", "нос", "ст ", "е с", " ни", " ќе", "ове", "аат", "аци", "ќе ", "со ", "ови", "про", "ј и", "тво", " ра", "ест", "што", " де", "т и", "акв", " ко", "раз", "гов", "его", "нег", "ани", "едн", "ако", "циј", "бра", "од ", "а з", "е б", "и о", "а б", "о п", "ват", " е ", " др", "ето", "ваа", "как", "ди ", "т с", " ка", " чо", "ени", "алн", "одн", "ено", " си", "чов", " шт", "а г", "а е", "вен", "нит", " ја", "де ", "оди", "е о", "ран", "и з", "сно", "нот", " ед", "тит", "лно", "ви ", "јат", "ден", "т н", "нац", " оп", " до", " ос", "и в", "осн", "кон", "дна", "е д", " ст", "век", "о о", "род", "сто", "сит", "еме", "ара", "дно", "обр", "ј н", "пшт", "еди", "опш", "за ", "ние", "аро", "нов", "а к", "вни", "дру", " ов", "тве", "жив", "ште", "д н", "ие ", " ме", "ед ", "иот", "и м", "о в", "ќи ", "дат", "шти", "јќи", "без", "бед", "ки ", "ков", "ко ", "а р", "нар", "чно", "дни", " вр", "ели", "нак", "ашт", "ичн", "ка ", "ема", "цел", "зем", "еду", "чув", "тес", "држ", "ник", "т п", "луч", "аа ", "деј", "нст", "не ", "а ч", "руг", "ода", "ивн", " це", "нив", "дин", "авн", " зе", "нио", "пор", "а м", "заш", "лас", "вит", "дек", "го ", "ине", "ело", "нет", "ез ", "тен", " ре", " из", "под", "раб", "або", "бот", "дув", "нув", " бе", "ење", "еде", "он ", "њет", "зов", "иту", "ван", "н и", "аѓа", "е в", "еѓу", "рем", "дел", "о к", "кот", "им ", " жи", "дос", "вре", "меѓ", "олн", "нап", " го", "емј", "кри", "уна", "нем", "оја", " су", "ита", "азо", "лит", "тор", "инс", "ора", "огл", "ипа", "пот", "слу", "кви"},
}

var arabicLangs = langProfileList{
	Arb: []string{" ال", "ية ", " في", "الح", "في ", " وا", "وال", " أو", "ة ا", "أو ", "الم", "الت", "لحق", "حق ", "لى ", "كل ", "ان ", "ة و", "الأ", " لك", "لكل", "ن ا", "ها ", "ق ف", "ات ", "مة ", "ون ", "أن ", "ما ", "اء ", "ته ", "و ا", "الع", "ي ا", "شخص", "ي أ", " أن", "الإ", "م ا", "حري", " عل", "ة ل", "من ", "الا", "حقو", "على", "قوق", "ت ا", "أي ", "رد ", " شخ", " لل", " أي", "ق ا", "لا ", "فرد", "رية", " ول", " من", "د ا", " كا", " إل", "خص ", "وق ", "ا ا", "ة أ", "ا ي", "ل ف", "ه ا", "نسا", "جتم", "ن ي", "امة", "كان", "دة ", " حق", "ام ", "الق", "ة م", " فر", "اية", "سان", "ل ش", "ين ", "ن ت", "إنس", "ا ل", " لا", "ذا ", "هذا", "ن أ", "لة ", "ي ح", " دو", "ه ل", "لك ", "ترا", "لتع", "اً ", "له ", "إلى", " عن", "ى ا", "ه و", "ع ا", "ماع", "د أ", "اسي", " حر", "ة ع", "مع ", "الد", "نون", " با", "لحر", "لعا", "ن و", "، و", "يات", "ي ت", "الج", " هذ", "ير ", "بال", "دول", "لإن", "عية", "الف", "ص ا", " وي", "الو", "لأس", " إن", "أسا", "ساس", "ماي", "حما", "رام", "سية", "انو", "مل ", "ي و", "عام", "ا و", "تما", " مت", "ة ت", "علي", "ع ب", "ك ا", " له", "ة ف", "قان", "ى أ", "ول ", "هم ", "الب", "ة ب", "ساو", "لقا", "الر", "لجم", "ا ك", "تمت", "ليه", "لتم", "لمت", "انت", " قد", "اد ", "ه أ", " يج", "ريا", "ق و", "ل ا", "ا ب", "ال ", "يه ", "اعي", "لدو", "ل و", "لإع", "لمي", "لمج", "لأم", "تع ", "دم ", "تسا", "عمل", "اته", "لاد", "رة ", "اة ", "غير", "قدم", "وز ", "جوز", "يجو", "عال", "لان", "متع", "مان", "فيه", "اجت", "م و", "يد ", "تعل", "ن ل", "ر ا", " يع", " كل", "مم ", "مجت", "تمع", "دون", " مع", "تمي", "ذلك", "كرا", "يها", " مس", "ميع", "إعل", "علا", " تم", " عا", "ملا", "اعا", "لاج", "ني ", "ليم", "متس", "ييز", "يم ", "اعت", "الش", " تع", "ميي", "عن ", "تنا", " بح", "لما", "ي ي", "يز ", "ود ", "أمم", "لات", "أسر", "شتر", "تي ", " جم", "ه ع", "ر و", "ي إ", "تحد", "حدة", " أس", "عة ", "ي م", "ة، ", "معي", "ن م", "لمس", "م ب", "اق ", "جمي", "لي ", "مية", "الض", "الس", "لضم", "ضما", "لفر", " وس", "لحم", "امل", "ق م", "را ", "ا ح", "نت ", " تن", "يته", " أم", "إلي", "واج", "د و", "لتي", " مر", "مرا", "متح", " ذل", " وأ", " تح", "ا ف", " به", " وم", " بم", "وية", "ولي", "لزو"},
	Urd: []string{"ور ", " او", "اور", " کی", "کے ", " کے", "یں ", " کا", "کی ", " حق", "ے ک", "ایٔ", "کا ", "یٔے", " کو", "یا ", "نے ", "سے ", " اس", "ٔے ", "میں", "کو ", " ہے", " می", "ے ا", " ان", "وں ", " کر", " ہو", "اس ", "ی ا", "ر ا", "شخص", " شخ", "حق ", " سے", " جا", "خص ", "ہر ", "ام ", "ے م", "ں ک", "ہیں", " یا", "سی ", "ادی", "آزا", " آز", "زاد", "ص ک", "ہ ا", "ہے ", "جای", "ا ح", "ر ش", "ت ک", "کہ ", "م ک", " پر", "ی ک", "ان ", "پر ", "۔ہر", "دی ", "یٔی", "س ک", "ا ج", "ر م", "ہے۔", "ق ہ", "ں ا", "ی ح", "و ا", "ار ", "ن ک", "قوق", "کسی", "حقو", "ری ", "وق ", "ے گ", " ہی", "ی ج", " مع", "سان", " نہ", " مل", " حا", "ٔی ", " جو", "نی ", "کرن", " لی", "تی ", "ی ت", "نسا", "ل ک", " کہ", "جو ", "انس", "اپن", "ے ب", "نہ ", " اپ", "یت ", "ا ا", "ہ ک", " کس", "ر ک", "رے ", "ے ہ", " ای", "می ", "ل ہ", "۔ ا", "ے ل", "ی ش", "رنے", "وہ ", "حاص", "ی م", "معا", "اصل", "صل ", "یں۔", "ویٔ", "نہی", "ملک", "ایس", "انہ", "ات ", "ی ب", "د ک", "ی ہ", " تع", "کیا", "ق ک", "ر ہ", "ا م", "دہ ", " من", " بن", " قو", "ے ج", "یہ ", "ں م", "اشر", "مل ", " دو", "عاش", "قوم", "ر ب", "انی", "وام", "قوا", "اقو", "لیٔ", "دار", " وہ", " و ", " عا", "ی س", "بر ", "علا", "اد ", "ہ م", "و ت", "ر ن", " جس", "ے۔ہ", "ے، ", "انو", " دی", "گی ", "لیم", "یوں", " قا", " یہ", "دوس", "ے۔ ", "ا ہ", "تعل", "یم ", "ر پ", "جس ", "ریق", "ے ح", " اق", "نیا", "لک ", " گی", "ین ", "یاد", " مس", "لاق", "، ا", "ی ن", "پنے", "وری", "م ا", " با", "علی", "یر ", "ی، ", "انے", "ون ", "ن ا", "ر ع", " بر", "ی آ", "ر ح", " رک", "ے پ", "کر ", "گا۔", " پی", "سب ", " گا", "نا ", " پو", "یسے", "رای", " مر", "اری", "قان", "نون", " مم", "ندگ", " اع", "دگی", "ہ و", " ہر", "ر س", " چا", "خلا", "ا پ", "ق ح", " بھ", "س م", " شا", "ہوگ", "ے خ", "وسر", "رتی", "ومی", " بی", "رکھ", " مت", "کوی", "ر آ", "پور", "اف ", " مح", "ے س", "ہوں", "نکہ", "ونک", "ت ا", " طر", "ے ع", "یٔد", "د ا", "ال ", "ں۔ ", "م م", "اں ", " مق", "غیر", "پنی", " ام", "ں، ", "من ", "ہو ", "ریع", "و ک", "ذری", " ذر", "عام", "، م", "دان", "ادا", "اعل", "مام", "تما", " عل", "دیو", "بھی", "ھی ", "بنی", "ے ی", "ا ک", "اوی", "ل م", " زن", "یاس", "لان", "عمل", " عم", "ت م", " بچ"},
	Skr: []string{"تے ", "اں ", " تے", "دے ", "دی ", "وں ", " دا", " حق", " کو", "ے ا", "کوں", " دے", "دا ", " دی", "یاں", " کی", "ے ۔", "یں ", "ہر ", " ۔ ", "کیت", "ہے ", " وچ", " ہے", "وچ ", " ان", " شخ", "شخص", "ادی", "ال ", " حا", "اصل", "حق ", "حاص", "ے م", "خص ", "صل ", "ں د", " نا", "یا ", " ای", "اتے", "ق ح", "ل ہ", "ے و", "ں ک", " ات", "ہیں", "سی ", " مل", "نال", "زاد", "ازا", "ی ت", " از", "قوق", "ار ", "ا ح", "حقو", " او", "ص ک", " ۔ہ", "۔ہر", "ر ش", "دیا", "ے ج", "وق ", "ندے", " کر", "یند", " یا", "نہ ", " جو", "کہی", "ئے ", "ی د", "سان", "نسا", "وند", "ی ا", "یتے", "انس", "ا ا", "ملک", "ے ح", "و ڄ", "ے ک", "ڻ د", " وی", "یسی", "ے ب", "ا و", " ہو", "ں ا", "ئی ", "ندی", "تی ", "آپڻ", "وڻ ", "ر ک", "ن ۔", " نہ", "انہ", "جو ", " کن", " آپ", " جی", "اون", "ویس", "ی ن", " تھ", " کہ", "ان ", "ری ", "ڻے ", " ڄئ", " ہر", "ے ن", "دہ ", "ام ", "ں م", "ے ہ", "تھی", "ں و", "۔ ا", "ں ت", "ی ۔", "کنو", "ی ح", "ی ک", "نوں", "رے ", "ہاں", " بچ", "ون ", "ے ت", "کو ", " من", "ی ہ", "اری", "ور ", "نہا", "ہکو", "یتا", "نی ", "یاد", "ت د", "ن د", " ون", "وام", "ی م", "قوا", "تا ", "ڄئے", "پڻے", " ہک", "می ", " قو", "ق ت", "ے د", "لے ", "اف ", "ل ک", "ل ت", " تع", "چ ا", "ین ", "خلا", "اے ", "علا", " سا", "جیا", "ئو ", "کرڻ", "ی و", "انی", "ہو ", "دار", " و ", "ی ج", " اق", "ن ا", "یت ", "ارے", "ے س", "لک ", "ق د", "ہوو", " ڋو", "ر ت", " اے", "ے خ", " چا", " خل", "لاف", "قنو", "نون", "پور", "ڻ ک", " پو", "ایہ", "بچئ", "چئو", "ات ", "الا", "ونڄ", "وری", "این", " وس", " لو", "و ا", "ہ د", " رک", "یب ", "سیب", "وسی", "یر ", "ا ک", "قوم", "ریا", "ں آ", " جا", "رکھ", "مل ", "کاں", "رڻ ", "اد ", "او ", "عزت", " قن", "ب د", "وئی", "ے ع", " عز", " ۔ک", " مع", "اقو", "ایں", "م م", "زت ", "ڻی ", "یوڻ", "ر ہ", " سم", "ں س", "لوک", " جھ", " سی", "جھی", "ت ت", "ل ا", "اوڻ", "کوئ", "ں ج", "ہی ", "حدہ", "تعل", "ے ذ", "وے ", "تحد", "متح", "لا ", "ا ت", "کار", " اع", "ے ر", " مت", "ر ا", "ا م", "ھین", "ھیو", "یہو", " مط", " سڱ", "ی س", "ڄے ", "نڄے", "سڱد", "لیم", "علی", "ے ق", " ذر", "م ت", " کھ", "ن ک", " کم", "ہ ا", "سار", "ائد", "ائی", "د ا", " ہن", "ہن ", "ی، ", "و ک", "ں ب", "ھیا", "ذری", "ں پ", "لی "},
	Uig: []string{" ئا", " ھە", "ىنى", "ە ئ", "نىڭ", "ىلى", " ۋە", "ىڭ ", "ۋە ", " ئى", " بو", "ھوق", "وقۇ", " ھو", "قۇق", "نى ", "بول", " ئە", "لىك", "قىل", "ىن ", "لىش", "شقا", "قا ", "ەن ", " قى", "ن ب", "ھەم", "ى ئ", "ئاد", "ىشى", "دەم", "ادە", "كى ", "لىق", "غان", "ىي ", "ىغا", "گە ", " بى", "دىن", "ىدى", "ەت ", "كىن", "ىكى", "ندا", "ۇق ", " تە", "نلى", "تىن", "ەم ", "لەت", "قان", "ىگە", "ىتى", "ىش ", "ھەر", "ئەر", " با", "ولۇ", "دۆل", "غا ", "اند", " دۆ", "اق ", "مە ", "لۇش", "دە ", "لۇق", " ئۆ", "ان ", " يا", "ەرق", "ۆلە", "ركى", " قا", "ەرك", "ەمم", "ا ئ", "ممە", "ۇقى", "ىق ", " بە", "رقا", "داق", "ارا", "ىلە", "رىم", "ىشق", "ى ۋ", "لغا", "مەن", "اكى", "ەر ", "ا ھ", "دۇ ", "ياك", "ۇقل", "ئار", "ق ئ", "ىنل", "لار", " ئې", "ى ب", "لىن", "ڭ ئ", "ئۆز", "ق ھ", "شى ", "ىمە", "قلۇ", "ن ئ", "لەر", "ەتل", "نىش", "ىك ", "ەھر", " مە", "ھرى", "لەن", "ىلا", "ار ", "بەھ", " ئۇ", "ە ق", "ئىي", "اسى", " مۇ", "رلى", " ئو", "بىر", "، ئ", "بىل", "ش ھ", "بار", "ى، ", "ۇ ھ", "ايد", "ۇشق", "شكە", "ە ب", "يەت", "ا ب", "رنى", "كە ", "ىسى", " كې", "ېلى", "الى", "ەك ", "م ئ", "ماي", "ولم", "تنى", "ىدا", "ارى", "يدۇ", "لىد", " قو", "ەشك", "تلە", "ك ھ", "انل", "ەمد", "مائ", "ئال", "ر ئ", "مدە", "ىيە", "ش ئ", "ە ھ", "لما", "ائى", "ئىگ", "دا ", "ي ئ", "ۇشى", "راۋ", "ا، ", "سىي", " تۇ", "كىل", "ە ت", "ىقى", "قى ", "ۆزى", "ېتى", "ىرى", "ىر ", "ىپ ", "ى ك", "ن، ", "ر ب", "لەش", "اسا", "اۋا", "ى ھ", "شلى", "ساس", "ادى", "تى ", "اشق", "ەتت", "قىغ", "ىما", "انى", " خى", "ۇرۇ", " خە", "ن ق", "منى", " خا", "چە ", "ى ق", " جە", "رقى", "تىد", " ھۆ", "باش", "ارل", "ئىش", "تۇر", " جى", "مۇش", "نۇن", "شۇ ", "انۇ", "ۇش ", "رەك", "ېرە", "كېر", " سا", "الغ", "ۇنى", "ئېل", "ىشل", "تەش", "خەل", "مەت", "اش ", "دىغ", "كەن", "ەلق", "تىش", "مىن", "ايى", "سىز", "ق ۋ", "نىي", "جىن", "رىش", "پ ق", " كى", "ېرى", "ئاس", "ەلى", " ما", "تتى", "ىرل", "ولى", " دە", "ارق", "سىت", "ە م", " قە", "شىل", " تى", "ەرن", "كىش", "ن ھ", "ەلگ", "ەمن", "ك ئ", " تو", "ى ي", "قتى", "ئاش", "تىم", "تەۋ", "ناي", "ىدە", "ىنا", " بۇ", "ىيا", "زىن", "امى", "قار", "شكى", "ىز ", " ئۈ", "ەۋە", "ۆرم", "ە خ", "شىش", "ىيى", "جتى", "ىجت", "ئىج", "نام", "تەر"},
	Pes: []string{" و ", " حق", " با", "که ", "ند ", " که", " در", "در ", "رد ", " دا", "دار", "از ", " از", "هر ", " هر", "یت ", "ر ک", "حق ", "د ه", "ای ", "د و", "ان ", " را", "ین ", "ود ", "یا ", " یا", "را ", "ارد", "ی و", "کس ", " کس", " بر", " آز", "باش", "ه ب", "آزا", "د ک", " خو", "ه ا", "د ب", "زاد", " اس", "ار ", " آن", "ق د", "شد ", "حقو", "قوق", "ی ب", "وق ", "ده ", "ه د", "ید ", "ی ک", "و ا", "ور ", "ر م", "رای", "اشد", "خود", "ادی", "تما", "ری ", " اج", "ام ", "دی ", "اید", "س ح", "است", "ر ا", "و م", " ان", "د ا", "نه ", " بی", "با ", " هم", " نم", "مای", " تا", "د، ", "ی ا", "انه", "ات ", "ون ", "ایت", "ا ب", "ست ", " کن", "برا", "انو", " بش", " مو", "این", " مر", "اسا", " مل", "وان", "ر ب", "جتم", " شو", " اع", "ن ا", "ورد", " می", " ای", "آن ", " به", "و آ", "ملل", "ا م", "ماع", "نی ", "ت ا", "، ا", "ت و", "ئی ", "عی ", "ائی", "اجت", "و ب", "های", "ن م", "ی ی", "بشر", "کند", "شود", " من", " زن", "ن و", "ی، ", "بای", "ی ر", " مس", "مل ", "مور", "ز آ", "توا", "دان", "اری", "علا", "گرد", "یگر", "کار", " گر", " بد", "ن ب", "ت ب", "ت م", "ی م", " مق", "د آ", "شور", "یه ", "اعی", " عم", "ر خ", "ن ح", " کش", "رند", "مین", " اح", "ن ت", "ی د", " مت", "ه م", "د ش", " حم", "و د", "دیگ", "لام", "کشو", "هٔ ", "ه و", "انی", "لی ", "ت ک", " مج", "ق م", "میت", " کا", " شد", "اه ", "نون", " آم", "اد ", "ادا", "اعل", "د م", "ق و", "ا ک", "می ", "ی ح", "لل ", "نجا", " مح", "ساس", "یده", " قا", "بعی", "قان", "ر ش", "مقا", "ا د", "هد ", "وی ", "نوا", "گی ", "ساو", "ر ت", "بر ", "اً ", "نمی", "اسی", "اده", "او ", " او", " دی", " هی", "هیچ", "ه‌ا", "‌ها", "یر ", "خوا", "د ت", "همه", "ا ه", "تی ", "حما", "دگی", "بین", "ع ا", "سان", "ر و", "شده", "ومی", " عق", " بع", "ز ح", "شر ", "مند", " شر", "ٔمی", "أم", "تأ", "انت", "اند", "اوی", "مسا", "ردد", "بهر", " بم", "ارن", "یتو", "ل م", "ران", "و ه", "ر د", "م م", "رار", "عقی", "سی ", "و ت", "زش ", " بو", "ا ا", "ی ن", "موم", "جا ", "عمو", "رفت", "عیت", " فر", "ندگ", "واه", "زند", "م و", "نما", "ه ح", "ا ر", "دیه", "جام", "مرد", "ت، ", "د ر", "مام", " تم", "ملی", "نند", "الم", "طور", "ی ت", "تخا", "ا ت", "امی", "امل", "دد ", " شخ", "شخص"},
}

var devanagariLangs = langProfileList{
	Hin: []string{"के ", "प्र", "और ", " और", " के", "ों ", " का", "कार", " प्", "का ", " को", "या ", "ं क", "ति ", "ार ", "को ", " है", "िका", "ने ", "है ", "्रत", "धिक", " अध", "अधि", "की ", "ा क", " कि", " की", " सम", "ें ", "व्य", "्ति", "क्त", "से ", " व्", "ा अ", "्यक", "में", "मान", "ि क", " स्", " मे", "सी ", "न्त", " हो", "े क", "ता ", "यक्", "क्ष", "ै ।", "िक ", "त्य", " कर", "्य ", " या", "भी ", " वि", "रत्", "र स", "ी स", " जा", "स्व", "रों", "्ये", "ेक ", "येक", "त्र", "िया", "ा ज", "क व", "र ह", "ित ", "्रा", "किस", " अन", "ा स", "िसी", "ा ह", "ना ", " से", " पर", "र क", " सा", "देश", "गा ", " । ", " अप", "्त्", "े स", "समा", "ान ", "ी क", "्त ", "वार", " ।प", "ा प", " रा", "षा ", "न क", "।प्", "ष्ट", "था ", "अन्", " मा", "्षा", "्वा", "ारो", "तन्", "वतन", "ट्र", "्वत", "प्त", "ाप्", "्ट्", "राष", "ाष्", " इस", "े अ", " उस", " सं", "राप", "कि ", "त ह", "हो ", "ं औ", "ार्", "ा ।", "किय", "े प", " दे", " भी", "करन", "री ", "जाए", "ी प", " न ", "र अ", "क स", "अपन", "े व", "ाओं", "्तर", "ओं ", " नि", "सभी", "रा ", " तथ", "तथा", "िवा", "यों", "पर ", " ऐस", "रता", "ारा", "्री", "सम्", " द्", "ीय ", "िए ", "व क", "सके", "द्व", "होग", " सभ", "ं म", "माज", "रने", "िक्", "्या", "ा व", "र प", " जि", "ो स", "र उ", "रक्", "े म", "पूर", " लि", "ाएग", " भा", "इस ", "त क", "ाव ", "स्थ", "पने", "ा औ", "द्ध", "श्य", "र्व", " घो", "घोष", "रूप", "भाव", "ाने", "कृत", "ो प", "े ल", "लिए", "शिक", "ूर्", " उन", "। इ", "ं स", "य क", "्ध ", "दी ", "ी र", "र्य", "णा ", "एगा", "न्य", "रीय", "ेश ", "रति", "े ब", " रू", "ूप ", "परा", "्र ", "तर्", " पा", " सु", "जिस", "तिक", "सार", "जो ", "ेशो", " शि", "ानव", "ी अ", "चित", "े औ", " पू", "ियो", "ा उ", "म क", "ी भ", "शों", " बु", "म्म", "स्त", "िश्", "्रो", "्म ", "ो क", " यह", "र द", "नव ", "चार", "दिय", "े य", "र्ण", "राध", "ोगा", "ले ", "नून", "ानू", "ोषण", "षणा", "विश", " जन", "ारी", "परि", "गी ", "वाह", "साम", "ाना", "रका", " जो", "ाज ", "ी ज", "ध क", "बन्", "ताओ", "ंकि", "ूंक", "ास ", "कर ", "चूं", "ी व", "य ह", "ा ग", "य स", "न स", "त र", "कोई", "ुक्", "ोई ", " ।क", "ं न", "हित", "निय", "याद", "ादी", "्मा", "्था", "ामा", "ाह ", "ी म", "े ज"},
	Mar: []string{"्या", "या ", "त्य", "याच", "चा ", " व ", "ण्य", "प्र", "कार", "ाचा", " प्", "धिक", "िका", " अध", "अधि", "च्य", "ार ", "आहे", " आह", "ा अ", "हे ", " स्", "्रत", "्ये", "ा क", "स्व", " कर", "्वा", "ता ", "ास ", "ा स", "ा व", "त्र", " त्", "वा ", "ांच", "यां", "िक ", "मान", " या", "्य ", " का", " अस", "रत्", "ष्ट", "र्य", "येक", "ल्य", "र आ", "ाहि", "क्ष", " को", "ामा", "कोण", " सं", "ाच्", "ात ", "ा न", " रा", "ंत्", "ून ", "ेका", " सा", "राष", "ाष्", "चे ", "्ट्", "ट्र", "तंत", " मा", "ने ", "किं", " कि", "व्य", "वात", "े स", "करण", "ंवा", "िंव", "ये ", "क्त", " सम", "ा प", "ना ", " मि", "कास", "ातं", "्र्", "र्व", "समा", "मिळ", " जा", "े प", "व स", "यास", "ोणत", "रण्", "काम", "ीय ", "ा आ", " दे", "े क", "ांन", "हि ", "रां", " व्", "्यक", "ा म", "िळण", "ही ", " पा", "्षण", "ार्", "ान ", "े अ", " आप", " वि", "ळण्", "ाही", "ची ", "े व", "्रा", "मा ", "ली ", "ंच्", "ारा", "ा द", " आण", " नि", "णे ", "द्ध", " नय", "ला ", "ा ह", "नये", " सर", "सर्", "्री", "बंध", "ी प", "आपल", "ले ", "ील ", "माज", " हो", "्त ", "त क", "ाचे", "्व ", "षण ", "ंना", "लेल", "ी अ", "देश", "आणि", "णि ", "ध्य", " शि", "ी स", "े ज", "शिक", "रीय", "ानव", "पाह", "हिज", "िजे", "जे ", "क स", "यक्", "न क", "व त", "ा ज", "यात", "पल्", "न्य", "वी ", "स्थ", "ज्य", " ज्", "े आ", "रक्", "त स", "िक्", "ंबं", "संब", " के", "क व", "केल", "असल", "य अ", "य क", "त व", "ीत ", "णत्", "त्व", "ाने", " उप", "्वत", "भाव", "े त", "करत", "याह", "रता", "िष्", "व म", "कां", "साम", "रति", "सार", "ंचा", "र व", "क आ", "याय", "ासा", "साठ", "ाठी", "्ती", "ठी ", "ेण्", "र्थ", "ीने", "े य", "जाह", "ोणा", "संर", "ायद", "च्छ", "स स", "ंरक", "तील", "ी व", "त आ", "ी आ", "ंधा", "ेशा", "ित ", " अश", "हीर", " हक", "हक्", "क्क", "य व", "शा ", "व आ", "तीन", "ण म", "ूर्", "ेल्", "द्य", "ेले", "ांत", "ा य", "ा ब", "ी म", "ंचे", "याव", "देण", "कृत", "ारण", "ेत ", "िवा", "वस्", "स्त", "ाची", "नवी", " अर", "थवा", "अथव", "ा त", " अथ", "अर्", "ती ", "पूर", "इतर", "र्ण", "ी क", "यत्", " इत", " शा", "रका", "तिष", "ण स", "तिक", "्रक", "्ध ", "रणा", " आल", "ेल ", "ाजि", " न्", "धात", "रून", "श्र", "असे", "ष्ठ", "ुक्", "ेश ", "तो ", "जिक", "े म"},
	Mai: []string{"ाक ", " आ ", "प्र", "कार", "िका", "धिक", "ार ", "्रत", "ेँ ", "क अ", "्यक", "िक ", "्ति", " अध", "व्य", "अधि", "क स", " प्", "क्त", " व्", "केँ", "यक्", "तिक", "न्त", " स्", "हि ", "क व", "मे ", "बाक", "मान", " सम", "त्य", "क्ष", " छै", "छैक", "ेक ", "स्व", "त्र", "रत्", "्ये", "ष्ट", " अप", "येक", "र छ", "सँ ", "वा ", " एह", "ैक।", "ित ", " वि", " जा", "ति ", "्त्", "ट्र", "िके", "राष", "ाष्", " हो", "्ट्", " रा", "्य ", " सा", " अन", " कर", "अपन", "।प्", "कोन", "अछि", "वतन", "्वत", "तन्", "क आ", " अछ", "ताक", "था ", " पर", " वा", " को", "ार्", "एहि", "पन ", "ा आ", "नहि", "नो ", "समा", " मा", "्री", "रता", " नि", " का", "देश", " नह", "्षा", "क प", " दे", " कए", "रक ", " सं", "ोनो", "ि क", "न्य", "आ स", "छि ", "्त ", "ल ज", "्वा", "ारक", "ा स", "तथा", "ान्", " तथ", "्या", "आ अ", "ना ", "ँ क", "ान ", " जे", "जाए", "वार", "ता ", "ीय ", "र आ", "क ह", "करब", "िवा", "ामा", "र्व", " आओ", "्रस", "परि", "त क", "स्थ", "ा प", "ानव", "रीय", "धार", "्तर", "अन्", "घोष", "साम", "माज", "आओर", "ारण", " एक", "कएल", "ँ अ", "ओर ", "एबा", "स्त", "द्ध", "्रा", "ँ स", "रण ", " सभ", "ोषण", "क।प", "ाहि", "रबा", "क ज", "ा अ", "चित", "यक ", "कर ", "पूर", "रक्", "नक ", " घो", "षा ", "िक्", "सम्", "एहन", " उप", "र प", " अव", "एल ", "ूर्", "षणा", " हे", "त अ", "शिक", "तु ", "ाधि", "ेतु", "हेत", "हन ", "िमे", "र अ", "वक ", "ँ ए", "जाह", " शि", "आ प", "भाव", "े स", "्ध ", "क क", "ि ज", "प्त", "रूप", "निर", "िर्", " सक", "च्छ", "होए", "रति", "अनु", "सभ ", "हो ", "ेल ", "त आ", "चार", "ण स", "रा ", "त ह", "जिक", "ाजि", "र्ण", "्रक", "एत।", "ि आ", "र्य", "सभक", "ैक ", "क उ", " जन", "त स", "ाप्", "न प", "श्य", "न अ", "कृत", "हु ", "रसं", "री ", "राप", "ा व", "जे ", "क ब", "ि घ", " भा", "उद्", "ाएत", "्ण ", "विव", " उद", "वाध", "िसँ", "आ व", "ि स", "न व", "ारा", "ोएत", " ओ ", "य आ", "कान", "िश्", "न क", " दो", "णाक", " द्", "हिम", " अथ", "अथव", "ामे", "द्व", "ेश ", "ओ व", "ि अ", "क ए", "वास", " पू", "षाक", "त्त", "य प", " बी", "यता", "धक ", "ए स", "थवा", "ि द", "पर ", " भे", "जेँ", " कि", "कि ", "क ल", " रू", "विश", "न स", " ले", "सार", "ाके", "िष्", "रिव", "क र", "ास ", "ेओ ", "्थि", "केओ", "राज"},
	Bho: []string{" के", "के ", "ार ", "े क", "कार", "धिक", "िका", "ओर ", "आओर", " आओ", " अध", "अधि", "े स", "ा क", "े अ", " हो", " सं", "र क", "र स", "ें ", " मे", "में", "िक ", " कर", "ा स", "र ह", " से", "से ", "रा ", "मान", " सम", "न क", "क्ष", "े ब", "नो ", " चा", "वे ", "ता ", "चाह", "ष्ट", " रा", "ति ", "्रा", "खे ", "राष", "ाष्", "प्र", " सा", " का", "ट्र", "े आ", " प्", " सक", " मा", "्ट्", " स्", "होख", " बा", "करे", "ि क", "ौनो", "त क", "था ", "कौन", "पन ", " जा", " कौ", "रे ", "ाति", "ला ", " ओक", "ेला", "तथा", "आपन", "्त ", " आप", "कर ", "हवे", "र म", " हव", " तथ", "सबह", "र आ", "ोखे", " ह।", "िर ", "े ओ", "केल", "सके", "हे ", " और", "ही ", "तिर", "त्र", "जा ", "ना ", "बहि", "।सब", "े च", " खा", "े म", " पर", "खात", "ान ", "र ब", "न स", "ावे", " लो", "षा ", "ाहे", "ी क", "ओकर", "ा आ", "माज", "ित ", "े ज", "ल ज", "मिल", "संग", "्षा", "ं क", " सब", "ा प", "और ", "रक्", "वे।", "िं ", "े ह", "ंत्", "ाज ", "स्व", "हिं", "नइख", "कान", "ो स", " जे", "समा", "क स", "लोग", "करा", "क्त", "्रत", "ला।", " नइ", "े। ", "ानव", "िया", "हु ", "इखे", "्र ", "रता", "्वत", "ानू", "े न", "ाम ", "नून", "ाही", "वतं", "पर ", "ी स", " ओ ", "े उ", "े व", "्री", "रीय", "स्थ", "तंत", "दी ", "ीय ", "े त", "र अ", "र प", "्य ", "साम", "बा।", " आद", "ून ", "। स", "व्य", "ा।स", "सभे", "भे ", "या ", " दे", "ा म", "े ख", " वि", " सु", "केह", "प्त", "योग", "ु क", "ोग ", "े द", "चार", "ादी", "ाप्", " दो", " या", "राप", "ल ह", "पूर", " मि", "तिक", "खल ", "यता", "्ति", " बि", "ए क", "आदि", "दिम", " ही", "हि ", "मी ", " नि", "र न", " इ ", "ेहु", "नवा", "ा ह", "री ", "ले ", " पा", "ाधि", " सह", " उप", "्या", " जर", "षण ", " सभ", "िमी", "देश", "े प", "म क", "जे ", "ाव ", " अप", "शिक", "ाजि", "जाद", "जिक", "े भ", "क आ", "्तर", "िक्", "ि म", "ेकर", "ुक्", "वाध", "गठन", " व्", "निय", "ठन ", "।के", "ामा", "रो ", " जी", "य क", "न म", "े ल", "न ह", "ास ", "ेश ", " शा", "घोष", "ंगठ", "िल ", " घो", "्षण", " पू", "े र", "ंरक", "संर", "उपय", "पयो", "हो ", "बा ", "ी ब", "्म ", "सब ", "दोस", "ा। ", " आज", "साथ", " शि", "आजा", " भी", " उच", "ने ", "चित", " अं", "र व", "ज क", "न आ", " ले", "नि ", "ार्", "कि ", "याह", "्थि"},
	Nep: []string{"को ", " र ", "कार", "प्र", "ार ", "ने ", "िका", "क्त", "धिक", "्यक", " गर", "व्य", "्रत", " प्", "अधि", "्ति", " अध", " व्", "यक्", "मा ", "िक ", "त्य", "ाई ", "लाई", "न्त", "मान", " सम", "त्र", "गर्", "र्न", "क व", " वा", "्ने", "वा ", " स्", "रत्", "र स", "्ये", "तिल", "येक", "ेक ", "छ ।", "ो स", "ा स", "हरू", " वि", "क्ष", "्त्", "िला", " । ", "स्व", "हुन", "ति ", " हु", "ले ", " रा", " मा", "ष्ट", "समा", "वतन", "तन्", " छ ", "र छ", " सं", "्ट्", "ट्र", "ाष्", "ो अ", "राष", "्वत", "ुने", "नेछ", "हरु", "ान ", "ता ", "े अ", "्र ", " का", "िने", "ाको", "गरि", "े छ", "ना ", " अन", " नि", "रता", "नै ", " सा", "ित ", "तिक", "क स", "र र", "रू ", "ा अ", "था ", "स्त", "कुन", "ा र", "ुनै", " छै", "्त ", "छैन", "ा प", "ार्", "वार", "ा व", " पर", "तथा", " तथ", "का ", "्या", "एको", "रु ", "्षा", "माज", "रक्", "परि", "द्ध", "। प", " ला", "सको", "ामा", " यस", "ाहर", "ेछ ", "धार", "्रा", "ो प", "नि ", "देश", "भाव", "िवा", "्य ", "र ह", "र व", "र म", "सबै", "न अ", "े र", "न स", "रको", "अन्", "ताक", "ंरक", "संर", "्वा", " त्", "सम्", "री ", "ो व", "ा भ", "रहर", " कु", "्रि", "त र", "रिन", "श्य", "पनि", "ै व", "यस्", "ारा", "ानव", " शि", "ा त", "लाग", "रा ", "शिक", " सब", "ाउन", "िक्", "्न ", "ारक", "ा न", "रिय", "्यस", "द्व", "रति", "चार", " सह", "्षण", " सु", "ारम", "ुक्", "ुद्", "साम", "षा ", "ैन ", " अप", " भए", "बाट", "ुन ", " उप", "ान्", "ो आ", "्तर", "िय ", "कान", "ि र", "रूक", "द्द", "र प", "ाव ", "ो ल", "तो ", " पन", "ैन।", " आव", "ा ग", "।प्", "बै ", "ूर्", "िएक", "र त", "निज", "त्प", " भे", "जिक", "ेछ।", "िको", "्तो", "वाह", "त स", "ाट ", " अर", "ाजि", "्ध ", " उस", "रमा", "ात्", "र्य", "नको", "ाय ", "जको", "ित्", "ागि", " अभ", "न ग", "गि ", "ा म", " आध", "स्थ", " पा", "ारह", "घोष", "त्व", "यता", "ा क", "र्द", " मत", "विध", " सक", "सार", "परा", "युक", "राध", " घो", "णको", "अपर", "े स", "ारी", "।कु", " दि", " जन", "भेद", "रिव", "उसक", "क र", "र अ", "ि स", "ानु", "ो ह", "रुद", " छ।", "ूको", "रका", "नमा", " भन", "र्म", "हित", "पूर", "न्य", "क अ", "ा ब", "ो भ", "राज", "अनु", "ोषण", "षणा", "य र", " मन", " बि", "्धा", " दे", "निर", "ताह", "र उ", "यस ", "उने", "रण ", "विक"},
}

var ethiopicLangs = langProfileList{
	Amh: []string{"፡መብ", "ሰው፡", "ት፡አ", "ብት፡", "መብት", "፡ሰው", "፡አለ", "፡ወይ", "ወይም", "ይም፡", "ነት፡", "ንዱ፡", "አለው", "ለው።", "ዳንዱ", "ያንዳ", "ንዳን", "እያን", "ዱ፡ሰ", "ት፡መ", "፡እን", "፡የመ", "።እያ", "እንዲ", "፡ነጻ", "፡የተ", "ም፡በ", "ው፡የ", "ም፡የ", "፡የሚ", "ና፡በ", "ን፡የ", "፡የማ", "፡አይ", "ነጻነ", "ና፡የ", "ው፡በ", "ቶች፡", "ው።፡", "ሆነ፡", "ት፡የ", "፡በሚ", "፡መን", "ው።እ", "ትና፡", "ኀብረ", "ትን፡", "ውም፡", "ንኛው", "እኩል", "ብቻ፡", "ኛውም", "ንም፡", "፡ለመ", "፡ያለ", "ም፡ሰ", "ማንኛ", "መብቶ", "፡አገ", "ት፡በ", "ራዊ፡", "፡እኩ", "፡ለማ", "ለት፡", "በት፡", "ሆን፡", "መንግ", "፡በተ", "ረት፡", "ብቶች", "ጋብቻ", "ዎች፡", "ህንነ", "ጻነት", "ም፡እ", "ወንጀ", "፡ልዩ", "ሰብ፡", "ማንም", "ጠበቅ", "ኩል፡", "ደህን", "።ማን", "ነጻ፡", "ግኘት", "ማግኘ", "።፡እ", "፡የሆ", "፡ሁሉ", "ች፡በ", "፡በመ", "ሥራ፡", "፡ደህ", "ፈጸም", "ል፡መ", "ተግባ", "፡ድር", "ት፡ወ", "ው።ማ", "ፍርድ", "ርድ፡", "፡በሆ", "ር፡ወ", "በትም", "ትም፡", "ይነት", "ቸው፡", "ብ፡የ", "ነትና", "ቱን፡", "ሕግ፡", "ንና፡", "፡ሥራ", "የማግ", "፡መሠ", "ኘት፡", "፡ጊዜ", "ጻነቶ", "ነቶች", "በር፡", "በኀብ", "ዩነት", "ልዩነ", "፡በኀ", "፡ዓይ", "ዓይነ", "ችና፡", "ግባር", "ባር፡", "፡ደረ", "ነው።", "፡ነው", "ደረጃ", "ም።እ", "ም፡መ", "፡ወን", "ይማኖ", "ማኀበ", "ሃይማ", "፡ኑሮ", "መሠረ", "ሁሉ፡", "ነቱ፡", "ሌሎች", "ንግሥ", "በቅ፡", "የሆነ", "፡ይህ", "ንዲጠ", "ገር፡", "ተባበ", "ትክክ", "ጸም፡", "ር፡የ", "ዲጠበ", "ትም።", "ው፡ከ", "፡እያ", "ሩት፡", "ድርጅ", "፡ብቻ", "ና፡ለ", "ይገባ", "የመኖ", "፡ማን", "ንነት", "ቤተሰ", "ርጅት", "ት፡ድ", "፡መሰ", "እንደ", "፡አላ", "ብሔራ", "ት፡ለ", "ሔራዊ", "ርት፡", "ህርት", "ውን፡", "የሚያ", "ል።እ", "ሆኑ፡", "ምህር", "ትምህ", "በት።", "ለበት", "አለበ", "፡አስ", "ሎች፡", "ች፡የ", "፡በሕ", "ብረ፡", "፡ከሚ", "ን፡አ", "ት፡እ", "ን፡ወ", "ረግ፡", "በሆነ", "የኀብ", "፡የኀ", "መሆን", "፡መሆ", "ን፡መ", "፡ውሳ", "ንጀል", "ፈላጊ", "ህም፡", "ረታዊ", "ክለኛ", "ክክለ", "ታዊ፡", "ጀል፡", "ኑሮ፡", "።፡ይ", "ዓዊ፡", "ዜግነ", "ንዲሁ", "ዲሁም", "፡ማኀ", "ገሩ፡", "ር፡በ", "ብዓዊ", "አገሩ", "ሁም፡", "ና፡ነ", "ሰብዓ", "የተባ", "ጅት፡", "ማኖት", "ር፡አ", "ንግስ", "ኖት፡", "በሕግ", "መኖር", "ው፡ያ", "መጠበ", "ረጃ፡", "፡በማ", "ነትን", "ብነት", "ገብነ", "፡ገብ", "መፈጸ", "፡ሁኔ", "ሁኔታ", "ን፡ለ", "ው፡ለ", "፡ተግ", "፡የአ", "፡ይገ", "፡በአ", "ችን፡", "፡ትም", "ነቱን", "፡ቢሆ", "ቢሆን", "ጊዜ፡", "ረ፡ሰ", "ት፡ጊ", "ሰቡ፡", "ምበት", "ላቸው", "አላቸ", "በነጻ", "፡በነ", "አንድ", "ቅ፡መ", "፡መጠ", "ት፡ይ", "መሰረ", "ጥ፡የ", "ስጥ፡", "ፈጸመ", "ውስጥ", "ንድ፡", "፡ውስ", "፡በግ", "፡ሆኖ", "ሉ፡በ", "፡ጋብ", "ንስ፡", "ንነቱ", "መው፡", "የሚፈ", "አይፈ", "ብረሰ", "ነ፡መ", "፡የሃ", "ም፡ከ", "ች፡እ", "ስት፡", "ሙሉ፡", "አገር", "ሆኖ፡", "ደረግ", "ኢንተ", "ንተር", "ተርና", "ርናሽ", "ናሽና", "ሽናል"},
	Tir: []string{" መሰ", " ሰብ", "ሰብ ", " ኦለ", "ትን ", "ኦለዎ", "ናይ ", " ናይ", " ኦብ", "ዎ፡፡", "ለዎ፡", "ሕድሕ", "ኦብ ", "ድሕድ", "ሕድ ", "መሰል", "ውን ", "ሰል ", "ድ ሰ", "ይ ም", "ል ኦ", "ካብ ", "፡ሕድ", "፡፡ሕ", " ወይ", "ወይ ", " መን", " ነፃ", "ን መ", "ዝኾነ", "፡፡ ", "ታት ", "ብ ዝ", "ነት ", "ን ነ", " ካብ", "መሰላ", "ነፃነ", " እዚ", "ብ መ", "ኦዊ ", "ታትን", "መንግ", "ዊ መ", " እን", "ብ ብ", "ንግስ", "ት ኦ", "ሰላት", "ን ም", "ኾነ ", "እዚ ", "ብኦዊ", "ሰብኦ", "ን ኦ", "ን፡፡", " ንክ", " ዝኾ", "ን ን", " ምር", "ኹን ", "ይኹን", " ይኹ", "ምርካ", "ርካብ", " ኦይ", " ሃገ", "ሕጊ ", "ራት ", "ሎም ", " ብሕ", "ነ ይ", " ከም", "ማዕሪ", "ይ ብ", " ንም", " ዝተ", "ርን ", "ን ብ", "ራዊ ", " ፣ ", "ብ ሕ", "ላትን", "ብ ኦ", "ማሕበ", "ነታት", " ኦድ", "ዕሪ ", " ማዕ", "ስታት", "ግስታ", "’ውን", "ት መ", "ን ዝ", "ታዊ ", "፣ ብ", " ማሕ", "ነትን", "ንጋገ", "ድንጋ", " ስለ", " ድን", "ስራሕ", "ኩሎም", "ሕበራ", "ኦት ", "ን ሰ", "ዓለም", "ፃነታ", " ብም", "ት ወ", "መሰሪ", " ስራ", "ፃነት", "ተሰብ", "ካልኦ", "ልኦት", "ን ሓ", "ዓት ", "ዋን ", "ቡራት", "ሕቡራ", " ሕቡ", "ብሕጊ", "ድብ ", "ውድብ", " ውድ", "ብን ", "ትምህ", "ነቱ ", "ዚ ድ", "፣ ኦ", "ሃገራ", " ኩሎ", "ለዎም", "ምህር", "ም፡፡", "ም መ", " ብዝ", "ምኡ’", "ኡ’ው", "እንት", " ዓለ", " ብዘ", "በራዊ", " ሓለ", "ሓለዋ", "ዎም፡", "ቱ ን", "ት ብ", "ጋገ ", "ነፃ ", " ምዃ", "ን ዘ", " ገበ", "ት፣ ", " ትም", "ኸውን", "ራሕ ", " ዘይ", "ህርቲ", "ርቲ ", "ከምኡ", "ሃይማ", " ምስ", "ነ፣ ", "እንተ", " ስር", "ስርዓ", "ርዓት", "ባት ", "ይማኖ", "ሰሪታ", "ን ና", " ክብ", "ልን ", " ብማ", "ገሩ ", " ህዝ", "ላት ", "ት ና", "ይ ኦ", "ዕሊ ", "ለዝኾ", "ስለዝ", "ሪተሰ", "ብሪተ", "ሕብሪ", " ሕብ", "ን ተ", "ኾነ፣", "በን ", "ሃገሩ", "ገ እ", "ኻዊ ", " ሃይ", "እን ", "ሪጋገ", " ምሕ", "ን እ", "ለኻዊ", "ር፣ ", " ብሓ", " ብሃ", " ክኸ", "ክኸው", "ብ ዘ", "ዃኑ ", "ዊ ክ", "ምን ", "ሓደ ", "ምዃኑ", "ም ን", "ት እ", "ዊ ወ", "ታውን", " ሕድ", "ብዘይ", " ሕጊ", "ት ን", " ልዕ", " ካል", "ን ካ", "ሰባት", "ን ስ", "ናን ", "ቤተሰ", "ሕን ", "ለምለ", "ት ስ", "ምለኻ", "፣ ከ", "ተደን", "ባል ", "ኦድላ", "እዋን", " እዋ", "ደቂ ", " ደቂ", " ሰባ", "ፃን ", "ነፃን", "ግስቲ", "፣ ን", "ዚ ብ", "ስቲ ", " ቤተ", "ምጥሓ", " ክሳ", " ነዚ", "ን ክ", "ነቲ ", " ነቲ", "ነዚ ", " ምእ", "ብነፃ", " ምዕ", "ምዕባ", "ዕባለ", "ክሳብ", " ብነ", "ል እ", "ዚ መ", "ልዕሊ", "ክብሩ", "ብማዕ", "ሳብ ", "ህይወ", "ኦቶም", "ምስ ", "ንገገ", "እምነ", " እም", "ድ ኦ", "ቶም ", "ቲ ክ", "ፍትሓ", "ለም ", " ፍት", "ብ ን", "ን ዓ", "ራውን", "ሓፈሻ", "ደንገ", "ም ብ", "ትዮን", " ዝሰ", "ዝተደ", "ሉ መ", "ብ ና", "ጊ ካ", "ልዎ ", "ኦባል", " ኦባ", "ድልዎ", "ን ድ", "ኦድል", "ዜግነ", "ላውን", " ድሕ"},
}

var hebrewLangs = langProfileList{
	Heb: []string{"ות ", "ים ", "כל ", "ת ה", " כל", "דם ", "אדם", "יות", " של", " זכ", "ל א", " אד", "של ", "ל ה", "אי ", "ויו", "כאי", "ת ו", "י ל", "זכא", " ול", "לא ", " וה", "רות", "זכו", "ית ", "ירו", "ין ", " או", "ם ז", " לא", " הח", "או ", " הא", " וב", " המ", "חיר", "ת ל", "יים", "ם ל", "את ", "ת ב", "ת ש", "רה ", "ון ", " לה", "נה ", "כוי", "ותי", "ה ש", "ו ל", "ו ב", " הו", "ת א", "ם ב", "ם ו", "תו ", " את", "לה ", "ני ", "אומ", " במ", "דה ", "א י", "ה ה", "ה ב", "על ", "ם ה", " על", "הוא", "וך ", "ה א", "בוד", "וד ", "ואי", "נות", "ה ו", "ת כ", "י ה", "יה ", "ם ש", "ו ו", " שה", "ם א", "ו כ", "ינו", "ן ה", " שו", "שוו", "החי", "כות", "לאו", "בות", "דות", "ה ל", "לית", "ה מ", " בי", "וה ", "וא ", " הי", " לפ", "ור ", " לב", "ל ב", "בחי", "הכר", "לו ", "ת מ", "ן ש", "החו", "ה כ", " בכ", "ומי", "בין", "ן ו", "ן ל", "רוי", "פלי", "ולה", "ליה", " הז", "חינ", " לע", " בנ", "יבו", "חוק", " אח", "חבר", " יה", " חי", "מי ", "ירה", " חו", "האד", "ווה", "חופ", "ופש", "וק ", "נו ", "יו ", "ל מ", "מדי", "כבו", " הע", "נוך", " הד", "י א", "י ו", " הכ", "בני", "עה ", "ו א", "רצו", "דינ", "בזכ", "מות", "יפו", " אל", "סוד", "לם ", "איש", "רך ", " אי", "הגנ", "הם ", "פי ", "ם כ", "חות", "ל ו", "איל", "ילי", "תיה", "כלל", "אלי", "יסו", "האו", "זש ", " בא", "ר א", "ו ה", "זו ", "אחר", " הפ", " בע", " בז", "משפ", " בה", " לח", "דרך", "ומו", " בח", " דר", " מע", "ל י", "תוך", "מנו", " בש", "לל ", "רבו", " למ", "פני", " לק", "תם ", "שה ", "שית", "ללא", "לפי", "היה", "מעש", "דו ", "שות", "להג", "וצי", "שוא", "אין", "וי ", "תי ", "ונו", "ליל", " לו", "חיי", "ל ז", " זו", "היא", "יא ", "נתו", "ה פ", "לת ", "ובי", " לכ", "ך ה", "יל ", "י ש", "שיו", "ן ב", "עול", "המד", "ודה", "ולם", " ומ", "א ה", "ולא", " בת", "הכל", " סו", " מש", " עב", "סוצ", "ארצ", " אר", "ציא", "ד א", "לחי", "הן ", "יחס", " יח", "יאל", "הזכ", "ם נ", " שר", "בו ", "עבו", "היס", " לי", "ת ז", "פול", "יהי", "גבל", "תיו", "המא", "שהי", "א ל", "מאו", " יו", "ותו", "ישי", "גנה", "פשי", "וחד", "יהם", "חרו", "לכל", "ידה", "עות", "ונה", "ום ", "חה ", "עם ", "שרי", "ם י", "שר ", "והח", " אש", " הג", "ק ב", "הפל", "נשו", "הגב", "ד ו"},
	Ydd: []string{" פֿ", "ון ", "ער ", "ן א", " אַ", "דער", "ט א", " או", "און", "אַר", "ען ", "פֿו", " אױ", " אי", "ן פ", "ֿון", "רעכ", " דע", " רע", "עכט", "פֿא", "ן ד", "כט ", " די", "די ", "אַ ", "אױף", "ױף ", "ֿאַ", " זײ", " גע", "אַל", "אָס", " אָ", "ונג", " הא", "האָ", "זײַ", " מע", "אָל", "נג ", "װאָ", "ַן ", "אַנ", "רײַ", " װא", "ָס ", "באַ", " יע", "יעד", "ניט", "ן ז", "ר א", "יט ", "אָט", "אָר", "עדע", "מען", "זאָ", "ָט ", "פֿר", "ײַן", " בא", "טן ", "אין", "ן ג", "ין ", "ן װ", "נאַ", "ֿרײ", "ר ה", " זא", "לעכ", "ע א", "אָד", "ַ ר", "ענט", "אַצ", "ַצי", "אָנ", " צו", " װע", "יז ", "מענ", "ָדע", "איז", "ן מ", "ַלע", "בן ", "ר מ", "טער", " מי", " פּ", "מיט", "טלע", "ָל ", "עכע", "ײט ", "ַנד", "ע פ", "לע ", "געז", "לאַ", "אַפ", "עזע", "ראַ", " ני", "ַפֿ", "רן ", "ײַנ", "נען", "טיק", "כע ", "פֿע", "יע ", "הײט", "ַהײ", "נטש", "ײַה", "ט ד", "ן ב", "לן ", "ן נ", "פֿט", "שאַ", "רונ", " זי", " װי", "ט פ", " דא", "טאָ", "דיק", "קן ", "ר פ", "ר ג", "יקן", "אָב", "ף א", "אַק", "קער", "ערע", "כער", "י פ", "ות ", "ַרב", "פּר", "קט ", "עם ", "יאָ", "ציע", "ציא", "יט־", "צו ", "ישע", " קײ", "ן ק", "סער", " גל", "דאָ", "ונט", "גן ", "ַרא", "יקע", " טא", "ענע", "לײַ", "שן ", "ַנע", "יק ", "טאַ", "ס א", "עט ", "נגע", "ט־א", "ָנא", "־אי", "יקט", "נטע", "ײנע", "־ני", "ָר ", "װער", "י א", "ן י", "יך ", "זיך", "ער־", "ערן", "אױס", "ָבן", "נדע", "ָסע", "װי ", "ֿעל", "ר־נ", "ן ה", " גר", "גלײ", " צי", "ראָ", "זעל", "עלק", "נד ", "לקע", "אָפ", " כּ", "ט װ", "ג א", " נא", "ט צ", "ר ד", "עס ", "דור", "גען", "קע ", "ג פ", "ֿט ", "ן ל", "שע ", "ר ז", "רע ", "ײטן", "פּע", "קלא", "קײט", "יטע", "ים ", "ס ז", "ײַ ", " דו", "אַט", " לא", "ר װ", "קײנ", "עלש", "י ד", "לשא", "יות", "נט ", "ַרז", "ע ר", "ל ז", "אַמ", "ן ש", " שו", "אינ", "נטל", " הי", "בעט", "ָפּ", "ף פ", "ײַכ", "בער", "ן צ", "מאָ", " שט", " לע", "גער", "ורך", "רך ", "נעם", "גרו", "פֿן", "לער", "װעל", "ע מ", "ום ", "שפּ", "ך א", "יונ", "רבע", "עפֿ", "טעט", "ן כ", "רעס", "ערצ", "ז א", "עמע", "ם א", "שטע", "כן ", "רט ", "י ג", "סן ", "נער", "ליט", "ט ז", "נעמ", "ּרא", "היו", "אַש", "ת װ", "אומ", "ק א", "יבע", "ֿן ", "ץ א", "פֿי", "ײן ", "ם ט"},
}
