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

type Payrollemployee struct {
	Employee_id int32 `json:"employee_id"`

	First_name string `json:"first_name"`

	Last_name string `json:"last_name"`

	Age int32 `json:"age"`

	Ssn string `json:"ssn"`

	Hourly_rate int32 `json:"hourly_rate"`

	Gender string `json:"gender"`

	Email string `json:"email"`
}

const PayrollemployeeAvroCRC64Fingerprint = "\xeb\x9b\xd6.\xbe\xec2\xce"

func NewPayrollemployee() Payrollemployee {
	r := Payrollemployee{}
	return r
}

func DeserializePayrollemployee(r io.Reader) (Payrollemployee, error) {
	t := NewPayrollemployee()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePayrollemployeeFromSchema(r io.Reader, schema string) (Payrollemployee, error) {
	t := NewPayrollemployee()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePayrollemployee(r Payrollemployee, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Employee_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.First_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Last_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Age, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Ssn, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Hourly_rate, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Gender, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	return err
}

func (r Payrollemployee) Serialize(w io.Writer) error {
	return writePayrollemployee(r, w)
}

func (r Payrollemployee) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"age\":32,\"email\":\"xcqn@mycompany.com\",\"employee_id\":1000,\"first_name\":\"Carleton\",\"gender\":\"female\",\"hourly_rate\":21,\"last_name\":\"Beagan\",\"ssn\":\"723-37-6885\"},{\"age\":42,\"email\":\"ymcw@mycompany.com\",\"employee_id\":1001,\"first_name\":\"Frieda\",\"gender\":\"male\",\"hourly_rate\":20,\"last_name\":\"Starcks\",\"ssn\":\"992-86-7676\"},{\"age\":38,\"email\":\"zkth@mycompany.com\",\"employee_id\":1002,\"first_name\":\"Trumann\",\"gender\":\"male\",\"hourly_rate\":17,\"last_name\":\"Murfill\",\"ssn\":\"980-75-4291\"},{\"age\":46,\"email\":\"yfxw@mycompany.com\",\"employee_id\":1003,\"first_name\":\"Zeb\",\"gender\":\"male\",\"hourly_rate\":11,\"last_name\":\"Murrish\",\"ssn\":\"557-94-0484\"},{\"age\":49,\"email\":\"ireg@mycompany.com\",\"employee_id\":1004,\"first_name\":\"Hayley\",\"gender\":\"female\",\"hourly_rate\":16,\"last_name\":\"Beatson\",\"ssn\":\"291-87-0461\"},{\"age\":47,\"email\":\"kzvq@mycompany.com\",\"employee_id\":1005,\"first_name\":\"Nolan\",\"gender\":\"female\",\"hourly_rate\":15,\"last_name\":\"Hallowell\",\"ssn\":\"798-35-7786\"},{\"age\":46,\"email\":\"tauo@mycompany.com\",\"employee_id\":1006,\"first_name\":\"Frieda\",\"gender\":\"male\",\"hourly_rate\":22,\"last_name\":\"Joblin\",\"ssn\":\"124-52-6784\"},{\"age\":46,\"email\":\"gexc@mycompany.com\",\"employee_id\":1007,\"first_name\":\"Kalila\",\"gender\":\"female\",\"hourly_rate\":8,\"last_name\":\"Foulsham\",\"ssn\":\"287-30-7370\"},{\"age\":38,\"email\":\"ryny@mycompany.com\",\"employee_id\":1008,\"first_name\":\"Nolan\",\"gender\":\"female\",\"hourly_rate\":13,\"last_name\":\"Rosensaft\",\"ssn\":\"666-02-9029\"},{\"age\":33,\"email\":\"haad@mycompany.com\",\"employee_id\":1009,\"first_name\":\"Isadora\",\"gender\":\"female\",\"hourly_rate\":23,\"last_name\":\"Ravenscroft\",\"ssn\":\"524-83-3550\"},{\"age\":39,\"email\":\"hgxz@mycompany.com\",\"employee_id\":1010,\"first_name\":\"Byrann\",\"gender\":\"female\",\"hourly_rate\":13,\"last_name\":\"Zuker\",\"ssn\":\"608-15-9492\"},{\"age\":46,\"email\":\"bnhl@mycompany.com\",\"employee_id\":1011,\"first_name\":\"Egon\",\"gender\":\"female\",\"hourly_rate\":24,\"last_name\":\"Bettles\",\"ssn\":\"455-48-4428\"},{\"age\":35,\"email\":\"yxpu@mycompany.com\",\"employee_id\":1012,\"first_name\":\"Faye\",\"gender\":\"female\",\"hourly_rate\":17,\"last_name\":\"Mcall\",\"ssn\":\"967-90-4356\"},{\"age\":43,\"email\":\"yofv@mycompany.com\",\"employee_id\":1013,\"first_name\":\"Lucais\",\"gender\":\"male\",\"hourly_rate\":24,\"last_name\":\"Mears\",\"ssn\":\"378-20-1991\"},{\"age\":41,\"email\":\"alln@mycompany.com\",\"employee_id\":1014,\"first_name\":\"Frieda\",\"gender\":\"female\",\"hourly_rate\":22,\"last_name\":\"Rocco\",\"ssn\":\"401-12-7473\"},{\"age\":34,\"email\":\"feym@mycompany.com\",\"employee_id\":1015,\"first_name\":\"Marchall\",\"gender\":\"male\",\"hourly_rate\":20,\"last_name\":\"Ottiwill\",\"ssn\":\"989-47-6666\"},{\"age\":36,\"email\":\"uflo@mycompany.com\",\"employee_id\":1016,\"first_name\":\"Raychel\",\"gender\":\"male\",\"hourly_rate\":17,\"last_name\":\"Cleugher\",\"ssn\":\"859-19-5275\"},{\"age\":39,\"email\":\"nkkb@mycompany.com\",\"employee_id\":1017,\"first_name\":\"Lea\",\"gender\":\"female\",\"hourly_rate\":18,\"last_name\":\"Hobben\",\"ssn\":\"848-02-9978\"},{\"age\":48,\"email\":\"ecec@mycompany.com\",\"employee_id\":1018,\"first_name\":\"Sonnie\",\"gender\":\"male\",\"hourly_rate\":21,\"last_name\":\"Woolfall\",\"ssn\":\"211-34-6636\"},{\"age\":31,\"email\":\"apun@mycompany.com\",\"employee_id\":1019,\"first_name\":\"Derk\",\"gender\":\"female\",\"hourly_rate\":10,\"last_name\":\"Labba\",\"ssn\":\"628-39-7556\"},{\"age\":44,\"email\":\"nyse@mycompany.com\",\"employee_id\":1020,\"first_name\":\"Johanna\",\"gender\":\"female\",\"hourly_rate\":22,\"last_name\":\"Wren\",\"ssn\":\"544-75-7207\"},{\"age\":38,\"email\":\"ygrg@mycompany.com\",\"employee_id\":1021,\"first_name\":\"Barnett\",\"gender\":\"male\",\"hourly_rate\":20,\"last_name\":\"Graveson\",\"ssn\":\"238-81-9399\"},{\"age\":34,\"email\":\"gidb@mycompany.com\",\"employee_id\":1022,\"first_name\":\"Trent\",\"gender\":\"male\",\"hourly_rate\":22,\"last_name\":\"Broadbridge\",\"ssn\":\"928-07-0451\"},{\"age\":47,\"email\":\"mkoz@mycompany.com\",\"employee_id\":1023,\"first_name\":\"Cyril\",\"gender\":\"female\",\"hourly_rate\":13,\"last_name\":\"Rylatt\",\"ssn\":\"100-38-8723\"},{\"age\":44,\"email\":\"yezj@mycompany.com\",\"employee_id\":1024,\"first_name\":\"Johanna\",\"gender\":\"male\",\"hourly_rate\":23,\"last_name\":\"Murrish\",\"ssn\":\"612-61-9477\"},{\"age\":39,\"email\":\"govs@mycompany.com\",\"employee_id\":1025,\"first_name\":\"Byrann\",\"gender\":\"female\",\"hourly_rate\":18,\"last_name\":\"Woolfall\",\"ssn\":\"889-10-6161\"},{\"age\":36,\"email\":\"cruw@mycompany.com\",\"employee_id\":1026,\"first_name\":\"Derk\",\"gender\":\"male\",\"hourly_rate\":24,\"last_name\":\"Baldi\",\"ssn\":\"929-39-7920\"},{\"age\":37,\"email\":\"mwic@mycompany.com\",\"employee_id\":1027,\"first_name\":\"Matt\",\"gender\":\"male\",\"hourly_rate\":10,\"last_name\":\"Vockins\",\"ssn\":\"174-48-7138\"},{\"age\":46,\"email\":\"vklt@mycompany.com\",\"employee_id\":1028,\"first_name\":\"Faye\",\"gender\":\"male\",\"hourly_rate\":17,\"last_name\":\"Barock\",\"ssn\":\"495-35-5641\"},{\"age\":41,\"email\":\"rbvg@mycompany.com\",\"employee_id\":1029,\"first_name\":\"Carleton\",\"gender\":\"male\",\"hourly_rate\":18,\"last_name\":\"Ridgley\",\"ssn\":\"172-35-4720\"},{\"age\":41,\"email\":\"mtnj@mycompany.com\",\"employee_id\":1030,\"first_name\":\"Derk\",\"gender\":\"female\",\"hourly_rate\":10,\"last_name\":\"Ridgley\",\"ssn\":\"317-43-5426\"},{\"age\":33,\"email\":\"ojkl@mycompany.com\",\"employee_id\":1031,\"first_name\":\"Trent\",\"gender\":\"female\",\"hourly_rate\":8,\"last_name\":\"Milazzo\",\"ssn\":\"435-51-6785\"},{\"age\":49,\"email\":\"zasw@mycompany.com\",\"employee_id\":1032,\"first_name\":\"Lucais\",\"gender\":\"male\",\"hourly_rate\":12,\"last_name\":\"Haskell\",\"ssn\":\"598-73-8837\"},{\"age\":32,\"email\":\"jyum@mycompany.com\",\"employee_id\":1033,\"first_name\":\"Avis\",\"gender\":\"female\",\"hourly_rate\":20,\"last_name\":\"Wren\",\"ssn\":\"215-96-0693\"},{\"age\":38,\"email\":\"gbce@mycompany.com\",\"employee_id\":1034,\"first_name\":\"Sibyl\",\"gender\":\"female\",\"hourly_rate\":11,\"last_name\":\"Guildford\",\"ssn\":\"811-14-9013\"},{\"age\":34,\"email\":\"bcgy@mycompany.com\",\"employee_id\":1035,\"first_name\":\"Raychel\",\"gender\":\"female\",\"hourly_rate\":19,\"last_name\":\"MacVay\",\"ssn\":\"991-62-5231\"},{\"age\":32,\"email\":\"fyzt@mycompany.com\",\"employee_id\":1036,\"first_name\":\"Winslow\",\"gender\":\"female\",\"hourly_rate\":19,\"last_name\":\"Beagan\",\"ssn\":\"884-20-3190\"},{\"age\":36,\"email\":\"erpx@mycompany.com\",\"employee_id\":1037,\"first_name\":\"Robyn\",\"gender\":\"female\",\"hourly_rate\":15,\"last_name\":\"Westhofer\",\"ssn\":\"433-18-5627\"},{\"age\":48,\"email\":\"dshv@mycompany.com\",\"employee_id\":1038,\"first_name\":\"Lurline\",\"gender\":\"female\",\"hourly_rate\":21,\"last_name\":\"Woolfall\",\"ssn\":\"153-88-4840\"},{\"age\":38,\"email\":\"vahe@mycompany.com\",\"employee_id\":1039,\"first_name\":\"Electra\",\"gender\":\"male\",\"hourly_rate\":15,\"last_name\":\"MacVay\",\"ssn\":\"906-28-8038\"},{\"age\":45,\"email\":\"grhs@mycompany.com\",\"employee_id\":1040,\"first_name\":\"Babara\",\"gender\":\"female\",\"hourly_rate\":18,\"last_name\":\"Rosensaft\",\"ssn\":\"732-77-6755\"},{\"age\":42,\"email\":\"rven@mycompany.com\",\"employee_id\":1041,\"first_name\":\"Gerda\",\"gender\":\"female\",\"hourly_rate\":9,\"last_name\":\"Cleugher\",\"ssn\":\"904-94-1504\"},{\"age\":42,\"email\":\"ynxr@mycompany.com\",\"employee_id\":1042,\"first_name\":\"Charley\",\"gender\":\"female\",\"hourly_rate\":17,\"last_name\":\"Lezemere\",\"ssn\":\"869-37-8441\"},{\"age\":49,\"email\":\"ocpi@mycompany.com\",\"employee_id\":1043,\"first_name\":\"Dur\",\"gender\":\"female\",\"hourly_rate\":18,\"last_name\":\"Klemt\",\"ssn\":\"091-49-9989\"},{\"age\":47,\"email\":\"xiyy@mycompany.com\",\"employee_id\":1044,\"first_name\":\"Sonnie\",\"gender\":\"male\",\"hourly_rate\":22,\"last_name\":\"Lezemere\",\"ssn\":\"458-96-3638\"},{\"age\":38,\"email\":\"jcex@mycompany.com\",\"employee_id\":1045,\"first_name\":\"Quint\",\"gender\":\"male\",\"hourly_rate\":11,\"last_name\":\"Beckwith\",\"ssn\":\"262-57-5599\"},{\"age\":35,\"email\":\"mvcx@mycompany.com\",\"employee_id\":1046,\"first_name\":\"Electra\",\"gender\":\"male\",\"hourly_rate\":11,\"last_name\":\"Bidgod\",\"ssn\":\"753-51-8847\"},{\"age\":37,\"email\":\"dlow@mycompany.com\",\"employee_id\":1047,\"first_name\":\"Frieda\",\"gender\":\"male\",\"hourly_rate\":15,\"last_name\":\"Dudill\",\"ssn\":\"364-67-4057\"},{\"age\":44,\"email\":\"uovu@mycompany.com\",\"employee_id\":1048,\"first_name\":\"Aryn\",\"gender\":\"male\",\"hourly_rate\":21,\"last_name\":\"Rocco\",\"ssn\":\"730-69-3872\"},{\"age\":44,\"email\":\"vqyh@mycompany.com\",\"employee_id\":1049,\"first_name\":\"Frieda\",\"gender\":\"male\",\"hourly_rate\":12,\"last_name\":\"Tuma\",\"ssn\":\"696-95-5025\"},{\"age\":31,\"email\":\"voel@mycompany.com\",\"employee_id\":1050,\"first_name\":\"Trev\",\"gender\":\"male\",\"hourly_rate\":22,\"last_name\":\"Napton\",\"ssn\":\"770-11-4815\"},{\"age\":43,\"email\":\"ugiz@mycompany.com\",\"employee_id\":1051,\"first_name\":\"Jere\",\"gender\":\"female\",\"hourly_rate\":9,\"last_name\":\"Rosensaft\",\"ssn\":\"593-18-2315\"},{\"age\":39,\"email\":\"lcfy@mycompany.com\",\"employee_id\":1052,\"first_name\":\"Sibyl\",\"gender\":\"female\",\"hourly_rate\":20,\"last_name\":\"Nayshe\",\"ssn\":\"229-81-4143\"},{\"age\":42,\"email\":\"hspi@mycompany.com\",\"employee_id\":1053,\"first_name\":\"Talbot\",\"gender\":\"male\",\"hourly_rate\":14,\"last_name\":\"Coll\",\"ssn\":\"397-54-1060\"},{\"age\":48,\"email\":\"dxwi@mycompany.com\",\"employee_id\":1054,\"first_name\":\"Adele\",\"gender\":\"male\",\"hourly_rate\":13,\"last_name\":\"Zuker\",\"ssn\":\"053-08-6877\"},{\"age\":48,\"email\":\"ovoz@mycompany.com\",\"employee_id\":1055,\"first_name\":\"Gustave\",\"gender\":\"male\",\"hourly_rate\":12,\"last_name\":\"Lezemere\",\"ssn\":\"632-36-0913\"},{\"age\":32,\"email\":\"cxrr@mycompany.com\",\"employee_id\":1056,\"first_name\":\"Cherrita\",\"gender\":\"male\",\"hourly_rate\":9,\"last_name\":\"Lezemere\",\"ssn\":\"920-73-3034\"},{\"age\":34,\"email\":\"ymov@mycompany.com\",\"employee_id\":1057,\"first_name\":\"Dur\",\"gender\":\"female\",\"hourly_rate\":22,\"last_name\":\"Beaument\",\"ssn\":\"367-26-8343\"},{\"age\":49,\"email\":\"atfe@mycompany.com\",\"employee_id\":1058,\"first_name\":\"Livvie\",\"gender\":\"female\",\"hourly_rate\":16,\"last_name\":\"Smewings\",\"ssn\":\"300-99-5551\"},{\"age\":49,\"email\":\"dwxm@mycompany.com\",\"employee_id\":1059,\"first_name\":\"Karena\",\"gender\":\"male\",\"hourly_rate\":23,\"last_name\":\"Bainbridge\",\"ssn\":\"960-91-6169\"},{\"age\":30,\"email\":\"qvfv@mycompany.com\",\"employee_id\":1060,\"first_name\":\"Trev\",\"gender\":\"female\",\"hourly_rate\":17,\"last_name\":\"Beaument\",\"ssn\":\"172-28-3691\"},{\"age\":48,\"email\":\"fpaz@mycompany.com\",\"employee_id\":1061,\"first_name\":\"Avis\",\"gender\":\"female\",\"hourly_rate\":9,\"last_name\":\"Borge\",\"ssn\":\"476-07-4335\"},{\"age\":31,\"email\":\"qzkg@mycompany.com\",\"employee_id\":1062,\"first_name\":\"Vergil\",\"gender\":\"male\",\"hourly_rate\":20,\"last_name\":\"Foulsham\",\"ssn\":\"713-65-6480\"},{\"age\":37,\"email\":\"ezdw@mycompany.com\",\"employee_id\":1063,\"first_name\":\"Robyn\",\"gender\":\"female\",\"hourly_rate\":21,\"last_name\":\"Haskell\",\"ssn\":\"209-24-6869\"},{\"age\":37,\"email\":\"ulct@mycompany.com\",\"employee_id\":1064,\"first_name\":\"Pattin\",\"gender\":\"male\",\"hourly_rate\":17,\"last_name\":\"Moyler\",\"ssn\":\"023-00-7370\"},{\"age\":40,\"email\":\"sbwf@mycompany.com\",\"employee_id\":1065,\"first_name\":\"Sonnie\",\"gender\":\"male\",\"hourly_rate\":11,\"last_name\":\"Tenbrug\",\"ssn\":\"739-69-6275\"},{\"age\":31,\"email\":\"juba@mycompany.com\",\"employee_id\":1066,\"first_name\":\"Talbot\",\"gender\":\"male\",\"hourly_rate\":20,\"last_name\":\"Milazzo\",\"ssn\":\"080-46-2275\"},{\"age\":31,\"email\":\"hcln@mycompany.com\",\"employee_id\":1067,\"first_name\":\"Trumann\",\"gender\":\"female\",\"hourly_rate\":12,\"last_name\":\"Coll\",\"ssn\":\"343-16-9551\"},{\"age\":37,\"email\":\"iwpx@mycompany.com\",\"employee_id\":1068,\"first_name\":\"Andriette\",\"gender\":\"male\",\"hourly_rate\":12,\"last_name\":\"Starcks\",\"ssn\":\"658-51-1241\"},{\"age\":39,\"email\":\"jfik@mycompany.com\",\"employee_id\":1069,\"first_name\":\"Lea\",\"gender\":\"male\",\"hourly_rate\":18,\"last_name\":\"Hobben\",\"ssn\":\"842-28-9022\"},{\"age\":47,\"email\":\"sqyg@mycompany.com\",\"employee_id\":1070,\"first_name\":\"Barnett\",\"gender\":\"female\",\"hourly_rate\":10,\"last_name\":\"Bidgod\",\"ssn\":\"860-14-7110\"},{\"age\":38,\"email\":\"yqkb@mycompany.com\",\"employee_id\":1071,\"first_name\":\"West\",\"gender\":\"female\",\"hourly_rate\":8,\"last_name\":\"Milazzo\",\"ssn\":\"449-28-4678\"},{\"age\":41,\"email\":\"oowf@mycompany.com\",\"employee_id\":1072,\"first_name\":\"Avis\",\"gender\":\"female\",\"hourly_rate\":12,\"last_name\":\"Mears\",\"ssn\":\"164-80-2955\"},{\"age\":41,\"email\":\"dmmi@mycompany.com\",\"employee_id\":1073,\"first_name\":\"Robyn\",\"gender\":\"male\",\"hourly_rate\":15,\"last_name\":\"Rylatt\",\"ssn\":\"062-56-4978\"},{\"age\":30,\"email\":\"ozkd@mycompany.com\",\"employee_id\":1074,\"first_name\":\"Giacobo\",\"gender\":\"male\",\"hourly_rate\":23,\"last_name\":\"Mayhow\",\"ssn\":\"960-65-6801\"},{\"age\":33,\"email\":\"nhlm@mycompany.com\",\"employee_id\":1075,\"first_name\":\"Giacobo\",\"gender\":\"female\",\"hourly_rate\":22,\"last_name\":\"MacVay\",\"ssn\":\"801-57-0746\"},{\"age\":41,\"email\":\"mzfb@mycompany.com\",\"employee_id\":1076,\"first_name\":\"Germayne\",\"gender\":\"female\",\"hourly_rate\":10,\"last_name\":\"Bainbridge\",\"ssn\":\"226-28-2222\"},{\"age\":43,\"email\":\"jbpl@mycompany.com\",\"employee_id\":1077,\"first_name\":\"Alis\",\"gender\":\"female\",\"hourly_rate\":24,\"last_name\":\"Vockins\",\"ssn\":\"997-17-4419\"},{\"age\":49,\"email\":\"pgwd@mycompany.com\",\"employee_id\":1078,\"first_name\":\"Caria\",\"gender\":\"male\",\"hourly_rate\":18,\"last_name\":\"Hobben\",\"ssn\":\"438-17-3464\"},{\"age\":38,\"email\":\"eska@mycompany.com\",\"employee_id\":1079,\"first_name\":\"Talbot\",\"gender\":\"female\",\"hourly_rate\":14,\"last_name\":\"Eringey\",\"ssn\":\"416-10-1590\"},{\"age\":43,\"email\":\"tbeg@mycompany.com\",\"employee_id\":1080,\"first_name\":\"Simona\",\"gender\":\"female\",\"hourly_rate\":19,\"last_name\":\"Gookey\",\"ssn\":\"742-79-2261\"},{\"age\":48,\"email\":\"jpbi@mycompany.com\",\"employee_id\":1081,\"first_name\":\"Sheridan\",\"gender\":\"male\",\"hourly_rate\":23,\"last_name\":\"Lezemere\",\"ssn\":\"674-68-7406\"},{\"age\":49,\"email\":\"ihxq@mycompany.com\",\"employee_id\":1082,\"first_name\":\"Nikolai\",\"gender\":\"male\",\"hourly_rate\":15,\"last_name\":\"Fewings\",\"ssn\":\"662-14-4672\"},{\"age\":33,\"email\":\"lzrf@mycompany.com\",\"employee_id\":1083,\"first_name\":\"Winslow\",\"gender\":\"male\",\"hourly_rate\":14,\"last_name\":\"Murfill\",\"ssn\":\"536-57-7029\"},{\"age\":32,\"email\":\"ygfg@mycompany.com\",\"employee_id\":1084,\"first_name\":\"Constantin\",\"gender\":\"male\",\"hourly_rate\":8,\"last_name\":\"Ballin\",\"ssn\":\"422-62-7803\"},{\"age\":49,\"email\":\"cfpl@mycompany.com\",\"employee_id\":1085,\"first_name\":\"Lea\",\"gender\":\"female\",\"hourly_rate\":9,\"last_name\":\"Wren\",\"ssn\":\"825-03-0834\"},{\"age\":44,\"email\":\"phju@mycompany.com\",\"employee_id\":1086,\"first_name\":\"Giacobo\",\"gender\":\"female\",\"hourly_rate\":13,\"last_name\":\"Hawyes\",\"ssn\":\"611-82-8402\"},{\"age\":31,\"email\":\"chst@mycompany.com\",\"employee_id\":1087,\"first_name\":\"Livvie\",\"gender\":\"female\",\"hourly_rate\":24,\"last_name\":\"Guildford\",\"ssn\":\"016-14-4716\"},{\"age\":42,\"email\":\"kpum@mycompany.com\",\"employee_id\":1088,\"first_name\":\"West\",\"gender\":\"female\",\"hourly_rate\":8,\"last_name\":\"Klemt\",\"ssn\":\"307-27-8393\"},{\"age\":30,\"email\":\"gpbe@mycompany.com\",\"employee_id\":1089,\"first_name\":\"Anatollo\",\"gender\":\"female\",\"hourly_rate\":13,\"last_name\":\"Henderson\",\"ssn\":\"547-89-0756\"},{\"age\":39,\"email\":\"awns@mycompany.com\",\"employee_id\":1090,\"first_name\":\"Caria\",\"gender\":\"male\",\"hourly_rate\":19,\"last_name\":\"Murfill\",\"ssn\":\"968-92-5521\"},{\"age\":49,\"email\":\"ickd@mycompany.com\",\"employee_id\":1091,\"first_name\":\"Trent\",\"gender\":\"female\",\"hourly_rate\":17,\"last_name\":\"Napton\",\"ssn\":\"131-10-2192\"},{\"age\":38,\"email\":\"hevl@mycompany.com\",\"employee_id\":1092,\"first_name\":\"Chelsae\",\"gender\":\"female\",\"hourly_rate\":21,\"last_name\":\"Ballin\",\"ssn\":\"929-64-1842\"},{\"age\":37,\"email\":\"kcms@mycompany.com\",\"employee_id\":1093,\"first_name\":\"Rossie\",\"gender\":\"female\",\"hourly_rate\":15,\"last_name\":\"Gwyther\",\"ssn\":\"675-97-3680\"},{\"age\":41,\"email\":\"xjkz@mycompany.com\",\"employee_id\":1094,\"first_name\":\"Sonnie\",\"gender\":\"male\",\"hourly_rate\":9,\"last_name\":\"Smewings\",\"ssn\":\"841-57-1504\"},{\"age\":35,\"email\":\"jpji@mycompany.com\",\"employee_id\":1095,\"first_name\":\"Jeremy\",\"gender\":\"female\",\"hourly_rate\":8,\"last_name\":\"MacFadzean\",\"ssn\":\"847-82-0071\"},{\"age\":47,\"email\":\"cevc@mycompany.com\",\"employee_id\":1096,\"first_name\":\"Jere\",\"gender\":\"female\",\"hourly_rate\":11,\"last_name\":\"Coll\",\"ssn\":\"427-07-7019\"},{\"age\":44,\"email\":\"ftpk@mycompany.com\",\"employee_id\":1097,\"first_name\":\"Andriette\",\"gender\":\"male\",\"hourly_rate\":16,\"last_name\":\"Broadbridge\",\"ssn\":\"342-03-4896\"},{\"age\":33,\"email\":\"bcrr@mycompany.com\",\"employee_id\":1098,\"first_name\":\"Brenna\",\"gender\":\"female\",\"hourly_rate\":10,\"last_name\":\"Cleugher\",\"ssn\":\"874-68-3535\"},{\"age\":37,\"email\":\"bmvb@mycompany.com\",\"employee_id\":1099,\"first_name\":\"Trumann\",\"gender\":\"male\",\"hourly_rate\":9,\"last_name\":\"Graveson\",\"ssn\":\"345-99-5761\"}]},\"fields\":[{\"name\":\"employee_id\",\"type\":\"int\"},{\"name\":\"first_name\",\"type\":\"string\"},{\"name\":\"last_name\",\"type\":\"string\"},{\"name\":\"age\",\"type\":\"int\"},{\"name\":\"ssn\",\"type\":\"string\"},{\"name\":\"hourly_rate\",\"type\":\"int\"},{\"name\":\"gender\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"}],\"name\":\"payroll.payrollemployee\",\"type\":\"record\"}"
}

