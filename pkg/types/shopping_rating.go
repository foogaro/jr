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

type ShoppingRating struct {
	Rating_id int64 `json:"rating_id"`

	User_id int32 `json:"user_id"`

	Stars int32 `json:"stars"`

	Route_id int32 `json:"route_id"`

	Rating_time int64 `json:"rating_time"`

	Channel string `json:"channel"`

	Message string `json:"message"`
}

const ShoppingRatingAvroCRC64Fingerprint = ")\x05\xa2:\xcb\xdd\xf4/"

func NewShoppingRating() ShoppingRating {
	r := ShoppingRating{}
	return r
}

func DeserializeShoppingRating(r io.Reader) (ShoppingRating, error) {
	t := NewShoppingRating()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeShoppingRatingFromSchema(r io.Reader, schema string) (ShoppingRating, error) {
	t := NewShoppingRating()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeShoppingRating(r ShoppingRating, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Rating_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.User_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Stars, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Route_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Rating_time, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Channel, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Message, w)
	if err != nil {
		return err
	}
	return err
}

func (r ShoppingRating) Serialize(w io.Writer) error {
	return writeShoppingRating(r, w)
}

func (r ShoppingRating) Schema() string {
	return "{\"fields\":[{\"name\":\"rating_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1}},\"type\":\"long\"}},{\"name\":\"user_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":20,\"min\":-1}},\"type\":\"int\"}},{\"name\":\"stars\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5,\"min\":1}},\"type\":\"int\"}},{\"name\":\"route_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":9999,\"min\":1}},\"type\":\"int\"}},{\"name\":\"rating_time\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1,\"step\":12}},\"format_as_time\":\"unix_long\",\"type\":\"long\"}},{\"name\":\"channel\",\"type\":{\"arg.properties\":{\"options\":[\"web\",\"iOS\",\"ios\",\"iOS-test\",\"android\"]},\"type\":\"string\"}},{\"name\":\"message\",\"type\":{\"arg.properties\":{\"options\":[\"worst. flight. ever. #neveragain\",\"Surprisingly good, maybe you are getting your mojo back at long last!\",\"Exceeded all my expectations. Thank you !\",\"meh\",\"(expletive deleted)\",\"is this as good as it gets? really ?\",\"airport refurb looks great, will fly outta here more!\",\"why is it so difficult to keep the bathrooms clean ?\",\"thank you for the most friendly, helpful experience today at your new lounge\",\"more peanuts please\",\"your team here rocks!\"]},\"type\":\"string\"}}],\"name\":\"shopping.ShoppingRating\",\"type\":\"record\"}"
}

func (r ShoppingRating) SchemaName() string {
	return "shopping.ShoppingRating"
}

func (_ ShoppingRating) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ShoppingRating) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ShoppingRating) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ShoppingRating) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ShoppingRating) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ShoppingRating) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ShoppingRating) SetString(v string)   { panic("Unsupported operation") }
func (_ ShoppingRating) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ShoppingRating) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Rating_id}

		return w

	case 1:
		w := types.Int{Target: &r.User_id}

		return w

	case 2:
		w := types.Int{Target: &r.Stars}

		return w

	case 3:
		w := types.Int{Target: &r.Route_id}

		return w

	case 4:
		w := types.Long{Target: &r.Rating_time}

		return w

	case 5:
		w := types.String{Target: &r.Channel}

		return w

	case 6:
		w := types.String{Target: &r.Message}

		return w

	}
	panic("Unknown field index")
}

func (r *ShoppingRating) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ShoppingRating) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ShoppingRating) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ShoppingRating) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ShoppingRating) HintSize(int)                     { panic("Unsupported operation") }
func (_ ShoppingRating) Finalize()                        {}

func (_ ShoppingRating) AvroCRC64Fingerprint() []byte {
	return []byte(ShoppingRatingAvroCRC64Fingerprint)
}

func (r ShoppingRating) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["rating_id"], err = json.Marshal(r.Rating_id)
	if err != nil {
		return nil, err
	}
	output["user_id"], err = json.Marshal(r.User_id)
	if err != nil {
		return nil, err
	}
	output["stars"], err = json.Marshal(r.Stars)
	if err != nil {
		return nil, err
	}
	output["route_id"], err = json.Marshal(r.Route_id)
	if err != nil {
		return nil, err
	}
	output["rating_time"], err = json.Marshal(r.Rating_time)
	if err != nil {
		return nil, err
	}
	output["channel"], err = json.Marshal(r.Channel)
	if err != nil {
		return nil, err
	}
	output["message"], err = json.Marshal(r.Message)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ShoppingRating) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["rating_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rating_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rating_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["user_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for user_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["stars"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Stars); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for stars")
	}
	val = func() json.RawMessage {
		if v, ok := fields["route_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Route_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for route_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["rating_time"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rating_time); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rating_time")
	}
	val = func() json.RawMessage {
		if v, ok := fields["channel"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Channel); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for channel")
	}
	val = func() json.RawMessage {
		if v, ok := fields["message"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Message); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for message")
	}
	return nil
}
