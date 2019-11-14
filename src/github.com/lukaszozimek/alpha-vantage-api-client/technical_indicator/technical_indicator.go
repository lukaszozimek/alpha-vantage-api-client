package technical_indicator

import (
	"fmt"
)

func GetSMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", SMA, symbol, interval, timePeriod, seriesType, apiKey), c)
}
func GetEMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", EMA, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetWMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", WMA, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetDEMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", DEMA, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTEMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", TEMA, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTRIMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", TRIMA, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetKAMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", KAMA, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMAMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MAMA, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetVWAPReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", VWAP, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetT3Report(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", T3, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMACDReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MACD, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMACDEXTReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MACDEXT, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSTOCHReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", STOCH, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSTOCHFReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", STOCHF, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetRSIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", RSI, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSTOCHRSIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", STOCHRSI, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetWILLRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", WILLR, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetADXReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", ADX, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetADXRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", ADXR, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetAPOReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", APO, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetPPOReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", PPO, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMOMReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MOM, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetBOPReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", BOP, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetCCIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", CCI, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetCMOReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", CMO, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetROCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", ROC, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetROCRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", ROCR, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetAROONReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", AROON, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetAROONOSCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", AROONOSC, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMFIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MFI, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTRIXReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", TRIX, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetULTOSCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", ULTOSC, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetDXReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", DX, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetminusDireport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MINUS_DI, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetplusDireport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", PLUS_DI, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetminusDmreport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MINUS_DM, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetplusDmreport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", PLUS_DM, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetBBANDSReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", BBANDS, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMIDPOINTReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MIDPOINT, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMIDPRICEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", MIDPRICE, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSARReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", SAR, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTRANGEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", TRANGE, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetATRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", ATR, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetNATRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", NATR, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetADReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", AD, symbol, interval, timePeriod, seriesType, apiKey), c)
}
func GetADOSCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", ADOSC, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetOBVReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", OBV, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GethtTrendlinereport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", HT_TRENDLINE, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GethtSinereport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", HT_SINE, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GethtTrendmodereport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", HT_TRENDMODE, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GethtDcperiodreport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", HT_DCPERIOD, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GethtDcphasereport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", HT_DCPHASE, symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GethtPhasorreport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=%v&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", HT_PHASOR, symbol, interval, timePeriod, seriesType, apiKey), c)

}
