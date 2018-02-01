package sib

/* Request Types */

// SMSCampaign defines attributes of SIB campaigns
type SMSCampaign struct {
	Name           string `json:"name"` // Mandatory
	Sender         string `json:"sender"`
	Content        string `json:"content"`
	Bat_sent       string `json:"bat_sent"`
	List_ids       []int  `json:"listid"` // Mandatory if Scheduled_date
	Exclude_list   []int  `json:"exclude_list"`
	Scheduled_date string `json:"scheduled_date"` // Format: YYYY-MM-DD 00:00:00
	Send_now       int    `json:"send_now"`       // 0 = campaign not ready to send, 1 = ready to send now
}

// SMSRequest defines attributes of an SMS request
type SMSRequest struct {
	To      string `json:"to"`   // Mobile Number (Mandatory)
	From    string `json:"from"` // No more than 11 alphanumeric characters (Mandatory)
	Text    string `json:"text"` // No more than 160 characters (Mandatory)
	Web_url string `json:"web_url"`
	Tag     string `json:"tag"`
	Type    string `json:"type"` // "marketing" (default) or "transactional"
}

// SMSTest ...
type SMSTest struct {
	To string `json:"to"`
}

/* Response Types */

// SMSCampaignData ...
type SMSCampaignData struct {
	Id int `json:"id"`
}

// SMSCampaignResponse ...
type SMSCampaignResponse struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    SMSCampaignData `json:"data"`
}

// SMSData defines attributes of an SMS transaction
type SMSData struct {
	Status           string       `json:"status"`
	Number_sent      int          `json:"number_sent"`
	To               string       `json:"to"`
	Sms_count        int          `json:"sms_count"`
	Credits_used     float64      `json:"credits_used"`
	Remaining_credit float64      `json:"remaining_credit"`
	Reference        SMSReference `json:"reference"`
	Description      string       `json:"description"`
	Reply            string       `json:"reply"`
	Bounce_type      string       `json:"bounce_type"`
	Error_code       int          `json:"error_code"`
}

// SMSReference ...
type SMSReference struct {
	One string `json:"1"`
}

// SMSResponse ...
type SMSResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    SMSData `json:"data"`
}
