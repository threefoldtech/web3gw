module sftpgo
import net.http


[noinit]
pub struct SFTPGoClient {
pub mut:
	address string
	header http.Header
}

[params]
pub struct SFTPGOClientArgs {
pub:
	address    string = 'http://localhost:8080/api/v2'
	key string 
}

pub fn new(args SFTPGOClientArgs) !SFTPGoClient {
	header := http.new_custom_header_from_map({'X-SFTPGO-API-KEY': args.key})!
	return SFTPGoClient{
		address: args.address,
		header: header
	}
}

