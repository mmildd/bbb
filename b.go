package main

import "fmt"

type WasherEntity struct {
	name string
	money int
}

// Fields of the covered
func (WasherEntity) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time"),
	}
}

// Edges of the covered
func (WasherEntity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resolution", Resolution.Type).Ref("WasherEntity").Unique(),
	}
}
x