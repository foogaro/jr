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

type Siemlogs struct {
	Hostname string `json:"hostname"`

	Action string `json:"action"`

	L4 string `json:"l4"`

	Access_group string `json:"access_group"`

	Source Source `json:"source"`

	Destination Destination `json:"destination"`
}

const SiemlogsAvroCRC64Fingerprint = "D\x02\x006\x13\xd5^-"

func NewSiemlogs() Siemlogs {
	r := Siemlogs{}
	r.Source = NewSource()

	r.Destination = NewDestination()

	return r
}

func DeserializeSiemlogs(r io.Reader) (Siemlogs, error) {
	t := NewSiemlogs()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSiemlogsFromSchema(r io.Reader, schema string) (Siemlogs, error) {
	t := NewSiemlogs()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSiemlogs(r Siemlogs, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Hostname, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Action, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.L4, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Access_group, w)
	if err != nil {
		return err
	}
	err = writeSource(r.Source, w)
	if err != nil {
		return err
	}
	err = writeDestination(r.Destination, w)
	if err != nil {
		return err
	}
	return err
}

func (r Siemlogs) Serialize(w io.Writer) error {
	return writeSiemlogs(r, w)
}

func (r Siemlogs) Schema() string {
	return "{\"fields\":[{\"name\":\"hostname\",\"type\":{\"arg.properties\":{\"options\":[\"204.149.41.63\",\"54.222.176.193\",\"122.120.109.66\",\"59.145.93.29\",\"222.85.75.48\",\"68.70.120.23\",\"8.243.217.97\",\"47.82.29.180\",\"182.23.41.165\",\"230.195.181.37\",\"45.206.35.54\",\"59.222.135.27\",\"178.74.178.202\",\"112.10.11.67\",\"231.62.195.226\",\"56.25.171.14\",\"72.206.143.236\",\"101.162.21.103\",\"98.24.15.221\",\"109.151.30.96\",\"67.143.145.10\",\"161.110.198.55\",\"56.186.158.4\",\"215.63.183.99\",\"173.218.229.93\",\"119.0.185.169\",\"13.206.166.235\",\"17.208.177.68\",\"230.234.177.180\",\"177.83.69.26\",\"109.78.222.30\",\"232.59.60.147\",\"250.173.99.68\",\"4.75.226.124\",\"174.185.97.9\",\"214.15.142.138\",\"84.46.49.248\",\"189.150.30.47\",\"226.178.139.46\",\"167.153.109.224\",\"126.147.106.47\",\"157.200.87.61\",\"131.255.186.13\",\"152.102.102.68\",\"214.177.4.5\",\"161.166.90.144\",\"98.160.215.50\",\"210.74.249.137\",\"244.64.29.244\",\"68.227.6.52\",\"141.228.33.77\",\"33.207.45.196\",\"60.35.239.48\",\"25.88.68.126\",\"239.126.133.195\",\"173.13.57.55\",\"134.23.137.5\",\"120.48.50.185\",\"215.90.55.197\",\"23.97.205.108\",\"117.250.191.255\",\"165.188.33.225\",\"153.41.174.133\",\"42.252.99.100\",\"12.42.99.54\",\"237.83.9.240\",\"137.148.240.223\",\"113.9.0.63\",\"73.126.21.73\",\"70.109.118.28\",\"121.166.215.228\",\"15.128.133.129\",\"156.64.22.26\",\"74.31.184.80\",\"32.110.165.153\",\"3.84.77.181\",\"106.45.159.209\",\"49.33.192.133\",\"66.95.63.195\",\"242.252.182.130\",\"245.91.183.42\",\"120.121.18.230\",\"144.81.168.91\",\"206.219.105.117\",\"225.242.249.169\",\"66.153.168.58\",\"160.199.67.92\",\"53.124.103.36\",\"94.51.180.226\",\"60.179.124.207\",\"4.126.94.49\",\"127.136.244.28\",\"171.169.227.52\",\"137.204.219.247\",\"20.195.159.248\",\"97.217.72.152\",\"152.30.133.188\",\"173.79.61.56\",\"173.119.145.171\",\"149.135.212.84\"]},\"type\":\"string\"}},{\"name\":\"action\",\"type\":{\"arg.properties\":{\"options\":[\"allow\",\"deny\"]},\"type\":\"string\"}},{\"name\":\"l4\",\"type\":{\"arg.properties\":{\"options\":[\"tcp\"]},\"type\":\"string\"}},{\"name\":\"access_group\",\"type\":{\"arg.properties\":{\"options\":[\"group-1\",\"group-2\",\"group-3\",\"group-4\",\"admin\"]},\"type\":\"string\"}},{\"name\":\"source\",\"type\":{\"fields\":[{\"name\":\"ip\",\"type\":{\"arg.properties\":{\"options\":[\"204.149.41.63\",\"54.222.176.193\",\"122.120.109.66\",\"59.145.93.29\",\"222.85.75.48\",\"68.70.120.23\",\"8.243.217.97\",\"47.82.29.180\",\"182.23.41.165\",\"230.195.181.37\",\"45.206.35.54\",\"59.222.135.27\",\"178.74.178.202\",\"112.10.11.67\",\"231.62.195.226\",\"56.25.171.14\",\"72.206.143.236\",\"101.162.21.103\",\"98.24.15.221\",\"109.151.30.96\",\"67.143.145.10\",\"161.110.198.55\",\"56.186.158.4\",\"215.63.183.99\",\"173.218.229.93\",\"119.0.185.169\",\"13.206.166.235\",\"17.208.177.68\",\"230.234.177.180\",\"177.83.69.26\",\"109.78.222.30\",\"232.59.60.147\",\"250.173.99.68\",\"4.75.226.124\",\"174.185.97.9\",\"214.15.142.138\",\"84.46.49.248\",\"189.150.30.47\",\"226.178.139.46\",\"167.153.109.224\",\"126.147.106.47\",\"157.200.87.61\",\"131.255.186.13\",\"152.102.102.68\",\"214.177.4.5\",\"161.166.90.144\",\"98.160.215.50\",\"210.74.249.137\",\"244.64.29.244\",\"68.227.6.52\",\"141.228.33.77\",\"33.207.45.196\",\"60.35.239.48\",\"25.88.68.126\",\"239.126.133.195\",\"173.13.57.55\",\"134.23.137.5\",\"120.48.50.185\",\"215.90.55.197\",\"23.97.205.108\",\"117.250.191.255\",\"165.188.33.225\",\"153.41.174.133\",\"42.252.99.100\",\"12.42.99.54\",\"237.83.9.240\",\"137.148.240.223\",\"113.9.0.63\",\"73.126.21.73\",\"70.109.118.28\",\"121.166.215.228\",\"15.128.133.129\",\"156.64.22.26\",\"74.31.184.80\",\"32.110.165.153\",\"3.84.77.181\",\"106.45.159.209\",\"49.33.192.133\",\"66.95.63.195\",\"242.252.182.130\",\"245.91.183.42\",\"120.121.18.230\",\"144.81.168.91\",\"206.219.105.117\",\"225.242.249.169\",\"66.153.168.58\",\"160.199.67.92\",\"53.124.103.36\",\"94.51.180.226\",\"60.179.124.207\",\"4.126.94.49\",\"127.136.244.28\",\"171.169.227.52\",\"137.204.219.247\",\"20.195.159.248\",\"97.217.72.152\",\"152.30.133.188\",\"173.79.61.56\",\"173.119.145.171\",\"149.135.212.84\"]},\"type\":\"string\"}},{\"name\":\"port\",\"type\":{\"arg.properties\":{\"range\":{\"max\":99999,\"min\":10000}},\"type\":\"int\"}}],\"name\":\"source\",\"type\":\"record\"}},{\"name\":\"destination\",\"type\":{\"fields\":[{\"name\":\"ip\",\"type\":{\"arg.properties\":{\"options\":[\"204.149.41.63\",\"54.222.176.193\",\"122.120.109.66\",\"59.145.93.29\",\"222.85.75.48\",\"68.70.120.23\",\"8.243.217.97\",\"47.82.29.180\",\"182.23.41.165\",\"230.195.181.37\",\"45.206.35.54\",\"59.222.135.27\",\"178.74.178.202\",\"112.10.11.67\",\"231.62.195.226\",\"56.25.171.14\",\"72.206.143.236\",\"101.162.21.103\",\"98.24.15.221\",\"109.151.30.96\",\"67.143.145.10\",\"161.110.198.55\",\"56.186.158.4\",\"215.63.183.99\",\"173.218.229.93\",\"119.0.185.169\",\"13.206.166.235\",\"17.208.177.68\",\"230.234.177.180\",\"177.83.69.26\",\"109.78.222.30\",\"232.59.60.147\",\"250.173.99.68\",\"4.75.226.124\",\"174.185.97.9\",\"214.15.142.138\",\"84.46.49.248\",\"189.150.30.47\",\"226.178.139.46\",\"167.153.109.224\",\"126.147.106.47\",\"157.200.87.61\",\"131.255.186.13\",\"152.102.102.68\",\"214.177.4.5\",\"161.166.90.144\",\"98.160.215.50\",\"210.74.249.137\",\"244.64.29.244\",\"68.227.6.52\",\"141.228.33.77\",\"33.207.45.196\",\"60.35.239.48\",\"25.88.68.126\",\"239.126.133.195\",\"173.13.57.55\",\"134.23.137.5\",\"120.48.50.185\",\"215.90.55.197\",\"23.97.205.108\",\"117.250.191.255\",\"165.188.33.225\",\"153.41.174.133\",\"42.252.99.100\",\"12.42.99.54\",\"237.83.9.240\",\"137.148.240.223\",\"113.9.0.63\",\"73.126.21.73\",\"70.109.118.28\",\"121.166.215.228\",\"15.128.133.129\",\"156.64.22.26\",\"74.31.184.80\",\"32.110.165.153\",\"3.84.77.181\",\"106.45.159.209\",\"49.33.192.133\",\"66.95.63.195\",\"242.252.182.130\",\"245.91.183.42\",\"120.121.18.230\",\"144.81.168.91\",\"206.219.105.117\",\"225.242.249.169\",\"66.153.168.58\",\"160.199.67.92\",\"53.124.103.36\",\"94.51.180.226\",\"60.179.124.207\",\"4.126.94.49\",\"127.136.244.28\",\"171.169.227.52\",\"137.204.219.247\",\"20.195.159.248\",\"97.217.72.152\",\"152.30.133.188\",\"173.79.61.56\",\"173.119.145.171\",\"149.135.212.84\"]},\"type\":\"string\"}},{\"name\":\"port\",\"type\":{\"arg.properties\":{\"range\":{\"max\":99999,\"min\":10000}},\"type\":\"int\"}}],\"name\":\"destination\",\"type\":\"record\"}}],\"name\":\"siem_logs.siemlogs\",\"type\":\"record\"}"
}

func (r Siemlogs) SchemaName() string {
	return "siem_logs.siemlogs"
}

func (_ Siemlogs) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Siemlogs) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Siemlogs) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Siemlogs) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Siemlogs) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Siemlogs) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Siemlogs) SetString(v string)   { panic("Unsupported operation") }
func (_ Siemlogs) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Siemlogs) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Hostname}

		return w

	case 1:
		w := types.String{Target: &r.Action}

		return w

	case 2:
		w := types.String{Target: &r.L4}

		return w

	case 3:
		w := types.String{Target: &r.Access_group}

		return w

	case 4:
		r.Source = NewSource()

		w := types.Record{Target: &r.Source}

		return w

	case 5:
		r.Destination = NewDestination()

		w := types.Record{Target: &r.Destination}

		return w

	}
	panic("Unknown field index")
}

