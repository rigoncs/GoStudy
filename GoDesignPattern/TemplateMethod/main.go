package templatemethod

func main() {
	smsOtp := &Sms{}
	o := Otp{
		iOtp: smsOtp,
	}
	o.genAndSendOTP(4)

	emailOtp := &Email{}
	o = Otp{iOtp: emailOtp}
	o.genAndSendOTP(4)
}
