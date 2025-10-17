package app

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	. "rpi-test/pkg/openapi/v1"
)

const (
	DiscreteClass     = "discrete"
	RangeClass        = "range"
	TrafficLightClass = "tl"
)

type ApiRequest struct {
	Command string                 `json:"cmd"`
	Params  map[string]interface{} `json:"params"`
}

type ApiResponse struct {
	Result interface{} `json:"result"`
	Error  *string     `json:"error"`
}

type TlRequest struct {
	Mode string
}

type TlRequestManual struct {
	Mode string
	Sn   string
	Ew   string
}

type TlMask uint8

const VIRTUAL_REGISTER = 0xAA

const (
	TlModeMask TlMask = 0b1100_0000
	TlSnMask   TlMask = 0b00_000_111
	TlEwMask   TlMask = 0b00_111_000
)

// TODO: get rid of reflection, it's used in every handler so it is very consuming
// It is used only on GET queries to abstract StrictHandler request objects.
// You can try to replace openapi StrictHandler by openapi default Handler interface.
func SetValue(body interface{}, value interface{}, fieldName string) {
	bodyVal := reflect.ValueOf(body)
	if bodyVal.Kind() == reflect.Ptr {
		bodyVal = bodyVal.Elem()
	}

	bodyType := bodyVal.Type()

	for i := 0; i < bodyType.NumField(); i++ {
		field := bodyVal.Field(i)
		fieldType := bodyType.Field(i)

		if fieldType.Name == fieldName {
			if field.CanAddr() {
				field = field.Addr().Elem()
				switch field.Kind() {
				case reflect.String:
					field.SetString(value.(string))
				case reflect.Int:
					field.SetInt(value.(int64))
				}
			}
		}
	}
}

func (app *App) processCommand(task, element, body interface{}) {
	fmt.Println("processClassCommand:", task, element, body)
	for _, t := range app.Tasks {
		if t.Name != task {
			continue
		}
		fmt.Println(t.config.Elements)
		for el, elconf := range t.config.Elements {
			if el != element {
				continue
			}
			app.processElement(elconf, body)
			return
		}
	}
}

func (app *App) processInfo(task, element, body interface{}) interface{} {
	fmt.Println("processClassCommand:", task, element, body)
	for _, t := range app.Tasks {
		if t.Name != task {
			continue
		}
		for el, elconf := range t.config.Elements {
			if el != element {
				continue
			}
			app.processElementInfo(elconf, body)
			break
		}
	}

	return nil
}

func (app *App) processElementInfo(elementConfig *ElementsConfig, body interface{}) {
	r, ok := elementConfig.Register.(int)
	if !ok {
		log.Println("Not ok: register!")
		return
	}

	query := ModBusReg{
		Addr: *elementConfig.task.config.Addr,
		Reg:  uint8(r),
	}

	result, err := app.getRegisterValue(&query)
	if err != nil {
		log.Println("ERROR: getRegisterValue", err)
		return
	}

	switch elementConfig.Class {
	case DiscreteClass:
		switch result {
		case 0:
			SetValue(body, string(DiscreteValueOff), "Value")
		case 1:
			SetValue(body, string(DiscreteValueOn), "Value")
		case 2:
			SetValue(body, string(DiscreteValueBlink), "Value")
		}

	case RangeClass:
		SetValue(body, int64(result), "Value")

	case TrafficLightClass:
		var mode string

		for m, v := range app.Classes[TrafficLightClass]["mode"].Values {
			log.Printf("m: %s v: %v, masked: %d", m, v, result&uint8(TlModeMask))
			if v.(int) == int(result&uint8(TlModeMask)) {
				log.Println("found mode!", v)
				mode = m
				break
			}
		}

		if mode == "" {
			log.Println("err no mode found")
			return
		}

		fmt.Println(mode)

		switch TlMode(mode) {
		case TlModeManual:
			var snColor TlSn
			var ewColor TlEw

			snValue := result & uint8(TlSnMask)
			ewValue := result & uint8(TlEwMask)

			for c, v := range app.Classes[TrafficLightClass]["sn"].Values {
				if v == snValue {
					snColor = TlSn(c)
					break
				}
			}

			for c, v := range app.Classes[TrafficLightClass]["ew"].Values {
				if v == ewValue {
					ewColor = TlEw(c)
					break
				}
			}

			SetValue(body, string(snColor), "Sn")
			SetValue(body, string(ewColor), "Ew")
			fallthrough
		default:
			SetValue(body, mode, "Mode")
		}
	}
}

func (app *App) processElement(elementConfig *ElementsConfig, body interface{}) {
	var value uint8

	r, ok := elementConfig.Register.(int)
	if !ok {
		log.Println("Not ok: register!")
		return
	}

	switch elementConfig.Class {
	case DiscreteClass:
		body, ok := body.(Discrete)
		if !ok {
			log.Println("Not ok: body!")
			return
		}

		v, ok := app.Classes[DiscreteClass]["value"].Values[string(body.Value)].(int)
		if !ok {
			log.Println("Not value: body!")
			return
		}

		value = uint8(v)

	case TrafficLightClass:
		body, ok := body.(Tl)
		if !ok {
			log.Println("Not ok: body!")
			return
		}
		mode, ok := app.Classes[TrafficLightClass]["mode"].Values[string(body.Mode)].(int)
		if !ok {
			log.Println("Not mody: body!")
			return
		}

		if body.Mode == TlModeManual {
			if body.Sn == nil || body.Ew == nil {
				log.Println("Error: now tl directions in request")
				return
			}

			sn, ok := app.Classes[TrafficLightClass]["sn"].Values[string(*body.Sn)].(int)
			if !ok {
				log.Println("Not sn: body!")
				return
			}

			ew, ok := app.Classes[TrafficLightClass]["ew"].Values[string(*body.Ew)].(int)
			if !ok {
				log.Println("Not sn: body!")
				return
			}

			value = uint8(mode) | uint8(ew) | uint8(sn)
		} else {
			value = uint8(mode)
		}

	case RangeClass:
		body, ok := body.(Range)
		if !ok {
			log.Println("Not ok: body!")
			return
		}

		value = uint8(body.Value)
	default:
		log.Println("Error no matching class handler")
		return
	}

	// TODO: get rid of hardcode inside if statement and implement logic for virtual registers
	if r == VIRTUAL_REGISTER {
		url := fmt.Sprintf("http://%s:%s/hacked", app.Config.Virtual.Addr, app.Config.Virtual.Port)
		fmt.Println(url)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Println(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		}

		if resp != nil {
			log.Println(resp.StatusCode)
		}
		return
	}

	fmt.Println(app.setRegister(&ModbusRegData{
		Addr:  *elementConfig.task.config.Addr,
		Reg:   uint8(r),
		Value: value,
	}))
}
