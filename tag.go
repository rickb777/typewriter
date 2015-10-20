package typewriter

type Tag struct {
	Name    string
	Values  []TagValue
	Negated bool
}

type TagValue struct {
	Name           string
	TypeParameters []Type
	typeParameters []item
}

type TagSlice []Tag
