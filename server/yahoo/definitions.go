package yahoo

import "go.mongodb.org/mongo-driver/bson/primitive"

type YFStock struct {
	ID      primitive.ObjectID
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    struct {
		Five2WeekChange              float64  `json:"52WeekChange"`
		SandP52WeekChange            float64  `json:"SandP52WeekChange"`
		Address1                     string   `json:"address1"`
		Address2                     string   `json:"address2"`
		Algorithm                    string   `json:"algorithm"`
		AnnualHoldingsTurnover       string   `json:"annualHoldingsTurnover"`
		AnnualReportExpenseRatio     string   `json:"annualReportExpenseRatio"`
		Ask                          int      `json:"ask"`
		AskSize                      int      `json:"askSize"`
		AverageDailyVolume10Day      int      `json:"averageDailyVolume10Day"`
		AverageVolume                int      `json:"averageVolume"`
		AverageVolume10Days          int      `json:"averageVolume10days"`
		Beta                         float64  `json:"beta"`
		Beta3Year                    string   `json:"beta3Year"`
		Bid                          float64  `json:"bid"`
		BidSize                      int      `json:"bidSize"`
		BookValue                    float64  `json:"bookValue"`
		Category                     string   `json:"category"`
		CirculatingSupply            string   `json:"circulatingSupply"`
		City                         string   `json:"city"`
		CoinMarketCapLink            string   `json:"coinMarketCapLink"`
		CompanyOfficers              []string `json:"companyOfficers"`
		Country                      string   `json:"country"`
		Currency                     string   `json:"currency"`
		CurrentPrice                 float64  `json:"currentPrice"`
		CurrentRatio                 float64  `json:"currentRatio"`
		DateShortInterest            string   `json:"dateShortInterest"`
		DayHigh                      float64  `json:"dayHigh"`
		DayLow                       int      `json:"dayLow"`
		DebtToEquity                 string   `json:"debtToEquity"`
		DividendRate                 string   `json:"dividendRate"`
		DividendYield                string   `json:"dividendYield"`
		EarningsGrowth               float64  `json:"earningsGrowth"`
		EarningsQuarterlyGrowth      float64  `json:"earningsQuarterlyGrowth"`
		Ebitda                       string   `json:"ebitda"`
		EbitdaMargins                int      `json:"ebitdaMargins"`
		EnterpriseToEbitda           string   `json:"enterpriseToEbitda"`
		EnterpriseToRevenue          float64  `json:"enterpriseToRevenue"`
		EnterpriseValue              int      `json:"enterpriseValue"`
		ExDividendDate               string   `json:"exDividendDate"`
		Exchange                     string   `json:"exchange"`
		ExchangeTimezoneName         string   `json:"exchangeTimezoneName"`
		ExchangeTimezoneShortName    string   `json:"exchangeTimezoneShortName"`
		ExpireDate                   string   `json:"expireDate"`
		FiftyDayAverage              float64  `json:"fiftyDayAverage"`
		FiftyTwoWeekHigh             float64  `json:"fiftyTwoWeekHigh"`
		FiftyTwoWeekLow              float64  `json:"fiftyTwoWeekLow"`
		FinancialCurrency            string   `json:"financialCurrency"`
		FiveYearAverageReturn        string   `json:"fiveYearAverageReturn"`
		FiveYearAvgDividendYield     string   `json:"fiveYearAvgDividendYield"`
		FloatShares                  int      `json:"floatShares"`
		ForwardEps                   float64  `json:"forwardEps"`
		ForwardPE                    float64  `json:"forwardPE"`
		FreeCashflow                 int      `json:"freeCashflow"`
		FromCurrency                 string   `json:"fromCurrency"`
		FullTimeEmployees            int      `json:"fullTimeEmployees"`
		FundFamily                   string   `json:"fundFamily"`
		FundInceptionDate            string   `json:"fundInceptionDate"`
		GmtOffSetMilliseconds        string   `json:"gmtOffSetMilliseconds"`
		GrossMargins                 int      `json:"grossMargins"`
		GrossProfits                 int      `json:"grossProfits"`
		HeldPercentInsiders          float64  `json:"heldPercentInsiders"`
		HeldPercentInstitutions      float64  `json:"heldPercentInstitutions"`
		ImpliedSharesOutstanding     int      `json:"impliedSharesOutstanding"`
		Industry                     string   `json:"industry"`
		IsEsgPopulated               bool     `json:"isEsgPopulated"`
		LastCapGain                  string   `json:"lastCapGain"`
		LastDividendDate             string   `json:"lastDividendDate"`
		LastDividendValue            string   `json:"lastDividendValue"`
		LastFiscalYearEnd            int      `json:"lastFiscalYearEnd"`
		LastMarket                   string   `json:"lastMarket"`
		LastSplitDate                string   `json:"lastSplitDate"`
		LastSplitFactor              string   `json:"lastSplitFactor"`
		LegalType                    string   `json:"legalType"`
		LogoURL                      string   `json:"logo_url"`
		LongBusinessSummary          string   `json:"longBusinessSummary"`
		LongName                     string   `json:"longName"`
		Market                       string   `json:"market"`
		MarketCap                    int      `json:"marketCap"`
		MaxAge                       int      `json:"maxAge"`
		MaxSupply                    string   `json:"maxSupply"`
		MessageBoardID               string   `json:"messageBoardId"`
		MorningStarOverallRating     string   `json:"morningStarOverallRating"`
		MorningStarRiskRating        string   `json:"morningStarRiskRating"`
		MostRecentQuarter            int      `json:"mostRecentQuarter"`
		NavPrice                     string   `json:"navPrice"`
		NetIncomeToCommon            int      `json:"netIncomeToCommon"`
		NextFiscalYearEnd            int      `json:"nextFiscalYearEnd"`
		NumberOfAnalystOpinions      int      `json:"numberOfAnalystOpinions"`
		Open                         float64  `json:"open"`
		OpenInterest                 string   `json:"openInterest"`
		OperatingCashflow            int      `json:"operatingCashflow"`
		OperatingMargins             float64  `json:"operatingMargins"`
		PayoutRatio                  int      `json:"payoutRatio"`
		PegRatio                     string   `json:"pegRatio"`
		Phone                        string   `json:"phone"`
		PreMarketPrice               string   `json:"preMarketPrice"`
		PreviousClose                float64  `json:"previousClose"`
		PriceHint                    int      `json:"priceHint"`
		PriceToBook                  float64  `json:"priceToBook"`
		PriceToSalesTrailing12Months float64  `json:"priceToSalesTrailing12Months"`
		ProfitMargins                float64  `json:"profitMargins"`
		QuickRatio                   float64  `json:"quickRatio"`
		QuoteType                    string   `json:"quoteType"`
		RecommendationKey            string   `json:"recommendationKey"`
		RecommendationMean           float64  `json:"recommendationMean"`
		RegularMarketDayHigh         float64  `json:"regularMarketDayHigh"`
		RegularMarketDayLow          int      `json:"regularMarketDayLow"`
		RegularMarketOpen            float64  `json:"regularMarketOpen"`
		RegularMarketPreviousClose   float64  `json:"regularMarketPreviousClose"`
		RegularMarketPrice           float64  `json:"regularMarketPrice"`
		RegularMarketVolume          int      `json:"regularMarketVolume"`
		ReturnOnAssets               float64  `json:"returnOnAssets"`
		ReturnOnEquity               float64  `json:"returnOnEquity"`
		RevenueGrowth                float64  `json:"revenueGrowth"`
		RevenuePerShare              float64  `json:"revenuePerShare"`
		RevenueQuarterlyGrowth       string   `json:"revenueQuarterlyGrowth"`
		Sector                       string   `json:"sector"`
		SharesOutstanding            int      `json:"sharesOutstanding"`
		SharesPercentSharesOut       string   `json:"sharesPercentSharesOut"`
		SharesShort                  string   `json:"sharesShort"`
		SharesShortPreviousMonthDate string   `json:"sharesShortPreviousMonthDate"`
		SharesShortPriorMonth        string   `json:"sharesShortPriorMonth"`
		ShortName                    string   `json:"shortName"`
		ShortPercentOfFloat          string   `json:"shortPercentOfFloat"`
		ShortRatio                   string   `json:"shortRatio"`
		StartDate                    string   `json:"startDate"`
		StrikePrice                  string   `json:"strikePrice"`
		Symbol                       string   `json:"symbol"`
		TargetHighPrice              float64  `json:"targetHighPrice"`
		TargetLowPrice               float64  `json:"targetLowPrice"`
		TargetMeanPrice              float64  `json:"targetMeanPrice"`
		TargetMedianPrice            float64  `json:"targetMedianPrice"`
		ThreeYearAverageReturn       string   `json:"threeYearAverageReturn"`
		ToCurrency                   string   `json:"toCurrency"`
		TotalAssets                  string   `json:"totalAssets"`
		TotalCash                    int      `json:"totalCash"`
		TotalCashPerShare            float64  `json:"totalCashPerShare"`
		TotalDebt                    int      `json:"totalDebt"`
		TotalRevenue                 int      `json:"totalRevenue"`
		Tradeable                    bool     `json:"tradeable"`
		TrailingAnnualDividendRate   int      `json:"trailingAnnualDividendRate"`
		TrailingAnnualDividendYield  int      `json:"trailingAnnualDividendYield"`
		TrailingEps                  float64  `json:"trailingEps"`
		TrailingPE                   float64  `json:"trailingPE"`
		TwoHundredDayAverage         float64  `json:"twoHundredDayAverage"`
		Volume                       int      `json:"volume"`
		Volume24Hr                   string   `json:"volume24Hr"`
		VolumeAllCurrencies          string   `json:"volumeAllCurrencies"`
		Website                      string   `json:"website"`
		Yield                        string   `json:"yield"`
		YtdReturn                    string   `json:"ytdReturn"`
		Zip                          string   `json:"zip"`
	} `json:"data"`
}