func (r Payrollemployee) SchemaName() string {
	return "payroll.payrollemployee"
}

func (_ Payrollemployee) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Payrollemployee) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Payrollemployee) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Payrollemployee) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Payrollemployee) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Payrollemployee) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Payrollemployee) SetString(v string)   { panic("Unsupported operation") }
func (_ Payrollemployee) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Payrollemployee) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Employee_id}

		return w

	case 1:
		w := types.String{Target: &r.First_name}

		return w

	case 2:
		w := types.String{Target: &r.Last_name}

		return w

	case 3:
		w := types.Int{Target: &r.Age}

		return w

	case 4:
		w := types.String{Target: &r.Ssn}

		return w

	case 5:
		w := types.Int{Target: &r.Hourly_rate}

		return w

	case 6:
		w := types.String{Target: &r.Gender}

		return w

	case 7:
		w := types.String{Target: &r.Email}

		return w

	}
	panic("Unknown field index")
}

func (r *Payrollemployee) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Payrollemployee) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Payrollemployee) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Payrollemployee) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Payrollemployee) HintSize(int)                     { panic("Unsupported operation") }
func (_ Payrollemployee) Finalize()                        {}

func (_ Payrollemployee) AvroCRC64Fingerprint() []byte {
	return []byte(PayrollemployeeAvroCRC64Fingerprint)
}

func (r Payrollemployee) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["employee_id"], err = json.Marshal(r.Employee_id)
	if err != nil {
		return nil, err
	}
	output["first_name"], err = json.Marshal(r.First_name)
	if err != nil {
		return nil, err
	}
	output["last_name"], err = json.Marshal(r.Last_name)
	if err != nil {
		return nil, err
	}
	output["age"], err = json.Marshal(r.Age)
	if err != nil {
		return nil, err
	}
	output["ssn"], err = json.Marshal(r.Ssn)
	if err != nil {
		return nil, err
	}
	output["hourly_rate"], err = json.Marshal(r.Hourly_rate)
	if err != nil {
		return nil, err
	}
	output["gender"], err = json.Marshal(r.Gender)
	if err != nil {
		return nil, err
	}
	output["email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Payrollemployee) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["employee_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Employee_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for employee_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["first_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.First_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for first_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["last_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Last_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for last_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["age"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Age); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for age")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ssn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ssn); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ssn")
	}
	val = func() json.RawMessage {
		if v, ok := fields["hourly_rate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Hourly_rate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for hourly_rate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["gender"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Gender); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for gender")
	}
	val = func() json.RawMessage {
		if v, ok := fields["email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for email")
	}
	return nil
}
