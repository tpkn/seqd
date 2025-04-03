package seqd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mockOutputWriter struct {
	result []string
}

func (o *mockOutputWriter) WriteString(s string) (int, error) {
	o.result = append(o.result, s)
	return 0, nil
}

func Test_PrintDateRange(t *testing.T) {
	got := mockOutputWriter{}
	err := GenerateDateRange(&got, &Args{StartDateTime: "", EndDateTime: ""})
	require.ErrorContains(t, err, "can't parse date")

	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "2025-01-03", IncreaseBySecond: true})
	require.ErrorContains(t, err, "flags '-h', '-m' and '-s' can only be used with")

	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-03", EndDateTime: "2025-01-01", IncreaseByDay: true})
	require.ErrorContains(t, err, "start date is greater than end date")

	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-03 23:59:59", EndDateTime: "2025-01-01", IncreaseByDay: true})
	require.ErrorContains(t, err, "start date and end date has different format")

	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-03", EndDateTime: "2025-01-01", IncreaseByDay: true, ReversedOrder: true})
	require.ErrorContains(t, err, "start date is greater than end date")

	err = GenerateDateRange(&got, &Args{StartDateTime: "eom", EndDateTime: "2025-01-01", IncreaseByDay: true})
	require.ErrorContains(t, err, "'eom' macros can only be used in end date")

	err = GenerateDateRange(&got, &Args{StartDateTime: "eoy", EndDateTime: "2025-01-01", IncreaseByDay: true})
	require.ErrorContains(t, err, "'eoy' macros can only be used in end date")

	// ------------------------------

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "2025-01-01", IncreaseByDay: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "2027-01-03", IncreaseByYear: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01", "2026-01-01", "2027-01-01"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-12-01", EndDateTime: "2026-02-03", IncreaseByMonth: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-12-01", "2026-01-01", "2026-02-01"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "2025-01-03", IncreaseByDay: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01", "2025-01-02", "2025-01-03"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 00:00:00", EndDateTime: "2025-01-01 02:00:00", IncreaseByHour: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01 00:00:00", "2025-01-01 01:00:00", "2025-01-01 02:00:00"}, got.result)

	// '00:59:59 + 1 hour' would go over '01:00:59' so the end time in range is '00:59:59'
	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 23:59:59", EndDateTime: "2025-01-02 01:00:59", IncreaseByHour: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01 23:59:59", "2025-01-02 00:59:59"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 23:59:00", EndDateTime: "2025-01-02 00:01:00", IncreaseByMinute: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01 23:59:00", "2025-01-02 00:00:00", "2025-01-02 00:01:00"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 23:59:59", EndDateTime: "2025-01-02 00:00:02", IncreaseBySecond: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01 23:59:59", "2025-01-02 00:00:00", "2025-01-02 00:00:01", "2025-01-02 00:00:02"}, got.result)

	// eom/eoy ----------------------

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "eom", IncreaseByYear: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "eom", IncreaseByMonth: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-29", EndDateTime: "eom", IncreaseByDay: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-29", "2025-01-30", "2025-01-31"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-31 21:34:59", EndDateTime: "eom", IncreaseByHour: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-31 21:34:59", "2025-01-31 22:34:59", "2025-01-31 23:34:59"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-31 23:57:00", EndDateTime: "eom", IncreaseByMinute: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-31 23:57:00", "2025-01-31 23:58:00", "2025-01-31 23:59:00"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-12-30", EndDateTime: "eoy", IncreaseByYear: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-12-30"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-12-29", EndDateTime: "eoy", IncreaseByDay: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-12-29", "2025-12-30", "2025-12-31"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-12-31 21:34:59", EndDateTime: "eoy", IncreaseByHour: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-12-31 21:34:59", "2025-12-31 22:34:59", "2025-12-31 23:34:59"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-12-31 23:57:00", EndDateTime: "eoy", IncreaseByMinute: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-12-31 23:57:00", "2025-12-31 23:58:00", "2025-12-31 23:59:00"}, got.result)

	// Reversed ---------------------

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "2025-01-01", IncreaseByDay: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "2027-01-03", IncreaseByYear: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2027-01-03", "2026-01-03", "2025-01-03"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-12-01", EndDateTime: "2026-02-03", IncreaseByMonth: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2026-02-03", "2026-01-03", "2025-12-03"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01", EndDateTime: "2025-01-03", IncreaseByDay: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-03", "2025-01-02", "2025-01-01"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 00:00:00", EndDateTime: "2025-01-01 02:00:00", IncreaseByHour: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-01 02:00:00", "2025-01-01 01:00:00", "2025-01-01 00:00:00"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 23:59:59", EndDateTime: "2025-01-02 01:00:59", IncreaseByHour: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-02 01:00:59", "2025-01-02 00:00:59"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 23:59:00", EndDateTime: "2025-01-02 00:01:00", IncreaseByMinute: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-02 00:01:00", "2025-01-02 00:00:00", "2025-01-01 23:59:00"}, got.result)

	got = mockOutputWriter{}
	err = GenerateDateRange(&got, &Args{StartDateTime: "2025-01-01 23:59:59", EndDateTime: "2025-01-02 00:00:02", IncreaseBySecond: true, ReversedOrder: true})
	require.NoError(t, err)
	require.Equal(t, []string{"2025-01-02 00:00:02", "2025-01-02 00:00:01", "2025-01-02 00:00:00", "2025-01-01 23:59:59"}, got.result)
}
