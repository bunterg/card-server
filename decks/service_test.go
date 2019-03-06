package decks

import (
	"reflect"
	"testing"

	"github.com/bunterg/card-server/cards"
)

func TestNewService(t *testing.T) {
	type args struct {
		dR Repository
		cR cards.Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.dR, tt.args.cR); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AddDeck(t *testing.T) {
	type fields struct {
		dR Repository
		cR cards.Repository
	}
	type args struct {
		ds []Deck
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				dR: tt.fields.dR,
				cR: tt.fields.cR,
			}
			s.AddDeck(tt.args.ds...)
		})
	}
}
