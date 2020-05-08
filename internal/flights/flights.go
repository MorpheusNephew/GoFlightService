package flights

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var airlineURL = "https://www.spirit.com"
var availableDays = []*Day{}

// GetAvailableFlights returns data about what flights are
// available from Medellin to South Florida.
func GetAvailableFlights() (string, error) {
	availableDays = nil

	fmt.Println()
	fmt.Println("Getting available flights")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.Flag("headless", false),
		chromedp.WindowSize(1600, 1200),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	err := chromedp.Run(taskCtx,
		chromedp.Navigate(airlineURL),
		nagivateToCurrentMonthForFlights(),
		captureScreenshotWithName("CurrentMonth.jpg"),
		getCalendarDays(),
		navigateToNextMonthForFlights(),
		captureScreenshotWithName("NextMonth.jpg"),
		getCalendarDays(),
	)

	dt := time.Now()

	return fmt.Sprint(dt.Format("Jan-02-06 03:04:05 PM"), "\n", alertIfAvailableDays()), err
}

func alertIfAvailableDays() string {
	if !(len(availableDays) > 0) {
		return "No flights available"
	}

	var text = []string{"Flight(s) available!!!"}

	for _, availableDay := range availableDays {
		text = append(text, string(availableDay.text))
	}

	return strings.Join(text, "\n")
}

func captureScreenshot() chromedp.Tasks {
	return captureScreenshotInternal("Spirit.jpg")
}

func captureScreenshotInternal(nameOfFile string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			var imageData []byte

			chromedp.Run(ctx, chromedp.CaptureScreenshot(&imageData))

			filePath := path.Join("/", "Users", "jelani.jackson", "Desktop", "SpiritImages", nameOfFile)

			err := ioutil.WriteFile(filePath, imageData, 0644)

			if err != nil {
				log.Fatalln(err)
			}

			return err
		}),
	}
}

func captureScreenshotWithName(nameOfFile string) chromedp.Tasks {
	return captureScreenshotInternal(nameOfFile)
}

func getCalendarDays() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			var daysText []*String
			jsText := jsGetText("app-low-fare-day > button > div.ng-star-inserted")

			err := chromedp.Run(ctx,
				chromedp.Evaluate(jsText, &daysText),
			)

			days := ConvertToDays(daysText)

			for _, day := range days {
				flightAvailable := day.FlightAvailable()

				if flightAvailable {
					availableDays = append(availableDays, day)
				}
			}

			return err
		}),
	}
}

func jsGetText(sel string) string {
	const jsFunc = `
		function getText(selector) {
			var text = [];
			var elements = document.body.querySelectorAll(selector);
			var elementsLength = elements.length;
			var index = 0;

			for (; index < elementsLength ; index++) {
				var elementString = elements[index].innerText.trim();

				if (elementString.length > 0) {
					text.push(elementString);
				}
			}

			return text;
		}`

	invokeJs := `var a = getText('` + sel + `'); a;`

	return strings.Join([]string{jsFunc, invokeJs}, " ")
}

func nagivateToCurrentMonthForFlights() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.Run(ctx,
				chromedp.WaitEnabled("radio-oneWay", chromedp.ByID),
				chromedp.Click("radio-oneWay", chromedp.ByID),
				chromedp.SendKeys("flight-OriginStationCode", "MDE", chromedp.ByID),
				chromedp.SendKeys("flight-DestinationStationCode", "FLL", chromedp.ByID),
				chromedp.WaitVisible(".btn.btn-primary.d-block", chromedp.ByQuery),
				captureScreenshot(),
				chromedp.Click(".btn.btn-primary.d-block", chromedp.ByQuery),
				captureScreenshot(),
				chromedp.WaitReady("button.btn.btn-primary.mt-1", chromedp.ByQuery),
				chromedp.Click("button.btn.btn-primary.mt-1", chromedp.ByQuery),
				captureScreenshot(),
				chromedp.WaitReady("input[value='viewType.Month']", chromedp.ByQuery),
				chromedp.Click("input[value='viewType.Month']", chromedp.ByQuery),
				chromedp.Sleep(2*time.Second),
			)
		}),
	}
}

func navigateToNextMonthForFlights() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.Run(ctx,
				chromedp.Click("button[aria-label='next page']", chromedp.ByQuery),
				chromedp.Sleep(2*time.Second),
			)
		}),
	}
}
