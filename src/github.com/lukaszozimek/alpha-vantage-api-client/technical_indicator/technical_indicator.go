package technical_indicator

import (
	"fmt"
)

func GetSMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=SMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)
}
func GetEMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=EMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetWMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=WMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetDEMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=DEMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTEMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=TEMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTRIMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=TRIMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetKAMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=KAMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMAMAReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MAMA&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetVWAPReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=VWAP&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetT3Report(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=T3&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMACDReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MACD&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMACDEXTReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MACDEXT&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSTOCHReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=STOCH&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSTOCHFReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=STOCHF&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetRSIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=RSI&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSTOCHRSIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=STOCHRSI&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetWILLRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=WILLR&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetADXReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=ADX&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetADXRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=ADXR&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetAPOReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=APO&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetPPOReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=PPO&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMOMReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MOM&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetBOPReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=BOP&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetCCIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=CCI&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetCMOReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=CMO&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetROCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=HT_DCPHASE&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetROCRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=ROCR&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetAROONReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=AROON&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetAROONOSCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=AROONOSC&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMFIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MFI&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTRIXReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=TRIX&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetULTOSCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=ULTOSC&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetDXReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=DX&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMINUS_DIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MINUS_DI&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetPLUS_DIReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=PLUS_DI&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMINUS_DMReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MINUS_DM&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetPLUS_DMReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=PLUS_DM&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetBBANDSReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=BBANDS&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMIDPOINTReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MIDPOINT&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetMIDPRICEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=MIDPRICE&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetSARReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=SAR&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetTRANGEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=TRANGE&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetATRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=ATR&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetNATRReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=NATR&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetADReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=AD&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)
}
func GetADOSCReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=ADOSC&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetOBVReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=OBV&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetHT_TRENDLINEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=HT_TRENDLINE&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetHT_SINEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=HT_SINE&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetHT_TRENDMODEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=HT_TRENDMODE&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetHT_DCPERIODReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=HT_DCPERIOD&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetHT_DCPHASEReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {
	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=HT_DCPHASE&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
func GetHT_PHASORReport(symbol string, interval string, timePeriod string, seriesType string, apiKey string, c *Client) *AlphaVantageTechnicalIndicatorResponse {
	return makeApiCallGet(fmt.Sprintf("https: //www.alphavantage.co/query?function=HT_PHASOR&symbol=%v&interval=%v&time_period=%v&series_type=%v&apikey=%v", symbol, interval, timePeriod, seriesType, apiKey), c)

}
