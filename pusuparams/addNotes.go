package pusuparams

import (
	"github.com/nickwells/param.mod/v6/param"
	"github.com/nickwells/pusu.mod/pusu"
)

const (
	noteBaseName = "pub/sub - "

	// NoteNameNamespaces is the name used when adding the note on
	// Namespaces. It should be used if you want to reference the note
	// elsewhere, for instance, in a call to param.SeeNote(...).
	NoteNameNamespaces = noteBaseName + "Namespaces"
	// NoteNameTopics is the name used when adding the note on Topics. It
	// should be used if you want to reference the note elsewhere, for
	// instance, in a call to param.SeeNote(...).
	NoteNameTopics = noteBaseName + "Topics"
)

// AddNoteNamespaces adds a note, that will appear in the standard help,
// Notes section, explaining what a publish/subscribe namespace is.
func AddNoteNamespaces() param.PSetOptFunc {
	return func(ps *param.PSet) error {
		ps.AddNote(NoteNameNamespaces, pusu.NoteTextNamespace)

		return nil
	}
}

// AddNoteTopics adds a note, that will appear in the standard help,
// Notes section, explaining what a publish/subscribe topic is.
func AddNoteTopics() param.PSetOptFunc {
	return func(ps *param.PSet) error {
		ps.AddNote(NoteNameTopics, pusu.NoteTextTopic)

		return nil
	}
}
