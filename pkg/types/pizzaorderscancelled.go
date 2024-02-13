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

type Pizzaorderscancelled struct {
	Store_id int32 `json:"store_id"`

	Store_order_id int32 `json:"store_order_id"`

	Date int32 `json:"date"`

	Status string `json:"status"`
}

const PizzaorderscancelledAvroCRC64Fingerprint = "Hh\xe7\x81\xddf\xdd\f"

func NewPizzaorderscancelled() Pizzaorderscancelled {
	r := Pizzaorderscancelled{}
	return r
}

func DeserializePizzaorderscancelled(r io.Reader) (Pizzaorderscancelled, error) {
	t := NewPizzaorderscancelled()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePizzaorderscancelledFromSchema(r io.Reader, schema string) (Pizzaorderscancelled, error) {
	t := NewPizzaorderscancelled()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePizzaorderscancelled(r Pizzaorderscancelled, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Store_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Store_order_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Date, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Status, w)
	if err != nil {
		return err
	}
	return err
}

func (r Pizzaorderscancelled) Serialize(w io.Writer) error {
	return writePizzaorderscancelled(r, w)
}

func (r Pizzaorderscancelled) Schema() string {
	return "{\"fields\":[{\"name\":\"store_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":10,\"min\":1}},\"type\":\"int\"}},{\"name\":\"store_order_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1001,\"step\":2}},\"type\":\"int\"}},{\"name\":\"date\",\"type\":{\"arg.properties\":{\"range\":{\"max\":19000,\"min\":18000}},\"logicalType\":\"date\",\"type\":\"int\"}},{\"name\":\"status\",\"type\":{\"arg.properties\":{\"options\":[\"cancelled\"]},\"type\":\"string\"}}],\"name\":\"pizza_orders.pizzaorderscancelled\",\"type\":\"record\"}"
}

func (r Pizzaorderscancelled) SchemaName() string {
	return "pizza_orders.pizzaorderscancelled"
}

func (_ Pizzaorderscancelled) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) SetString(v string)   { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Pizzaorderscancelled) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Store_id}

		return w

	case 1:
		w := types.Int{Target: &r.Store_order_id}

		return w

	case 2:
		w := types.Int{Target: &r.Date}

		return w

	case 3:
		w := types.String{Target: &r.Status}

		return w

	}
	panic("Unknown field index")
}

func (r *Pizzaorderscancelled) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Pizzaorderscancelled) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Pizzaorderscancelled) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) HintSize(int)                     { panic("Unsupported operation") }
func (_ Pizzaorderscancelled) Finalize()                        {}

func (_ Pizzaorderscancelled) AvroCRC64Fingerprint() []byte {
	return []byte(PizzaorderscancelledAvroCRC64Fingerprint)
}

func (r Pizzaorderscancelled) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["store_id"], err = json.Marshal(r.Store_id)
	if err != nil {
		return nil, err
	}
	output["store_order_id"], err = json.Marshal(r.Store_order_id)
	if err != nil {
		return nil, err
	}
	output["date"], err = json.Marshal(r.Date)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Pizzaorderscancelled) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["store_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["store_order_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_order_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_order_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["date"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Date); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for date")
	}
	val = func() json.RawMessage {
		if v, ok := fields["status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for status")
	}
	return nil
}