func (r *Siemlogs) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Siemlogs) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Siemlogs) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Siemlogs) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Siemlogs) HintSize(int)                     { panic("Unsupported operation") }
func (_ Siemlogs) Finalize()                        {}

func (_ Siemlogs) AvroCRC64Fingerprint() []byte {
	return []byte(SiemlogsAvroCRC64Fingerprint)
}

func (r Siemlogs) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["hostname"], err = json.Marshal(r.Hostname)
	if err != nil {
		return nil, err
	}
	output["action"], err = json.Marshal(r.Action)
	if err != nil {
		return nil, err
	}
	output["l4"], err = json.Marshal(r.L4)
	if err != nil {
		return nil, err
	}
	output["access_group"], err = json.Marshal(r.Access_group)
	if err != nil {
		return nil, err
	}
	output["source"], err = json.Marshal(r.Source)
	if err != nil {
		return nil, err
	}
	output["destination"], err = json.Marshal(r.Destination)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Siemlogs) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["hostname"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Hostname); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for hostname")
	}
	val = func() json.RawMessage {
		if v, ok := fields["action"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Action); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for action")
	}
	val = func() json.RawMessage {
		if v, ok := fields["l4"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.L4); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for l4")
	}
	val = func() json.RawMessage {
		if v, ok := fields["access_group"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Access_group); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for access_group")
	}
	val = func() json.RawMessage {
		if v, ok := fields["source"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Source); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for source")
	}
	val = func() json.RawMessage {
		if v, ok := fields["destination"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destination); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destination")
	}
	return nil
}
