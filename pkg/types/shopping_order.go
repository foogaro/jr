// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     csv_product.avsc
 *     csv_user.avsc
 *     finance_stock_trade.avsc
 *     fleetmgmt_description.avsc
 *     fleetmgmt_location.avsc
 *     fleetmgmt_sensor.avsc
 *     gaming_game.avsc
 *     gaming_player.avsc
 *     gaming_player_activity.avsc
 *     genericstore_purchase.avsc
 *     insurance_customer.avsc
 *     insurance_customer_activity.avsc
 *     insurance_offer.avsc
 *     inventorymgmt_inventory.avsc
 *     inventorymgmt_product.avsc
 *     iot_device_information.avsc
 *     map_dumb_schema.avsc
 *     marketing_campaign_finance.avsc
 *     net_device.avsc
 *     payment_credit_card.avsc
 *     payment_transaction.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizzastore_order.avsc
 *     pizzastore_order_cancelled.avsc
 *     pizzastore_order_completed.avsc
 *     shoestore_clickstream.avsc
 *     shoestore_customer.avsc
 *     shoestore_order.avsc
 *     shoestore_shoe.avsc
 *     shopping_order.avsc
 *     shopping_rating.avsc
 *     siem_log.avsc
 *     store.avsc
 *     syslog_log.avsc
 *     user.avsc
 *     users.avsc
 *     users_array_map.avsc
 *     webanalytics_clickstream.avsc
 *     webanalytics_code.avsc
 *     webanalytics_page_view.avsc
 *     webanalytics_user.avsc
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

type ShoppingOrder struct {
	Ordertime int64 `json:"ordertime"`

	Orderid int32 `json:"orderid"`

	Itemid string `json:"itemid"`

	Orderunits float64 `json:"orderunits"`

	Address Address `json:"address"`
}

const ShoppingOrderAvroCRC64Fingerprint = "\x14ʵ\x1d\xd3\xc7g\x8d"

func NewShoppingOrder() ShoppingOrder {
	r := ShoppingOrder{}
	r.Address = NewAddress()

	return r
}

func DeserializeShoppingOrder(r io.Reader) (ShoppingOrder, error) {
	t := NewShoppingOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeShoppingOrderFromSchema(r io.Reader, schema string) (ShoppingOrder, error) {
	t := NewShoppingOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeShoppingOrder(r ShoppingOrder, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Ordertime, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Orderid, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Itemid, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Orderunits, w)
	if err != nil {
		return err
	}
	err = writeAddress(r.Address, w)
	if err != nil {
		return err
	}
	return err
}

func (r ShoppingOrder) Serialize(w io.Writer) error {
	return writeShoppingOrder(r, w)
}

func (r ShoppingOrder) Schema() string {
	return "{\"fields\":[{\"name\":\"ordertime\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1519273364600,\"min\":1487715775521}},\"type\":\"long\"}},{\"name\":\"orderid\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"int\"}},{\"name\":\"itemid\",\"type\":{\"arg.properties\":{\"regex\":\"Item_[1-9][0-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"orderunits\",\"type\":{\"arg.properties\":{\"range\":{\"max\":10,\"min\":0.1}},\"type\":\"double\"}},{\"name\":\"address\",\"type\":{\"fields\":[{\"name\":\"city\",\"type\":{\"arg.properties\":{\"regex\":\"City_[1-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"state\",\"type\":{\"arg.properties\":{\"regex\":\"State_[1-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"zipcode\",\"type\":{\"arg.properties\":{\"range\":{\"max\":99999,\"min\":10000}},\"type\":\"long\"}}],\"name\":\"address\",\"type\":\"record\"}}],\"name\":\"shopping.ShoppingOrder\",\"type\":\"record\"}"
}

func (r ShoppingOrder) SchemaName() string {
	return "shopping.ShoppingOrder"
}

func (_ ShoppingOrder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ShoppingOrder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ShoppingOrder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ShoppingOrder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ShoppingOrder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ShoppingOrder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ShoppingOrder) SetString(v string)   { panic("Unsupported operation") }
func (_ ShoppingOrder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ShoppingOrder) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Ordertime}

		return w

	case 1:
		w := types.Int{Target: &r.Orderid}

		return w

	case 2:
		w := types.String{Target: &r.Itemid}

		return w

	case 3:
		w := types.Double{Target: &r.Orderunits}

		return w

	case 4:
		r.Address = NewAddress()

		w := types.Record{Target: &r.Address}

		return w

	}
	panic("Unknown field index")
}

func (r *ShoppingOrder) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ShoppingOrder) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ShoppingOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ShoppingOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ShoppingOrder) HintSize(int)                     { panic("Unsupported operation") }
func (_ ShoppingOrder) Finalize()                        {}

func (_ ShoppingOrder) AvroCRC64Fingerprint() []byte {
	return []byte(ShoppingOrderAvroCRC64Fingerprint)
}

func (r ShoppingOrder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ordertime"], err = json.Marshal(r.Ordertime)
	if err != nil {
		return nil, err
	}
	output["orderid"], err = json.Marshal(r.Orderid)
	if err != nil {
		return nil, err
	}
	output["itemid"], err = json.Marshal(r.Itemid)
	if err != nil {
		return nil, err
	}
	output["orderunits"], err = json.Marshal(r.Orderunits)
	if err != nil {
		return nil, err
	}
	output["address"], err = json.Marshal(r.Address)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ShoppingOrder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ordertime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ordertime); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ordertime")
	}
	val = func() json.RawMessage {
		if v, ok := fields["orderid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Orderid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for orderid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["itemid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Itemid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for itemid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["orderunits"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Orderunits); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for orderunits")
	}
	val = func() json.RawMessage {
		if v, ok := fields["address"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Address); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for address")
	}
	return nil
}
