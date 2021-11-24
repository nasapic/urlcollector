package nasa

import (
	"testing"
	"time"

	"gitlab.com/QWRyaWFuIEdvR29BcHBzIE5BU0E/base"
	"gitlab.com/QWRyaWFuIEdvR29BcHBzIE5BU0E/urlcollector/pkg/collector"
)

type (
	setupData struct {
		API  *API
		From time.Time
		To   time.Time
	}

	assertionData struct {
		actual   assertionItem
		expected assertionItem
	}

	assertionItem struct {
		Result collector.Result
		Err    error
	}

	testCase struct {
		Name       string
		SetupData  *setupData
		Setup      func(f *setupData)
		AssertFn   func(t *testing.T, ad assertionData)
		AssertData assertionData
	}
)

var (
	opts = Options{
		APIKey:        "DEMO_KEY",
		TimeoutInSecs: 5,
		MaxConcurrent: 5,
	}

	log = base.NewLogger("error", "urlcollector-test", "json")
)

func TestSomething(t *testing.T) {
	api := NewAPI(opts, log)

	testCases := []testCase{
		{
			Name: "TestValidSmallRange",
			SetupData: &setupData{
				API:  api,
				From: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
			},
			AssertFn: verifyAssertion,
			AssertData: assertionData{
				expected: assertionItem{
					Result: &Result{
						list: []string{
							"https://apod.nasa.gov/apod/image/2001/OrionTrees123019Westlake1024.jpg",
							"https://apod.nasa.gov/apod/image/2001/BetelgeuseImagined_EsoCalcada_960.jpg",
						},
					},
					Err: nil,
				},
			},
		},
		// {
		// 	Name: "TestValidLongRange",
		// 	SetupData: setupData{
		// 		API: api,
		// 	},
		// 	AssertFn: verifyAssertion,
		// 	AssertData: &assertionData{
		// 		expected: &assertionItem{
		// 			Result: collector.Result{},
		// 			Err:    nil,
		// 		},
		// 	},
		// },
		// {
		// 	Name: "TestToDateBeforeFromDate",
		// 	SetupData: setupData{
		// 		API: api,
		// 	},
		// 	AssertFn: verifyAssertion,
		// 	AssertData: &assertionData{
		// 		expected: &assertionItem{
		// 			Result: collector.Result{},
		// 			Err:    nil,
		// 		},
		// 	},
		// },
	}
	runTests(t, testCases)
}

func runTests(t *testing.T, tcs []testCase) {
	for _, test := range tcs {
		runTest(t, test)
	}
}

func runTest(t *testing.T, tc testCase) {
	t.Run(tc.Name, func(t *testing.T) {
		api := tc.SetupData.API
		sd := tc.SetupData

		if tc.Setup != nil {
			tc.Setup(sd)
		}

		result, err := api.GetBetweenDates(sd.From, sd.To)

		tc.AssertData.actual = assertionItem{
			Result: result,
			Err:    err,
		}

		tc.AssertFn(t, tc.AssertData)
	})
}

func verifyAssertion(t *testing.T, ad assertionData) {
	t.Helper()

	if !(assertExpected(ad)) {
		t.Errorf("received value '%+v' does not match expected '%+v'\n", ad.actual, ad.expected)
	}
}

func assertExpected(ad assertionData) (ok bool) {
	eval0 := containsSameElements(ad.actual.Result.GetList(), ad.expected.Result.GetList())
	eval1 := ad.actual.Err == ad.expected.Err

	return eval0 && eval1
}

func containsSameElements(list, toCompare []string) bool {
	if len(list) != len(toCompare) {
		return false
	}

	exists := make(map[string]bool)
	for _, value := range list {
		exists[value] = true
	}

	for _, value := range toCompare {
		if !exists[value] {
			return false
		}
	}

	return true
}
