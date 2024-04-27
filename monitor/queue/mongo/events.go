package mongo

import "time"

type EventName string

const (
	USBillUpdatedEventName EventName = "us-bill-updated"
	USBillCreatedEventName EventName = "us-bill-created"
)

type USBillUpdatedEvent struct {
	PackageID   string    `bson:"package_id"`
	UpdatedAt   time.Time `bson:"updated_at"`
	PackageLink string    `bson:"package_link"`
}

type USBillCreatedEvent struct {
	PackageID   string    `bson:"package_id"`
	UpdatedAt   time.Time `bson:"updated_at"`
	PackageLink string    `bson:"package_link"`
}
