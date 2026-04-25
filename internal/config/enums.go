package config

// OnErrorAction controls multi-statement behavior after a statement fails.
type OnErrorAction string

const (
	// OnErrorResume continues executing remaining statements after an error.
	OnErrorResume OnErrorAction = "RESUME"
	// OnErrorStop stops execution immediately when a statement fails.
	OnErrorStop OnErrorAction = "STOP"
)

func (a OnErrorAction) isValid() bool {
	switch a {
	case OnErrorResume, OnErrorStop:
		return true
	default:
		return false
	}
}

type TableColor string

const (
	FgDefault   TableColor = "default" // default is cyan for header, white for column and caption
	FgBlack     TableColor = "black"
	FgRed       TableColor = "red"
	FgGreen     TableColor = "green"
	FgYellow    TableColor = "yellow"
	FgBlue      TableColor = "blue"
	FgMagenta   TableColor = "magenta"
	FgCyan      TableColor = "cyan"
	FgWhite     TableColor = "white"
	FgHiBlack   TableColor = "black+"
	FgHiRed     TableColor = "red+"
	FgHiGreen   TableColor = "green+"
	FgHiYellow  TableColor = "yellow+"
	FgHiBlue    TableColor = "blue+"
	FgHiMagenta TableColor = "magenta+"
	FgHiCyan    TableColor = "cyan+"
	FgHiWhite   TableColor = "white+"
)

func (c TableColor) isValid() bool {
	_, ok := validTableColors[c]
	return ok
}

var validTableColors = map[TableColor]struct{}{
	FgDefault:   {},
	FgBlack:     {},
	FgRed:       {},
	FgGreen:     {},
	FgYellow:    {},
	FgBlue:      {},
	FgMagenta:   {},
	FgCyan:      {},
	FgWhite:     {},
	FgHiBlack:   {},
	FgHiRed:     {},
	FgHiGreen:   {},
	FgHiYellow:  {},
	FgHiBlue:    {},
	FgHiMagenta: {},
	FgHiCyan:    {},
	FgHiWhite:   {},
}

type TableStyle string

const (
	StyleNone        TableStyle = "none"
	StyleASCII       TableStyle = "ascii"
	StyleLight       TableStyle = "light"
	StyleHeavy       TableStyle = "heavy"
	StyleDouble      TableStyle = "double"
	StyleDoubleLong  TableStyle = "double_long"
	StyleLightHeavy  TableStyle = "light_heavy"
	StyleHeavyLight  TableStyle = "heavy_light"
	StyleLightDouble TableStyle = "light_double"
	StyleDoubleLight TableStyle = "double_light"
	StyleRounded     TableStyle = "rounded"
	StyleMarkdown    TableStyle = "markdown"
	StyleGraphical   TableStyle = "graphical"
	StyleMerger      TableStyle = "merger"
	StyleDefault     TableStyle = "default"
	StyleDotted      TableStyle = "dotted"
	StyleArrow       TableStyle = "arrow"
	StyleStarry      TableStyle = "starry"
	StyleHearts      TableStyle = "hearts"
	StyleCircuit     TableStyle = "circuit"
	StyleNature      TableStyle = "nature"
	StyleArtistic    TableStyle = "artistic"
	Style8Bit        TableStyle = "8bit"
	StyleChaos       TableStyle = "chaos"
	StyleDots        TableStyle = "dots"
	StyleBlocks      TableStyle = "blocks"
	StyleZen         TableStyle = "zen"
	StyleVintage     TableStyle = "vintage"
	StyleSketch      TableStyle = "sketch"
	StyleArrowDouble TableStyle = "arrow_double"
	StyleCelestial   TableStyle = "celestial"
	StyleCyber       TableStyle = "cyber"
	StyleRunic       TableStyle = "runic"
	StyleIndustrial  TableStyle = "industrial"
	StyleInk         TableStyle = "ink"
	StyleArcade      TableStyle = "arcade"
	StyleBlossom     TableStyle = "blossom"
	StyleFrosted     TableStyle = "frosted"
	StyleMosaic      TableStyle = "mosaic"
	StyleUFO         TableStyle = "ufo"
	StyleSteampunk   TableStyle = "steampunk"
	StyleGalaxy      TableStyle = "galaxy"
	StyleJazz        TableStyle = "jazz"
	StylePuzzle      TableStyle = "puzzle"
	StyleHypno       TableStyle = "hypno"
)

func (s TableStyle) isValid() bool {
	_, ok := validTableStyles[s]
	return ok
}

var validTableStyles = map[TableStyle]struct{}{
	StyleNone:        {},
	StyleASCII:       {},
	StyleLight:       {},
	StyleHeavy:       {},
	StyleDouble:      {},
	StyleDoubleLong:  {},
	StyleLightHeavy:  {},
	StyleHeavyLight:  {},
	StyleLightDouble: {},
	StyleDoubleLight: {},
	StyleRounded:     {},
	StyleMarkdown:    {},
	StyleGraphical:   {},
	StyleMerger:      {},
	StyleDefault:     {},
	StyleDotted:      {},
	StyleArrow:       {},
	StyleStarry:      {},
	StyleHearts:      {},
	StyleCircuit:     {},
	StyleNature:      {},
	StyleArtistic:    {},
	Style8Bit:        {},
	StyleChaos:       {},
	StyleDots:        {},
	StyleBlocks:      {},
	StyleZen:         {},
	StyleVintage:     {},
	StyleSketch:      {},
	StyleArrowDouble: {},
	StyleCelestial:   {},
	StyleCyber:       {},
	StyleRunic:       {},
	StyleIndustrial:  {},
	StyleInk:         {},
	StyleArcade:      {},
	StyleBlossom:     {},
	StyleFrosted:     {},
	StyleMosaic:      {},
	StyleUFO:         {},
	StyleSteampunk:   {},
	StyleGalaxy:      {},
	StyleJazz:        {},
	StylePuzzle:      {},
	StyleHypno:       {},
}
