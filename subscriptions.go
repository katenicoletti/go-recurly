package recurly

import (
	"encoding/xml"

	"github.com/blacklightcms/go-recurly/types"
)

type (
	// Subscription represents an individual subscription.
	Subscription struct {
		XMLName                xml.Name            `xml:"subscription"`
		Plan                   NestedPlan          `xml:"plan,omitempty"`
		AccountCode            string              `xml:"-"`
		InvoiceNumber          int                 `xml:"-"`
		UUID                   string              `xml:"uuid,omitempty"`
		State                  string              `xml:"state,omitempty"`
		UnitAmountInCents      int                 `xml:"unit_amount_in_cents,omitempty"`
		Currency               string              `xml:"currency,omitempty"`
		Quantity               int                 `xml:"quantity,omitempty"`
		ActivatedAt            types.NullTime      `xml:"activated_at,omitempty"`
		CanceledAt             types.NullTime      `xml:"canceled_at,omitempty"`
		ExpiresAt              types.NullTime      `xml:"expires_at,omitempty"`
		CurrentPeriodStartedAt types.NullTime      `xml:"current_period_started_at,omitempty"`
		CurrentPeriodEndsAt    types.NullTime      `xml:"current_period_ends_at,omitempty"`
		TrialStartedAt         types.NullTime      `xml:"trial_started_at,omitempty"`
		TrialEndsAt            types.NullTime      `xml:"trial_ends_at,omitempty"`
		TaxInCents             int                 `xml:"tax_in_cents,omitempty"`
		TaxType                string              `xml:"tax_type,omitempty"`
		TaxRegion              string              `xml:"tax_region,omitempty"`
		TaxRate                float64             `xml:"tax_rate,omitempty"`
		PONumber               string              `xml:"po_number,omitempty"`
		NetTerms               types.NullInt       `xml:"net_terms,omitempty"`
		SubscriptionAddOns     []SubscriptionAddOn `xml:"subscription_add_ons>subscription_add_on,omitempty"`
	}

	NestedPlan struct {
		Code string `xml:"plan_code,omitempty"`
		Name string `xml:"name,omitempty"`
	}

	// SubscriptionAddOn are add ons to subscriptions.
	// https://docs.recurly.com/api/subscriptions/subscription-add-ons
	SubscriptionAddOn struct {
		XMLName           xml.Name `xml:"subscription_add_on"`
		Type              string   `xml:"add_on_type,omitempty"`
		Code              string   `xml:"add_on_code"`
		UnitAmountInCents int      `xml:"unit_amount_in_cents"`
		Quantity          int      `xml:"quantity,omitempty"`
	}

	// NewSubscription is used to create new subscriptions.
	NewSubscription struct {
		XMLName                 xml.Name             `xml:"subscription"`
		PlanCode                string               `xml:"plan_code"`
		Account                 Account              `xml:"account"`
		SubscriptionAddOns      *[]SubscriptionAddOn `xml:"subscription_add_ons>subscription_add_on,omitempty"`
		CouponCode              string               `xml:"coupon_code,omitempty"`
		UnitAmountInCents       int                  `xml:"unit_amount_in_cents,omitempty"`
		Currency                string               `xml:"currency"`
		Quantity                int                  `xml:"quantity,omitempty"`
		TrialEndsAt             types.NullTime       `xml:"trial_ends_at,omitempty"`
		StartsAt                types.NullTime       `xml:"starts_at,omitempty"`
		TotalBillingCycles      int                  `xml:"total_billing_cycles,omitempty"`
		FirstRenewalDate        types.NullTime       `xml:"first_renewal_date,omitempty"`
		CollectionMethod        string               `xml:"collection_method,omitempty"`
		NetTerms                types.NullInt        `xml:"net_terms,omitempty"`
		PONumber                string               `xml:"po_number,omitempty"`
		Bulk                    bool                 `xml:"bulk,omitempty"`
		TermsAndConditions      string               `xml:"terms_and_conditions,omitempty"`
		CustomerNotes           string               `xml:"customer_notes,omitempty"`
		VATReverseChargeNotes   string               `xml:"vat_reverse_charge_notes,omitempty"`
		BankAccountAuthorizedAt types.NullTime       `xml:"bank_account_authorized_at,omitempty"`
	}

	// UpdateSubscription is used to update subscriptions
	UpdateSubscription struct {
		XMLName            xml.Name             `xml:"subscription"`
		Timeframe          string               `xml:"timeframe,omitempty"`
		PlanCode           string               `xml:"plan_code,omitempty"`
		Quantity           int                  `xml:"quantity,omitempty"`
		UnitAmountInCents  int                  `xml:"unit_amount_in_cents,omitempty"`
		CollectionMethod   string               `xml:"collection_method,omitempty"`
		NetTerms           types.NullInt        `xml:"net_terms,omitempty"`
		PONumber           string               `xml:"po_number,omitempty"`
		SubscriptionAddOns *[]SubscriptionAddOn `xml:"subscription_add_ons>subscription_add_on,omitempty"`
	}

	// SubscriptionNotes is used to update a subscription's notes.
	SubscriptionNotes struct {
		XMLName               xml.Name `xml:"subscription"`
		TermsAndConditions    string   `xml:"terms_and_conditions,omitempty"`
		CustomerNotes         string   `xml:"customer_notes,omitempty"`
		VATReverseChargeNotes string   `xml:"vat_reverse_charge_notes,omitempty"`
	}
)

