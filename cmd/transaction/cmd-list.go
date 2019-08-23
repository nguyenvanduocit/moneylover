package transaction

import (
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/nguyenvanduocit/moneylover/core"
	"regexp"
	"strconv"
	"time"
)

var since *time.Time
var until *time.Time
var token string

var ListCmd = &cobra.Command{
	Use:   "list",
	Example: "moneylover transaction list: lấy các transaction trong vòng 1 tháng\nmoneylover transaction list --since 2018: lấy transaction bắt đầu từ ngày 1 tháng 1 năm 2018 cho đến nay\nmoneylover transaction list --since 2018-01-01 --until 2019-08-01: đoán xem",
	Short: "List all transaction",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		token = viper.GetString("token")
		if token == "" {
			return errors.New("token is required")
		}

		inputUntil, err := cmd.Flags().GetString("until")
		if err != nil {
			return err
		}
		if inputUntil == "" {
			until = timePointer(time.Now())
		} else {
			until, err = parseTime(inputUntil)
			if err != nil {
				return err
			}
		}

		inputSince, err := cmd.Flags().GetString("since")
		if err != nil {
			return err
		}
		if inputSince == "" {
			since = timePointer(until.AddDate(0, -1, 0))
		} else {
			since, err = parseTime(inputSince)
			if err != nil {
				return err
			}
		}

		until = timePointer(until.Add(23*time.Hour + 59 * time.Minute + 59 * time.Second))

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string)error{
		transactions, err := core.GetTransactions(token, since, until)
		if err != nil {
			return err
		}
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.SetHeader([]string{"Date", "Category", "Amount", "Note"})
		table.SetCaption(true, fmt.Sprintf("From %s to %s", since.Format("2006-01-02"), until.Format("2006-01-02")))
		total := 0
		for _, transaction := range transactions {
			table.Append([]string{
				transaction.DisplayDate.Format("2006-01-02"),
				transaction.Category.Name,
				strconv.Itoa(transaction.Amount),
				transaction.Note,
			})
			total = total + transaction.Amount
		}
		table.SetFooter([]string{"", "Total", strconv.Itoa(total), ""})
		table.Render()
		return nil
	},
}

func init () {
	RootCmd.AddCommand(ListCmd)
	ListCmd.Flags().String("since", "", "YYYY-MM-DD hoặc YYYY-MM hoặc YYYY")
	ListCmd.Flags().String("until", "", "YYYY-MM-DD hoặc YYYY-MM hoặc YYYY")
}

func parseTime (datetime string) (*time.Time, error) {
	now := time.Now()
	var re = regexp.MustCompile(`(?m)^([\d]{4})-?([\d]{0,2})-?([\d]{0,2})$`)
	matches := re.FindAllStringSubmatch(datetime, -1)

	if len(matches) == 0 || len(matches[0]) < 2 {
		return nil, errors.New("can not parse " + datetime +" as time")
	}

	datetimeParts := matches[0]
	datetimeLength := len(datetimeParts)

	var year, month, day int

	if datetimeLength >= 2 && datetimeParts[1] != ""{
		 year64, err := strconv.ParseInt(datetimeParts[1], 10, 32)
		 if err != nil {
		 	return nil, err
		 }
		year = int(year64)
	} else {
		return nil, errors.New("year is required")
	}

	if datetimeLength >= 3 && datetimeParts[2] != ""{
		month64, err := strconv.ParseInt(datetimeParts[2], 10, 32)
		if err != nil {
			return nil, err
		}
		month = int(month64)
	} else {
		month = 1
	}

	if datetimeLength >= 4 && datetimeParts[3] != ""{
		day64, err := strconv.ParseInt(datetimeParts[3], 10, 32)
		if err != nil {
			return nil, err
		}
		day = int(day64)
	} else {
		day = 1
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, now.Location())
	return &date, nil
}

func timePointer(time time.Time) *time.Time {
	return &time
}