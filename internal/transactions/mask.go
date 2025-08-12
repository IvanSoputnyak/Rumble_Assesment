package transactions

import "strconv"

func maskPAN(pan int) string {
	s := strconv.FormatInt(int64(pan), 10)
	n := len(s)
	if n <= 4 {
		return s
	}
	maskLen := n - 4
	buf := make([]byte, 0, n)
	for i := 0; i < maskLen; i++ {
		buf = append(buf, '*')
	}
	return string(buf) + s[n-4:]
}

func MaskAll(input []Transaction) []TransactionMasked {
	out := make([]TransactionMasked, len(input))
	for i, t := range input {
		out[i] = TransactionMasked{
			ID:                  t.ID,
			Amount:              t.Amount,
			MessageType:         t.MessageType,
			CreatedAt:           t.CreatedAt,
			TransactionID:       t.TransactionID,
			PAN:                 maskPAN(t.PAN),
			TransactionCategory: t.TransactionCategory,
			PostedTimeStamp:     t.PostedTimeStamp,
			TransactionType:     t.TransactionType,
			SendingAccount:      t.SendingAccount,
			ReceivingAccount:    t.ReceivingAccount,
			TransactionNote:     t.TransactionNote,
		}
	}
	return out
}
