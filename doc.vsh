import os
import freeflowuniverse.crystallib.pathlib

const docs_dir = dir(@FILE) + '/docs'

fn sh(cmd string) {
	$if debug {
		println('❯ ${cmd}')
		print(execute(cmd).output)
	}
	execute(cmd)
}

// creates a new dir at path, removing existing
fn new(path string) ! {
	if exists(path) {
		rmdir_all(path)!
	}
	mkdir_all(path)!
}

fn build_manual() ! {
	println('Building manual')
	new('docs/_docs')!
	manual_dir := dir(@FILE) + '/manual'
	sh('bash $manual_dir/run.sh')
}

fn build_openrpc_docs() ! {
	println('Building OpenRPC Docs')
	mkdir('docs/openrpc')!
	sh('v ~/.vmodules/freeflowuniverse/crystallib/openrpc/cli')
	cli := '~/.vmodules/freeflowuniverse/crystallib/openrpc/cli/cli'
	sh('$cli docgen -t "Web3Proxy JSON-RPC API" -p -o docs/openrpc lib')
	mut lib_path := pathlib.get('lib')
	clients := lib_path.dir_list(recursive: true)!
	for client in clients {
		client_name := client.path.all_after_last('/')
		sh('mkdir docs/openrpc/$client_name')
		doc_title := "$client_name JSON-RPC API"
		sh('$cli docgen -exclude_dirs threelang -t "$doc_title" -p -o docs/openrpc/$client_name $client.path')
	}
}

// builds the inspector module used by the playground
fn build_inspector() ! {
	println('Building OpenRPC Inspector')
	sh('git clone https://github.com/timurgordon/inspector.git $docs_dir/_inspector')
	chdir('$docs_dir/_inspector')!
	sh('npm install')
	sh('npm run build')
	sh('npm run build:package')
	chdir('$docs_dir')!
	
	// mount built dirs to docs/inspector
	new('$docs_dir/inspector')!
	cp_all('$docs_dir/_inspector/build', '$docs_dir/inspector', true)!
	new('$docs_dir/inspector/package')!
	cp_all('$docs_dir/_inspector/package', '$docs_dir/inspector/package', true)!
	rmdir_all('$docs_dir/_inspector')! // cleanup
}

// builds the inspector module used by the playground
fn build_playground(url string) ! {
	println('Building OpenRPC Playground')
	sh('git clone https://github.com/open-rpc/playground.git $docs_dir/_playground')
	chdir('$docs_dir/_playground')!
	sh('npm install')
	// replace inspector module with forked module package
	sh('mv -f $docs_dir/inspector/package $docs_dir/_playground/node_modules/@open-rpc/inspector/package')
	// modify homepage	
	package_json := read_file('$docs_dir/_playground/package.json')!
	modified_package_json := package_json.replace('"homepage": "https://playground.open-rpc.org/"', '"homepage": "$url/playground"')
	write_file('$docs_dir/_playground/package.json', modified_package_json)!

	sh('npm run build')
	chdir('$docs_dir')!
	new('$docs_dir/playground')!
	cp_all('$docs_dir/_playground/build', '$docs_dir/playground', true)!
	rmdir_all('$docs_dir/_playground')! // cleanup
}

url := if os.args.len == 2 {
	'${os.args[1]}'
} else {''}

new(docs_dir)!
build_manual()!
build_openrpc_docs()!
build_inspector()!
build_playground(url)!