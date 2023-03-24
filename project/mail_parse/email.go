package mail_parse

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

const (
	Addr     string = "imap.qq.com:993"
	UserName string = "1414818093@qq.com" // 邮箱地址
	Password string = ""                  // 这里的密码是使用开启 imap 协议后对应的服务商给到的密码，不是邮箱账号密码
)

// IMAP（Internet Message Access Protocol）是一种用于在互联网上访问电子邮件的协议。
// 它允许用户通过 Internet 访问他们在邮件服务器上存储的电子邮件。
// Go 语言的 go-imap 库是一个用于从 IMAP 服务器获取电子邮件的库，它可以帮助你在 Go 代码中访问 IMAP 协议

func ReadEmail() {
	log.Println("开始连接服务器")

	// 建立与 IMAP 服务器的连接
	c, err := client.DialTLS(Addr, nil)
	if err != nil {
		log.Fatalf("连接 IMAP 服务器失败: %v \n", err)
	}
	log.Println("连接成功！")
	// 最后退出登录
	defer c.Logout()

	// 登录
	if err := c.Login(UserName, Password); err != nil {
		log.Fatalf("登录失败: %v \n", err)
	}
	log.Println("登录成功！")

	// 列出邮箱
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1) // 记录错误的 chan
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}
	if err := <-done; err != nil {
		log.Fatalf("列出邮箱列表时，出现错误：%v \n", err)
	}

	// 选择收件箱
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatalf("选择邮件箱失败: %v \n", err)
	}
	log.Println("邮箱的标记为：", mbox.Flags)
	if mbox.Messages == 0 {
		log.Fatal("没有邮件")
	}

	// 获取最后 4 条消息
	from := uint32(1)
	to := mbox.Messages
	log.Printf("总共有 %v 封邮件 \n", to)
	if mbox.Messages > 3 {
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)
	// seqset.AddNum(mbox.Messages)

	// 获取整个消息正文
	var section imap.BodySectionName
	// []imap.FetchItem{imap.FetchEnvelope, imap.FetchBody}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	log.Println("开始读取内容")
	// document link: https://github.com/emersion/go-imap/wiki/Fetching-messages
	for msg := range messages {
		// log.Println("邮件主题： " + msg.Envelope.Subject)
		// spew.Dump(msg)

		log.Println("==============")

		r := msg.GetBody(&section)
		if r == nil {
			log.Fatal("服务器没有返回消息内容")
		}

		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Printf("邮件读取时出现错误： %v \n", err)
			continue
		}
		header := mr.Header
		if date, err := header.Date(); err == nil {
			log.Println("Date:", date)
		}
		if from, err := header.AddressList("From"); err == nil {
			log.Println("From:", from)
		}
		if to, err := header.AddressList("To"); err == nil {
			log.Println("To:", to)
		}
		if subject, err := header.Subject(); err == nil {
			log.Println("Subject:", subject)
		}

		// Process each message's part
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				// This is the message's text (can be plain-text or HTML)
				b, _ := ioutil.ReadAll(p.Body)
				log.Printf("Got text: %v \n", string(b))
			case *mail.AttachmentHeader:
				// This is an attachment
				filename, _ := h.Filename()
				log.Printf("Got attachment: %v \n", filename)
			}
		}

		log.Println("一条读取完毕")

	}

	if err := <-done; err != nil {
		log.Fatalf("读取邮件内容出现错误：%v \n", err)
	}

	log.Println("Done!")

}
