package naturaldate

import (
	"log"
	"testing"
	"time"

	"github.com/tj/assert"
)

// base time.
var base = time.Unix(1574687238, 0).UTC()

// pastCases are test cases for the past direction.
var pastCases = []struct {
	Input    string
	Output   string
	ExprType ExprType
}{
	// now
	{`now`, `2019-11-25 13:07:18 +0000 UTC`, ExprTypeTime | ExprTypeDate},
	{`right now`, `2019-11-25 13:07:18 +0000 UTC`, ExprTypeTime | ExprTypeDate},
	{`  right  now  `, `2019-11-25 13:07:18 +0000 UTC`, ExprTypeTime | ExprTypeDate},

	// minutes
	{`1 minute`, `2019-11-25 13:06:18 +0000 UTC`, ExprTypeTime},
	{`next minute`, `2019-11-25 13:08:18 +0000 UTC`, ExprTypeTime},
	{`last minute`, `2019-11-25 13:06:18 +0000 UTC`, ExprTypeTime},
	{`one minute`, `2019-11-25 13:06:18 +0000 UTC`, ExprTypeTime},
	{`1 minute ago`, `2019-11-25 13:06:18 +0000 UTC`, ExprTypeTime},
	{`5 minutes ago`, `2019-11-25 13:02:18 +0000 UTC`, ExprTypeTime},
	{`five minutes ago`, `2019-11-25 13:02:18 +0000 UTC`, ExprTypeTime},
	{`   5    minutes  ago   `, `2019-11-25 13:02:18 +0000 UTC`, ExprTypeTime},
	{`2 minutes from now`, `2019-11-25 13:09:18 +0000 UTC`, ExprTypeTime},
	{`two minutes from now`, `2019-11-25 13:09:18 +0000 UTC`, ExprTypeTime},
	{`Message me in 2 minutes`, `2019-11-25 13:09:18 +0000 UTC`, ExprTypeTime},
	{`Message me in 2 minutes from now`, `2019-11-25 13:09:18 +0000 UTC`, ExprTypeTime},

	// hours
	{`1 hour`, `2019-11-25 12:07:18 +0000 UTC`, ExprTypeTime},
	{`last hour`, `2019-11-25 12:07:18 +0000 UTC`, ExprTypeTime},
	{`next hour`, `2019-11-25 14:07:18 +0000 UTC`, ExprTypeTime},
	{`1 hour ago`, `2019-11-25 12:07:18 +0000 UTC`, ExprTypeTime},
	{`6 hours ago`, `2019-11-25 07:07:18 +0000 UTC`, ExprTypeTime},
	{`1 hour from now`, `2019-11-25 14:07:18 +0000 UTC`, ExprTypeTime},
	{`Remind me in 1 hour`, `2019-11-25 14:07:18 +0000 UTC`, ExprTypeTime},
	{`Remind me in 1 hour from now`, `2019-11-25 14:07:18 +0000 UTC`, ExprTypeTime},
	{`Remind me in 1 hour and 3 minutes from now`, `2019-11-25 14:10:18 +0000 UTC`, ExprTypeTime},
	{`Remind me in an hour`, `2019-11-25 14:07:18 +0000 UTC`, ExprTypeTime},
	{`Remind me in an hour from now`, `2019-11-25 14:07:18 +0000 UTC`, ExprTypeTime},

	// days
	{`1 day`, `2019-11-24 00:00:00 +0000 UTC`, ExprTypeDate},
	{`next day`, `2019-11-26 00:00:00 +0000 UTC`, ExprTypeDate},
	{`1 day ago`, `2019-11-24 00:00:00 +0000 UTC`, ExprTypeDate},
	{`3 days ago`, `2019-11-22 00:00:00 +0000 UTC`, ExprTypeDate},
	{`3 days ago at 11:25am`, `2019-11-22 11:25:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`1 day from now`, `2019-11-26 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me one day from now`, `2019-11-26 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in a day`, `2019-11-26 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in one day`, `2019-11-26 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in one day from now`, `2019-11-26 13:07:18 +0000 UTC`, ExprTypeDate},

	// weeks
	{`1 week`, `2019-11-18 00:00:00 +0000 UTC`, ExprTypeDate},
	{`1 week ago`, `2019-11-18 00:00:00 +0000 UTC`, ExprTypeDate},
	{`2 weeks ago`, `2019-11-11 00:00:00 +0000 UTC`, ExprTypeDate},
	{`2 weeks ago at 8am`, `2019-11-11 08:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`next week`, `2019-12-02 00:00:00 +0000 UTC`, ExprTypeDate},
	{`Message me in a week`, `2019-12-02 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Message me in one week`, `2019-12-02 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Message me in one week from now`, `2019-12-02 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Message me in two weeks from now`, `2019-12-09 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Message me two weeks from now`, `2019-12-09 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Message me in two weeks`, `2019-12-09 13:07:18 +0000 UTC`, ExprTypeDate},

	// months
	{`1 month ago`, `2019-10-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`last month`, `2019-10-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`next month`, `2019-12-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`1 month ago at 9:30am`, `2019-10-25 09:30:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`2 months ago`, `2019-09-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`12 months ago`, `2018-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`1 month from now`, `2019-12-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`next 2 months`, `2020-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`2 months from now`, `2020-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`12 months from now at 6am`, `2020-11-25 06:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`Remind me in 12 months from now at 6am`, `2020-11-25 06:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`Remind me in a month`, `2019-12-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in 2 months`, `2020-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in a month from now`, `2019-12-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in 2 months from now`, `2020-01-25 13:07:18 +0000 UTC`, ExprTypeDate},

	// years
	{`last year`, `2018-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`next year`, `2020-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`one year ago`, `2018-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`one year from now`, `2020-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`two years ago`, `2017-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`2 years ago`, `2017-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in one year from now`, `2020-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in a year`, `2020-11-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in a year from now`, `2020-11-25 13:07:18 +0000 UTC`, ExprTypeDate},

	// today
	{`today`, `2019-11-25 00:00:00 +0000 UTC`, ExprTypeDate},
	{`today at 10am`, `2019-11-25 10:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},

	// yesterday
	{`yesterday`, `2019-11-24 00:00:00 +0000 UTC`, ExprTypeDate},
	{`yesterday 10am`, `2019-11-24 10:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`yesterday at 10am`, `2019-11-24 10:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`yesterday at 10:15am`, `2019-11-24 10:15:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},

	// tomorrow
	{`tomorrow`, `2019-11-26 00:00:00 +0000 UTC`, ExprTypeDate},
	{`tomorrow 10am`, `2019-11-26 10:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`tomorrow at 10am`, `2019-11-26 10:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`tomorrow at 10:15am`, `2019-11-26 10:15:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},

	// past weekdays
	{`sunday`, `2019-11-24 00:00:00 +0000 UTC`, ExprTypeDate},
	{`monday`, `2019-11-18 00:00:00 +0000 UTC`, ExprTypeDate},
	{`tuesday`, `2019-11-19 00:00:00 +0000 UTC`, ExprTypeDate},
	{`wednesday`, `2019-11-20 00:00:00 +0000 UTC`, ExprTypeDate},
	{`thursday`, `2019-11-21 00:00:00 +0000 UTC`, ExprTypeDate},
	{`friday`, `2019-11-22 00:00:00 +0000 UTC`, ExprTypeDate},
	{`saturday`, `2019-11-23 00:00:00 +0000 UTC`, ExprTypeDate},

	{`last sunday`, `2019-11-24 00:00:00 +0000 UTC`, ExprTypeDate},
	{`past sunday`, `2019-11-24 00:00:00 +0000 UTC`, ExprTypeDate},
	{`last monday`, `2019-11-18 00:00:00 +0000 UTC`, ExprTypeDate},
	{`last tuesday`, `2019-11-19 00:00:00 +0000 UTC`, ExprTypeDate},
	{`last wednesday`, `2019-11-20 00:00:00 +0000 UTC`, ExprTypeDate},
	{`last thursday`, `2019-11-21 00:00:00 +0000 UTC`, ExprTypeDate},
	{`last friday`, `2019-11-22 00:00:00 +0000 UTC`, ExprTypeDate},
	{`last saturday`, `2019-11-23 00:00:00 +0000 UTC`, ExprTypeDate},

	// future weekdays
	{`next tuesday`, `2019-11-26 00:00:00 +0000 UTC`, ExprTypeDate},
	{`next wednesday`, `2019-11-27 00:00:00 +0000 UTC`, ExprTypeDate},
	{`next thursday`, `2019-11-28 00:00:00 +0000 UTC`, ExprTypeDate},
	{`next friday`, `2019-11-29 00:00:00 +0000 UTC`, ExprTypeDate},
	{`next saturday`, `2019-11-30 00:00:00 +0000 UTC`, ExprTypeDate},
	{`next sunday`, `2019-12-01 00:00:00 +0000 UTC`, ExprTypeDate},
	{`next monday`, `2019-12-02 00:00:00 +0000 UTC`, ExprTypeDate},

	// months
	{`last january`, `2019-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`next january`, `2020-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`january`, `2019-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`february`, `2019-02-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`march`, `2019-03-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`april`, `2019-04-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`may`, `2019-05-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`june`, `2019-06-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`july`, `2019-07-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`august`, `2019-08-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`september`, `2019-09-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`october`, `2019-10-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`november`, `2018-11-25 13:07:18 +0000 UTC`, ExprTypeDate},

	// ordinal dates
	{`november 15th`, `2018-11-15 13:07:18 +0000 UTC`, ExprTypeDate},
	{`december 1st`, `2018-12-01 13:07:18 +0000 UTC`, ExprTypeDate},
	{`december 2nd`, `2018-12-02 13:07:18 +0000 UTC`, ExprTypeDate},
	{`december 3rd`, `2018-12-03 13:07:18 +0000 UTC`, ExprTypeDate},
	{`december 4th`, `2018-12-04 13:07:18 +0000 UTC`, ExprTypeDate},
	{`december 15th`, `2018-12-15 13:07:18 +0000 UTC`, ExprTypeDate},
	{`december 23rd`, `2018-12-23 13:07:18 +0000 UTC`, ExprTypeDate},
	{`december 23rd 5pm`, `2018-12-23 17:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`december 23rd at 5pm`, `2018-12-23 17:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`december 23rd at 5:25pm`, `2018-12-23 17:25:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},

	// 12-hour clock
	{`10am`, `2019-11-25 10:00:00 +0000 UTC`, ExprTypeTime},
	{`10 am`, `2019-11-25 10:00:00 +0000 UTC`, ExprTypeTime},
	{`5pm`, `2019-11-25 17:00:00 +0000 UTC`, ExprTypeTime},
	{`10:25am`, `2019-11-25 10:25:00 +0000 UTC`, ExprTypeTime},
	{`1:05pm`, `2019-11-25 13:05:00 +0000 UTC`, ExprTypeTime},
	{`10:25:10am`, `2019-11-25 10:25:10 +0000 UTC`, ExprTypeTime},
	{`1:05:10pm`, `2019-11-25 13:05:10 +0000 UTC`, ExprTypeTime},

	// 24-hour clock
	{`10`, `2019-11-25 10:00:00 +0000 UTC`, ExprTypeTime},
	{`10:25`, `2019-11-25 10:25:00 +0000 UTC`, ExprTypeTime},
	{`10:25:30`, `2019-11-25 10:25:30 +0000 UTC`, ExprTypeTime},
	{`17`, `2019-11-25 17:00:00 +0000 UTC`, ExprTypeTime},
	{`17:25:30`, `2019-11-25 17:25:30 +0000 UTC`, ExprTypeTime},

	// case sensitivity
	{`December 23rd AT 5:25 PM`, `2018-12-23 17:25:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`next December 23rd AT 5:25 PM`, `2019-12-23 17:25:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},

	// QA
	{`Restart the server in 2 days from now`, `2019-11-27 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me on the 5th of next month`, `2019-12-05 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me on the 5th of next month at 7am`, `2019-12-05 07:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`Remind me at 7am on the 5th of next month`, `2019-12-05 07:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`Remind me in one month from now`, `2019-12-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me in one month from now at 7am`, `2019-12-25 07:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},

	// errors
	{`10:am`, "\nparse error near PegText (line 1 symbol 1 - line 1 symbol 3):\n\"10\"\n", ExprTypeInvalid},
}

// futureCases are test cases for the future direction.
var futureCases = []struct {
	Input    string
	Output   string
	ExprType ExprType
}{
	{`now`, `2019-11-25 13:07:18 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`1 minute`, `2019-11-25 13:08:18 +0000 UTC`, ExprTypeTime},
	{`1 hour`, `2019-11-25 14:07:18 +0000 UTC`, ExprTypeTime},
	{`1 day`, `2019-11-26 00:00:00 +0000 UTC`, ExprTypeDate},
	{`1 week`, `2019-12-02 00:00:00 +0000 UTC`, ExprTypeDate},
	{`previous tuesday`, `2019-11-19 00:00:00 +0000 UTC`, ExprTypeDate},
	{`tuesday`, `2019-11-26 00:00:00 +0000 UTC`, ExprTypeDate},
	{`wednesday`, `2019-11-27 00:00:00 +0000 UTC`, ExprTypeDate},
	{`thursday`, `2019-11-28 00:00:00 +0000 UTC`, ExprTypeDate},
	{`friday`, `2019-11-29 00:00:00 +0000 UTC`, ExprTypeDate},
	{`saturday`, `2019-11-30 00:00:00 +0000 UTC`, ExprTypeDate},
	{`sunday`, `2019-12-01 00:00:00 +0000 UTC`, ExprTypeDate},
	{`monday`, `2019-12-02 00:00:00 +0000 UTC`, ExprTypeDate},
	{`last january`, `2019-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`january`, `2020-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`next january`, `2020-01-25 13:07:18 +0000 UTC`, ExprTypeDate},
	{`Remind me on the December 25th at 7am`, `2019-12-25 07:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`Remind me at 7am on December 25th`, `2019-12-25 07:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`Remind me on the 25th of December at 7am`, `2019-12-25 07:00:00 +0000 UTC`, ExprTypeDate | ExprTypeTime},
	{`Check logs in the past 5 minutes`, `2019-11-25 13:02:18 +0000 UTC`, ExprTypeTime},
}

// Test parsing with past direction.
func TestParse_past(t *testing.T) {
	for _, c := range pastCases {
		t.Run(c.Input, func(t *testing.T) {
			v, exprType, err := Parse(c.Input, base)
			if err != nil {
				assert.Equal(t, c.Output, err.Error())
				return
			}
			assert.Equal(t, c.ExprType, exprType)
			assert.Equal(t, c.Output, v.UTC().String())
		})
	}
}

// Test parsing with future direction.
func TestParse_future(t *testing.T) {
	for _, c := range futureCases {
		t.Run(c.Input, func(t *testing.T) {
			v, exprType, err := Parse(c.Input, base, WithDirection(Future))
			if err != nil {
				assert.Equal(t, c.Output, err.Error())
				return
			}
			assert.Equal(t, c.ExprType, exprType)
			assert.Equal(t, c.Output, v.UTC().String())
		})
	}
}

// Benchmark parsing.
func BenchmarkParse(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		_, _, err := Parse(`december 23rd at 5:25pm`, base)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}
}
