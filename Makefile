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
