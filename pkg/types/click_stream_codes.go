// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
 *     click_stream.avsc
 *     click_stream_users.avsc
 *     clickstreamcodes.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurancecustomeractivity.avsc
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payrollbonus.avsc
 *     payrollemployee.avsc
 *     payrollemployeelocation.avsc
 *     pizzaorders.avsc
 *     pizzaorderscancelled.avsc
 *     pizzaorderscompleted.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siemlogs.avsc
 *     stock_trades.avsc
 *     stores.avsc
 *     sysloglogs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ClickStreamCodes struct {
	Code int32 `json:"code"`

	Definition string `json:"definition"`
}

const ClickStreamCodesAvroCRC64Fingerprint = "+d\xd5}\xf6\xc5F "

func NewClickStreamCodes() ClickStreamCodes {
	r := ClickStreamCodes{}
	return r
}

func DeserializeClickStreamCodes(r io.Reader) (ClickStreamCodes, error) {
	t := NewClickStreamCodes()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeClickStreamCodesFromSchema(r io.Reader, schema string) (ClickStreamCodes, error) {
	t := NewClickStreamCodes()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeClickStreamCodes(r ClickStreamCodes, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Code, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Definition, w)
	if err != nil {
		return err
	}
	return err
}

func (r ClickStreamCodes) Serialize(w io.Writer) error {
	return writeClickStreamCodes(r, w)
}

func (r ClickStreamCodes) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"code\":200,\"definition\":\"Successful\"},{\"code\":302,\"definition\":\"Redirect\"},{\"code\":404,\"definition\":\"Page not found\"},{\"code\":405,\"definition\":\"Method not allowed\"},{\"code\":406,\"definition\":\"Not acceptable\"},{\"code\":407,\"definition\":\"Proxy authentication required\"}]},\"fields\":[{\"name\":\"code\",\"type\":\"int\"},{\"name\":\"definition\",\"type\":\"string\"}],\"name\":\"clickstream.ClickStreamCodes\",\"type\":\"record\"}"
}

func (r ClickStreamCodes) SchemaName() string {
	return "clickstream.ClickStreamCodes"
}

func (_ ClickStreamCodes) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ClickStreamCodes) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ClickStreamCodes) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ClickStreamCodes) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ClickStreamCodes) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ClickStreamCodes) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ClickStreamCodes) SetString(v string)   { panic("Unsupported operation") }
func (_ ClickStreamCodes) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ClickStreamCodes) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Code}

		return w

	case 1:
		w := types.String{Target: &r.Definition}

		return w

	}
	panic("Unknown field index")
}

func (r *ClickStreamCodes) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ClickStreamCodes) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ClickStreamCodes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ClickStreamCodes) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ClickStreamCodes) HintSize(int)                     { panic("Unsupported operation") }
func (_ ClickStreamCodes) Finalize()                        {}

func (_ ClickStreamCodes) AvroCRC64Fingerprint() []byte {
	return []byte(ClickStreamCodesAvroCRC64Fingerprint)
}

func (r ClickStreamCodes) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["code"], err = json.Marshal(r.Code)
	if err != nil {
		return nil, err
	}
	output["definition"], err = json.Marshal(r.Definition)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ClickStreamCodes) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["code"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Code); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for code")
	}
	val = func() json.RawMessage {
		if v, ok := fields["definition"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Definition); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for definition")
	}
	return nil
}
