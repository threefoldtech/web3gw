module zos

pub struct PublicIP {
	v4 bool
	v6 bool
}

pub fn (p PublicIP) challenge() string {
	mut output := ''
	output += '${p.v4}'
	output += '${p.v6}'

	return output
}
