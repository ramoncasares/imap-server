package imap

import "time"

// Mailstore is an interface to be implemented to provide mail storage
type Mailstore interface {
	// Attempt to authenticate a user with given credentials,
	// and return the user if successful
	Authenticate(username string, password string) (User, error)
}

// User represents a user in the mail storage system
type User interface {
	// Return a list of mailboxes belonging to this user
	Mailboxes() []Mailbox

	MailboxByName(name string) (Mailbox, error)
}

// Mailbox represents a mailbox belonging to a user in the mail storage system
type Mailbox interface {
	// The name of the mailbox
	Name() string

	// The unique identifier that will LIKELY be assigned
	// to the next mail that is added to this mailbox
	NextUID() uint32

	// Number of recent messages in the mailbox
	Recent() uint32

	// Number of messages in the mailbox
	Messages() uint32

	// Number messages that do not have the Unseen flag set yet
	Unseen() uint32

	// Get a message by its sequence number
	MessageBySequenceNumber(seqno uint32) Message

	// Get a message by its uid number
	MessageByUID(uidno uint32) Message

	// Get messages that belong to a set of ranges of UIDs
	MessageSetByUID(set SequenceSet) []Message

	// Get messages that belong to a set of ranges of sequence numbers
	MessageSetBySequenceNumber(set SequenceSet) []Message
}

// Message represents a standard email message
type Message interface {
	// Return the message's MIME headers as a map in format
	// key: value
	Header() MIMEHeader

	// Return the unique id of the email
	UID() uint32

	// Return the sequence number of the email
	SequenceNumber() uint32

	// Return the RFC822 size of the message
	Size() uint32

	// Return the date the email was received by the server
	// (This is not the date on the envelope of the email)
	InternalDate() time.Time

	// Return the body of the email
	Body() string

	// Return the list of custom keywords/flags for this message
	Keywords() []string

	IsSeen() bool
	IsAnswered() bool
	IsFlagged() bool
	IsDeleted() bool
	IsDraft() bool
	IsRecent() bool
}
