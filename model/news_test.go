package model

import (
	_ "github.com/go-sql-driver/mysql"
)

// func TestNewsItem_InsertNewsItem(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	item := NewsItem{
// 		Title:       null.StringFrom("东芝存储芯片出售难产 债权银行施压其尽快完成交易"),
// 		Link:        null.StringFrom("http://www.chinaflashmarket.com/Producer/TOSHIBA/News/160860"),
// 		Type:        null.StringFrom("厂商动态"),
// 		Description: null.StringFrom("据《金融时报》北京时间4月9日报道，银行业人士消息称，尽管有维权股东认为东芝存储芯片业务实际价值是出售价的两倍以上，但是东芝主要债权银行正督促东芝推进这笔2万亿日元(约合187亿美元)的出售交易。"),
// 		Body:        null.StringFrom(""),
// 		Source:      null.StringFrom("来源：凤凰科技"),
// 	}

// 	mock.ExpectExec("Insert news item").WillReturnResult(sqlmock.NewResult(1, 1))

// 	_, err = item.InsertNewsItem(db)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
