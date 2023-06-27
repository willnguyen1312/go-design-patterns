package patterns

import (
	"strings"
	"testing"
)

func TestGetRenderReactUI(t *testing.T) {
	ui, err := GetRenderUI(ReactElement)
	if err != nil {
		t.Fatal("A ui of type 'ReactElement' must exist")
	}

	msg := ui.Request("react")
	if !strings.Contains(msg, "Request URL react by using React") {
		t.Error("The React ui message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetRenderVueUI(t *testing.T) {
	ui, err := GetRenderUI(VueElement)
	if err != nil {
		t.Fatal("A ui of type 'VueElement' must exist")
	}

	msg := ui.Request("vue")
	if !strings.Contains(msg, "Request URL vue by using Vue") {
		t.Error("The Vue ui message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetRenderAngularUI(t *testing.T) {
	ui, err := GetRenderUI(AngularElement)
	if err != nil {
		t.Fatal("A ui of type 'AngularElement' must exist")
	}

	msg := ui.Request("angular")
	if !strings.Contains(msg, "Request URL angular by using Angular") {
		t.Error("The Angular ui message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetRenderUI(20)
	if err == nil {
		t.Error("A ui with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
