module sftpgo
import net.http


[noinit]
pub struct SFTPGoClient {
pub mut:
	url string
	jwt string
	header http.Header
}

[params]
pub struct SFTPGOClientArgs {
pub:
	url    string = 'http://localhost:8080/api/v2'
	jwt string 
}

pub fn new(args SFTPGOClientArgs) SFTPGoClient {
	return SFTPGoClient{
		url: args.url,
		jwt: args.jwt
	}
}

pub fn (mut cl SFTPGoClient) construct_header() http.Header{
	header_config := http.HeaderConfig{
		key: http.CommonHeader.authorization
		value: 'bearer ${cl.jwt}'
	}
	return http.new_header(header_config)
}


