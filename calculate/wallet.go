package calculate

import (
	"errors"
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WalletCalculate struct {
	Repo     *repository.Repositories
	Limit    int64
	Stopped  chan os.Signal
	Interval *time.Ticker
}

func NewWalletCalculate(repos *repository.Repositories, limit int64, timer *time.Ticker) *WalletCalculate {
	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	return &WalletCalculate{
		Repo:     repos,
		Limit:    limit,
		Stopped:  quitChannel,
		Interval: timer,
	}
}

func (cw *WalletCalculate) DoPostEvent(w, t int64, event string) (int64, error) {
	var result int64
	var err error

	switch event {
	case "win":
		result = w + t
	case "lost":
		result = w - t
	}

	if result < 0 {
		return 0, errors.New("ERROR - User wallet can't be negative")
	} else {
		return result, nil
	}

	return result, err
}

func (cw *WalletCalculate) BackgroundWork() {
	for {
		select {
		case t := <-cw.Interval.C:
			logrus.Debugf("Work start at %v", t.String())
			users, err := cw.Repo.Users.Collection(nil, nil)
			if err != nil {
				logrus.Debugf("Background work failed ", err)
				break
			}
			if users != nil {
				for _, user := range users {
					cw.CancelledTransactions(user)
				}
				break
			} else {
				logrus.Debugf("No users for background work")
				break
			}
		case <-cw.Stopped:
			logrus.Debugf("Exit background")
			break
		}
	}
}

func (cw *WalletCalculate) CancelledTransactions(user *repository.UsersModel) {
	tQ := repository.NewTransactionQuery()
	tQ.UserGUID(user.GUID.String)
	tQ.Done()
	tQ.OrderById(true)

	allT, err := cw.Repo.Transactions.FindOdd(tQ, &cw.Limit)
	if err != nil {
		logrus.Debugf("For user %v not exist last %v transaction - try later", user.GUID.String, cw.Limit)
		return
	} else if int64(len(allT)) != cw.Limit {
		logrus.Debugf("For user %v not exist last %v transaction - try later", user.GUID.String, cw.Limit)
		return
	} else {
		newW, ok := checkDiff(allT, user.Wallet.Int64)
		if ok {
			if ok := cw.updateTransactions(allT); ok {
				logrus.Debugf("Transactions cancelled succeed %v", user.GUID.String)
				user.Wallet = null.IntFrom(cast.ToInt64(newW))
				_, err := cw.Repo.Users.Update(user.GUID.String, user)
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

func (cw *WalletCalculate) updateTransactions(t []*repository.TransactionModel) (ok bool) {
	ok = true
	for _, item := range t {
		item.Done = null.BoolFrom(false)
		_, err := cw.Repo.Transactions.Update(item.GUID.String, item)
		if err != nil {
			ok = false
			return
		}
	}
	return
}
