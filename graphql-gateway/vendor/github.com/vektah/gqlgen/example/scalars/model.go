package scalars

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/vektah/gqlgen/graphql"
)

type Banned bool

type User struct {
	ID       string
	Name     string
	Location Point     // custom scalar types
	Created  time.Time // direct binding to builtin types with external Marshal/Unmarshal methods
	IsBanned Banned    // aliased primitive
}

// Point is serialized as a simple array, eg [1, 2]
type Point struct {
	X int
	Y int
}

func (p *Point) UnmarshalGQL(v interface{}) error {
	pointStr, ok := v.(string)
	if !ok {
		return fmt.Errorf("points must be strings")
	}

	parts := strings.Split(pointStr, ",")

	if len(parts) != 2 {
		return fmt.Errorf("points must have 2 parts")
	}

	var err error
	if p.X, err = strconv.Atoi(parts[0]); err != nil {
		return err
	}
	if p.Y, err = strconv.Atoi(parts[1]); err != nil {
		return err
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (p Point) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%d,%d"`, p.X, p.Y)
}

// if the type referenced in types.json is a function that returns a marshaller we can use it to encode and decode
// onto any existing go type.
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

// Unmarshal{Typename} is only required if the scalar appears as an input. The raw values have already been decoded
// from json into int/float64/bool/nil/map[string]interface/[]interface
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}

type SearchArgs struct {
	Location     *Point
	CreatedAfter *time.Time
	IsBanned     Banned
}
