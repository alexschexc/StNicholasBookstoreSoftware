with import <nixpkgs> {};
mkShell {
	buildInputs = [
		go
		postgresql
		fyne
	];

}
