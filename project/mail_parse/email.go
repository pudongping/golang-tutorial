package mail_parse

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

const (
	Addr          string = "imap.qq.com:993"
	UserName      string = "123456789@qq.com" // 邮箱地址
	Password      string = ""                 // 这里的密码是使用开启 imap 协议后对应的服务商给到的密码，不是邮箱账号密码
	Folder        string = "INBOX"            // 邮箱文件夹，比如： INBOX 收件箱、Sent Messages 发件箱、Drafts 草稿箱、Trash、Junk 垃圾箱
	ReadBatchSize int    = 2                  // 每次读取的邮件数量
)

// IMAP（Internet Message Access Protocol）是一种用于在互联网上访问电子邮件的协议。
// 它允许用户通过 Internet 访问他们在邮件服务器上存储的电子邮件。
// Go 语言的 go-imap 库是一个用于从 IMAP 服务器获取电子邮件的库，它可以帮助你在 Go 代码中访问 IMAP 协议

func ReadEmail() {
	log.Println("开始连接服务器")

	// 建立与 IMAP 服务器的连接
	c, err := client.DialTLS(Addr, nil)
	if err != nil {
		log.Fatalf("连接 IMAP 服务器失败: %+v \n", err)
	}
	log.Println("连接成功！")
	// 最后一定不要忘记退出登录
	defer c.Logout()

	// 登录
	if err := c.Login(UserName, Password); err != nil {
		log.Fatalf("邮箱[%s] 登录失败: %v \n", Addr, err)
	}
	log.Printf("邮箱[%s] 登录成功！\n", UserName)

	// 列出当前邮箱中的文件夹
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1) // 记录错误的 chan
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	log.Println("-->当前邮箱的文件夹 Mailboxes:")
	var folderExists bool
	for m := range mailboxes {
		log.Println("* ", m.Name)
		if m.Name == Folder {
			folderExists = true
		}
	}
	if err := <-done; err != nil {
		log.Fatalf("列出邮箱列表时，出现错误：%v \n", err)
	}
	log.Println("-->列出邮箱列表完毕！")
	if !folderExists {
		log.Fatalf("文件夹[%s] 不存在 \n", Folder)
	}

	// 选择指定的文件夹
	mbox, err := c.Select(Folder, false)
	if err != nil {
		log.Fatalf("选择邮件箱失败: %v \n", err)
	}
	log.Printf("mbox %+v \n", mbox)
	log.Printf("当前文件夹[%s]中，总共有 %d 封邮件 \n", Folder, mbox.Messages)
	if mbox.Messages == 0 {
		log.Fatalf("当前文件夹[%s]中没有邮件", Folder)
	}

	// 创建一个序列集，用于批量读取邮件
	seqset := new(imap.SeqSet)

	// 假设需要获取最后4封邮件时
	// from := uint32(1)
	// to := mbox.Messages // 此文件下的邮件总数
	// if mbox.Messages > 3 {
	// 	from = mbox.Messages - 3
	// }
	// seqset.AddRange(from, to) // 添加指定范围内的邮件编号

	// 搜索指定状态的邮件
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag} // 未读邮件标记
	// criteria.WithFlags = []string{imap.SeenFlag} // 已读邮件标记
	uids, err := c.Search(criteria)
	// 在这里也可以使用 UidSearch 方法，但是用了 UidSearch 方法后，下面的很多方法都需要使用 Uid 开头的方法
	// 也就是说 Fetch -> UidFetch，Store -> UidStore，Copy -> UidCopy，Move -> UidMove，Search -> UidSearch
	// uids, err := c.UidSearch(criteria)
	// 关于 Store 方法和 UidStore 方法
	// Store 和 UidStore 方法都是用于在 IMAP 中更新邮件标志的，但它们有一些区别：
	//
	// Store：使用的是消息序列号（message sequence number）来标识邮件。序列号是动态的，每次邮件删除或添加时，序列号可能会改变。序列号从1开始，按邮件在邮箱中的位置进行排序。
	// UidStore：使用的是消息的唯一标识符（UID）来标识邮件。UID 是固定的，不会因为邮件的添加或删除而改变，适合于需要确保唯一标识邮件的操作。
	// 在标记为已读时，使用 UidStore 方法更为安全和可靠，因为它使用邮件的唯一标识符，可以避免由于序列号变化导致的潜在问题。
	if err != nil {
		log.Fatalf("搜索邮件时出现错误：%v \n", err)
	}
	log.Printf("搜索到的邮件 uids: %+v \n", uids)
	if len(uids) == 0 {
		log.Println("没有搜索到邮件")
		return
	}
	log.Printf("搜索到的邮件总共有 %v 封 %+v \n", len(uids), uids)

	// 获取整个消息正文
	// imap.FetchEnvelope：请求获取邮件的信封数据（例如发件人、收件人、主题等元数据）。
	// imap.FetchRFC822：请求获取完整的邮件内容，包括所有头部和正文。
	items := []imap.FetchItem{imap.FetchFlags, imap.FetchEnvelope, imap.FetchRFC822}

	for i, uidsCount := 0, len(uids); i < uidsCount; i += ReadBatchSize {
		// 清空序列集中的所有邮件编号，以便添加新的邮件编号。每次循环开始时调用此方法，确保序列集中只有当前批次的邮件编号
		seqset.Clear()

		// 添加一批邮件到序列集中
		if i+ReadBatchSize < uidsCount {
			seqset.AddNum(uids[i : i+ReadBatchSize]...) // 添加指定范围内的邮件编号
		} else {
			seqset.AddNum(uids[i:]...) // 添加剩余的邮件编号
		}

		// 获取邮件内容 Start
		messages := make(chan *imap.Message, ReadBatchSize) // 创建一个通道，用于接收邮件消息
		fetchDone := make(chan error, 1)                    // 创建一个通道，用于接收错误消息
		go func() {
			// Fetch方法用于从服务器获取邮件数据，这里请求了邮件的信封和完整内容
			fetchDone <- c.Fetch(seqset, items, messages)
		}()
		log.Println("开始读取邮件内容")
		for msg := range messages {
			readEveryMsg(msg)
		}
		if err := <-fetchDone; err != nil {
			log.Fatalf("获取邮件信息出现错误：%v \n", err)
		}
		// 获取邮件内容 End

		// 给邮件打标记 Start
		item := imap.FormatFlagsOp(imap.AddFlags, true) // 标记为已读
		// item := imap.FormatFlagsOp(imap.RemoveFlags, true) // 标记为未读
		flags := []interface{}{imap.SeenFlag}
		log.Printf("即将给这些邮件 [%s] 打标记 \n", seqset)
		if err := c.Store(seqset, item, flags, nil); err != nil {
			log.Fatalf("给邮件打标记失败：%v \n", err)
		}
		// 给邮件打标记 End

		time.Sleep(time.Second * 10) // 休眠10秒
	}

	log.Println("读取了所有邮件，完毕！")

}

