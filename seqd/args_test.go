package seqd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseArgs(t *testing.T) {
	_, err := ParseArgs([]string{"seqd"})
	require.ErrorContains(t, err, "wrong amount of arguments")

	_, err = ParseArgs([]string{"seqd", "--kek"})
	require.ErrorContains(t, err, "wrong argument")

	_, err = ParseArgs([]string{"seqd", "--kek", "--kek", "--kek", "--kek", "--kek"})
	require.ErrorContains(t, err, "wrong amount of arguments")

	_, err = ParseArgs([]string{"seqd", "-S", "2023-12-31", "2024-01-01"})
	require.ErrorContains(t, err, "unknown flag -S")

	_, err = ParseArgs([]string{"seqd", "-D", "2023-12-31", "2024-01-01", "-R"})
	require.ErrorContains(t, err, "unknown flag -R")

	// ------------------------------

	got, err := ParseArgs([]string{"seqd", "-Y", "2023-12-31", "2024-01-01"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByYear: true, ReversedOrder: false})

	got, err = ParseArgs([]string{"seqd", "-M", "2023-12-31", "2024-01-01"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByMonth: true, ReversedOrder: false})

	got, err = ParseArgs([]string{"seqd", "-D", "2023-12-31", "2024-01-01"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByDay: true, ReversedOrder: false})

	got, err = ParseArgs([]string{"seqd", "-h", "2023-12-31", "2024-01-01"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByHour: true, ReversedOrder: false})

	got, err = ParseArgs([]string{"seqd", "-m", "2023-12-31", "2024-01-01"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByMinute: true, ReversedOrder: false})

	got, err = ParseArgs([]string{"seqd", "-s", "2023-12-31", "2024-01-01"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseBySecond: true, ReversedOrder: false})

	got, err = ParseArgs([]string{"seqd", "-Y", "2023-12-31", "2024-01-01", "-r"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByYear: true, ReversedOrder: true})

	got, err = ParseArgs([]string{"seqd", "-Yr", "2023-12-31", "2024-01-01"})
	require.NoError(t, err)
	require.Equal(t, got, Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByYear: true, ReversedOrder: true})

	got, err = ParseArgs([]string{"seqd", "--help"})
	require.NoError(t, err)
	require.Equal(t, got, Args{Help: true})

	got, err = ParseArgs([]string{"seqd", "--version"})
	require.NoError(t, err)
	require.Equal(t, got, Args{Version: true})
}
