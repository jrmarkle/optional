package optional

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

type TestStruct struct {
	TestString string `json:"test_string,omitempty"`
	TestInt    int    `json:"test_int"`
}

type TestInnerStruct struct {
	I int
}

type TestStructWithOptional struct {
	OptionalString Optional[string]          `json:"optional_string,omitempty"`
	OptionalStruct Optional[TestInnerStruct] `json:"optional_struct,omitempty"`
}

func TestMarshallJSON(t *testing.T) {
	{
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(None[TestStruct]())
		if err != nil {
			t.Error(err)
		}
		if strings.TrimSpace(buffer.String()) != `{}` {
			t.Error(buffer.String())
		}
	}
	{
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(Some[TestStruct](TestStruct{}))
		if err != nil {
			t.Error(err)
		}
		if strings.TrimSpace(buffer.String()) != `{"test_int":0}` {
			t.Error(buffer.String())
		}
	}
	{
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(Some[TestStruct](TestStruct{TestString: "foo", TestInt: 42}))
		if err != nil {
			t.Error(err)
		}
		if strings.TrimSpace(buffer.String()) != `{"test_string":"foo","test_int":42}` {
			t.Error(buffer.String())
		}
	}

	{
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(TestStructWithOptional{OptionalString: None[string]()})
		if err != nil {
			t.Error(err)
		}

		if strings.TrimSpace(buffer.String()) != `{"optional_string":"","optional_struct":{}}` {
			t.Error(buffer.String())
		}
	}
	{
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(TestStructWithOptional{
			OptionalString: Some[string]("foo"),
			OptionalStruct: Some[TestInnerStruct](TestInnerStruct{I: 42}),
		})
		if err != nil {
			t.Error(err)
		}
		if strings.TrimSpace(buffer.String()) != `{"optional_string":"foo","optional_struct":{"I":42}}` {
			t.Error(buffer.String())
		}
	}
}

func TestUnmarshallJSON(t *testing.T) {
	{
		var value Optional[TestStruct]
		err := json.NewDecoder(strings.NewReader(`{}`)).Decode(&value)
		if err != nil {
			t.Error(err)
		}
		if !value.Is() {
			t.Error("not set by unmarshall")
		}
		if value.Get().TestInt != 0 ||
			value.Get().TestString != "" {
			t.Error("wrong value")
		}
	}
	{
		var value Optional[TestStruct]
		err := json.NewDecoder(strings.NewReader(`{"test_int":0}`)).Decode(&value)
		if err != nil {
			t.Error(err)
		}
		if !value.Is() {
			t.Error("not set by unmarshall")
		}
		if value.Get().TestInt != 0 ||
			value.Get().TestString != "" {
			t.Error("wrong value")
		}
	}
	{
		var value Optional[TestStruct]
		err := json.NewDecoder(strings.NewReader(`{"test_string":"foo","test_int":42}`)).Decode(&value)
		if err != nil {
			t.Error(err)
		}
		if !value.Is() {
			t.Error("not set by unmarshall")
		}
		if value.Get().TestInt != 42 ||
			value.Get().TestString != "foo" {
			t.Error("wrong value")
		}
	}

	{
		var value TestStructWithOptional
		err := json.NewDecoder(strings.NewReader(`{}`)).Decode(&value)
		if err != nil {
			t.Error(err)
		}
		if value.OptionalString.Is() {
			t.Error("string set by unmarshall")
		}
		if value.OptionalStruct.Is() {
			t.Error("struct set by unmarshall")
		}
	}
	{
		var value TestStructWithOptional
		err := json.NewDecoder(strings.NewReader(`{"optional_string":""}`)).Decode(&value)
		if err != nil {
			t.Error(err)
		}
		if !value.OptionalString.Is() {
			t.Error("string not set by unmarshall")
		}
		if value.OptionalString.Get() != "" {
			t.Error("wrong value")
		}
		if value.OptionalStruct.Is() {
			t.Error("struct set by unmarshall")
		}
	}
	{
		var value TestStructWithOptional
		err := json.NewDecoder(strings.NewReader(`{"optional_string":"","optional_struct":{}}`)).Decode(&value)
		if err != nil {
			t.Error(err)
		}
		if !value.OptionalString.Is() {
			t.Error("string not set by unmarshall")
		}
		if value.OptionalString.Get() != "" {
			t.Error("wrong value")
		}
		if !value.OptionalStruct.Is() {
			t.Error("struct not set by unmarshall")
		}
		if value.OptionalStruct.Get().I != 0 {
			t.Error("wrong value")
		}
	}
	{
		var value TestStructWithOptional
		err := json.NewDecoder(strings.NewReader(`{"optional_string":"foo","optional_struct":{"I":42}}`)).Decode(&value)
		if err != nil {
			t.Error(err)
		}
		if !value.OptionalString.Is() {
			t.Error("string not set by unmarshall")
		}
		if value.OptionalString.Get() != "foo" {
			t.Error("wrong value")
		}
		if !value.OptionalStruct.Is() {
			t.Error("struct not set by unmarshall")
		}
		if value.OptionalStruct.Get().I != 42 {
			t.Error("wrong value")
		}
	}
}
