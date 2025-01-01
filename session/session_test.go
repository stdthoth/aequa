package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession(t *testing.T) {
	ses := &Session{
		Name: "AEQUA",
		Lifetime: "100",
		Persists: "true",
		Domain: "localhost",
		SessionType: "cookie",
	}

	var scsmanager *scs.SessionManager

	sesInit := ses.NewSession()

	var sessionKind reflect.Kind
	var sessionType reflect.Type

	rv:= reflect.ValueOf(sesInit)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("For loop:", rv.Kind(), rv.Type(),rv)
		sessionKind = rv.Kind()
		sessionType = rv.Type()

		rv = rv.Elem()

	}

	if !rv.IsValid() {
		t.Error("invalid type or kind:", rv.Kind(), "type:", rv.Type())
	}

	if sessionKind != reflect.ValueOf(scsmanager).Kind() {
		t.Error("wrong kind of session testing cookie",reflect.ValueOf(scsmanager).Kind(), "and got:", sessionKind)
	}

	if sessionType != reflect.ValueOf(scsmanager).Type() {
		t.Error("wrong type of session testing cookie",reflect.ValueOf(scsmanager).Type(), "and got:", sessionType)
	}

	
}