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

bounds:
	cue eval -i bounds.cue

bounds2:
	cue eval -i bounds2.cue

lists:
	cue eval -i lists.cue

templates:
	cue eval -i templates.cue

scopes:
	cue eval scopes.cue

selectors:
	cue eval selectors.cue

alias:
	cue eval alias.cue

emit:
	cue eval emit.cue

cycle:
	cue eval -i -c cycle.cue

cycleref:
	cue eval cycleref.cue

hidden:
	cue export hidden.cue
