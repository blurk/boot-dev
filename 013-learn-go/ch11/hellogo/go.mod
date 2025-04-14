module example.com/username/hellogo

go 1.24.2

replace example.com/username/mystrings v0.0.0 => ../mystrings

require (
	example.com/username/mystrings v0.0.0
)
