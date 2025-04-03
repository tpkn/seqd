package seqd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_parseDate(t *testing.T) {
	_, _, err := parseDate("")
	require.ErrorContains(t, err, "can't parse date")

	_, _, err = parseDate("abcde")
	require.ErrorContains(t, err, "can't parse date")

	_, _, err = parseDate("2023")
	require.ErrorContains(t, err, "can't parse date")

	_, _, err = parseDate("20/12 12:30:00PM")
	require.ErrorContains(t, err, "can't parse date")

	_, _, err = parseDate("2025-13-32")
	require.ErrorContains(t, err, "can't parse date")

	// ------------------------------

	test_date, _ := time.Parse(time.DateTime, "2025-02-01 12:34:56")

	got, format, err := parseDate("2025-02-01")
	require.NoError(t, err)
	require.Equal(t, got.Format(time.DateOnly), test_date.Format(time.DateOnly))
	require.Equal(t, format, time.DateOnly)

	got, format, err = parseDate("2025-02-01 12:34:56")
	require.NoError(t, err)
	require.Equal(t, got.Format(time.DateTime), test_date.Format(time.DateTime))
	require.Equal(t, format, time.DateTime)
}

func Test_getDateRangeBounds(t *testing.T) {
	_, err := getDateRangeBounds("", "")
	require.Error(t, err)

	_, err = getDateRangeBounds("2024-01-01", "")
	require.Error(t, err)

	_, err = getDateRangeBounds("", "2024-01-01")
	require.Error(t, err)

	_, err = getDateRangeBounds("2024-01-01", "2024")
	require.Error(t, err)

	_, err = getDateRangeBounds("9000-01-01", "2023-02-02")
	require.ErrorContains(t, err, "start date is greater than end date")

	_, err = getDateRangeBounds("2024-01-01", "2024-02-02 23:59:59")
	require.ErrorContains(t, err, "start date and end date has different format")

	_, err = getDateRangeBounds("eom", "2023-02-02")
	require.ErrorContains(t, err, "'eom' macros can only be used in end date")

	_, err = getDateRangeBounds("eoy", "2023-02-02")
	require.ErrorContains(t, err, "'eoy' macros can only be used in end date")

	// ------------------------------

	got, err := getDateRangeBounds("2024-01-01", "2024-02-02")
	require.NoError(t, err)
	require.Equal(t, got, DateRange{Start: mockParseDate("2024-01-01", time.DateOnly), End: mockParseDate("2024-02-02", time.DateOnly), Format: time.DateOnly})

	got, err = getDateRangeBounds("2024-01-01 05:13:59", "2024-02-02 23:59:59")
	require.NoError(t, err)
	require.Equal(t, got, DateRange{Start: mockParseDate("2024-01-01 05:13:59", time.DateTime), End: mockParseDate("2024-02-02 23:59:59", time.DateTime), Format: time.DateTime})

	got, err = getDateRangeBounds("2024-02-01", "eom")
	require.NoError(t, err)
	require.Equal(t, got, DateRange{Start: mockParseDate("2024-02-01", time.DateOnly), End: mockParseDate("2024-02-29", time.DateOnly), Format: time.DateOnly})

	got, err = getDateRangeBounds("2024-02-01 05:13:59", "eom")
	require.NoError(t, err)
	require.Equal(t, got, DateRange{Start: mockParseDate("2024-02-01 05:13:59", time.DateTime), End: mockParseDate("2024-02-29 23:59:59", time.DateTime), Format: time.DateTime})

	got, err = getDateRangeBounds("2024-01-01", "eoy")
	require.NoError(t, err)
	require.Equal(t, got, DateRange{Start: mockParseDate("2024-01-01", time.DateOnly), End: mockParseDate("2024-12-31", time.DateOnly), Format: time.DateOnly})

	got, err = getDateRangeBounds("2024-01-01 05:13:59", "eoy")
	require.NoError(t, err)
	require.Equal(t, got, DateRange{Start: mockParseDate("2024-01-01 05:13:59", time.DateTime), End: mockParseDate("2024-12-31 23:59:59", time.DateTime), Format: time.DateTime})
}

func Test_getEndOfMonthDate(t *testing.T) {
	test_date, _ := time.Parse(time.DateTime, "2025-02-01 12:34:56")

	got := getEndOfMonthDate(test_date).Format(time.DateOnly)
	require.Equal(t, "2025-02-28", got)

	got = getEndOfMonthDate(test_date).Format(time.DateTime)
	require.Equal(t, "2025-02-28 23:59:59", got)
}

func Test_getAndOfYearDate(t *testing.T) {
	test_date, _ := time.Parse(time.DateTime, "2025-01-01 12:34:56")

	got := getAndOfYearDate(test_date).Format(time.DateOnly)
	require.Equal(t, "2025-12-31", got)

	got = getAndOfYearDate(test_date).Format(time.DateTime)
	require.Equal(t, "2025-12-31 23:59:59", got)
}

func Benchmark_parseDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseDate("2023-12-01 22:04:44")
	}
}

func Benchmark_getDateRangeBounds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDateRangeBounds("2023-12-01 22:04:44", "2024-12-01 22:04:44")
	}
}

func mockParseDate(s string, f string) time.Time {
	date, _ := time.Parse(f, s)
	return date
}
