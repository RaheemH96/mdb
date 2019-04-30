package main

import (
	"context"
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Trainer is the structure of our documents in collection Trainers
type Trainers struct {
	Name string
	Age  float64
	City string
}

type Tri struct {
	Symbol                                           string  `bson:"Symbol"`
	Name                                             string  `bson:"Name"`
	PE                                               float64 `bson:"P/E"`
	PriceToBook                                      float64 `bson:"Price to Book"`
	PriceToCF                                        float64 `bson:"Price to CF"`
	PEG                                              float64 `bson:"PEG"`
	DividendAdjustedPEG                              float64 `bson:"Dividend-Adjusted PEG"`
	ROE                                              float64 `bson:"ROE"`
	MarketCapMil                                     float64 `bson:"Market Cap (mil)"`
	EnterpriseValMil                                 float64 `bson:"Enterprise Val (mil)"`
	Five2WkHigh                                      float64 `bson:"52-Wk High"`
	Five2WkLow                                       float64 `bson:"52-Wk Low"`
	Five0DayAvgVol1000S                              float64 `bson:"50-Day Avg Vol (1000s)"`
	Volume1000S                                      float64 `bson:"Volume (1000s)"`
	CurrentPrice                                     float64 `bson:"Current Price"`
	PriceChg                                         float64 `bson:"Price $ Chg"`
	PriceChgPct                                      float64 `bson:"Price % Chg"`
	ChgCurWeek                                       float64 `bson:"% Chg Cur Week"`
	Chg1Month                                        float64 `bson:"% Chg 1 Month"`
	Chg3Months                                       float64 `bson:"% Chg 3 Months"`
	Chg6Months                                       float64 `bson:"% Chg 6 Months"`
	Chg12Months                                      float64 `bson:"% Chg 12 Months"`
	ChgYTD                                           float64 `bson:"% Chg YTD"`
	EPSLstRptd                                       string  `bson:"EPS Lst Rptd"`
	EPSDueDate                                       string  `bson:"EPS Due Date"`
	Alpha                                            float64 `bson:"Alpha"`
	Beta                                             float64 `bson:"Beta"`
	AvgTrueRange                                     float64 `bson:"Avg True Range"`
	UpDownVol                                        float64 `bson:"Up/Down Vol"`
	Yield                                            float64 `bson:"Yield %"`
	Debt                                             float64 `bson:"Debt %"`
	CurrentRati0                                     float64 `bson:"Current Rati0"`
	EVT0FCF                                          float64 `bson:"EV t0 FCF"`
	LTDebtT0W0RkingCap                               float64 `bson:"LT Debt t0 W0rking Cap"`
	Exchange                                         string  `bson:"Exchange"`
	ADR                                              float64 `bson:"ADR"`
	ETF                                              float64 `bson:"ETF"`
	Options                                          float64 `bson:"Options"`
	IPODate                                          string  `bson:"IPO Date"` //Changed string to int
	IncorpDate                                       float64 `bson:"Incorp Date"`
	CompanyDescription                               string  `bson:"Company Description"`
	ExDividendDate                                   string  `bson:"Ex-Dividend Date"` //Changed string to int
	SharesInFloat1000S                               float64 `bson:"Shares in Float (1000s)"`
	AnnualSalesMil                                   float64 `bson:"Annual Sales (mil)"`
	SalesChgLstQtr                                   float64 `bson:"Sales % Chg Lst Qtr"`
	SalesChgLstYr                                    float64 `bson:"Sales % Chg Lst Yr"`
	SalesAccel2Qtrs                                  float64 `bson:"Sales Accel 2 Qtrs"`
	SalesAccel3Qtrs                                  float64 `bson:"Sales Accel 3 Qtrs"`
	SalesGrowth3Yr                                   float64 `bson:"Sales Growth 3 Yr"`
	SalesGrowth5Yr                                   float64 `bson:"Sales Growth 5 Yr"`
	AvgSalesChg2Q                                    float64 `bson:"Avg Sales % Chg 2Q"`
	AvgSalesChg3Q                                    float64 `bson:"Avg Sales % Chg 3Q"`
	AvgSalesChg4Q                                    float64 `bson:"Avg Sales % Chg 4Q"`
	AvgSalesChg5Q                                    float64 `bson:"Avg Sales % Chg 5Q"`
	AvgSalesChg6Q                                    float64 `bson:"Avg Sales % Chg 6Q"`
	EPSChgLastQtr                                    float64 `bson:"EPS % Chg Last Qtr"`
	EPSChgLastQtrPN                                  float64 `bson:"EPS % Chg Last Qtr (-/+)"`
	EPSChg1QAgo                                      float64 `bson:"EPS % Chg 1 Q Ago"`
	EPSChg1QAgoPN                                    float64 `bson:"EPS % Chg 1 Q Ago (-/+)"`
	EPSChg2QAgo                                      float64 `bson:"EPS % Chg 2 Q Ago"`
	EPSChg2QAgoPN                                    float64 `bson:"EPS % Chg 2 Q Ago (-/+)"`
	EPSChg3QAgo                                      float64 `bson:"EPS % Chg 3 Q Ago"`
	EPSChg3QAgoPN                                    float64 `bson:"EPS % Chg 3 Q Ago (-/+)"`
	EPSTrailing4Qtrs                                 float64 `bson:"EPS Trailing 4 Qtrs"`
	FiscalEPSLstYr                                   float64 `bson:"Fiscal EPS Lst Yr"`
	FiscalEPS1YrAgo                                  float64 `bson:"Fiscal EPS 1 Yr Ago"`
	FiscalEPS2YrsAgo                                 float64 `bson:"Fiscal EPS 2 Yrs Ago"`
	FiscalEPS3YrsAgo                                 float64 `bson:"Fiscal EPS 3 Yrs Ago"`
	FiscalEPS4YrsAgo                                 float64 `bson:"Fiscal EPS 4 Yrs Ago"`
	FiscalEPS5YrsAgo                                 float64 `bson:"Fiscal EPS 5 Yrs Ago"`
	FiscalEPS6YrsAgo                                 float64 `bson:"Fiscal EPS 6 Yrs Ago"`
	EPSChgLstYr                                      float64 `bson:"EPS % Chg Lst Yr"`
	EPSChg1YrAgo                                     float64 `bson:"EPS % Chg 1 Yr Ago"`
	EPSEstCurQtr                                     float64 `bson:"EPS Est Cur Qtr %"`
	EPSEstCurYr                                      float64 `bson:"EPS Est Cur Yr %"`
	EPSEstNextYr                                     float64 `bson:"EPS Est Next Yr %"`
	SustainableGrowth                                float64 `bson:"Sustainable Growth %"`
	EPSAccel3Qtrs                                    float64 `bson:"EPS Accel 3 Qtrs"`
	EPSGrowth1Yr                                     float64 `bson:"EPS % Growth 1 Yr"`
	EPSGrowth3Yr                                     float64 `bson:"EPS % Growth 3 Yr"`
	EPSGrowth5Yr                                     float64 `bson:"EPS % Growth 5 Yr"`
	EPSGrowth5YrPctRnk                               float64 `bson:"EPS % Growth 5 Yr Pct Rnk"`
	EarningsStability                                float64 `bson:"Earnings Stability"`
	AvgEPSChg2Q                                      float64 `bson:"Avg EPS % Chg 2Q"`
	AvgEPSChg3Q                                      float64 `bson:"Avg EPS % Chg 3Q"`
	AvgEPSChg4Q                                      float64 `bson:"Avg EPS % Chg 4Q"`
	AvgEPSChg5Q                                      float64 `bson:"Avg EPS % Chg 5Q"`
	AvgEPSChg6Q                                      float64 `bson:"Avg EPS % Chg 6Q"`
	EPSSurprise                                      float64 `bson:"EPS Surprise"`
	EPSTrl4QGtrEPS4YrsAgo                            float64 `bson:"EPS Trl 4Q Gtr EPS 4 Yrs Ago"`
	EPSLstYrGtrEPS4YrsAgo                            float64 `bson:"EPS Lst Yr Gtr EPS 4 Yrs Ago"`
	EPSTrl4QGeqEPSLstFiscalYr                        float64 `bson:"EPS Trl 4Q Geq EPS Lst Fiscal Yr"`
	EPSChgLstQGtr3YrGrowth                           float64 `bson:"EPS % Chg Lst Q Gtr 3-Yr Growth"`
	ThreeYrEPSGrowthGeq5Yr                           float64 `bson:"3-Yr EPS Growth Geq 5-Yr"`
	Sector                                           string  `bson:"Sector"`
	IndustryName                                     string  `bson:"Industry Name"`
	Funds                                            float64 `bson:"Funds %"`
	FundsIncrease                                    float64 `bson:"Funds % Increase"`
	NumberOfFunds                                    float64 `bson:"Number of Funds"`
	Mgmt                                             float64 `bson:"Mgmt %"`
	InstOwnershipLssMedian                           float64 `bson:"Inst Ownership Lss Median"`
	OffHigh                                          float64 `bson:"% Off High"`
	Trl26WkPerfVsSP500                               float64 `bson:"Trl 26 Wk % Perf vs S&P 500"`
	PriceVs50Day                                     float64 `bson:"Price vs 50-Day"`
	PriceVs200Day                                    float64 `bson:"Price vs 200-Day"`
	VolChgVs50Day                                    float64 `bson:"Vol % Chg vs 50-Day"`
	EstPE                                            float64 `bson:"Est P/E"`
	PEVsSP500PE                                      float64 `bson:"P/E vs S&P 500 P/E (%)"`
	LiabToAssetsLssIndMedian                         float64 `bson:"Liab to Assets Lss Ind Median"`
	PriceVs10Day                                     float64 `bson:"Price vs 10-Day"`
	PriceVs21Day                                     float64 `bson:"Price vs 21-Day"`
	PriceVs150Day                                    float64 `bson:"Price vs 150-Day"`
	VolChgVs10Week                                   float64 `bson:"Vol % Chg vs 10-Week"`
	One0Day21Day50Day                                float64 `bson:"10 Day > 21 Day > 50 Day"`
	Five0Day150Day200Day                             float64 `bson:"50-Day > 150-Day > 200-Day"`
	CurrentDaySVolumeGreaterThanPrevious5DaysVolume  float64 `bson:"Current day's Volume greater than previous 5 days' Volume"`
	CurrentDaySVolumeGreaterThanPrevious10DaysVolume float64 `bson:"Current day's Volume greater than previous 10 days' Volume"`
	CurrentDaySVolumeGreaterThanPrevious20DaysVolume float64 `bson:"Current day's Volume greater than previous 20 days' Volume"`
	ShortVolume                                      float64 `bson:"Short Volume"`
	Shrtfloat64OfFloat                               float64 `bson:"Shrt float64 % of Float"`
	DaysVolShortCurrent                              float64 `bson:"Days Vol Short Current"`
	DaysVolShort1PeriodAgo                           float64 `bson:"Days Vol Short 1 Period Ago"`
	DaysVolShort2PeriodsAgo                          float64 `bson:"Days Vol Short 2 Periods Ago"`
	Shrtfloat64Chg                                   float64 `bson:"Shrt float64 % Chg"`
}

const startloopcount = 1195
const row_input = 2083

var startingpoint int
var endingpoint int
var cell = string("A2")

// func (t Tri) iporeceiver() string {
// 	return string(t.IPODate)
// }

func main() {
	// open excel file
	//f := excelize.NewFile()
	f, err := excelize.OpenFile("./TSDB.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set Client Options
	clientOptions := options.Client().ApplyURI("mongodb://192.168.1.62:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Ping Server to COnfirm Connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	fmt.Printf("Please input where youd like to start in loop = ")
	fmt.Scanf("%d", &startingpoint)
	//fmt.Printf("Please input endingpoint                      = ")
	//fmt.Scanf("%d", &endingpoint)
	//fmt.Println(startingpoint, endingpoint)

	//Crreate connection object to collection of choice.
	collections := client.Database("Test").Collection("Trisight")

	/*SAMPLE DATA*/
	//ash := Trainer{"Ash", 10, "Pallet Town", []string{"In Pallet", "Out Pallet", "Around Pallet"}}
	//misty := Trainer{"Misty", 10, "Cerulean City", []string{"In Pallet", "Out Pallet", "Around Pallet"}}
	//brock := Trainer{"Brock", 15, "Pewter City", []string{"In Pallet", "Out Pallet", "Around Pallet"}}

	// insertResult, err := collections.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Prfloat64ln("Inserted a single document: ", insertResult.InsertedID)

	// //Insert Multiple Documents
	// trainers := []float64erface{}{misty, brock}

	// insertMultiResult, err := collections.InsertMany(context.TODO(), trainers)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Prfloat64ln("Insert multiple documents: ", insertMultiResult.InsertedIDs)

	//Updating Documents
	//filter := bson.D{{"name", "Ash"}}
	for z := startingpoint; z < row_input; z++ {
		// Get value from cell by given worksheet name and axis.
		cell := fmt.Sprintf("A"+"%v", z)
		fmt.Println(cell)
		celldbinput, err := f.GetCellValue("Export", cell)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("This is cell=", celldbinput)

		filter := bson.D{{"Symbol", celldbinput}}
		// update := bson.D{
		// 	{"$set", bson.D{
		// 		{"test2", []string{"new words", "old words", "middle words"}},
		// 	}},
		// }
		// updateResult, err := collections.UpdateOne(context.TODO(), filter, update)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Prfloat64f("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

		//Retrieve the value of a Document
		var result Tri
		err = collections.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found a single document: %+v\n", result.Name)
	}
	// err = client.Disconnect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Prfloat64ln("Connection to MongoDB closed.")

}
