get:
	cue get go . --local

openapi:
	cue def --out openapi

filetypes:
	cue filetypes


validate:
	cue vet ranges.json ranges.cue
	cue vet ranges.json ranges.v2.cue
	cue vet ranges.json ranges*.cue

validate-list:
	cue vet ranges.list.json ranges.v2.cue ranges.list.cue
