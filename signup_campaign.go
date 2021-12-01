package main

import (
	"github.com/mailgun/mailgun-go/v4"
	"gomodules.xyz/mailer"
	"google.golang.org/api/sheets/v4"
	"time"
)

func NewEnterpriseSignupCampaign(srv *sheets.Service, mg mailgun.Mailgun) *mailer.DripCampaign {
	return &mailer.DripCampaign{
		Name: "New Signup",
		Steps: []mailer.CampaignStep{
			{
				WaitTime: 0,
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "Welcome to {{.Product}}",
					Body:            "Hey {{.Name}}, Thanks for using {{.Product}}!",
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
			{
				WaitTime: 10 * time.Second,
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "How is it going with {{.ProductDisplayName}}?",
					Body:            `Hi {{.Name}},

I hope you are doing well. Just wanted to check your progress with {{.ProductDisplayName}} so far. We want to make sure your journey with {{.ProductDisplayName}} is going smoothly so far.

For support, please mail us at support@appscode.com. Our engineers are be happy to help you during the evaluation period. 

Warm Regards,
Team AppsCode

[![Website](https://cdn.appscode.com/images/website.png)](https://appscode.com) [![Linkedin](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/ln.png)](https://www.linkedin.com/company/appscode/) [![Twitter](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/tt.png)](https://twitter.com/AppsCodeHQ) [![Youtube](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/yt.png)](https://www.youtube.com/c/AppsCodeInc)
`,
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
			{
				WaitTime: 5 * 24 * time.Hour, // 5 days
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "Keeping Up With {{.ProductDisplayName}}",
					Body:           `Hi {{.Name}},

Thanks again for trying {{.ProductDisplayName}}. Did you know we have a [LinkedIn](https://www.linkedin.com/company/appscode/) and [Twitter](https://twitter.com/{{.TwitterHandle}}) handle? Connect with us to be up to date with AppsCode. You can also subscribe to our [YouTube](https://www.youtube.com/c/appscodeinc) channel.

Warm Regards,
Team AppsCode

[![Website](https://cdn.appscode.com/images/website.png)](https://appscode.com) [![Linkedin](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/ln.png)](https://www.linkedin.com/company/appscode/) [![Twitter](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/tt.png)](https://twitter.com/AppsCodeHQ) [![Youtube](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/yt.png)](https://www.youtube.com/c/AppsCodeInc)
`,
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
			{
				WaitTime: 25 * 24 * time.Hour, // 25 days
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "Your {{.ProductDisplayName}} trial ending soon",
					Body:            `Hi {{.Name}},

You are almost at the end of your trial period for {{.Plan}}. Would you like to extend your trial? Or maybe you would like to get a full Enterprise license? For an Enterprise price quote, please reach us [here](https://appscode.com/contact/).

Or you can also email us at sales@appscode.com.

Warm Regards,
Team AppsCode

[![Website](https://cdn.appscode.com/images/website.png)](https://appscode.com) [![Linkedin](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/ln.png)](https://www.linkedin.com/company/appscode/) [![Twitter](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/tt.png)](https://twitter.com/AppsCodeHQ) [![Youtube](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/yt.png)](https://www.youtube.com/c/AppsCodeInc)
`,
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
			{
				WaitTime: 30 * 24 * time.Hour, // 30 days
				Mailer: mailer.Mailer{
					Sender:          MailHello,
					BCC:             MailLicenseTracker,
					ReplyTo:         MailHello,
					Subject:         "{{.ProductDisplayName}} next steps",
					Body:            `Hi {{.Name}},

Congratulations on reaching the end of your trial period with {{.ProductDisplayName}}. Here on, we would love to hear about your journey with our product. And how would you like to move forward with us? For an Enterprise price quote, please reach us [here](https://appscode.com/contact/).

Or email us at sales@appscode.com.

Warm Regards,
Team AppsCode

[![Website](https://cdn.appscode.com/images/website.png)](https://appscode.com) [![Linkedin](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/ln.png)](https://www.linkedin.com/company/appscode/) [![Twitter](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/tt.png)](https://twitter.com/AppsCodeHQ) [![Youtube](https://codetwocdn.azureedge.net/images/mail-signatures/generator-dm/pad-box/yt.png)](https://www.youtube.com/c/AppsCodeInc)
`,
					Params:          nil,
					AttachmentBytes: nil,
					GDriveFiles:     nil,
					GoogleDocIds:    nil,
					EnableTracking:  true,
				},
			},
		},
		M:             mg,
		SheetService:  srv,
		SpreadsheetId: spreadsheetId,
		SheetName:     sheetName,
	}
}
