module threelang

fn test_parse_content() ! {
	mut t := parse(
		content: "
!!tfgrid.login
	mnemonic: 'PUT YOUR MNEMONICS HERE'
	network: dev

!!tfgrid.gateway_name.create 
	name: hamadagateway
	backend: http://1.1.1.1:9000

!!tfgrid.gateway_name.get
	name: hamadagateway

!!tfgrid.gateway_name.delete
	name: hamadagateway
"
	)!

	t.execute()!
}

fn test_parse_invalid_module() {
	_ := parse(
			content: '
!!invalid.ns.op
	key1:"value1"
'
	)!
}

fn test_parse_invalid_name(){
	_ := parse(
		content: '
!!tfgrid.ns.op.too.many
	key1:"value1"
'
	)!
}

fn test_invalid_doc(){
	_ := parse(
		content: '
!!tfgrid.ns.op.too.many
	key1:"value1"

this is some invalid text inside the doc
'
	)!
}

fn test_parse_valid_file(){
	parse(path: './test_doc.md')!
}
