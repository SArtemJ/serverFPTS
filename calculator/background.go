package calculator

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BackgroundWorker struct {
	Repo     *repository.Repositories
	Limit    int64
	Stopped  chan os.Signal
	Interval *time.Ticker
}

func NewBackgroundWorker(repos *repository.Repositories, limit int64, timer *time.Ticker) *BackgroundWorker {
	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	return &BackgroundWorker{
		Repo:     repos,
		Limit:    limit,
		Stopped:  quitChannel,
		Interval: timer,
	}
}

func (bw *BackgroundWorker) Run() {
	for {
		select {
		case t := <-bw.Interval.C:
			logrus.Debugf("Work start at %v", t.String())
			users, err := bw.Repo.Users.Collection(nil, nil)
			if err != nil {
				logrus.Debugf("Background work failed - %v", err)
				break
			}
			if users != nil {
				for _, user := range users {
					bw.CancelledTransactions(user)
				}
				break
			} else {
				logrus.Debug("No users for background work")
				break
			}
		case <-bw.Stopped:
			logrus.Debug("Exit background")
			break
		}
	}
}

func (bw *BackgroundWorker) CancelledTransactions(user *repository.UsersModel) {
	tQ := repository.NewTransactionQuery()
	tQ.UserGUID(user.GUID.String)
	tQ.Done()
	tQ.OrderById(true)

	allT, err := bw.Repo.Transactions.FindOdd(tQ, &bw.Limit)
	if err != nil {
		logrus.Debugf("For user %v not exist last %v transaction - try later", user.GUID.String, bw.Limit)
		return
	} else if int64(len(allT)) != bw.Limit {
		logrus.Debugf("For user %v not exist last %v transaction - try later", user.GUID.String, bw.Limit)
		return
	} else {
		newW, ok := checkDiff(allT, user.Wallet.Int64)
		if ok {
			if ok := bw.updateTransactions(allT); ok {
				logrus.Debugf("Transactions cancelled succeed %v", user.GUID.String)
				user.Wallet = null.IntFrom(cast.ToInt64(newW))
				_, err := bw.Repo.Users.Update(user.GUID.String, user)
				if err != nil {
					logrus.Debugf("Failed update user wallet %v", user.GUID.String)
					return
				}
			}
		}
	}
}

func checkDiff(allT []*repository.TransactionModel, u int64) (r *int64, ok bool) {
	var allLost int64
	var allWin int64
	for _, item := range allT {
		switch item.State.String {
		case "win":
			allWin += item.Amount.Int64
		case "lost":
			allLost += item.Amount.Int64
		}
	}

	result := u + allLost - allWin
	if result < 0 {
		return
	} else {
		r = &result
		ok = true
	}
	return
}

func (bw *BackgroundWorker) updateTransactions(t []*repository.TransactionModel) (ok bool) {
	ok = true
	for _, item := range t {
		item.Done = null.BoolFrom(false)
		_, err := bw.Repo.Transactions.Update(item.GUID.String, item)
		if err != nil {
			ok = false
			return
		}
	}
	return
}
