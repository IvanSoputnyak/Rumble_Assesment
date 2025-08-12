package transactions

import (
	"sort"
	"time"
)

func OrderByPostedTimestampDesc(transactions []Transaction) []Transaction {
	sortedTransactions := append([]Transaction(nil), transactions...)

	sort.SliceStable(sortedTransactions, func(i, j int) bool {
		timeI, parseErrI := time.Parse(time.RFC3339, sortedTransactions[i].PostedTimeStamp)
		timeJ, parseErrJ := time.Parse(time.RFC3339, sortedTransactions[j].PostedTimeStamp)

		switch {
		case parseErrI == nil && parseErrJ == nil:
			return timeI.After(timeJ)

		case parseErrI == nil && parseErrJ != nil:
			return true

		case parseErrI != nil && parseErrJ == nil:
			return false

		default:
			return false
		}
	})

	return sortedTransactions
}
