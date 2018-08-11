package structure

import (
	"encoding/json"
	"time"
)

// Platform the ad's publishing platform which those are Facebook, Google and Instagram at this moment.
type Platform struct {
	Status                CampaignStatus `json:"status" bson:"status"`
	TotalBudget           int            `json:"total_budget" bson:"totalBudget"`
	RemainingBudget       int            `json:"remaining_budget" bson:"remainingBudget"`
	MillisecondsStartDate int64          `json:"start_date"`
	StartDate             time.Time      `bson:"startDate"`
	MillisecondsEndDate   int64          `json:"end_date"`
	EndDate               time.Time      `bson:"endDate"`
	Audiance              Audiance       `json:"target_audiance" bson:"targetAudiance"` //typo?
	Creatives             Creatives      `json:"creatives" bson:"Creatives"`
	Insights              Insights       `json:"insights" bson:"insights"`
}

// We need our own custom marshaler/unmarshalers both in direction of bson and json this because of `time.Unix` function accepts it's argument only in seconds or nano seconds
// when creating a new 'Time' which we need for our BSON. But our JSON data comes in millisecond format.
// So a conversion must be done some where. For creating a simpler API ( for a decent Domain Driven Development)
// I thought the best place would be the Unmarshaling.

// UnmarshalJSON is json.Unmarshaler implementation for Platform.
func (p *Platform) UnmarshalJSON(b []byte) (err error) {

	p2 := struct {
		Status                CampaignStatus `json:"status" bson:"status"`
		TotalBudget           int            `json:"total_budget" bson:"totalBudget"`
		RemainingBudget       int            `json:"remaining_budget" bson:"remainingBudget"`
		MillisecondsStartDate int64          `json:"start_date"`
		StartDate             time.Time      `bson:"startDate"`
		MillisecondsEndDate   int64          `json:"end_date"`
		EndDate               time.Time      `bson:"endDate"`
		Audiance              Audiance       `json:"target_audiance" bson:"targetAudiance"` //typo?
		Creatives             Creatives      `json:"creatives" bson:"Creatives"`
		Insights              Insights       `json:"insights" bson:"insights"`
	}{}

	err = json.Unmarshal(b, &p2)
	if err != nil {
		return
	}

	p2.StartDate = time.Unix(0, p2.MillisecondsStartDate*int64(time.Millisecond/time.Nanosecond))
	p2.EndDate = time.Unix(0, p2.MillisecondsEndDate*int64(time.Millisecond/time.Nanosecond))

	p.Status = p2.Status
	p.TotalBudget = p2.TotalBudget
	p.RemainingBudget = p2.RemainingBudget
	p.MillisecondsStartDate = p2.MillisecondsStartDate
	p.StartDate = p2.StartDate
	p.MillisecondsEndDate = p2.MillisecondsEndDate
	p.EndDate = p2.EndDate
	p.Audiance = p2.Audiance
	p.Creatives = p2.Creatives
	p.Insights = p2.Insights

	return
}

// MarshalJSON is json.Marshaler implementation.
func (p *Platform) MarshalJSON() (b []byte, err error) {

	p.MillisecondsStartDate = p.StartDate.Unix() * int64(time.Second/time.Millisecond)
	p.MillisecondsEndDate = p.EndDate.Unix() * int64(time.Second/time.Millisecond)

	b, err = json.Marshal(p)
	return
}
