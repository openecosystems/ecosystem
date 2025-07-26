package sdkv2betalib

import "golang.org/x/text/language"

// Package locale provides a static list of common BCP 47 language tags.
var (
	English            = language.English         // "en"
	EnglishUS          = language.AmericanEnglish // "en-US"
	EnglishUK          = language.BritishEnglish  // "en-GB"
	French             = language.French          // "fr"
	FrenchFrance       = language.Make("fr-FR")   // France
	FrenchCanada       = language.Make("fr-CA")   // Canada
	Spanish            = language.Spanish         // "es"
	SpanishMexico      = language.Make("es-MX")
	Portuguese         = language.Portuguese // "pt"
	PortugueseBrazil   = language.Make("pt-BR")
	German             = language.German // "de"
	GermanGermany      = language.Make("de-DE")
	Japanese           = language.Japanese // "ja"
	JapaneseJapan      = language.Make("ja-JP")
	ChineseSimplified  = language.Make("zh-CN")
	ChineseTraditional = language.Make("zh-TW")
	Korean             = language.Korean // "ko"
	KoreanSouthKorea   = language.Make("ko-KR")
	Arabic             = language.Arabic  // "ar"
	Russian            = language.Russian // "ru"
	Hindi              = language.Make("hi-IN")
)

var (
	LocaleEnglishUS = EnglishUS.String()
	LocaleEnglishUK = EnglishUK.String()
)

var SupportedLanguages = []language.Tag{
	EnglishUS,
	FrenchFrance,
}
