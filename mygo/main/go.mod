module example.com/main

replace example.com/hello => ../hello

go 1.16

require example.com/hello v0.0.0-00010101000000-000000000000 // indirect
