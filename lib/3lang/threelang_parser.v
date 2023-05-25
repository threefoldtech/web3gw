module threelang

pub struct ThreeLangParser{
	tfgrid_parser TFGridParser
	tfchain_parser TFChainParser
	// other modules parsers
}

// parse takes an md file path as input, preprocesses it, returns a ThreeLangParser instance
pub parse(path string)!ThreeLangParser{

}

// execute performs all actions specified inside the md file
pub fn(t ThreeLangParser) execute()!{

}