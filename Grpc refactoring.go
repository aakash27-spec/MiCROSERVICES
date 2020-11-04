package data

import "github.com/hashicorp/go-hclog"

type ExchangeRates struct {
    log hclog.Logger
    rate map[string]float64
}

func Newrates(l hclog.Logger) (*ExchangeRates, error){
    er := &ExchangeRates{log: l,rate: map[string]float64{}}
	 
	err :=er.getrates()

    return er,nil
}

func (e*ExchangeRates) getRates() error {
    resp, err := http.defaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
    if err !=nil {
        return nil
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.errorf("Expected error code 200 got %d",resp.StatusCode)
    }
    defer resp.Body.Close()

    md := &cubes{}
    xml.NewDecoder(resp.Body).Decode(&md)

    for _,c:=range md.CubeData {
        r, err := strconv.ParseFloat(c.Rate, 64)
        if err !=nil {
            return err
        }
        e.rates[c.Currency]=r

    }

    return nil

    }

    type Cubes struct {
        CubeData []Cube 'xml:"Cube>cube>cube'
    }

    type Cube stuct{
        Currency string 'xml:"currency,attr"'
        Rate string 'xml:"rate,attr"'
    }
}