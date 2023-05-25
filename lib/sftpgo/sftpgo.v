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
	jwt string 
}

pub fn new(args SFTPGOClientArgs) SFTPGoClient {
	header := construct_header(args.jwt)
	return SFTPGoClient{
		address: args.address,
		header: header
	}
}

pub fn construct_header(jwt string) http.Header{
	header_config := http.HeaderConfig{
		key: http.CommonHeader.authorization
		value: 'bearer ${jwt}'
	}
	return http.new_header(header_config)
}