// document link: https://github.com/emersion/go-imap/wiki/Fetching-messages
func readEveryMsg(msg *imap.Message) {
	log.Printf("每一封邮件的消息序列号 %+v \n", msg.SeqNum)
	log.Println("-------------------------")
	// 获取邮件正文
	r := msg.GetBody(&imap.BodySectionName{})
	if r == nil {
		log.Fatal("服务器没有返回消息内容")
	}

	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatalf("邮件读取时出现错误： %v \n", err)
	}
	if date, err := mr.Header.Date(); err == nil {
		log.Println("收件时间 Date:", date)
	}
	if from, err := mr.Header.AddressList("From"); err == nil {
		log.Println("发件人 From:", from)
	}
	if to, err := mr.Header.AddressList("To"); err == nil {
		log.Println("收件人 To:", to)
	}
	if subject, err := mr.Header.Subject(); err == nil {
		log.Println("邮件主题 Subject:", subject)
	}
	log.Printf("抄送 Cc: %+v \n", msg.Envelope.Cc)

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("读取邮件内容时出现错误：%v \n", err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// 这是消息的文本（可以是纯文本或 HTML）
			contentType := h.Get("Content-Type")
			b, _ := ioutil.ReadAll(p.Body)
			if strings.HasPrefix(contentType, "text/plain") {
				log.Printf("得到正文 -> TEXT: %v \n", string(b))
			} else if strings.HasPrefix(contentType, "text/html") {
				log.Printf("得到正文 -> HTML: %v \n", len(b))
			}
			break
		case *mail.AttachmentHeader:
			// 这是一个附件
			filename, _ := h.Filename()
			log.Printf("得到附件: %v \n", filename)
			break
		}
	}

	log.Println("一封邮件读取完毕")
	log.Printf("------------------------- \n\n")
}
