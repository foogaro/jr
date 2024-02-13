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

type Clickstream struct {
	Ip string `json:"ip"`

	Userid int32 `json:"userid"`

	Remote_user string `json:"remote_user"`

	Time string `json:"time"`

	Logtime int64 `json:"_logtime"`

	Request string `json:"request"`

	Status string `json:"status"`

	Bytes string `json:"bytes"`

	Referrer string `json:"referrer"`

	Agent string `json:"agent"`
}

const ClickstreamAvroCRC64Fingerprint = "\x7f\r6\xa7Z\x17i*"

func NewClickstream() Clickstream {
	r := Clickstream{}
	return r
}

func DeserializeClickstream(r io.Reader) (Clickstream, error) {
	t := NewClickstream()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeClickstreamFromSchema(r io.Reader, schema string) (Clickstream, error) {
	t := NewClickstream()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeClickstream(r Clickstream, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Ip, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Userid, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Remote_user, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Time, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Logtime, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Request, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Status, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Bytes, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Referrer, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Agent, w)
	if err != nil {
		return err
	}
	return err
}

func (r Clickstream) Serialize(w io.Writer) error {
	return writeClickstream(r, w)
}

func (r Clickstream) Schema() string {
	return "{\"fields\":[{\"name\":\"ip\",\"type\":{\"arg.properties\":{\"options\":[\"111.152.45.45\",\"111.203.236.146\",\"111.168.57.122\",\"111.249.79.93\",\"111.168.57.122\",\"111.90.225.227\",\"111.173.165.103\",\"111.145.8.144\",\"111.245.174.248\",\"111.245.174.111\",\"222.152.45.45\",\"222.203.236.146\",\"222.168.57.122\",\"222.249.79.93\",\"222.168.57.122\",\"222.90.225.227\",\"222.173.165.103\",\"222.145.8.144\",\"222.245.174.248\",\"222.245.174.222\",\"122.152.45.245\",\"122.203.236.246\",\"122.168.57.222\",\"122.249.79.233\",\"122.168.57.222\",\"122.90.225.227\",\"122.173.165.203\",\"122.145.8.244\",\"122.245.174.248\",\"122.245.174.122\",\"233.152.245.45\",\"233.203.236.146\",\"233.168.257.122\",\"233.249.279.93\",\"233.168.257.122\",\"233.90.225.227\",\"233.173.215.103\",\"233.145.28.144\",\"233.245.174.248\",\"233.245.174.233\"]},\"session\":\"true\",\"type\":\"string\"}},{\"name\":\"userid\",\"type\":{\"arg.properties\":{\"range\":{\"max\":40,\"min\":-1}},\"session-sibling-int-hash\":\"true\",\"type\":\"int\"}},{\"name\":\"remote_user\",\"type\":{\"arg.properties\":{\"options\":[\"-\"]},\"type\":\"string\"}},{\"name\":\"time\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1,\"step\":10}},\"format_as_time\":\"dd/MMM/yyyy:HH:mm:ss Z\",\"type\":\"string\"}},{\"name\":\"_logtime\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1,\"step\":10}},\"format_as_time\":\"unix_long\",\"type\":\"long\"}},{\"name\":\"request\",\"type\":{\"arg.properties\":{\"options\":[\"GET /index.html HTTP/1.1\",\"GET /site/user_status.html HTTP/1.1\",\"GET /site/login.html HTTP/1.1\",\"GET /site/user_status.html HTTP/1.1\",\"GET /images/track.png HTTP/1.1\",\"GET /images/logo-small.png HTTP/1.1\"]},\"type\":\"string\"}},{\"name\":\"status\",\"type\":{\"arg.properties\":{\"options\":[\"200\",\"302\",\"404\",\"405\",\"406\",\"407\"]},\"type\":\"string\"}},{\"name\":\"bytes\",\"type\":{\"arg.properties\":{\"options\":[\"278\",\"1289\",\"2048\",\"4096\",\"4006\",\"4196\",\"14096\"]},\"type\":\"string\"}},{\"name\":\"referrer\",\"type\":{\"arg.properties\":{\"options\":[\"-\"]},\"type\":\"string\"}},{\"name\":\"agent\",\"type\":{\"arg.properties\":{\"options\":[\"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)\",\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36\"]},\"type\":\"string\"}}],\"name\":\"clickstream.clickstream\",\"type\":\"record\"}"
}

func (r Clickstream) SchemaName() string {
	return "clickstream.clickstream"
}

func (_ Clickstream) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Clickstream) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Clickstream) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Clickstream) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Clickstream) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Clickstream) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Clickstream) SetString(v string)   { panic("Unsupported operation") }
func (_ Clickstream) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Clickstream) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Ip}

		return w

	case 1:
		w := types.Int{Target: &r.Userid}

		return w

	case 2:
		w := types.String{Target: &r.Remote_user}

		return w

	case 3:
		w := types.String{Target: &r.Time}

		return w

	case 4:
		w := types.Long{Target: &r.Logtime}

		return w

	case 5:
		w := types.String{Target: &r.Request}

		return w

	case 6:
		w := types.String{Target: &r.Status}

		return w

	case 7:
		w := types.String{Target: &r.Bytes}

		return w

	case 8:
		w := types.String{Target: &r.Referrer}

		return w

	case 9:
		w := types.String{Target: &r.Agent}

		return w

	}
	panic("Unknown field index")
}

func (r *Clickstream) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Clickstream) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Clickstream) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Clickstream) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Clickstream) HintSize(int)                     { panic("Unsupported operation") }
func (_ Clickstream) Finalize()                        {}

func (_ Clickstream) AvroCRC64Fingerprint() []byte {
	return []byte(ClickstreamAvroCRC64Fingerprint)
}

func (r Clickstream) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ip"], err = json.Marshal(r.Ip)
	if err != nil {
		return nil, err
	}
	output["userid"], err = json.Marshal(r.Userid)
	if err != nil {
		return nil, err
	}
	output["remote_user"], err = json.Marshal(r.Remote_user)
	if err != nil {
		return nil, err
	}
	output["time"], err = json.Marshal(r.Time)
	if err != nil {
		return nil, err
	}
	output["_logtime"], err = json.Marshal(r.Logtime)
	if err != nil {
		return nil, err
	}
	output["request"], err = json.Marshal(r.Request)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	output["bytes"], err = json.Marshal(r.Bytes)
	if err != nil {
		return nil, err
	}
	output["referrer"], err = json.Marshal(r.Referrer)
	if err != nil {
		return nil, err
	}
	output["agent"], err = json.Marshal(r.Agent)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Clickstream) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ip"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ip); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ip")
	}
	val = func() json.RawMessage {
		if v, ok := fields["userid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Userid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for userid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["remote_user"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remote_user); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for remote_user")
	}
	val = func() json.RawMessage {
		if v, ok := fields["time"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Time); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for time")
	}
	val = func() json.RawMessage {
		if v, ok := fields["_logtime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Logtime); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for _logtime")
	}
	val = func() json.RawMessage {
		if v, ok := fields["request"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Request); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for request")
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
	val = func() json.RawMessage {
		if v, ok := fields["bytes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bytes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for bytes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["referrer"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Referrer); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for referrer")
	}
	val = func() json.RawMessage {
		if v, ok := fields["agent"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Agent); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for agent")
	}
	return nil
}
