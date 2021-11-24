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
		actual   *assertionItem
		expected *assertionItem
	}

	assertionItem struct {
		Result collector.Result
		Err    error
	}

	testCase struct {
		Name       string
		SetupData  *setupData
		Setup      func(f *setupData)
		AssertFn   func(t *testing.T, ad *assertionData)
		AssertData *assertionData
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
	api := NewAPI(opt, log)

	testCases := []testCase{
		{
			Name: "TestValidSmallRange",
			SetupData: setupData{
				API: api,
			},
			AssertFn: verifyAssertion,
			AssertData: &assertionData{
				expected: &assertionItem{
					Result: collector.Result{},
					Err:    nil,
				},
			},
		},
		{
			Name: "TestValidLongRange",
			SetupData: setupData{
				API: api,
			},
			AssertFn: verifyAssertion,
			AssertData: &assertionData{
				expected: &assertionItem{
					Result: collector.Result{},
					Err:    nil,
				},
			},
		},
		{
			Name: "TestToDateBefoereFromDate",
			SetupData: setupData{
				API: api,
			},
			AssertFn: verifyAssertion,
			AssertData: &assertionData{
				expected: &assertionItem{
					Result: collector.Result{},
					Err:    nil,
				},
			},
		},
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
			tc.Setup(&sd)
		}

		result, err := api.GetBetweenDates(sd.From, sd.To)

		tc.AssertFn(t, response)
	})
}

func verifyAssertion(t *testing.T, ad *assertionData) {
	t.Helper()

	eval0 := ad.actual.Result != ar.expected.Result
	eval1 := ad.actual.Err != ar.expected.Err

	if !(eval0 && eval1) {
		t.Errorf("received value '%+v' does not match expected '%+v'\n", ad.actual, ad.expected)
	}
}
