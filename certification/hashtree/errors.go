package hashtree

import (
	"fmt"
	"strings"
	"unicode"
    "encoding/hex"
)

// pathToString converts a path to a string, by joining the (string) labels with a slash.
func pathToString(path []Label) string {
	var sb strings.Builder
	for i, p := range path {
		if i > 0 {
			sb.WriteByte('/')
		}
		sb.WriteString(string(p))
	}
	return sb.String()
}

type LookupError struct {
	Type LookupResultType
	Path []Label
}

// NewLookupAbsentError returns a new LookupError with type LookupResultAbsent.
func NewLookupAbsentError(path ...Label) LookupError {
	return LookupError{
		Type: LookupResultAbsent,
		Path: path,
	}
}

// NewLookupError returns a new LookupError with type LookupResultError.
func NewLookupError(path ...Label) LookupError {
	return LookupError{
		Type: LookupResultError,
		Path: path,
	}
}

// NewLookupUnknownError returns a new LookupError with type LookupResultUnknown.
func NewLookupUnknownError(path ...Label) LookupError {
	return LookupError{
		Type: LookupResultUnknown,
		Path: path,
	}
}

func isASCII(s string) bool {
   for i := 0; i < len(s); i++ {
        if s[i] > unicode.MaxASCII {
            return false
        }
    }
    return true
}

func prettyLabel(label Label) string {
    str := string(label[:])
    if isASCII(str) {
        return str
    }

    return ""
}

func PrettyLabels(labels []Label) string {
    var ret string

    for i := 0; i < len(labels); i ++ {
        ret += "/"
        label := labels[i]
        pretty := string(label[:]);
        if isASCII(pretty) {
            ret += pretty
        } else {
            ret += hex.EncodeToString(label[:])
        }
    }

    return ret
}

func (l LookupError) Error() string {
	return fmt.Sprintf("lookup error (path: %s): %s", PrettyLabels(l.Path), l.error())
}

func (l LookupError) error() string {
	switch l.Type {
	case LookupResultAbsent:
		return "not found, not present in the tree"
	case LookupResultUnknown:
		return "not found, could be pruned"
	case LookupResultError:
		return "error, can not exist in the tree"
	default:
		return "unknown lookup error"
	}
}

// LookupResultType is the type of the lookup result.
// It indicates whether the result is guaranteed to be absent, unknown or is an invalid tree.
type LookupResultType int

const (
	// LookupResultAbsent means that the result is guaranteed to be absent.
	LookupResultAbsent LookupResultType = iota
	// LookupResultUnknown means that the result is unknown, some leaves were pruned.
	LookupResultUnknown
	// LookupResultError means that the result is an error, the path is not valid in this context.
	LookupResultError
)
