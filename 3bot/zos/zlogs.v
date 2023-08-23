module zos

pub struct ZLogs {
	zmachine string
	output   string
}

pub fn (z ZLogs) challenge() string {
	mut output := ''
	output += z.zmachine
	output += z.output

	return output
}
