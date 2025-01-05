package twilioOTP

import (
	"fmt"

	"github.com/twilio/twilio-go"

	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type SmsOtpClient interface {
	SendOtp(phone string) error
	VerifyOtp(phone string, otp string) (bool, error)
}

type TwilioClient struct {
	bypassMode bool
	client     *twilio.RestClient
	serviceSid string
}

func NewTwilioClient(accountSid, authToken, serviceSid string, byPassTwilio bool) SmsOtpClient {
	return &TwilioClient{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		}),
		serviceSid: serviceSid,
		bypassMode: byPassTwilio,
	}
}

func (tc *TwilioClient) SendOtp(phone string) error {
	if tc.bypassMode {
		fmt.Println("ByPass mode is turned on in environment. Skipping sending of otp")
		return nil
	}
	fmt.Println("Sending OTP")
	params := &verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")
	resp, err := tc.client.VerifyV2.CreateVerification(tc.serviceSid, params)
	if err != nil {
		return err
	} else {
		if resp.Status != nil { //?
			switch *resp.Status {
			case "pending":
				fmt.Println("OTP had been sent already")
			case "max_attempts_reached":
				fmt.Println("Max attempts reached")
			case "failed":
				fmt.Println("Failed")
				return fmt.Errorf("Failed to send OTP")
			}

			if *resp.Status == "pending" {
				// fmt.Println("OTP had been sent already")
				fmt.Print("") //blank code //implement later if needed
			}
		}
		return nil
	}
}

func (tc *TwilioClient) VerifyOtp(phone string, otp string) (bool, error) {
	if tc.bypassMode {
		fmt.Println("ByPass mode is turned on in environment. Skipping sending of otp")
		return true, nil
	}
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(otp)

	resp, err := tc.client.VerifyV2.CreateVerificationCheck(tc.serviceSid, params)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	} else {
		if resp.Status != nil {
			switch *resp.Status {
			//pending, approved, canceled, max_attempts_reached, deleted, failed or expired.
			case "approved":
				fmt.Println("OTP verification approved")
				return true, nil
			default:
				fmt.Println("OTP verification failed: ", *resp.Status)
				return false, nil
			}
		} else {
			fmt.Println("resp status:", resp.Status)
		}
		return true, nil
	}
}
