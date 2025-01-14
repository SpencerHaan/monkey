package tracer

import (
	"fmt"
	"strings"
)

const traceIdentPlaceholder string = "\t"

var head *Tracer

type Tracer struct {
	parent *Tracer
	level  int
	label  string
}

func Begin(label string) *Tracer {
	var level int
	if head == nil {
		level = 0
	} else {
		level = head.level + 1
	}

	head = &Tracer{
		parent: head,
		level:  level,
		label:  label,
	}

	// head.print("BEGIN")
	return head
}

func (t *Tracer) End() {
	// head.print("END")
	head = t.parent
}

func (t *Tracer) print(msg string) {
	ident := strings.Repeat(traceIdentPlaceholder, t.level)
	fmt.Printf("%s%s\n", ident, msg+" "+t.label)
}
