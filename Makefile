basic:
	cue export json.cue

duplicate:
	cue eval dup.cue

constraint:
	cue eval check.cue

schema:
	cue export schema.cue

validate:
	cue vet validation.cue data.yaml

order:
	cue eval -i order.cue

fold:
	cue export fold.cue

types:
	cue eval types.cue

bottom:
	cue eval -i bottom.cue

numbers:
	cue eval -i numbers.cue

stringlit:
	cue export stringlit.cue

stringraw:
	cue eval stringraw.cue

bytes:
	cue export bytes.cue

structs:
	cue export -i structs.cue

defs:
	cue eval -ic defs.cue

structs2:
	cue eval -c structs2.cue

disjunctions:
	cue eval disjunctions.cue

defaults:
	cue eval defaults.cue

sumstruct:
	cue eval sumstruct.cue
