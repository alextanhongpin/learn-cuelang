// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/alextanhongpin/learn-cuelang/go

package main

// User represents a user object.
#User: {
	name: string            @go(Name)
	age:  int & (>0 & <100) @go(Age)
}
