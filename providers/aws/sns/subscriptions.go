package sns
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)


type Subscription struct {
	ID uint `gorm:"primarykey"`
	AccountID string
	Region string
	Endpoint *string
	Owner *string
	Protocol *string
	SubscriptionArn *string
	TopicArn *string
}

func (Subscription) TableName() string {
	return "aws_sns_subscriptions"
}

func (c *Client) transformSubscriptions(values []*sns.Subscription) []*Subscription {
	var tValues []*Subscription
	for _, value := range values {
		tValue := Subscription {
			AccountID: c.accountID,
			Region: c.region,
			Endpoint: value.Endpoint,
			Owner: value.Owner,
			Protocol: value.Protocol,
			SubscriptionArn: value.SubscriptionArn,
			TopicArn: value.TopicArn,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
type SubscriptionConfig struct {
	Filter string
}

var SubscriptionTables = []interface{} {
	&Subscription{},
}

func (c *Client)subscriptions(gConfig interface{}) error {
	var config sns.ListSubscriptionsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(SubscriptionTables...)

	for {
		output, err := c.svc.ListSubscriptions(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformSubscriptions(output.Subscriptions))
		c.log.Info("Fetched resources", zap.String("resource", "sns.subscriptions"), zap.Int("count", len(output.Subscriptions)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}

