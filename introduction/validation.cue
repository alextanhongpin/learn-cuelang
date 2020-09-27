#Language: {
	tag: string
	// https://stackoverflow.com/questions/25977309/what-does-this-regexp-mean-plu
	name: =~"^\\p{Lu}" // Must start with uppercase letter.
}
languages: [...#Language]