const (
	// SubscriptionStateActive represents subscriptions that are valid for the
	// current time. This includes subscriptions in a trial period
	SubscriptionStateActive = "active"

	// SubscriptionStateCanceled are subscriptions that are valid for
	// the current time but will not renew because a cancelation was requested
	SubscriptionStateCanceled = "canceled"

	// SubscriptionStateExpired are subscriptions that have expired and are no longer valid
	SubscriptionStateExpired = "expired"

	// SubscriptionStateFuture are subscriptions that will start in the
	// future, they are not active yet
	SubscriptionStateFuture = "future"

	// SubscriptionStateInTrial are subscriptions that are active or canceled
	// and are in a trial period
	SubscriptionStateInTrial = "in_trial"

	// SubscriptionStateLive are all subscriptions that are not expired
	SubscriptionStateLive = "live"

	// SubscriptionStatePastDue are subscriptions that are active or canceled
	// and have a past-due invoice
	SubscriptionStatePastDue = "past_due"
)

// UnmarshalXML unmarshals transactions and handles intermediary state during unmarshaling
// for types like href.
func (s *Subscription) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName                xml.Name            `xml:"subscription"`
		Plan                   NestedPlan          `xml:"plan,omitempty"`
		AccountCode            types.HrefString    `xml:"account"`
		InvoiceNumber          types.HrefInt       `xml:"invoice"`
		UUID                   string              `xml:"uuid,omitempty"`
		State                  string              `xml:"state,omitempty"`
		UnitAmountInCents      int                 `xml:"unit_amount_in_cents,omitempty"`
		Currency               string              `xml:"currency,omitempty"`
		Quantity               int                 `xml:"quantity,omitempty"`
		ActivatedAt            types.NullTime      `xml:"activated_at,omitempty"`
		CanceledAt             types.NullTime      `xml:"canceled_at,omitempty"`
		ExpiresAt              types.NullTime      `xml:"expires_at,omitempty"`
		CurrentPeriodStartedAt types.NullTime      `xml:"current_period_started_at,omitempty"`
		CurrentPeriodEndsAt    types.NullTime      `xml:"current_period_ends_at,omitempty"`
		TrialStartedAt         types.NullTime      `xml:"trial_started_at,omitempty"`
		TrialEndsAt            types.NullTime      `xml:"trial_ends_at,omitempty"`
		TaxInCents             int                 `xml:"tax_in_cents,omitempty"`
		TaxType                string              `xml:"tax_type,omitempty"`
		TaxRegion              string              `xml:"tax_region,omitempty"`
		TaxRate                float64             `xml:"tax_rate,omitempty"`
		PONumber               string              `xml:"po_number,omitempty"`
		NetTerms               types.NullInt       `xml:"net_terms,omitempty"`
		SubscriptionAddOns     []SubscriptionAddOn `xml:"subscription_add_ons>subscription_add_on,omitempty"`
	}
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	*s = Subscription{
		XMLName:                v.XMLName,
		Plan:                   v.Plan,
		AccountCode:            string(v.AccountCode),
		InvoiceNumber:          int(v.InvoiceNumber),
		UUID:                   v.UUID,
		State:                  v.State,
		UnitAmountInCents:      v.UnitAmountInCents,
		Currency:               v.Currency,
		Quantity:               v.Quantity,
		ActivatedAt:            v.ActivatedAt,
		CanceledAt:             v.CanceledAt,
		ExpiresAt:              v.ExpiresAt,
		CurrentPeriodStartedAt: v.CurrentPeriodStartedAt,
		CurrentPeriodEndsAt:    v.CurrentPeriodEndsAt,
		TrialStartedAt:         v.TrialStartedAt,
		TrialEndsAt:            v.TrialEndsAt,
		TaxInCents:             v.TaxInCents,
		TaxType:                v.TaxType,
		TaxRegion:              v.TaxRegion,
		TaxRate:                v.TaxRate,
		PONumber:               v.PONumber,
		NetTerms:               v.NetTerms,
		SubscriptionAddOns:     v.SubscriptionAddOns,
	}

	return nil
}

// MakeUpdate creates an UpdateSubscription with values that need to be passed
// on update to be retained (meaning nil/zero values will delete that value).
// After calling MakeUpdate you should modify the struct with your updates.
// Once you're ready you can call client.Subscriptions.Update
func (s Subscription) MakeUpdate() UpdateSubscription {
	return UpdateSubscription{
		// NetTerms need to be copied over because on update they default to 0.
		// This ensures the NetTerms don't get overridden.
		NetTerms:           s.NetTerms,
		SubscriptionAddOns: &s.SubscriptionAddOns,
	}
}
