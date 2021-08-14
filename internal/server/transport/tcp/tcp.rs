use std::net::TcpListener;

fn main() {
  let connection_tcp = TcpListener::bind("127.0.0.1:7878").unwrap();

  for stream in connection_tcp.incoming() {
    let stream = stream.unwrap();
    println!("Listen TCP connections at the address 127.0.0.1:7878");
  }
}